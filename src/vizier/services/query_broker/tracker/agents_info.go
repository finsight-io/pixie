package tracker

import (
	"fmt"

	uuid "github.com/satori/go.uuid"

	"pixielabs.ai/pixielabs/src/carnot/planner/distributedpb"
	"pixielabs.ai/pixielabs/src/utils"
	"pixielabs.ai/pixielabs/src/vizier/services/metadata/metadatapb"
)

// AgentsInfo tracks information about the distributed state of the system.
type AgentsInfo interface {
	ClearState()
	UpdateAgentsInfo(agentUpdates []*metadatapb.AgentUpdate, schemaInfos []*distributedpb.SchemaInfo) error
	DistributedState() *distributedpb.DistributedState
}

// AgentsInfoImpl implements AgentsInfo to track information about the distributed state of the system.
type AgentsInfoImpl struct {
	ds *distributedpb.DistributedState
}

// NewAgentsInfo creates an empty agents info.
func NewAgentsInfo() AgentsInfo {
	return &AgentsInfoImpl{
		&distributedpb.DistributedState{
			SchemaInfo: []*distributedpb.SchemaInfo{},
			CarnotInfo: []*distributedpb.CarnotInfo{},
		},
	}
}

// NewTestAgentsInfo creates an agents info from a passed in distributed state.
func NewTestAgentsInfo(ds *distributedpb.DistributedState) AgentsInfo {
	return &AgentsInfoImpl{
		ds,
	}
}

// ClearState clears the agents info state.
func (a *AgentsInfoImpl) ClearState() {
	a.ds = &distributedpb.DistributedState{
		SchemaInfo: []*distributedpb.SchemaInfo{},
		CarnotInfo: []*distributedpb.CarnotInfo{},
	}
}

// UpdateAgentsInfo creates a new agent info.
func (a *AgentsInfoImpl) UpdateAgentsInfo(agentUpdates []*metadatapb.AgentUpdate, schemaInfos []*distributedpb.SchemaInfo) error {
	if len(schemaInfos) > 0 {
		a.ds.SchemaInfo = schemaInfos
	}
	carnotInfoMap := make(map[uuid.UUID]*distributedpb.CarnotInfo)
	for _, carnotInfo := range a.ds.CarnotInfo {
		agentUUID, err := utils.UUIDFromProto(carnotInfo.AgentID)
		if err != nil {
			return err
		}
		carnotInfoMap[agentUUID] = carnotInfo
	}

	for _, agentUpdate := range agentUpdates {
		agentUUID, err := utils.UUIDFromProto(agentUpdate.AgentID)
		if err != nil {
			return err
		}
		// case 1: agent info update
		agent := agentUpdate.GetAgent()
		if agent != nil {
			if agent.Info.Capabilities == nil || agent.Info.Capabilities.CollectsData {
				var metadataInfo *distributedpb.MetadataInfo
				if carnotInfo, present := carnotInfoMap[agentUUID]; present {
					metadataInfo = carnotInfo.MetadataInfo
				}
				// this is a PEM
				carnotInfoMap[agentUUID] = makeAgentCarnotInfo(agentUUID, agent.ASID, metadataInfo)
			} else {
				// this is a Kelvin
				kelvinGRPCAddress := agent.Info.IPAddress
				carnotInfoMap[agentUUID] = makeKelvinCarnotInfo(agentUUID, kelvinGRPCAddress, agent.ASID)
			}
		}
		// case 2: agent data info update
		dataInfo := agentUpdate.GetDataInfo()
		if dataInfo != nil {
			carnotInfo, present := carnotInfoMap[agentUUID]
			if !present || carnotInfo == nil {
				return fmt.Errorf("Could not update agent table metadata of unknown agent %+v", agentUUID)
			}
			if dataInfo.MetadataInfo != nil {
				carnotInfo.MetadataInfo = dataInfo.MetadataInfo
			}
		}
		// case 3: agent deleted
		if agentUpdate.GetDeleted() {
			delete(carnotInfoMap, agentUUID)
		}
	}

	// reset the array and recreate.
	a.ds.CarnotInfo = []*distributedpb.CarnotInfo{}
	for _, carnotInfo := range carnotInfoMap {
		a.ds.CarnotInfo = append(a.ds.CarnotInfo, carnotInfo)
	}
	return nil
}

// DistributedState returns the current distributed state. Will return nil if not existent.
func (a *AgentsInfoImpl) DistributedState() *distributedpb.DistributedState {
	return a.ds
}

func makeAgentCarnotInfo(agentID uuid.UUID, asid uint32, agentMetadata *distributedpb.MetadataInfo) *distributedpb.CarnotInfo {
	return &distributedpb.CarnotInfo{
		QueryBrokerAddress:   agentID.String(),
		AgentID:              utils.ProtoFromUUID(&agentID),
		ASID:                 asid,
		HasGRPCServer:        false,
		HasDataStore:         true,
		ProcessesData:        true,
		AcceptsRemoteSources: false,
		MetadataInfo:         agentMetadata,
	}
}

func makeKelvinCarnotInfo(agentID uuid.UUID, grpcAddress string, asid uint32) *distributedpb.CarnotInfo {
	return &distributedpb.CarnotInfo{
		QueryBrokerAddress:   agentID.String(),
		AgentID:              utils.ProtoFromUUID(&agentID),
		ASID:                 asid,
		HasGRPCServer:        true,
		GRPCAddress:          grpcAddress,
		HasDataStore:         false,
		ProcessesData:        true,
		AcceptsRemoteSources: true,
		// When we support persistent storage, Kelvins will also have MetadataInfo.
		MetadataInfo: nil,
	}
}
