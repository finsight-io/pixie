import clsx from 'clsx';
import * as React from 'react';

import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';

import { scrollbarStyles } from 'mui-theme';
import useIsMounted from 'utils/use-is-mounted';

import {
  CompletionId,
  CompletionItem,
  CompletionItems,
  CompletionTitle,
  Completions,
} from './completions';
import { Input } from './input';
import { Key } from './key';

const useStyles = makeStyles((theme: Theme) => createStyles({
  root: {
    backgroundColor: theme.palette.background.three,
    cursor: 'text',
    display: 'flex',
    flexDirection: 'column',
    // TODO(malthus): remove this once the scrollbar theme is set at global level.
    ...scrollbarStyles(theme),
  },
  input: {
    backgroundColor: theme.palette.background.two,
  },
  completions: {
    flex: 1,
    minHeight: 0,
  },
}));

export interface AutocompleteContextProps {
  /** @default true */
  allowTyping?: boolean;
  /** @default true */
  requireCompletion?: boolean;
  /**
   * What to do when an associated input is "opened" - like first being created, or being revealed in a dropdown.
   * 'none': don't respond to the input being opened
   * 'focus': moves tab focus to the input
   * 'select': moves tab focus to the input, then selects its contents as if Ctrl/Cmd+A had been pressed.
   * 'clear': moves tab focus to the input, and sets its value to the empty string.
   * @default focus
   */
  openMode?: 'none'|'focus'|'select'|'clear';
  inputRef?: React.MutableRefObject<HTMLInputElement>;
}

export const AutocompleteContext = React.createContext<
AutocompleteContextProps
>({
  allowTyping: true,
  requireCompletion: true,
  openMode: 'none',
});

interface AutoCompleteProps {
  onSelection: (id: CompletionId) => void;
  getCompletions: (input: string) => Promise<CompletionItems>;
  placeholder?: string;
  prefix?: React.ReactNode;
  className?: string;
}

type ItemsMap = Map<CompletionId, { title: CompletionTitle; index: number }>;

// Scrolls through the list, ignoring autoSelectPriority as the user is interacting directly.
function findNextItem(
  activeItem: CompletionId,
  itemsMap: ItemsMap,
  completions: CompletionItems,
  direction = 1,
): CompletionId {
  if (!activeItem || completions.length === 0) {
    return '';
  }
  const { index } = itemsMap.get(activeItem);
  const { length } = completions;
  for (let i = 1; i < length; i++) {
    const nextIndex = (index + i * direction + length) % length;
    const next = completions[nextIndex] as CompletionItem;
    if (next.title && next.id) {
      return next.id;
    }
  }
  return activeItem;
}

// Honors autoSelectPriority
function autoSelectItem(completions: CompletionItems): CompletionItem|null {
  const items: CompletionItem[] = completions.filter((c) => c.type === 'item') as CompletionItem[];
  const highestPriority = items.reduce((hi, cmpl) => Math.max(hi, cmpl.autoSelectPriority ?? 0), -Infinity);
  const selectable = items.filter((c) => (c.autoSelectPriority ?? 0) === highestPriority);
  return selectable[0] ?? null;
}

export const Autocomplete: React.FC<AutoCompleteProps> = ({
  onSelection,
  getCompletions,
  placeholder,
  prefix,
  className,
}) => {
  const classes = useStyles();
  const mounted = useIsMounted();
  const {
    allowTyping,
    requireCompletion,
    inputRef,
  } = React.useContext(AutocompleteContext);
  const [inputValue, setInputValue] = React.useState('');
  const [completions, setCompletions] = React.useState([]);
  const [activeItem, setActiveItem] = React.useState<CompletionId>('');
  const itemsMap = React.useMemo(() => {
    const map: ItemsMap = new Map();
    completions.forEach((item, index) => {
      if (!item.header) {
        map.set(item.id, { title: item.title, index });
      }
    });
    return map;
  }, [completions]);

  React.useEffect(() => {
    getCompletions(inputValue).then((cmpls) => {
      if (!mounted.current) return;
      setCompletions(cmpls);
      const selection = autoSelectItem(cmpls);
      if (selection?.title && selection?.id) setActiveItem(selection.id);
    });
    // `mounted` is not in this array because changing it should ONLY affect whether the resolved promise acts.
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [inputValue, getCompletions]);

  const handleSelection = React.useCallback(
    (id) => {
      const item = itemsMap.get(id);
      if (!item && requireCompletion) {
        return;
      }
      if (item) {
        onSelection(id);
        setInputValue(item.title);
      } else {
        onSelection(inputValue);
      }
    },
    [inputValue, requireCompletion, itemsMap, onSelection],
  );

  const handleKey = (key: Key) => {
    switch (key) {
      case 'UP':
        setActiveItem(findNextItem(activeItem, itemsMap, completions, -1));
        break;
      case 'DOWN':
        setActiveItem(findNextItem(activeItem, itemsMap, completions));
        break;
      case 'ENTER':
        handleSelection(activeItem);
        break;
      default:
        break;
    }
  };

  return (
    <div className={clsx(classes.root, className)}>
      {allowTyping !== false && (
        <Input
          className={classes.input}
          onChange={setInputValue}
          onKey={handleKey}
          value={inputValue}
          placeholder={placeholder}
          prefix={prefix}
          suggestion={itemsMap.get(activeItem)?.title || ''}
          customRef={inputRef}
        />
      )}
      <Completions
        className={classes.completions}
        items={completions}
        onActiveChange={setActiveItem}
        onSelection={handleSelection}
        activeItem={activeItem}
      />
    </div>
  );
};
