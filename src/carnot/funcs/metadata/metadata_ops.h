#pragma once

#include <string>
#include <utility>
#include <vector>

#include "src/carnot/udf/registry.h"
#include "src/shared/metadata/metadata_state.h"
#include "src/shared/types/types.h"

namespace pl {
namespace carnot {
namespace funcs {
namespace metadata {

using ScalarUDF = pl::carnot::udf::ScalarUDF;
using FunctionContext = pl::carnot::udf::FunctionContext;

inline const pl::md::AgentMetadataState* GetMetadataState(FunctionContext* ctx) {
  DCHECK(ctx != nullptr);
  auto md = ctx->metadata_state();
  DCHECK(md != nullptr);
  return md;
}

class ASIDUDF : public ScalarUDF {
 public:
  types::Int64Value Exec(FunctionContext* ctx) {
    auto md = GetMetadataState(ctx);
    return md->asid();
  }
};

class PodIDToPodNameUDF : public ScalarUDF {
 public:
  types::StringValue Exec(FunctionContext* ctx, types::StringValue pod_id) {
    auto md = GetMetadataState(ctx);

    const auto* pod_info = md->k8s_metadata_state().PodInfoByID(pod_id);
    if (pod_info != nullptr) {
      return absl::Substitute("$0/$1", pod_info->ns(), pod_info->name());
    }

    return "";
  }
};

class PodNameToPodIDUDF : public ScalarUDF {
 public:
  types::StringValue Exec(FunctionContext* ctx, types::StringValue pod_name) {
    auto md = GetMetadataState(ctx);
    // This UDF expects the pod name to be in the format of "<ns>/<pod-name>".
    std::vector<std::string_view> name_parts = absl::StrSplit(pod_name, "/");
    if (name_parts.size() != 2) {
      return "";
    }

    auto pod_name_view = std::make_pair(name_parts[0], name_parts[1]);
    auto pod_id = md->k8s_metadata_state().PodIDByName(pod_name_view);

    return pod_id;
  }
};

class UPIDToContainerIDUDF : public ScalarUDF {
 public:
  types::StringValue Exec(FunctionContext* ctx, types::UInt128Value upid_value) {
    auto md = GetMetadataState(ctx);

    auto upid_uint128 = absl::MakeUint128(upid_value.High64(), upid_value.Low64());
    auto upid = md::UPID(upid_uint128);
    auto pid = md->GetPIDByUPID(upid);
    if (pid == nullptr) {
      return "";
    }
    return pid->cid();
  }
};

class UPIDToPodIDUDF : public ScalarUDF {
 public:
  types::StringValue Exec(FunctionContext* ctx, types::UInt128Value upid_value) {
    auto md = GetMetadataState(ctx);

    auto upid_uint128 = absl::MakeUint128(upid_value.High64(), upid_value.Low64());
    auto upid = md::UPID(upid_uint128);
    auto pid = md->GetPIDByUPID(upid);
    if (pid == nullptr) {
      return "";
    }
    auto container_info = md->k8s_metadata_state().ContainerInfoByID(pid->cid());
    if (container_info == nullptr) {
      return "";
    }
    return std::string(container_info->pod_id());
  }
};

class UPIDToPodNameUDF : public ScalarUDF {
 public:
  types::StringValue Exec(FunctionContext* ctx, types::UInt128Value upid_value) {
    auto md = GetMetadataState(ctx);

    auto upid_uint128 = absl::MakeUint128(upid_value.High64(), upid_value.Low64());
    auto upid = md::UPID(upid_uint128);
    auto pid = md->GetPIDByUPID(upid);
    if (pid == nullptr) {
      return "";
    }
    auto container_info = md->k8s_metadata_state().ContainerInfoByID(pid->cid());
    if (container_info == nullptr) {
      return "";
    }
    const auto* pod_info = md->k8s_metadata_state().PodInfoByID(container_info->pod_id());
    if (pod_info == nullptr) {
      return "";
    }
    return absl::Substitute("$0/$1", pod_info->ns(), pod_info->name());
  }
};

void RegisterMetadataOpsOrDie(pl::carnot::udf::ScalarUDFRegistry* registry);

}  // namespace metadata
}  // namespace funcs
}  // namespace carnot
}  // namespace pl
