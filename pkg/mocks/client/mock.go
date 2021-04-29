// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bakito/adguardhome-sync/pkg/client (interfaces: Client)

// Package mock_client is a generated GoMock package.
package mock_client

import (
	reflect "reflect"

	types "github.com/bakito/adguardhome-sync/pkg/types"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AccessList mocks base method.
func (m *MockClient) AccessList() (*types.AccessList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessList")
	ret0, _ := ret[0].(*types.AccessList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccessList indicates an expected call of AccessList.
func (mr *MockClientMockRecorder) AccessList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessList", reflect.TypeOf((*MockClient)(nil).AccessList))
}

// AddClients mocks base method.
func (m *MockClient) AddClients(arg0 ...types.Client) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddClients", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddClients indicates an expected call of AddClients.
func (mr *MockClientMockRecorder) AddClients(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddClients", reflect.TypeOf((*MockClient)(nil).AddClients), arg0...)
}

// AddFilters mocks base method.
func (m *MockClient) AddFilters(arg0 bool, arg1 ...types.Filter) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddFilters", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFilters indicates an expected call of AddFilters.
func (mr *MockClientMockRecorder) AddFilters(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFilters", reflect.TypeOf((*MockClient)(nil).AddFilters), varargs...)
}

// AddRewriteEntries mocks base method.
func (m *MockClient) AddRewriteEntries(arg0 ...types.RewriteEntry) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddRewriteEntries", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRewriteEntries indicates an expected call of AddRewriteEntries.
func (mr *MockClientMockRecorder) AddRewriteEntries(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRewriteEntries", reflect.TypeOf((*MockClient)(nil).AddRewriteEntries), arg0...)
}

// Clients mocks base method.
func (m *MockClient) Clients() (*types.Clients, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clients")
	ret0, _ := ret[0].(*types.Clients)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Clients indicates an expected call of Clients.
func (mr *MockClientMockRecorder) Clients() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clients", reflect.TypeOf((*MockClient)(nil).Clients))
}

// DNSConfig mocks base method.
func (m *MockClient) DNSConfig() (*types.DNSConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DNSConfig")
	ret0, _ := ret[0].(*types.DNSConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DNSConfig indicates an expected call of DNSConfig.
func (mr *MockClientMockRecorder) DNSConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DNSConfig", reflect.TypeOf((*MockClient)(nil).DNSConfig))
}

// DeleteClients mocks base method.
func (m *MockClient) DeleteClients(arg0 ...types.Client) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteClients", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteClients indicates an expected call of DeleteClients.
func (mr *MockClientMockRecorder) DeleteClients(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClients", reflect.TypeOf((*MockClient)(nil).DeleteClients), arg0...)
}

// DeleteFilters mocks base method.
func (m *MockClient) DeleteFilters(arg0 bool, arg1 ...types.Filter) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteFilters", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFilters indicates an expected call of DeleteFilters.
func (mr *MockClientMockRecorder) DeleteFilters(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFilters", reflect.TypeOf((*MockClient)(nil).DeleteFilters), varargs...)
}

// DeleteRewriteEntries mocks base method.
func (m *MockClient) DeleteRewriteEntries(arg0 ...types.RewriteEntry) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteRewriteEntries", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRewriteEntries indicates an expected call of DeleteRewriteEntries.
func (mr *MockClientMockRecorder) DeleteRewriteEntries(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRewriteEntries", reflect.TypeOf((*MockClient)(nil).DeleteRewriteEntries), arg0...)
}

// Filtering mocks base method.
func (m *MockClient) Filtering() (*types.FilteringStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filtering")
	ret0, _ := ret[0].(*types.FilteringStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filtering indicates an expected call of Filtering.
func (mr *MockClientMockRecorder) Filtering() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filtering", reflect.TypeOf((*MockClient)(nil).Filtering))
}

// Host mocks base method.
func (m *MockClient) Host() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Host")
	ret0, _ := ret[0].(string)
	return ret0
}

// Host indicates an expected call of Host.
func (mr *MockClientMockRecorder) Host() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Host", reflect.TypeOf((*MockClient)(nil).Host))
}

// Parental mocks base method.
func (m *MockClient) Parental() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parental")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parental indicates an expected call of Parental.
func (mr *MockClientMockRecorder) Parental() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parental", reflect.TypeOf((*MockClient)(nil).Parental))
}

// QueryLogConfig mocks base method.
func (m *MockClient) QueryLogConfig() (*types.QueryLogConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryLogConfig")
	ret0, _ := ret[0].(*types.QueryLogConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryLogConfig indicates an expected call of QueryLogConfig.
func (mr *MockClientMockRecorder) QueryLogConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryLogConfig", reflect.TypeOf((*MockClient)(nil).QueryLogConfig))
}

// RefreshFilters mocks base method.
func (m *MockClient) RefreshFilters(arg0 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshFilters", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshFilters indicates an expected call of RefreshFilters.
func (mr *MockClientMockRecorder) RefreshFilters(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshFilters", reflect.TypeOf((*MockClient)(nil).RefreshFilters), arg0)
}

// RewriteList mocks base method.
func (m *MockClient) RewriteList() (*types.RewriteEntries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RewriteList")
	ret0, _ := ret[0].(*types.RewriteEntries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RewriteList indicates an expected call of RewriteList.
func (mr *MockClientMockRecorder) RewriteList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RewriteList", reflect.TypeOf((*MockClient)(nil).RewriteList))
}

// SafeBrowsing mocks base method.
func (m *MockClient) SafeBrowsing() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SafeBrowsing")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SafeBrowsing indicates an expected call of SafeBrowsing.
func (mr *MockClientMockRecorder) SafeBrowsing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SafeBrowsing", reflect.TypeOf((*MockClient)(nil).SafeBrowsing))
}

// SafeSearch mocks base method.
func (m *MockClient) SafeSearch() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SafeSearch")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SafeSearch indicates an expected call of SafeSearch.
func (mr *MockClientMockRecorder) SafeSearch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SafeSearch", reflect.TypeOf((*MockClient)(nil).SafeSearch))
}

// Services mocks base method.
func (m *MockClient) Services() (types.Services, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Services")
	ret0, _ := ret[0].(types.Services)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Services indicates an expected call of Services.
func (mr *MockClientMockRecorder) Services() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Services", reflect.TypeOf((*MockClient)(nil).Services))
}

// SetAccessList mocks base method.
func (m *MockClient) SetAccessList(arg0 *types.AccessList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAccessList", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAccessList indicates an expected call of SetAccessList.
func (mr *MockClientMockRecorder) SetAccessList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccessList", reflect.TypeOf((*MockClient)(nil).SetAccessList), arg0)
}

// SetCustomRules mocks base method.
func (m *MockClient) SetCustomRules(arg0 types.UserRules) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCustomRules", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCustomRules indicates an expected call of SetCustomRules.
func (mr *MockClientMockRecorder) SetCustomRules(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCustomRules", reflect.TypeOf((*MockClient)(nil).SetCustomRules), arg0)
}

// SetDNSConfig mocks base method.
func (m *MockClient) SetDNSConfig(arg0 *types.DNSConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDNSConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDNSConfig indicates an expected call of SetDNSConfig.
func (mr *MockClientMockRecorder) SetDNSConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDNSConfig", reflect.TypeOf((*MockClient)(nil).SetDNSConfig), arg0)
}

// SetQueryLogConfig mocks base method.
func (m *MockClient) SetQueryLogConfig(arg0 bool, arg1 int, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetQueryLogConfig", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetQueryLogConfig indicates an expected call of SetQueryLogConfig.
func (mr *MockClientMockRecorder) SetQueryLogConfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetQueryLogConfig", reflect.TypeOf((*MockClient)(nil).SetQueryLogConfig), arg0, arg1, arg2)
}

// SetServices mocks base method.
func (m *MockClient) SetServices(arg0 types.Services) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetServices", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetServices indicates an expected call of SetServices.
func (mr *MockClientMockRecorder) SetServices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetServices", reflect.TypeOf((*MockClient)(nil).SetServices), arg0)
}

// SetStatsConfig mocks base method.
func (m *MockClient) SetStatsConfig(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStatsConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStatsConfig indicates an expected call of SetStatsConfig.
func (mr *MockClientMockRecorder) SetStatsConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatsConfig", reflect.TypeOf((*MockClient)(nil).SetStatsConfig), arg0)
}

// Setup mocks base method.
func (m *MockClient) Setup() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Setup")
	ret0, _ := ret[0].(error)
	return ret0
}

