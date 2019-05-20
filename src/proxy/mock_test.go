// Automatically generated by MockGen. DO NOT EDIT!
// Source: handlers.go

package proxy

import (
	hash "github.com/Fantom-foundation/go-lachesis/src/hash"
	inter "github.com/Fantom-foundation/go-lachesis/src/inter"
	poset "github.com/Fantom-foundation/go-lachesis/src/poset"
	gomock "github.com/golang/mock/gomock"
)

// Mock of App interface
type MockApp struct {
	ctrl     *gomock.Controller
	recorder *_MockAppRecorder
}

// Recorder for MockApp (not exported)
type _MockAppRecorder struct {
	mock *MockApp
}

func NewMockApp(ctrl *gomock.Controller) *MockApp {
	mock := &MockApp{ctrl: ctrl}
	mock.recorder = &_MockAppRecorder{mock}
	return mock
}

func (_m *MockApp) EXPECT() *_MockAppRecorder {
	return _m.recorder
}

func (_m *MockApp) CommitHandler(block poset.Block) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "CommitHandler", block)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAppRecorder) CommitHandler(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CommitHandler", arg0)
}

func (_m *MockApp) SnapshotHandler(blockIndex int64) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "SnapshotHandler", blockIndex)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAppRecorder) SnapshotHandler(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SnapshotHandler", arg0)
}

func (_m *MockApp) RestoreHandler(snapshot []byte) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "RestoreHandler", snapshot)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAppRecorder) RestoreHandler(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RestoreHandler", arg0)
}

// Mock of Node interface
type MockNode struct {
	ctrl     *gomock.Controller
	recorder *_MockNodeRecorder
}

// Recorder for MockNode (not exported)
type _MockNodeRecorder struct {
	mock *MockNode
}

func NewMockNode(ctrl *gomock.Controller) *MockNode {
	mock := &MockNode{ctrl: ctrl}
	mock.recorder = &_MockNodeRecorder{mock}
	return mock
}

func (_m *MockNode) EXPECT() *_MockNodeRecorder {
	return _m.recorder
}

func (_m *MockNode) GetID() hash.Peer {
	ret := _m.ctrl.Call(_m, "GetID")
	ret0, _ := ret[0].(hash.Peer)
	return ret0
}

func (_mr *_MockNodeRecorder) GetID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetID")
}

func (_m *MockNode) AddInternalTxn(_param0 inter.InternalTransaction) {
	_m.ctrl.Call(_m, "AddInternalTxn", _param0)
}

func (_mr *_MockNodeRecorder) AddInternalTxn(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddInternalTxn", arg0)
}

func (_m *MockNode) GetInternalTxns() []*inter.InternalTransaction {
	ret := _m.ctrl.Call(_m, "GetInternalTxns")
	ret0, _ := ret[0].([]*inter.InternalTransaction)
	return ret0
}

func (_mr *_MockNodeRecorder) GetInternalTxns() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetInternalTxns")
}

// Mock of Consensus interface
type MockConsensus struct {
	ctrl     *gomock.Controller
	recorder *_MockConsensusRecorder
}

// Recorder for MockConsensus (not exported)
type _MockConsensusRecorder struct {
	mock *MockConsensus
}

func NewMockConsensus(ctrl *gomock.Controller) *MockConsensus {
	mock := &MockConsensus{ctrl: ctrl}
	mock.recorder = &_MockConsensusRecorder{mock}
	return mock
}

func (_m *MockConsensus) EXPECT() *_MockConsensusRecorder {
	return _m.recorder
}

func (_m *MockConsensus) GetBalanceOf(peer hash.Peer) uint64 {
	ret := _m.ctrl.Call(_m, "GetBalanceOf", peer)
	ret0, _ := ret[0].(uint64)
	return ret0
}

func (_mr *_MockConsensusRecorder) GetBalanceOf(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetBalanceOf", arg0)
}