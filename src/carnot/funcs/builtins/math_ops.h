#pragma once

#include <cmath>
#include <limits>

#include "src/carnot/udf/registry.h"
#include "src/carnot/udf/type_inference.h"
#include "src/shared/types/types.h"

namespace pl {
namespace carnot {
namespace builtins {

udf::ScalarUDFDocBuilder AddDoc();
template <typename TReturn, typename TArg1, typename TArg2>
class AddUDF : public udf::ScalarUDF {
 public:
  TReturn Exec(FunctionContext*, TArg1 b1, TArg2 b2) { return b1.val + b2.val; }
  static udf::InfRuleVec SemanticInferenceRules() {
    return {
        udf::InheritTypeFromArgs<AddUDF>::Create({types::ST_BYTES, types::ST_THROUGHPUT_PER_NS,
                                                  types::ST_THROUGHPUT_BYTES_PER_NS,
                                                  types::ST_DURATION_NS}),
    };
  }

  static udf::ScalarUDFDocBuilder Doc() { return AddDoc(); }
};

template <>
class AddUDF<types::StringValue, types::StringValue, types::StringValue> : public udf::ScalarUDF {
 public:
  types::StringValue Exec(FunctionContext*, types::StringValue b1, types::StringValue b2) {
    return b1 + b2;
  }

  static udf::ScalarUDFDocBuilder Doc() { return AddDoc(); }
};

template <typename TReturn, typename TArg1, typename TArg2>
class SubtractUDF : public udf::ScalarUDF {
 public:
  TReturn Exec(FunctionContext*, TArg1 b1, TArg2 b2) { return b1.val - b2.val; }
  static udf::InfRuleVec SemanticInferenceRules() {
    return {
        udf::InheritTypeFromArgs<SubtractUDF>::Create({types::ST_BYTES, types::ST_THROUGHPUT_PER_NS,
                                                       types::ST_THROUGHPUT_BYTES_PER_NS,
                                                       types::ST_DURATION_NS}),
    };
  }
  static udf::ScalarUDFDocBuilder Doc() {
    return udf::ScalarUDFDocBuilder(
               "Arithmetically subtract the first argument by the second argument.")
        .Details("This function is implicitly invoked by the - operator.")
        .Example(R"doc(# Implicit call.
        | df.subtracted = df.a - df.b
        | # Explicit call.
        | df.subtracted = px.subtract(df.a, df.b)
        )doc")
        .Arg("a", "The value to be subtracted from.")
        .Arg("b", "The value to subtract.")
        .Returns("`a` with `b` subtracted from it.");
  }
};

template <typename TReturn, typename TArg1, typename TArg2>
class DivideUDF : public udf::ScalarUDF {
  using ReturnValueType = typename types::ValueTypeTraits<TReturn>::native_type;

 public:
  TReturn Exec(FunctionContext*, TArg1 b1, TArg2 b2) {
    return ReturnValueType(b1.val) / ReturnValueType(b2.val);
  }

  static udf::InfRuleVec SemanticInferenceRules() {
    return {udf::ExplicitRule::Create<DivideUDF>(types::ST_THROUGHPUT_PER_NS,
                                                 {types::ST_NONE, types::ST_DURATION_NS}),
            udf::ExplicitRule::Create<DivideUDF>(types::ST_THROUGHPUT_BYTES_PER_NS,
                                                 {types::ST_BYTES, types::ST_DURATION_NS})};
  }

