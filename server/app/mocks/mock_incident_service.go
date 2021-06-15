// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-incident-collaboration/server/app (interfaces: IncidentService)

// Package mock_app is a generated GoMock package.
package mock_app

import (
	gomock "github.com/golang/mock/gomock"
	app "github.com/mattermost/mattermost-plugin-incident-collaboration/server/app"
	model "github.com/mattermost/mattermost-server/v5/model"
	reflect "reflect"
	time "time"
)

// MockIncidentService is a mock of IncidentService interface
type MockIncidentService struct {
	ctrl     *gomock.Controller
	recorder *MockIncidentServiceMockRecorder
}

// MockIncidentServiceMockRecorder is the mock recorder for MockIncidentService
type MockIncidentServiceMockRecorder struct {
	mock *MockIncidentService
}

// NewMockIncidentService creates a new mock instance
func NewMockIncidentService(ctrl *gomock.Controller) *MockIncidentService {
	mock := &MockIncidentService{ctrl: ctrl}
	mock.recorder = &MockIncidentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIncidentService) EXPECT() *MockIncidentServiceMockRecorder {
	return m.recorder
}

// AddChecklistItem mocks base method
func (m *MockIncidentService) AddChecklistItem(arg0, arg1 string, arg2 int, arg3 app.ChecklistItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddChecklistItem", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddChecklistItem indicates an expected call of AddChecklistItem
func (mr *MockIncidentServiceMockRecorder) AddChecklistItem(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChecklistItem", reflect.TypeOf((*MockIncidentService)(nil).AddChecklistItem), arg0, arg1, arg2, arg3)
}

// AddPostToTimeline mocks base method
func (m *MockIncidentService) AddPostToTimeline(arg0, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPostToTimeline", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPostToTimeline indicates an expected call of AddPostToTimeline
func (mr *MockIncidentServiceMockRecorder) AddPostToTimeline(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPostToTimeline", reflect.TypeOf((*MockIncidentService)(nil).AddPostToTimeline), arg0, arg1, arg2, arg3)
}

// CancelRetrospective mocks base method
func (m *MockIncidentService) CancelRetrospective(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelRetrospective", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelRetrospective indicates an expected call of CancelRetrospective
func (mr *MockIncidentServiceMockRecorder) CancelRetrospective(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelRetrospective", reflect.TypeOf((*MockIncidentService)(nil).CancelRetrospective), arg0, arg1)
}

// ChangeCreationDate mocks base method
func (m *MockIncidentService) ChangeCreationDate(arg0 string, arg1 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeCreationDate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeCreationDate indicates an expected call of ChangeCreationDate
func (mr *MockIncidentServiceMockRecorder) ChangeCreationDate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeCreationDate", reflect.TypeOf((*MockIncidentService)(nil).ChangeCreationDate), arg0, arg1)
}

// ChangeOwner mocks base method
func (m *MockIncidentService) ChangeOwner(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeOwner", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeOwner indicates an expected call of ChangeOwner
func (mr *MockIncidentServiceMockRecorder) ChangeOwner(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeOwner", reflect.TypeOf((*MockIncidentService)(nil).ChangeOwner), arg0, arg1, arg2)
}

// CheckAndSendMessageOnJoin mocks base method
func (m *MockIncidentService) CheckAndSendMessageOnJoin(arg0, arg1, arg2 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAndSendMessageOnJoin", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckAndSendMessageOnJoin indicates an expected call of CheckAndSendMessageOnJoin
func (mr *MockIncidentServiceMockRecorder) CheckAndSendMessageOnJoin(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAndSendMessageOnJoin", reflect.TypeOf((*MockIncidentService)(nil).CheckAndSendMessageOnJoin), arg0, arg1, arg2)
}

// CreateIncident mocks base method
func (m *MockIncidentService) CreateIncident(arg0 *app.Incident, arg1 *app.Playbook, arg2 string, arg3 bool) (*app.Incident, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIncident", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*app.Incident)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIncident indicates an expected call of CreateIncident
func (mr *MockIncidentServiceMockRecorder) CreateIncident(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIncident", reflect.TypeOf((*MockIncidentService)(nil).CreateIncident), arg0, arg1, arg2, arg3)
}

// EditChecklistItem mocks base method
func (m *MockIncidentService) EditChecklistItem(arg0, arg1 string, arg2, arg3 int, arg4, arg5, arg6 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditChecklistItem", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditChecklistItem indicates an expected call of EditChecklistItem
func (mr *MockIncidentServiceMockRecorder) EditChecklistItem(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditChecklistItem", reflect.TypeOf((*MockIncidentService)(nil).EditChecklistItem), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// GetChecklistAutocomplete mocks base method
func (m *MockIncidentService) GetChecklistAutocomplete(arg0 string) ([]model.AutocompleteListItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChecklistAutocomplete", arg0)
	ret0, _ := ret[0].([]model.AutocompleteListItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChecklistAutocomplete indicates an expected call of GetChecklistAutocomplete
func (mr *MockIncidentServiceMockRecorder) GetChecklistAutocomplete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChecklistAutocomplete", reflect.TypeOf((*MockIncidentService)(nil).GetChecklistAutocomplete), arg0)
}

// GetChecklistItemAutocomplete mocks base method
func (m *MockIncidentService) GetChecklistItemAutocomplete(arg0 string) ([]model.AutocompleteListItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChecklistItemAutocomplete", arg0)
	ret0, _ := ret[0].([]model.AutocompleteListItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChecklistItemAutocomplete indicates an expected call of GetChecklistItemAutocomplete
func (mr *MockIncidentServiceMockRecorder) GetChecklistItemAutocomplete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChecklistItemAutocomplete", reflect.TypeOf((*MockIncidentService)(nil).GetChecklistItemAutocomplete), arg0)
}

// GetIncident mocks base method
func (m *MockIncidentService) GetIncident(arg0 string) (*app.Incident, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncident", arg0)
	ret0, _ := ret[0].(*app.Incident)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncident indicates an expected call of GetIncident
func (mr *MockIncidentServiceMockRecorder) GetIncident(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncident", reflect.TypeOf((*MockIncidentService)(nil).GetIncident), arg0)
}

// GetIncidentIDForChannel mocks base method
func (m *MockIncidentService) GetIncidentIDForChannel(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncidentIDForChannel", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncidentIDForChannel indicates an expected call of GetIncidentIDForChannel
func (mr *MockIncidentServiceMockRecorder) GetIncidentIDForChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncidentIDForChannel", reflect.TypeOf((*MockIncidentService)(nil).GetIncidentIDForChannel), arg0)
}

// GetIncidentMetadata mocks base method
func (m *MockIncidentService) GetIncidentMetadata(arg0 string) (*app.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncidentMetadata", arg0)
	ret0, _ := ret[0].(*app.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncidentMetadata indicates an expected call of GetIncidentMetadata
func (mr *MockIncidentServiceMockRecorder) GetIncidentMetadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncidentMetadata", reflect.TypeOf((*MockIncidentService)(nil).GetIncidentMetadata), arg0)
}

// GetIncidents mocks base method
func (m *MockIncidentService) GetIncidents(arg0 app.RequesterInfo, arg1 app.IncidentFilterOptions) (*app.GetIncidentsResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncidents", arg0, arg1)
	ret0, _ := ret[0].(*app.GetIncidentsResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncidents indicates an expected call of GetIncidents
func (mr *MockIncidentServiceMockRecorder) GetIncidents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncidents", reflect.TypeOf((*MockIncidentService)(nil).GetIncidents), arg0, arg1)
}

// GetOwners mocks base method
func (m *MockIncidentService) GetOwners(arg0 app.RequesterInfo, arg1 app.IncidentFilterOptions) ([]app.OwnerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOwners", arg0, arg1)
	ret0, _ := ret[0].([]app.OwnerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOwners indicates an expected call of GetOwners
func (mr *MockIncidentServiceMockRecorder) GetOwners(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwners", reflect.TypeOf((*MockIncidentService)(nil).GetOwners), arg0, arg1)
}

// HandleReminder mocks base method
func (m *MockIncidentService) HandleReminder(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleReminder", arg0)
}

// HandleReminder indicates an expected call of HandleReminder
func (mr *MockIncidentServiceMockRecorder) HandleReminder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleReminder", reflect.TypeOf((*MockIncidentService)(nil).HandleReminder), arg0)
}

// IsOwner mocks base method
func (m *MockIncidentService) IsOwner(arg0, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOwner", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsOwner indicates an expected call of IsOwner
func (mr *MockIncidentServiceMockRecorder) IsOwner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOwner", reflect.TypeOf((*MockIncidentService)(nil).IsOwner), arg0, arg1)
}

// ModifyCheckedState mocks base method
func (m *MockIncidentService) ModifyCheckedState(arg0, arg1, arg2 string, arg3, arg4 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyCheckedState", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyCheckedState indicates an expected call of ModifyCheckedState
func (mr *MockIncidentServiceMockRecorder) ModifyCheckedState(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyCheckedState", reflect.TypeOf((*MockIncidentService)(nil).ModifyCheckedState), arg0, arg1, arg2, arg3, arg4)
}

// MoveChecklistItem mocks base method
func (m *MockIncidentService) MoveChecklistItem(arg0, arg1 string, arg2, arg3, arg4 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveChecklistItem", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveChecklistItem indicates an expected call of MoveChecklistItem
func (mr *MockIncidentServiceMockRecorder) MoveChecklistItem(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveChecklistItem", reflect.TypeOf((*MockIncidentService)(nil).MoveChecklistItem), arg0, arg1, arg2, arg3, arg4)
}

// NukeDB mocks base method
func (m *MockIncidentService) NukeDB() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NukeDB")
	ret0, _ := ret[0].(error)
	return ret0
}

// NukeDB indicates an expected call of NukeDB
func (mr *MockIncidentServiceMockRecorder) NukeDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NukeDB", reflect.TypeOf((*MockIncidentService)(nil).NukeDB))
}

// OpenAddChecklistItemDialog mocks base method
func (m *MockIncidentService) OpenAddChecklistItemDialog(arg0, arg1 string, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenAddChecklistItemDialog", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenAddChecklistItemDialog indicates an expected call of OpenAddChecklistItemDialog
func (mr *MockIncidentServiceMockRecorder) OpenAddChecklistItemDialog(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenAddChecklistItemDialog", reflect.TypeOf((*MockIncidentService)(nil).OpenAddChecklistItemDialog), arg0, arg1, arg2)
}

// OpenAddToTimelineDialog mocks base method
func (m *MockIncidentService) OpenAddToTimelineDialog(arg0 app.RequesterInfo, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenAddToTimelineDialog", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenAddToTimelineDialog indicates an expected call of OpenAddToTimelineDialog
func (mr *MockIncidentServiceMockRecorder) OpenAddToTimelineDialog(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenAddToTimelineDialog", reflect.TypeOf((*MockIncidentService)(nil).OpenAddToTimelineDialog), arg0, arg1, arg2, arg3)
}

// OpenCreateIncidentDialog mocks base method
func (m *MockIncidentService) OpenCreateIncidentDialog(arg0, arg1, arg2, arg3, arg4 string, arg5 []app.Playbook, arg6 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenCreateIncidentDialog", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenCreateIncidentDialog indicates an expected call of OpenCreateIncidentDialog
func (mr *MockIncidentServiceMockRecorder) OpenCreateIncidentDialog(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenCreateIncidentDialog", reflect.TypeOf((*MockIncidentService)(nil).OpenCreateIncidentDialog), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// OpenUpdateStatusDialog mocks base method
func (m *MockIncidentService) OpenUpdateStatusDialog(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUpdateStatusDialog", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenUpdateStatusDialog indicates an expected call of OpenUpdateStatusDialog
func (mr *MockIncidentServiceMockRecorder) OpenUpdateStatusDialog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUpdateStatusDialog", reflect.TypeOf((*MockIncidentService)(nil).OpenUpdateStatusDialog), arg0, arg1)
}

// PublishRetrospective mocks base method
func (m *MockIncidentService) PublishRetrospective(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishRetrospective", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishRetrospective indicates an expected call of PublishRetrospective
func (mr *MockIncidentServiceMockRecorder) PublishRetrospective(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishRetrospective", reflect.TypeOf((*MockIncidentService)(nil).PublishRetrospective), arg0, arg1, arg2)
}

// RemoveChecklistItem mocks base method
func (m *MockIncidentService) RemoveChecklistItem(arg0, arg1 string, arg2, arg3 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveChecklistItem", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveChecklistItem indicates an expected call of RemoveChecklistItem
func (mr *MockIncidentServiceMockRecorder) RemoveChecklistItem(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveChecklistItem", reflect.TypeOf((*MockIncidentService)(nil).RemoveChecklistItem), arg0, arg1, arg2, arg3)
}

// RemoveReminder mocks base method
func (m *MockIncidentService) RemoveReminder(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveReminder", arg0)
}

// RemoveReminder indicates an expected call of RemoveReminder
func (mr *MockIncidentServiceMockRecorder) RemoveReminder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveReminder", reflect.TypeOf((*MockIncidentService)(nil).RemoveReminder), arg0)
}

// RemoveReminderPost mocks base method
func (m *MockIncidentService) RemoveReminderPost(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveReminderPost", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveReminderPost indicates an expected call of RemoveReminderPost
func (mr *MockIncidentServiceMockRecorder) RemoveReminderPost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveReminderPost", reflect.TypeOf((*MockIncidentService)(nil).RemoveReminderPost), arg0)
}

// RemoveTimelineEvent mocks base method
func (m *MockIncidentService) RemoveTimelineEvent(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTimelineEvent", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTimelineEvent indicates an expected call of RemoveTimelineEvent
func (mr *MockIncidentServiceMockRecorder) RemoveTimelineEvent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTimelineEvent", reflect.TypeOf((*MockIncidentService)(nil).RemoveTimelineEvent), arg0, arg1, arg2)
}

// RunChecklistItemSlashCommand mocks base method
func (m *MockIncidentService) RunChecklistItemSlashCommand(arg0, arg1 string, arg2, arg3 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunChecklistItemSlashCommand", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunChecklistItemSlashCommand indicates an expected call of RunChecklistItemSlashCommand
func (mr *MockIncidentServiceMockRecorder) RunChecklistItemSlashCommand(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunChecklistItemSlashCommand", reflect.TypeOf((*MockIncidentService)(nil).RunChecklistItemSlashCommand), arg0, arg1, arg2, arg3)
}

// SetAssignee mocks base method
func (m *MockIncidentService) SetAssignee(arg0, arg1, arg2 string, arg3, arg4 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAssignee", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAssignee indicates an expected call of SetAssignee
func (mr *MockIncidentServiceMockRecorder) SetAssignee(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAssignee", reflect.TypeOf((*MockIncidentService)(nil).SetAssignee), arg0, arg1, arg2, arg3, arg4)
}

// SetReminder mocks base method
func (m *MockIncidentService) SetReminder(arg0 string, arg1 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReminder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReminder indicates an expected call of SetReminder
func (mr *MockIncidentServiceMockRecorder) SetReminder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReminder", reflect.TypeOf((*MockIncidentService)(nil).SetReminder), arg0, arg1)
}

// ToggleCheckedState mocks base method
func (m *MockIncidentService) ToggleCheckedState(arg0, arg1 string, arg2, arg3 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToggleCheckedState", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ToggleCheckedState indicates an expected call of ToggleCheckedState
func (mr *MockIncidentServiceMockRecorder) ToggleCheckedState(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToggleCheckedState", reflect.TypeOf((*MockIncidentService)(nil).ToggleCheckedState), arg0, arg1, arg2, arg3)
}

// UpdateRetrospective mocks base method
func (m *MockIncidentService) UpdateRetrospective(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRetrospective", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRetrospective indicates an expected call of UpdateRetrospective
func (mr *MockIncidentServiceMockRecorder) UpdateRetrospective(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRetrospective", reflect.TypeOf((*MockIncidentService)(nil).UpdateRetrospective), arg0, arg1, arg2)
}

// UpdateStatus mocks base method
func (m *MockIncidentService) UpdateStatus(arg0, arg1 string, arg2 app.StatusUpdateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockIncidentServiceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockIncidentService)(nil).UpdateStatus), arg0, arg1, arg2)
}

// UserHasJoinedChannel mocks base method
func (m *MockIncidentService) UserHasJoinedChannel(arg0, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UserHasJoinedChannel", arg0, arg1, arg2)
}

// UserHasJoinedChannel indicates an expected call of UserHasJoinedChannel
func (mr *MockIncidentServiceMockRecorder) UserHasJoinedChannel(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserHasJoinedChannel", reflect.TypeOf((*MockIncidentService)(nil).UserHasJoinedChannel), arg0, arg1, arg2)
}

// UserHasLeftChannel mocks base method
func (m *MockIncidentService) UserHasLeftChannel(arg0, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UserHasLeftChannel", arg0, arg1, arg2)
}

// UserHasLeftChannel indicates an expected call of UserHasLeftChannel
func (mr *MockIncidentServiceMockRecorder) UserHasLeftChannel(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserHasLeftChannel", reflect.TypeOf((*MockIncidentService)(nil).UserHasLeftChannel), arg0, arg1, arg2)
}