// Setup indicates an expected call of Setup.
func (mr *MockClientMockRecorder) Setup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockClient)(nil).Setup))
}

// StatsConfig mocks base method.
func (m *MockClient) StatsConfig() (*types.IntervalConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatsConfig")
	ret0, _ := ret[0].(*types.IntervalConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatsConfig indicates an expected call of StatsConfig.
func (mr *MockClientMockRecorder) StatsConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatsConfig", reflect.TypeOf((*MockClient)(nil).StatsConfig))
}

// Status mocks base method.
func (m *MockClient) Status() (*types.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(*types.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockClientMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockClient)(nil).Status))
}

// ToggleFiltering mocks base method.
func (m *MockClient) ToggleFiltering(arg0 bool, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToggleFiltering", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ToggleFiltering indicates an expected call of ToggleFiltering.
func (mr *MockClientMockRecorder) ToggleFiltering(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToggleFiltering", reflect.TypeOf((*MockClient)(nil).ToggleFiltering), arg0, arg1)
}

// ToggleParental mocks base method.
func (m *MockClient) ToggleParental(arg0 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToggleParental", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ToggleParental indicates an expected call of ToggleParental.
func (mr *MockClientMockRecorder) ToggleParental(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToggleParental", reflect.TypeOf((*MockClient)(nil).ToggleParental), arg0)
}

// ToggleProtection mocks base method.
func (m *MockClient) ToggleProtection(arg0 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToggleProtection", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ToggleProtection indicates an expected call of ToggleProtection.
func (mr *MockClientMockRecorder) ToggleProtection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToggleProtection", reflect.TypeOf((*MockClient)(nil).ToggleProtection), arg0)
}

// ToggleSafeBrowsing mocks base method.
func (m *MockClient) ToggleSafeBrowsing(arg0 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToggleSafeBrowsing", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ToggleSafeBrowsing indicates an expected call of ToggleSafeBrowsing.
func (mr *MockClientMockRecorder) ToggleSafeBrowsing(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToggleSafeBrowsing", reflect.TypeOf((*MockClient)(nil).ToggleSafeBrowsing), arg0)
}

// ToggleSafeSearch mocks base method.
func (m *MockClient) ToggleSafeSearch(arg0 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToggleSafeSearch", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ToggleSafeSearch indicates an expected call of ToggleSafeSearch.
func (mr *MockClientMockRecorder) ToggleSafeSearch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToggleSafeSearch", reflect.TypeOf((*MockClient)(nil).ToggleSafeSearch), arg0)
}

// UpdateClients mocks base method.
func (m *MockClient) UpdateClients(arg0 ...types.Client) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateClients", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClients indicates an expected call of UpdateClients.
func (mr *MockClientMockRecorder) UpdateClients(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClients", reflect.TypeOf((*MockClient)(nil).UpdateClients), arg0...)
}

// UpdateFilters mocks base method.
func (m *MockClient) UpdateFilters(arg0 bool, arg1 ...types.Filter) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateFilters", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFilters indicates an expected call of UpdateFilters.
func (mr *MockClientMockRecorder) UpdateFilters(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFilters", reflect.TypeOf((*MockClient)(nil).UpdateFilters), varargs...)
}