  static udf::ScalarUDFDocBuilder Doc() {
    return udf::ScalarUDFDocBuilder("Arithmetically divide the two arguments.")
        .Details("This function is implicitly invoked by the / operator.")
        .Example(R"doc(# Implicit call.
        | df.div = df.a / df.b
        | # Explicit call.
        | df.div = px.divide(df.a, df.b))doc")
        .Arg("arg1", "The value to divide.")
        .Arg("arg2", "The value to divide the first argument by.")
        .Returns("The value of arg1 divided by arg2.");
  }
};

template <typename TReturn, typename TArg1, typename TArg2>
class MultiplyUDF : public udf::ScalarUDF {
 public:
  TReturn Exec(FunctionContext*, TArg1 b1, TArg2 b2) { return b1.val * b2.val; }
  static udf::ScalarUDFDocBuilder Doc() {
    return udf::ScalarUDFDocBuilder("Multiplies the arguments.")
        .Details("Multiplies the two values together. Accessible using the `*` operator syntax.")
        .Example(R"doc(# Implicit call.
        | df.mult = df.duration * 2
        | # Explicit call.
        | df.mult = px.multiply(df.duration, 2)
        )doc")
        .Arg("a", "The first value to multiply.")
        .Arg("b", "The second value to multiply.")
        .Returns("The product of `a` and `b`.");
  }
};

template <typename TReturn, typename TArg1, typename TArg2>
class ModuloUDF : public udf::ScalarUDF {
 public:
  TReturn Exec(FunctionContext*, TArg1 b1, TArg2 b2) { return b1.val % b2.val; }
  static udf::ScalarUDFDocBuilder Doc() {
    return udf::ScalarUDFDocBuilder("Calculates the remainder of the division of the two numbers")
        .Details(
            "Calculates the remainder of dividing the first argument by the second argument. "
            "Same "
            "results as the C++ modulo operator. Accessible using the `%` syntax.")
        .Example(R"doc(# Implicit call.
        | df.duration_mod_5s = df.duration % px.seconds(5)
        | # Explicit call.
        | df.duration_mod_5s = px.modulo(df.duration, px.seconds(5))
        )doc")
        .Arg("a", "The dividend.")
        .Arg("n", "The divisor.")
        .Returns("The remainder of dividing `a` by `n`");
  }
};

template <typename TArg1, typename TArg2>
class LogicalOrUDF : public udf::ScalarUDF {
 public:
  BoolValue Exec(FunctionContext*, TArg1 b1, TArg2 b2) { return b1.val || b2.val; }
  static udf::ScalarUDFDocBuilder Doc() {
    return udf::ScalarUDFDocBuilder("Boolean ORs the passed in values.")
        .Example(R"doc(# Implicit call.
        | df.can_filter_or_has_data = df.can_filter or df.has_data
        | # Explicit call.
        | df.can_filter_or_has_data = px.logicalOr(df.can_filter, df.has_data)
        )doc")
        .Arg("b1", "Left side of the OR.")
        .Arg("b2", "Right side of the OR.")
        .Returns(
            "True if either expression is Truthy or both expressions are Truthy, otherwise "
            "False.");
  }
};

template <typename TArg1, typename TArg2>
class LogicalAndUDF : public udf::ScalarUDF {
 public:
  BoolValue Exec(FunctionContext*, TArg1 b1, TArg2 b2) { return b1.val && b2.val; }
  static udf::ScalarUDFDocBuilder Doc() {
    return udf::ScalarUDFDocBuilder("Boolean ANDs the passed in values.")
        .Example(R"doc(# Implicit call.
        | df.can_filter_and_has_data = df.can_filter and df.has_data
        | # Explicit call.
        | df.can_filter_and_has_data = px.logicalAnd(df.can_filter, df.has_data)
        )doc")
        .Arg("b1", "Left side of the AND.")
        .Arg("b2", "Right side of the AND.")
        .Returns("True if both expressions are Truthy otherwise False.");
  }
};

template <typename TArg1>
class LogicalNotUDF : public udf::ScalarUDF {
 public:
  BoolValue Exec(FunctionContext*, TArg1 b1) { return !b1.val; }
  static udf::ScalarUDFDocBuilder Doc() {
    return udf::ScalarUDFDocBuilder("Boolean NOTs the passed in value.")
        .Example(R"doc(# Implicit call.
        | df.not_can_filter = not df.can_filter
        | # Explicit call.
        | df.not_can_filter = px.logicalNot(df.can_filter)
        )doc")
        .Arg("b1", "The value to Invert.")
        .Returns("True if input is Falsey otherwise False.");
  }
};

template <typename TArg1>
class NegateUDF : public udf::ScalarUDF {
 public:
  TArg1 Exec(FunctionContext*, TArg1 b1) { return -b1.val; }
  static udf::ScalarUDFDocBuilder Doc() {
    return udf::ScalarUDFDocBuilder("Negates the passed in value.")
        .Example(R"doc(# Implicit call.
        | df.negative_latency_ms = -df.latency_ms
        | # Explicit call.
        | df.negative_latency_ms = px.negate(df.latency_ms))doc")
        .Arg("b1", "The value to negate.")
        .Returns("`b1` with a flipped negative sign.");
  }
};

template <typename TArg1>
class InvertUDF : public udf::ScalarUDF {
 public:
  TArg1 Exec(FunctionContext*, TArg1 b1) { return ~b1.val; }

