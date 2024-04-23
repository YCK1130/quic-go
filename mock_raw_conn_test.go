// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/YCK1130/quic-go (interfaces: RawConn)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/YCK1130/quic-go -destination mock_raw_conn_test.go github.com/YCK1130/quic-go RawConn
//

// Package quic is a generated GoMock package.
package quic

import (
	net "net"
	reflect "reflect"
	time "time"

	protocol "github.com/YCK1130/quic-go/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockRawConn is a mock of RawConn interface.
type MockRawConn struct {
	ctrl     *gomock.Controller
	recorder *MockRawConnMockRecorder
}

// MockRawConnMockRecorder is the mock recorder for MockRawConn.
type MockRawConnMockRecorder struct {
	mock *MockRawConn
}

// NewMockRawConn creates a new mock instance.
func NewMockRawConn(ctrl *gomock.Controller) *MockRawConn {
	mock := &MockRawConn{ctrl: ctrl}
	mock.recorder = &MockRawConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRawConn) EXPECT() *MockRawConnMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRawConn) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRawConnMockRecorder) Close() *MockRawConnCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRawConn)(nil).Close))
	return &MockRawConnCloseCall{Call: call}
}

// MockRawConnCloseCall wrap *gomock.Call
type MockRawConnCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRawConnCloseCall) Return(arg0 error) *MockRawConnCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRawConnCloseCall) Do(f func() error) *MockRawConnCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRawConnCloseCall) DoAndReturn(f func() error) *MockRawConnCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LocalAddr mocks base method.
func (m *MockRawConn) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr.
func (mr *MockRawConnMockRecorder) LocalAddr() *MockRawConnLocalAddrCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockRawConn)(nil).LocalAddr))
	return &MockRawConnLocalAddrCall{Call: call}
}

// MockRawConnLocalAddrCall wrap *gomock.Call
type MockRawConnLocalAddrCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRawConnLocalAddrCall) Return(arg0 net.Addr) *MockRawConnLocalAddrCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRawConnLocalAddrCall) Do(f func() net.Addr) *MockRawConnLocalAddrCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRawConnLocalAddrCall) DoAndReturn(f func() net.Addr) *MockRawConnLocalAddrCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReadPacket mocks base method.
func (m *MockRawConn) ReadPacket() (receivedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPacket")
	ret0, _ := ret[0].(receivedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadPacket indicates an expected call of ReadPacket.
func (mr *MockRawConnMockRecorder) ReadPacket() *MockRawConnReadPacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPacket", reflect.TypeOf((*MockRawConn)(nil).ReadPacket))
	return &MockRawConnReadPacketCall{Call: call}
}

// MockRawConnReadPacketCall wrap *gomock.Call
type MockRawConnReadPacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRawConnReadPacketCall) Return(arg0 receivedPacket, arg1 error) *MockRawConnReadPacketCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRawConnReadPacketCall) Do(f func() (receivedPacket, error)) *MockRawConnReadPacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRawConnReadPacketCall) DoAndReturn(f func() (receivedPacket, error)) *MockRawConnReadPacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetReadDeadline mocks base method.
func (m *MockRawConn) SetReadDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline.
func (mr *MockRawConnMockRecorder) SetReadDeadline(arg0 any) *MockRawConnSetReadDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockRawConn)(nil).SetReadDeadline), arg0)
	return &MockRawConnSetReadDeadlineCall{Call: call}
}

// MockRawConnSetReadDeadlineCall wrap *gomock.Call
type MockRawConnSetReadDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRawConnSetReadDeadlineCall) Return(arg0 error) *MockRawConnSetReadDeadlineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRawConnSetReadDeadlineCall) Do(f func(time.Time) error) *MockRawConnSetReadDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRawConnSetReadDeadlineCall) DoAndReturn(f func(time.Time) error) *MockRawConnSetReadDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WritePacket mocks base method.
func (m *MockRawConn) WritePacket(arg0 []byte, arg1 net.Addr, arg2 []byte, arg3 uint16, arg4 protocol.ECN) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WritePacket", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WritePacket indicates an expected call of WritePacket.
func (mr *MockRawConnMockRecorder) WritePacket(arg0, arg1, arg2, arg3, arg4 any) *MockRawConnWritePacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WritePacket", reflect.TypeOf((*MockRawConn)(nil).WritePacket), arg0, arg1, arg2, arg3, arg4)
	return &MockRawConnWritePacketCall{Call: call}
}

// MockRawConnWritePacketCall wrap *gomock.Call
type MockRawConnWritePacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRawConnWritePacketCall) Return(arg0 int, arg1 error) *MockRawConnWritePacketCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRawConnWritePacketCall) Do(f func([]byte, net.Addr, []byte, uint16, protocol.ECN) (int, error)) *MockRawConnWritePacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRawConnWritePacketCall) DoAndReturn(f func([]byte, net.Addr, []byte, uint16, protocol.ECN) (int, error)) *MockRawConnWritePacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// capabilities mocks base method.
func (m *MockRawConn) capabilities() connCapabilities {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "capabilities")
	ret0, _ := ret[0].(connCapabilities)
	return ret0
}

// capabilities indicates an expected call of capabilities.
func (mr *MockRawConnMockRecorder) capabilities() *MockRawConncapabilitiesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "capabilities", reflect.TypeOf((*MockRawConn)(nil).capabilities))
	return &MockRawConncapabilitiesCall{Call: call}
}

// MockRawConncapabilitiesCall wrap *gomock.Call
type MockRawConncapabilitiesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRawConncapabilitiesCall) Return(arg0 connCapabilities) *MockRawConncapabilitiesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRawConncapabilitiesCall) Do(f func() connCapabilities) *MockRawConncapabilitiesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRawConncapabilitiesCall) DoAndReturn(f func() connCapabilities) *MockRawConncapabilitiesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
