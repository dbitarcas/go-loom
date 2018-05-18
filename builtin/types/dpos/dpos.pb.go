// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/dpos/dpos.proto

/*
Package dpos is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/go-loom/builtin/types/dpos/dpos.proto

It has these top-level messages:
	Params
	State
	Voter
	Witness
	Candidate
	CandidateSet
	Vote
	VoteSet
	DPOSInitRequest
	RegisterCandidateRequest
	RegisterCandidateResponse
	UnregisterCandidateRequest
	UnregisterCandidateResponse
	VoteRequest
	VoteResponse
	ProxyVoteRequest
	ProxyVoteResponse
	UnproxyVoteRequest
	UnproxyVoteResponse
	ElectRequest
	ElectResponse
	ListWitnessesRequest
	ListWitnessesResponse
*/
package dpos

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Params struct {
	WitnessCount        uint64         `protobuf:"varint,1,opt,name=witness_count,json=witnessCount,proto3" json:"witness_count,omitempty"`
	VoteAllocation      uint64         `protobuf:"varint,2,opt,name=vote_allocation,json=voteAllocation,proto3" json:"vote_allocation,omitempty"`
	ElectionCycleLength uint64         `protobuf:"varint,3,opt,name=election_cycle_length,json=electionCycleLength,proto3" json:"election_cycle_length,omitempty"`
	MinPowerFraction    uint64         `protobuf:"varint,4,opt,name=min_power_fraction,json=minPowerFraction,proto3" json:"min_power_fraction,omitempty"`
	CoinContractAddress *types.Address `protobuf:"bytes,5,opt,name=coin_contract_address,json=coinContractAddress" json:"coin_contract_address,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (m *Params) String() string            { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{0} }

func (m *Params) GetWitnessCount() uint64 {
	if m != nil {
		return m.WitnessCount
	}
	return 0
}

func (m *Params) GetVoteAllocation() uint64 {
	if m != nil {
		return m.VoteAllocation
	}
	return 0
}

func (m *Params) GetElectionCycleLength() uint64 {
	if m != nil {
		return m.ElectionCycleLength
	}
	return 0
}

func (m *Params) GetMinPowerFraction() uint64 {
	if m != nil {
		return m.MinPowerFraction
	}
	return 0
}

func (m *Params) GetCoinContractAddress() *types.Address {
	if m != nil {
		return m.CoinContractAddress
	}
	return nil
}

type State struct {
	Params       *Params    `protobuf:"bytes,1,opt,name=params" json:"params,omitempty"`
	LastElection uint64     `protobuf:"varint,2,opt,name=last_election,json=lastElection,proto3" json:"last_election,omitempty"`
	Witnesses    []*Witness `protobuf:"bytes,3,rep,name=witnesses" json:"witnesses,omitempty"`
}

func (m *State) Reset()                    { *m = State{} }
func (m *State) String() string            { return proto.CompactTextString(m) }
func (*State) ProtoMessage()               {}
func (*State) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{1} }

func (m *State) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *State) GetLastElection() uint64 {
	if m != nil {
		return m.LastElection
	}
	return 0
}

func (m *State) GetWitnesses() []*Witness {
	if m != nil {
		return m.Witnesses
	}
	return nil
}

type Voter struct {
	Address            *types.Address   `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Balance            uint64           `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	ProxyAddress       *types.Address   `protobuf:"bytes,3,opt,name=proxy_address,json=proxyAddress" json:"proxy_address,omitempty"`
	PrincipalAddresses []*types.Address `protobuf:"bytes,4,rep,name=principal_addresses,json=principalAddresses" json:"principal_addresses,omitempty"`
}

func (m *Voter) Reset()                    { *m = Voter{} }
func (m *Voter) String() string            { return proto.CompactTextString(m) }
func (*Voter) ProtoMessage()               {}
func (*Voter) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{2} }

func (m *Voter) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Voter) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Voter) GetProxyAddress() *types.Address {
	if m != nil {
		return m.ProxyAddress
	}
	return nil
}

func (m *Voter) GetPrincipalAddresses() []*types.Address {
	if m != nil {
		return m.PrincipalAddresses
	}
	return nil
}

type Witness struct {
	PubKey     []byte `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	VoteTotal  uint64 `protobuf:"varint,2,opt,name=vote_total,json=voteTotal,proto3" json:"vote_total,omitempty"`
	PowerTotal uint64 `protobuf:"varint,3,opt,name=power_total,json=powerTotal,proto3" json:"power_total,omitempty"`
}