  static udf::ScalarUDFDocBuilder Doc() {
    return udf::ScalarUDFDocBuilder("Invert the bits of the given value.")
        .Example("df.inverted = px.invert(df.a)")
        .Arg("arg1", "The value to invert.")
        .Returns("The inverted form of arg1.");
  }
};

class CeilUDF : public udf::ScalarUDF {
 public:
  Int64Value Exec(FunctionContext*, Float64Value b1) {
    return static_cast<int64_t>(std::ceil(b1.val));
  }

  static udf::ScalarUDFDocBuilder Doc() {
    return udf::ScalarUDFDocBuilder("Compute the ceiling of the value.")
        .Example("df.b = px.ceil(df.a)")
        .Arg("arg1", "The value to take the ceiling of.")
        .Returns("The ceiling of arg1.");
  }
};

class FloorUDF : public udf::ScalarUDF {
 public:
  Int64Value Exec(FunctionContext*, Float64Value b1) {
    return static_cast<int64_t>(std::floor(b1.val));
  }

  static udf::ScalarUDFDocBuilder Doc() {
    return udf::ScalarUDFDocBuilder("Compute the floor of the value.")
        .Example("df.b = px.floor(df.a)")
        .Arg("arg1", "The value to take the floor of.")
        .Returns("The floor of arg1.");
  }
};

template <typename TArg1, typename TArg2>
class EqualUDF : public udf::ScalarUDF {
 public:
  BoolValue Exec(FunctionContext*, TArg1 b1, TArg2 b2) { return b1 == b2; }
  static udf::ScalarUDFDocBuilder Doc() {
    return udf::ScalarUDFDocBuilder("Returns whether the values are equal.")
        .Details(
            "Determines whether the values are equal. Passing in floating-point values might "
            "lead "
            "to false negatives. Use `px.approxEqual(b1, b2)` to compare floats instead.")
        .Example(R"doc(# Implicit call.
        | df.success_http = df.http_status == 200
        | # Explicit call.
        | df.success_http = px.equal(df.http_status, 200))doc")
        .Arg("b1", "")
        .Arg("b2", "")
        .Returns("True if `b1` is equal to `b2`, False otherwise.");
  }
};

template <typename TArg1, typename TArg2>
class NotEqualUDF : public udf::ScalarUDF {
 public:
  BoolValue Exec(FunctionContext*, TArg1 b1, TArg2 b2) { return b1 != b2; }
  static udf::ScalarUDFDocBuilder Doc() {
    return udf::ScalarUDFDocBuilder("Returns whether the values are not equal.")
        .Details(
            "Determines whether the values are not equal. Passing in floating-point values might "
            "lead "
            "to false positives. Use `not px.approxEqual(b1, b2)` to compare floats instead.")
        .Example(R"doc(# Implicit call.
        | df.failed_http = df.http_status != 200
        | # Explicit call.
        | df.failed_http = px.notEqual(df.http_status, 200))doc")
        .Arg("b1", "")
        .Arg("b2", "")
        .Returns("True if `b1` is not equal to `b2`, False otherwise.");
  }
};

template <typename TArg1, typename TArg2>
class ApproxEqualUDF : public udf::ScalarUDF {
 public:
  BoolValue Exec(FunctionContext*, TArg1 b1, TArg2 b2) {
    return std::abs(b1.val - b2.val) < std::numeric_limits<double>::epsilon();
  }

  static udf::ScalarUDFDocBuilder Doc() {
    return udf::ScalarUDFDocBuilder("Return whether the two values are approximately equal.")
        .Details(
            "Returns whether or not the two given values are approximately equal to each other, "
            "where the values have an absolute difference of less than 1E-9.")
        .Example("df.equals = px.approxEqual(df.a, df.b)")
        .Arg("arg1", "The value to be compared to.")
        .Arg("arg2", "The value which should be compared to the first argument.")
        .Returns("Boolean of whether the two arguments are approximately equal.");
  }
};

template <typename TArg1, typename TArg2>
class ApproxNotEqualUDF : public udf::ScalarUDF {
 public:
  BoolValue Exec(FunctionContext*, TArg1 b1, TArg2 b2) {
    return std::abs(b1.val - b2.val) > std::numeric_limits<double>::epsilon();
  }
};

template <typename TArg1, typename TArg2>
class GreaterThanUDF : public udf::ScalarUDF {
 public:
  BoolValue Exec(FunctionContext*, TArg1 b1, TArg2 b2) { return b1 > b2; }

