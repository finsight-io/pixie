/*
 * Copyright 2018- The Pixie Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

#include "src/stirling/source_connectors/socket_tracer/protocols/http/utils.h"

#include <utility>

namespace px {
namespace stirling {
namespace protocols {
namespace http {

bool MatchesHTTPHeaders(const HeadersMap& http_headers, const HTTPHeaderFilter& filter) {
  if (!filter.inclusions.empty()) {
    bool included = false;
    for (auto [http_header, substr] : filter.inclusions) {
      auto http_header_iter = http_headers.find(std::string(http_header));
      if (http_header_iter != http_headers.end() &&
          absl::StrContains(http_header_iter->second, substr)) {
        included = true;
        break;
      }
    }
    if (!included) {
      return false;
    }
  }
  // For symmetry with the above if block and safety in case of copy-paste, we put exclusions search
  // also inside a if statement, which is not needed for correctness.
  if (!filter.exclusions.empty()) {
    bool excluded = false;
    for (auto [http_header, substr] : filter.exclusions) {
      auto http_header_iter = http_headers.find(std::string(http_header));
      if (http_header_iter != http_headers.end() &&
          absl::StrContains(http_header_iter->second, substr)) {
        excluded = true;
        break;
      }
    }
    if (excluded) {
      return false;
    }
  }
  return true;
}

HTTPHeaderFilter ParseHTTPHeaderFilters(std::string_view filters) {
  HTTPHeaderFilter result;
  for (std::string_view header_filter : absl::StrSplit(filters, ",", absl::SkipEmpty())) {
    std::pair<std::string_view, std::string_view> header_substr =
        absl::StrSplit(header_filter, absl::MaxSplits(":", 1));
    if (absl::StartsWith(header_substr.first, "-")) {
      header_substr.first.remove_prefix(1);
      result.exclusions.emplace(header_substr);
    } else {
      result.inclusions.emplace(header_substr);
    }
  }
  return result;
}

// Lookup table for fast conversion of hex character to decimal value
static constexpr unsigned char hextable[] = {
    0, 1,  2,  3,  4,  5,  6,  7, 8, 9, 0, 0, 0, 0, 0, 0, /* 0x30 - 0x3f */
    0, 10, 11, 12, 13, 14, 15, 0, 0, 0, 0, 0, 0, 0, 0, 0, /* 0x40 - 0x4f */
    0, 0,  0,  0,  0,  0,  0,  0, 0, 0, 0, 0, 0, 0, 0, 0, /* 0x50 - 0x5f */
    0, 10, 11, 12, 13, 14, 15                             /* 0x60 - 0x66 */
};

#define HEX_TO_DEC(x) hextable[x - '0']

// Checks to see if the char is a valid hex digit [A-Fa-f0-9]
// Used to check that HEX_TO_DEC is safe to apply for the given char
bool IsHexDigit(char c) {
  return (c >= 0x30 && c <= 0x39) || (c >= 0x41 && c <= 0x46) || (c >= 0x61 && c <= 0x66);
}

std::string HTTPUrlDecode(const std::string_view input) {
  std::string output;
  output.reserve(input.size());
  size_t pos = 0;
  size_t end = input.size();
  while (pos < end) {
    auto c = input.at(pos);
    if (c == '+') {
      output.push_back(' ');
      pos += 1;
    } else if (c == '%' && (end - pos >= 2) && IsHexDigit(input[pos + 1]) &&
               IsHexDigit(input[pos + 2])) {
      output.push_back((unsigned char)(HEX_TO_DEC(input[pos + 1]) << 4) |
                       HEX_TO_DEC(input[pos + 2]));
      pos += 3;
    } else {
      output.push_back(c);
      pos += 1;
    }
  }
  return output;
}

bool IsJSONContent(const Message& message) {
  auto content_type_iter = message.headers.find(kContentType);
  if (content_type_iter == message.headers.end()) {
    return false;
  }
  if (absl::StrContains(content_type_iter->second, "json")) {
    return true;
  }
  return false;
}

}  // namespace http
}  // namespace protocols
}  // namespace stirling
}  // namespace px