func (m *Witness) Reset()                    { *m = Witness{} }
func (m *Witness) String() string            { return proto.CompactTextString(m) }
func (*Witness) ProtoMessage()               {}
func (*Witness) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{3} }

func (m *Witness) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *Witness) GetVoteTotal() uint64 {
	if m != nil {
		return m.VoteTotal
	}
	return 0
}

func (m *Witness) GetPowerTotal() uint64 {
	if m != nil {
		return m.PowerTotal
	}
	return 0
}

type Candidate struct {
	Address *types.Address `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	PubKey  []byte         `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
}

func (m *Candidate) Reset()                    { *m = Candidate{} }
func (m *Candidate) String() string            { return proto.CompactTextString(m) }
func (*Candidate) ProtoMessage()               {}
func (*Candidate) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{4} }

func (m *Candidate) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Candidate) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

type CandidateSet struct {
	Candidates map[string]*Candidate `protobuf:"bytes,1,rep,name=candidates" json:"candidates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CandidateSet) Reset()                    { *m = CandidateSet{} }
func (m *CandidateSet) String() string            { return proto.CompactTextString(m) }
func (*CandidateSet) ProtoMessage()               {}
func (*CandidateSet) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{5} }

func (m *CandidateSet) GetCandidates() map[string]*Candidate {
	if m != nil {
		return m.Candidates
	}
	return nil
}

type Vote struct {
	VoterAddress     *types.Address `protobuf:"bytes,1,opt,name=voter_address,json=voterAddress" json:"voter_address,omitempty"`
	CandidateAddress *types.Address `protobuf:"bytes,2,opt,name=candidate_address,json=candidateAddress" json:"candidate_address,omitempty"`
	Amount           uint64         `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *Vote) Reset()                    { *m = Vote{} }
func (m *Vote) String() string            { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()               {}
func (*Vote) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{6} }

func (m *Vote) GetVoterAddress() *types.Address {
	if m != nil {
		return m.VoterAddress
	}
	return nil
}

func (m *Vote) GetCandidateAddress() *types.Address {
	if m != nil {
		return m.CandidateAddress
	}
	return nil
}

func (m *Vote) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type VoteSet struct {
	Votes map[string]*Vote `protobuf:"bytes,1,rep,name=votes" json:"votes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *VoteSet) Reset()                    { *m = VoteSet{} }
func (m *VoteSet) String() string            { return proto.CompactTextString(m) }
func (*VoteSet) ProtoMessage()               {}
func (*VoteSet) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{7} }

func (m *VoteSet) GetVotes() map[string]*Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

type DPOSInitRequest struct {
	Params     *Params            `protobuf:"bytes,1,opt,name=params" json:"params,omitempty"`
	Validators []*types.Validator `protobuf:"bytes,2,rep,name=validators" json:"validators,omitempty"`
}

func (m *DPOSInitRequest) Reset()                    { *m = DPOSInitRequest{} }
func (m *DPOSInitRequest) String() string            { return proto.CompactTextString(m) }
func (*DPOSInitRequest) ProtoMessage()               {}
func (*DPOSInitRequest) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{8} }

func (m *DPOSInitRequest) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *DPOSInitRequest) GetValidators() []*types.Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

type RegisterCandidateRequest struct {
	PubKey []byte `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
}

func (m *RegisterCandidateRequest) Reset()                    { *m = RegisterCandidateRequest{} }
func (m *RegisterCandidateRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterCandidateRequest) ProtoMessage()               {}
func (*RegisterCandidateRequest) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{9} }

func (m *RegisterCandidateRequest) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

type RegisterCandidateResponse struct {
}

func (m *RegisterCandidateResponse) Reset()                    { *m = RegisterCandidateResponse{} }
func (m *RegisterCandidateResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterCandidateResponse) ProtoMessage()               {}
func (*RegisterCandidateResponse) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{10} }

type UnregisterCandidateRequest struct {
}

func (m *UnregisterCandidateRequest) Reset()                    { *m = UnregisterCandidateRequest{} }
func (m *UnregisterCandidateRequest) String() string            { return proto.CompactTextString(m) }
func (*UnregisterCandidateRequest) ProtoMessage()               {}
func (*UnregisterCandidateRequest) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{11} }