  static udf::ScalarUDFDocBuilder Doc() {
    return udf::ScalarUDFDocBuilder(
               "Compare whether the first argument is greater than the second argument.")
        .Details("This function is implicitly invoked by the > operator.")
        .Example(R"doc(# Implict call.
        | df.gt = df.a > df.b
        | Explicit call.
        | df.gt = px.greaterThan(df.a, df.b))doc")
        .Arg("arg1", "The value to be compared to.")
        .Arg("arg2", "The value to check if it is greater than the first argument.")
        .Returns("Boolean of whether arg1 is greater than arg2.");
  }
};

template <typename TArg1, typename TArg2>
class GreaterThanEqualUDF : public udf::ScalarUDF {
 public:
  BoolValue Exec(FunctionContext*, TArg1 b1, TArg2 b2) { return b1 >= b2; }

  static udf::ScalarUDFDocBuilder Doc() {
    return udf::ScalarUDFDocBuilder(
               "Compare whether the first argument is greater than or equal to the second "
               "argument.")
        .Details("This function is implicitly invoked by the >= operator.")
        .Example(R"doc(# Implict call.
        | df.gte = df.a >= df.b
        | Explicit call.
        | df.gte = px.greaterThanEqual(df.a, df.b))doc")
        .Arg("arg1", "The value to be compared to.")
        .Arg("arg2", "The value to check if it is greater than or equal to the first argument.")
        .Returns("Boolean of whether arg1 is greater than or equal to arg2.");
  }
};

template <typename TArg1, typename TArg2>
class LessThanUDF : public udf::ScalarUDF {
 public:
  BoolValue Exec(FunctionContext*, TArg1 b1, TArg2 b2) { return b1 < b2; }
  static udf::ScalarUDFDocBuilder Doc() {
    return udf::ScalarUDFDocBuilder("Returns which value is less than the other.")
        .Example(R"doc(# Implict call.
        | df.lt = df.a < df.b
        | Explicit call.
        | df.lt = px.lessThan(df.a, df.b))doc")
        .Arg("b1", "Left side of the expression.")
        .Arg("b2", "Right side of the expression.")
        .Returns("Whether the left side is less than the right.");
  }
};

template <typename TArg1, typename TArg2>
class LessThanEqualUDF : public udf::ScalarUDF {
 public:
  BoolValue Exec(FunctionContext*, TArg1 b1, TArg2 b2) { return b1 <= b2; }
  static udf::ScalarUDFDocBuilder Doc() {
    return udf::ScalarUDFDocBuilder("Returns which value is less than or equal to the the other.")
        .Example(R"doc(
        | # Implicit call.
        | df.lte = df.cpu1 <= df.cpu2
        | # Explicit call.
        | df.lte = px.lessThanOrEqual(df.cpu1, df.cup2)
        )doc")
        .Arg("b1", "Left side of the expression.")
        .Arg("b2", "Right side of the expression.")
        .Returns("Whether the left side is less than or equal to the right.");
  }
};

udf::ScalarUDFDocBuilder BinDoc();

template <typename TReturn, typename TArg1, typename TArg2>
class BinUDF : public udf::ScalarUDF {
 public:
  TReturn Exec(FunctionContext*, TArg1 b1, TArg2 b2) { return b1.val - (b1.val % b2.val); }
  static udf::ScalarUDFDocBuilder Doc() { return BinDoc(); }
};

// Special instantitization to handle float to int conversion
template <typename TReturn, typename TArg2>
class BinUDF<TReturn, types::Float64Value, TArg2> : public udf::ScalarUDF {
 public:
  TReturn Exec(FunctionContext*, Float64Value b1, Int64Value b2) {
    return static_cast<int64_t>(b1.val) - (static_cast<int64_t>(b1.val) % b2.val);
  }
  static udf::ScalarUDFDocBuilder Doc() { return BinDoc(); }
};

// TODO(philkuz, nserrino) Move decimal places to be a constructor arg after PL-1048 is done.
class RoundUDF : public udf::ScalarUDF {
 public:
  StringValue Exec(FunctionContext*, Float64Value value, Int64Value decimal_places) {
    return absl::StrFormat("%.*f", decimal_places.val, value.val);
  }
  static udf::ScalarUDFDocBuilder Doc() {
    return udf::ScalarUDFDocBuilder("Rounds the float to the nearest decimal places.")
        .Details(
            "Rounds the float to the nearest decimal place and returns value as a string. Used "
            "to "
            "clean up the data shown in tables. Set decimals to `0` if you want to round to the "
            "nearest int value.")
        .Example(R"doc(
        | df.cpu1 = 0.248
        | df.cpu1 = px.round(df.cpu1, 2) # 0.25
        )doc")
        .Arg("value", "The value to round.")
        .Arg("decimals", "Number of decimal places to round to. `0` => round to nearest int.")
        .Returns("Float rounded to the specified decimal place as a string.");
  }
};

template <typename TArg>
class MeanUDA : public udf::UDA {
 public:
  void Update(FunctionContext*, TArg arg) {
    info_.size++;
    info_.count += arg.val;
  }
  void Merge(FunctionContext*, const MeanUDA& other) {
    info_.size += other.info_.size;
    info_.count += other.info_.count;
  }
  Float64Value Finalize(FunctionContext*) { return info_.count / info_.size; }

  static udf::InfRuleVec SemanticInferenceRules() {
    return {udf::InheritTypeFromArgs<MeanUDA>::Create({types::ST_BYTES, types::ST_THROUGHPUT_PER_NS,
                                                       types::ST_THROUGHPUT_BYTES_PER_NS,
                                                       types::ST_DURATION_NS, types::ST_PERCENT})};
  }

  StringValue Serialize(FunctionContext*) {
    return StringValue(reinterpret_cast<char*>(&info_), sizeof(info_));
  }

  Status Deserialize(FunctionContext*, const StringValue& data) {
    info_ = *reinterpret_cast<const MeanInfo*>(data.data());
    return Status::OK();
  }
  static udf::UDADocBuilder Doc() {
    return udf::UDADocBuilder("Calculate the arithmetic mean.")
        .Details(
            "Calculates the arithmetic mean by summing the values then dividing by the number of "
            "values.")
        .Example("df = df.agg(mean=('latency_ms', px.mean))")
        .Arg("arg", "The data to average.")
        .Returns("The mean of the data.");
  }

 protected:
  struct MeanInfo {
    uint64_t size = 0;
    double count = 0;
  };

  MeanInfo info_;
};

template <typename TArg>
class SumUDA : public udf::UDA {
 public:
  void Update(FunctionContext*, TArg arg) { sum_ = sum_.val + arg.val; }
  void Merge(FunctionContext*, const SumUDA& other) { sum_ = sum_.val + other.sum_.val; }
  TArg Finalize(FunctionContext*) { return sum_; }
  static udf::InfRuleVec SemanticInferenceRules() {
    return {udf::InheritTypeFromArgs<SumUDA>::Create(
        {types::ST_BYTES, types::ST_THROUGHPUT_PER_NS, types::ST_THROUGHPUT_BYTES_PER_NS})};
  }
  StringValue Serialize(FunctionContext*) {
    return StringValue(reinterpret_cast<char*>(&sum_), sizeof(sum_));
  }

  Status Deserialize(FunctionContext*, const StringValue& data) {
    sum_ =
        *reinterpret_cast<const typename types::ValueTypeTraits<TArg>::native_type*>(data.data());
    return Status::OK();
  }

  static udf::UDADocBuilder Doc() {
    return udf::UDADocBuilder("Calculate the arithmetic sum of the grouped values.")
        .Example("df = df.agg(sum=('latency_ms', px.sum))")
        .Arg("arg", "The group to sum.")
        .Returns("The sum of the data.");
  }