type UnregisterCandidateResponse struct {
}

func (m *UnregisterCandidateResponse) Reset()                    { *m = UnregisterCandidateResponse{} }
func (m *UnregisterCandidateResponse) String() string            { return proto.CompactTextString(m) }
func (*UnregisterCandidateResponse) ProtoMessage()               {}
func (*UnregisterCandidateResponse) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{12} }

type VoteRequest struct {
	CandidateAddress *types.Address `protobuf:"bytes,1,opt,name=candidate_address,json=candidateAddress" json:"candidate_address,omitempty"`
	Amount           int64          `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *VoteRequest) Reset()                    { *m = VoteRequest{} }
func (m *VoteRequest) String() string            { return proto.CompactTextString(m) }
func (*VoteRequest) ProtoMessage()               {}
func (*VoteRequest) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{13} }

func (m *VoteRequest) GetCandidateAddress() *types.Address {
	if m != nil {
		return m.CandidateAddress
	}
	return nil
}

func (m *VoteRequest) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type VoteResponse struct {
}

func (m *VoteResponse) Reset()                    { *m = VoteResponse{} }
func (m *VoteResponse) String() string            { return proto.CompactTextString(m) }
func (*VoteResponse) ProtoMessage()               {}
func (*VoteResponse) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{14} }

type ProxyVoteRequest struct {
	ProxyAddress *types.Address `protobuf:"bytes,1,opt,name=proxy_address,json=proxyAddress" json:"proxy_address,omitempty"`
}

func (m *ProxyVoteRequest) Reset()                    { *m = ProxyVoteRequest{} }
func (m *ProxyVoteRequest) String() string            { return proto.CompactTextString(m) }
func (*ProxyVoteRequest) ProtoMessage()               {}
func (*ProxyVoteRequest) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{15} }

func (m *ProxyVoteRequest) GetProxyAddress() *types.Address {
	if m != nil {
		return m.ProxyAddress
	}
	return nil
}

type ProxyVoteResponse struct {
}

func (m *ProxyVoteResponse) Reset()                    { *m = ProxyVoteResponse{} }
func (m *ProxyVoteResponse) String() string            { return proto.CompactTextString(m) }
func (*ProxyVoteResponse) ProtoMessage()               {}
func (*ProxyVoteResponse) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{16} }

type UnproxyVoteRequest struct {
}

func (m *UnproxyVoteRequest) Reset()                    { *m = UnproxyVoteRequest{} }
func (m *UnproxyVoteRequest) String() string            { return proto.CompactTextString(m) }
func (*UnproxyVoteRequest) ProtoMessage()               {}
func (*UnproxyVoteRequest) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{17} }

type UnproxyVoteResponse struct {
}

func (m *UnproxyVoteResponse) Reset()                    { *m = UnproxyVoteResponse{} }
func (m *UnproxyVoteResponse) String() string            { return proto.CompactTextString(m) }
func (*UnproxyVoteResponse) ProtoMessage()               {}
func (*UnproxyVoteResponse) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{18} }

type ElectRequest struct {
}

func (m *ElectRequest) Reset()                    { *m = ElectRequest{} }
func (m *ElectRequest) String() string            { return proto.CompactTextString(m) }
func (*ElectRequest) ProtoMessage()               {}
func (*ElectRequest) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{19} }

type ElectResponse struct {
}

func (m *ElectResponse) Reset()                    { *m = ElectResponse{} }
func (m *ElectResponse) String() string            { return proto.CompactTextString(m) }
func (*ElectResponse) ProtoMessage()               {}
func (*ElectResponse) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{20} }

type ListWitnessesRequest struct {
}

func (m *ListWitnessesRequest) Reset()                    { *m = ListWitnessesRequest{} }
func (m *ListWitnessesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListWitnessesRequest) ProtoMessage()               {}
func (*ListWitnessesRequest) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{21} }

type ListWitnessesResponse struct {
	Witnesses []*Witness `protobuf:"bytes,1,rep,name=witnesses" json:"witnesses,omitempty"`
}

func (m *ListWitnessesResponse) Reset()                    { *m = ListWitnessesResponse{} }
func (m *ListWitnessesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListWitnessesResponse) ProtoMessage()               {}
func (*ListWitnessesResponse) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{22} }

func (m *ListWitnessesResponse) GetWitnesses() []*Witness {
	if m != nil {
		return m.Witnesses
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "Params")
	proto.RegisterType((*State)(nil), "State")
	proto.RegisterType((*Voter)(nil), "Voter")
	proto.RegisterType((*Witness)(nil), "Witness")
	proto.RegisterType((*Candidate)(nil), "Candidate")
	proto.RegisterType((*CandidateSet)(nil), "CandidateSet")
	proto.RegisterType((*Vote)(nil), "Vote")
	proto.RegisterType((*VoteSet)(nil), "VoteSet")
	proto.RegisterType((*DPOSInitRequest)(nil), "DPOSInitRequest")
	proto.RegisterType((*RegisterCandidateRequest)(nil), "RegisterCandidateRequest")
	proto.RegisterType((*RegisterCandidateResponse)(nil), "RegisterCandidateResponse")
	proto.RegisterType((*UnregisterCandidateRequest)(nil), "UnregisterCandidateRequest")
	proto.RegisterType((*UnregisterCandidateResponse)(nil), "UnregisterCandidateResponse")
	proto.RegisterType((*VoteRequest)(nil), "VoteRequest")
	proto.RegisterType((*VoteResponse)(nil), "VoteResponse")
	proto.RegisterType((*ProxyVoteRequest)(nil), "ProxyVoteRequest")
	proto.RegisterType((*ProxyVoteResponse)(nil), "ProxyVoteResponse")
	proto.RegisterType((*UnproxyVoteRequest)(nil), "UnproxyVoteRequest")
	proto.RegisterType((*UnproxyVoteResponse)(nil), "UnproxyVoteResponse")
	proto.RegisterType((*ElectRequest)(nil), "ElectRequest")
	proto.RegisterType((*ElectResponse)(nil), "ElectResponse")
	proto.RegisterType((*ListWitnessesRequest)(nil), "ListWitnessesRequest")
	proto.RegisterType((*ListWitnessesResponse)(nil), "ListWitnessesResponse")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/dpos/dpos.proto", fileDescriptorDpos)
}

var fileDescriptorDpos = []byte{
	// 790 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x5d, 0x4f, 0xe3, 0x46,
	0x14, 0x95, 0xf3, 0xd9, 0xdc, 0x18, 0x12, 0x26, 0x84, 0xba, 0xa1, 0x88, 0xc8, 0x95, 0x5a, 0x5a,
	0x95, 0xa4, 0x0a, 0xaa, 0xd4, 0x56, 0x54, 0x28, 0x4a, 0xa9, 0x8a, 0x8a, 0xd4, 0xc8, 0x14, 0x78,
	0xa9, 0x6a, 0x4d, 0x9c, 0xd9, 0x60, 0xe1, 0xcc, 0x78, 0xed, 0x31, 0x6c, 0xde, 0xf7, 0x47, 0xec,
	0xbf, 0xd8, 0x7f, 0xb7, 0xcf, 0xab, 0x19, 0xcf, 0x38, 0x26, 0x9b, 0x2c, 0xbb, 0x2f, 0xc1, 0xf7,
	0x9c, 0x73, 0xef, 0xdc, 0x7b, 0xae, 0x3d, 0xc0, 0xe9, 0xcc, 0xe7, 0x77, 0xc9, 0xa4, 0xe7, 0xb1,
	0x79, 0x3f, 0x60, 0x6c, 0x4e, 0x09, 0x7f, 0x64, 0xd1, 0x7d, 0x7f, 0xc6, 0x8e, 0x45, 0xd8, 0x9f,
	0x24, 0x7e, 0xc0, 0x7d, 0xda, 0xe7, 0x8b, 0x90, 0xc4, 0xfd, 0x69, 0xc8, 0xd2, 0x9f, 0x5e, 0x18,
	0x31, 0xce, 0x3a, 0x3f, 0x3d, 0x93, 0x9d, 0x66, 0xc9, 0xdf, 0x34, 0xc3, 0x7e, 0x67, 0x40, 0x65,
	0x8c, 0x23, 0x3c, 0x8f, 0xd1, 0x37, 0xb0, 0xf5, 0xe8, 0x73, 0x4a, 0xe2, 0xd8, 0xf5, 0x58, 0x42,
	0xb9, 0x65, 0x74, 0x8d, 0xa3, 0x92, 0x63, 0x2a, 0x70, 0x24, 0x30, 0xf4, 0x1d, 0x34, 0x1e, 0x18,
	0x27, 0x2e, 0x0e, 0x02, 0xe6, 0x61, 0xee, 0x33, 0x6a, 0x15, 0xa4, 0x6c, 0x5b, 0xc0, 0xc3, 0x0c,
	0x45, 0x03, 0x68, 0x93, 0x80, 0x78, 0xe2, 0xd9, 0xf5, 0x16, 0x5e, 0x40, 0xdc, 0x80, 0xd0, 0x19,
	0xbf, 0xb3, 0x8a, 0x52, 0xde, 0xd2, 0xe4, 0x48, 0x70, 0x97, 0x92, 0x42, 0x3f, 0x02, 0x9a, 0xfb,
	0xd4, 0x0d, 0xd9, 0x23, 0x89, 0xdc, 0x17, 0x11, 0x96, 0x02, 0xab, 0x24, 0x13, 0x9a, 0x73, 0x9f,
	0x8e, 0x05, 0xf1, 0xa7, 0xc2, 0xd1, 0x29, 0xb4, 0x3d, 0xe6, 0x53, 0xd7, 0x63, 0x94, 0x0b, 0xcc,
	0xc5, 0xd3, 0x69, 0x44, 0xe2, 0xd8, 0x2a, 0x77, 0x8d, 0xa3, 0xfa, 0xe0, 0x8b, 0xde, 0x30, 0x8d,
	0x9d, 0x96, 0x90, 0x8d, 0x94, 0x4a, 0x81, 0x76, 0x02, 0xe5, 0x2b, 0x8e, 0x39, 0x41, 0x87, 0x50,
	0x09, 0xa5, 0x01, 0x72, 0xde, 0xfa, 0xa0, 0xda, 0x4b, 0xfd, 0x70, 0x14, 0x2c, 0x7c, 0x09, 0x70,
	0xcc, 0x5d, 0xdd, 0xb1, 0x1a, 0xd8, 0x14, 0xe0, 0xb9, 0xc2, 0xd0, 0xb7, 0x50, 0x53, 0x3e, 0x91,
	0xd8, 0x2a, 0x76, 0x8b, 0xb2, 0x81, 0xdb, 0x14, 0x71, 0x96, 0x94, 0xfd, 0xd6, 0x80, 0xf2, 0x0d,
	0xe3, 0x24, 0x42, 0x36, 0x54, 0x75, 0xc3, 0xc6, 0x4a, 0xc3, 0x9a, 0x40, 0x16, 0x54, 0x27, 0x38,
	0xc0, 0xd4, 0x23, 0xea, 0x50, 0x1d, 0xa2, 0x63, 0xd8, 0x0a, 0x23, 0xf6, 0x6a, 0x91, 0x0d, 0x5d,
	0x5c, 0xa9, 0x61, 0x4a, 0x5a, 0x45, 0xe8, 0x57, 0x68, 0x85, 0x91, 0x4f, 0x3d, 0x3f, 0xc4, 0x81,
	0x4e, 0x21, 0xb1, 0x55, 0x52, 0x8d, 0xea, 0x24, 0x94, 0x89, 0x86, 0x5a, 0x63, 0x4f, 0xa0, 0xaa,
	0xe6, 0x40, 0x5f, 0x42, 0x35, 0x4c, 0x26, 0xee, 0x3d, 0x59, 0xc8, 0x96, 0x4d, 0xa7, 0x12, 0x26,
	0x93, 0xbf, 0xc9, 0x02, 0x1d, 0x00, 0xc8, 0xb7, 0x82, 0x33, 0x8e, 0x03, 0xd5, 0x6a, 0x4d, 0x20,
	0xff, 0x0a, 0x00, 0x1d, 0x42, 0x3d, 0xdd, 0x69, 0xca, 0xa7, 0x6f, 0x00, 0x48, 0x48, 0x0a, 0xec,
	0xbf, 0xa0, 0x36, 0xc2, 0x74, 0xea, 0x4f, 0xc5, 0x42, 0x3e, 0xc5, 0x98, 0x5c, 0x27, 0x85, 0x7c,
	0x27, 0xf6, 0x1b, 0x03, 0xcc, 0xac, 0xd4, 0x15, 0xe1, 0xe8, 0x77, 0x00, 0x4f, 0xc7, 0xa2, 0xa0,
	0x18, 0xf8, 0xa0, 0x97, 0x97, 0x2c, 0x83, 0xf8, 0x9c, 0xf2, 0x68, 0xe1, 0xe4, 0x12, 0x3a, 0x17,
	0xd0, 0x58, 0xa1, 0x51, 0x13, 0x8a, 0xda, 0x81, 0x9a, 0x23, 0x1e, 0x51, 0x17, 0xca, 0x0f, 0x38,
	0x48, 0xd2, 0x25, 0xd5, 0x07, 0xb0, 0xac, 0xe8, 0xa4, 0xc4, 0x6f, 0x85, 0x5f, 0x0c, 0xfb, 0xb5,
	0x01, 0x25, 0xb1, 0x7a, 0xb1, 0x3b, 0xe1, 0x4d, 0xe4, 0x6e, 0x1a, 0xd3, 0x94, 0xb4, 0xde, 0xdd,
	0xcf, 0xb0, 0x93, 0x35, 0x94, 0xa5, 0x14, 0x56, 0x52, 0x9a, 0x99, 0x44, 0xa7, 0xed, 0x41, 0x05,
	0xcf, 0xe5, 0x77, 0x9c, 0xfa, 0xad, 0x22, 0x3b, 0x81, 0xaa, 0xe8, 0x42, 0x78, 0xf3, 0x3d, 0x94,
	0xc5, 0x49, 0xda, 0x96, 0x56, 0x4f, 0x11, 0xf2, 0xaf, 0x32, 0x23, 0x55, 0x74, 0xce, 0x00, 0x96,
	0xe0, 0x1a, 0x0b, 0xf6, 0x9f, 0x5a, 0x50, 0x96, 0x25, 0xf2, 0xd3, 0xff, 0x0f, 0x8d, 0x3f, 0xc6,
	0xff, 0x5c, 0x5d, 0x50, 0x9f, 0x3b, 0xe4, 0x65, 0x42, 0x62, 0xfe, 0xfc, 0x97, 0xf7, 0x03, 0xc0,
	0x03, 0x0e, 0xc4, 0x54, 0x2c, 0x12, 0x23, 0x17, 0xa5, 0xb9, 0x37, 0x1a, 0x72, 0x72, 0xac, 0x7d,
	0x02, 0x96, 0x43, 0x66, 0x7e, 0xcc, 0x49, 0xb4, 0x74, 0x5f, 0x1d, 0xb4, 0xe9, 0xbd, 0xb5, 0xf7,
	0xe1, 0xab, 0x35, 0x49, 0x71, 0xc8, 0x68, 0x4c, 0xec, 0xaf, 0xa1, 0x73, 0x4d, 0xa3, 0x0d, 0x35,
	0xed, 0x03, 0xd8, 0x5f, 0xcb, 0xaa, 0xe4, 0xff, 0xa0, 0x2e, 0x1d, 0x50, 0x1d, 0xac, 0xdd, 0xa1,
	0xf1, 0x19, 0x3b, 0x14, 0xb6, 0x16, 0xb3, 0x1d, 0x6e, 0x83, 0x99, 0x56, 0x57, 0xa7, 0x0d, 0xa1,
	0x39, 0x16, 0x9f, 0x7b, 0xfe, 0xc8, 0x0f, 0x6e, 0x08, 0xe3, 0x63, 0x37, 0x84, 0xdd, 0x82, 0x9d,
	0x5c, 0x09, 0x55, 0x77, 0x17, 0xd0, 0x35, 0x0d, 0x57, 0x2a, 0xdb, 0x6d, 0x68, 0x3d, 0x41, 0x95,
	0x78, 0x1b, 0x4c, 0x79, 0x1d, 0x6a, 0x59, 0x03, 0xb6, 0x54, 0xac, 0x04, 0x7b, 0xb0, 0x7b, 0xe9,
	0xc7, 0xfc, 0x56, 0x5f, 0x86, 0x5a, 0x78, 0x06, 0xed, 0x15, 0x3c, 0x4d, 0x78, 0x7a, 0xa9, 0x1a,
	0x1b, 0x2f, 0xd5, 0x49, 0x45, 0xfe, 0x2f, 0x3b, 0x79, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x46,
	0x78, 0x5f, 0x3d, 0x07, 0x00, 0x00,
}