 protected:
  TArg sum_ = 0;
};

template <typename TArg>
class MaxUDA : public udf::UDA {
 public:
  void Update(FunctionContext*, TArg arg) {
    if (max_.val < arg.val) {
      max_ = arg;
    }
  }
  void Merge(FunctionContext*, const MaxUDA& other) {
    if (other.max_.val > max_.val) {
      max_ = other.max_;
    }
  }
  TArg Finalize(FunctionContext*) { return max_; }

  static udf::InfRuleVec SemanticInferenceRules() {
    return {udf::InheritTypeFromArgs<MaxUDA>::Create({types::ST_BYTES, types::ST_THROUGHPUT_PER_NS,
                                                      types::ST_THROUGHPUT_BYTES_PER_NS,
                                                      types::ST_DURATION_NS, types::ST_PERCENT})};
  }

  static udf::UDADocBuilder Doc() {
    return udf::UDADocBuilder("Returns the maximum in the group.")
        .Example("df = df.agg(max_latency=('latency_ms', px.max))")
        .Arg("arg", "The data on which to apply the function.")
        .Returns("The maximum value in the group.");
  }

  StringValue Serialize(FunctionContext*) {
    return StringValue(reinterpret_cast<char*>(&max_), sizeof(max_));
  }

  Status Deserialize(FunctionContext*, const StringValue& data) {
    max_ =
        *reinterpret_cast<const typename types::ValueTypeTraits<TArg>::native_type*>(data.data());
    return Status::OK();
  }

 protected:
  TArg max_ = std::numeric_limits<typename types::ValueTypeTraits<TArg>::native_type>::min();
};

template <typename TArg>
class MinUDA : public udf::UDA {
 public:
  void Update(FunctionContext*, TArg arg) {
    if (min_.val > arg.val) {
      min_ = arg;
    }
  }
  void Merge(FunctionContext*, const MinUDA& other) {
    if (other.min_.val < min_.val) {
      min_ = other.min_;
    }
  }
  TArg Finalize(FunctionContext*) { return min_; }

  static udf::InfRuleVec SemanticInferenceRules() {
    return {udf::InheritTypeFromArgs<MinUDA>::Create({types::ST_BYTES, types::ST_THROUGHPUT_PER_NS,
                                                      types::ST_THROUGHPUT_BYTES_PER_NS,
                                                      types::ST_DURATION_NS, types::ST_PERCENT})};
  }

  StringValue Serialize(FunctionContext*) {
    return StringValue(reinterpret_cast<char*>(&min_), sizeof(min_));
  }

  Status Deserialize(FunctionContext*, const StringValue& data) {
    min_ =
        *reinterpret_cast<const typename types::ValueTypeTraits<TArg>::native_type*>(data.data());
    return Status::OK();
  }
  static udf::UDADocBuilder Doc() {
    return udf::UDADocBuilder("Returns the minimum in the group.")
        .Example("df = df.agg(min_latency=('latency_ms', px.min))")
        .Arg("arg", "The data on which to apply the function.")
        .Returns("The minimum.");
  }

 protected:
  TArg min_ = std::numeric_limits<typename types::ValueTypeTraits<TArg>::native_type>::max();
};

template <typename TArg>
class CountUDA : public udf::UDA {
 public:
  void Update(FunctionContext*, TArg) { count_++; }
  void Merge(FunctionContext*, const CountUDA& other) { count_ += other.count_; }
  Int64Value Finalize(FunctionContext*) { return count_; }

  StringValue Serialize(FunctionContext*) {
    return StringValue(reinterpret_cast<char*>(&count_), sizeof(count_));
  }

  Status Deserialize(FunctionContext*, const StringValue& data) {
    count_ = *reinterpret_cast<const uint64_t*>(data.data());
    return Status::OK();
  }

  static udf::UDADocBuilder Doc() {
    return udf::UDADocBuilder("Returns number of rows in the aggregate group.")
        .Details(
            "This function counts the number of rows in the aggregate group. The count of "
            "rows is independent of which column this function is applied to. Each column should "
            "return the same number of rows.")
        .Example("df = df.agg(throughput_total=('latency_ms', px.count))")
        .Arg("arg", "The data on which to apply the function.")
        .Returns("The count.");
  }

 protected:
  uint64_t count_ = 0;
};

void RegisterMathOpsOrDie(udf::Registry* registry);
}  // namespace builtins
}  // namespace carnot
}  // namespace pl
