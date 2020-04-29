// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/logs.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	opsmngr "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	io "io"
	reflect "reflect"
)

// MockLogsDownloader is a mock of LogsDownloader interface
type MockLogsDownloader struct {
	ctrl     *gomock.Controller
	recorder *MockLogsDownloaderMockRecorder
}

// MockLogsDownloaderMockRecorder is the mock recorder for MockLogsDownloader
type MockLogsDownloaderMockRecorder struct {
	mock *MockLogsDownloader
}

// NewMockLogsDownloader creates a new mock instance
func NewMockLogsDownloader(ctrl *gomock.Controller) *MockLogsDownloader {
	mock := &MockLogsDownloader{ctrl: ctrl}
	mock.recorder = &MockLogsDownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogsDownloader) EXPECT() *MockLogsDownloaderMockRecorder {
	return m.recorder
}

// DownloadLog mocks base method
func (m *MockLogsDownloader) DownloadLog(arg0, arg1, arg2 string, arg3 io.Writer, arg4 *mongodbatlas.DateRangetOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadLog", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// DownloadLog indicates an expected call of DownloadLog
func (mr *MockLogsDownloaderMockRecorder) DownloadLog(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadLog", reflect.TypeOf((*MockLogsDownloader)(nil).DownloadLog), arg0, arg1, arg2, arg3, arg4)
}

// MockLogJobsDownloader is a mock of LogJobsDownloader interface
type MockLogJobsDownloader struct {
	ctrl     *gomock.Controller
	recorder *MockLogJobsDownloaderMockRecorder
}

// MockLogJobsDownloaderMockRecorder is the mock recorder for MockLogJobsDownloader
type MockLogJobsDownloaderMockRecorder struct {
	mock *MockLogJobsDownloader
}

// NewMockLogJobsDownloader creates a new mock instance
func NewMockLogJobsDownloader(ctrl *gomock.Controller) *MockLogJobsDownloader {
	mock := &MockLogJobsDownloader{ctrl: ctrl}
	mock.recorder = &MockLogJobsDownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogJobsDownloader) EXPECT() *MockLogJobsDownloaderMockRecorder {
	return m.recorder
}

// DownloadLogJob mocks base method
func (m *MockLogJobsDownloader) DownloadLogJob(arg0, arg1 string, arg2 io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadLogJob", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DownloadLogJob indicates an expected call of DownloadLogJob
func (mr *MockLogJobsDownloaderMockRecorder) DownloadLogJob(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadLogJob", reflect.TypeOf((*MockLogJobsDownloader)(nil).DownloadLogJob), arg0, arg1, arg2)
}

// MockLogCollector is a mock of LogCollector interface
type MockLogCollector struct {
	ctrl     *gomock.Controller
	recorder *MockLogCollectorMockRecorder
}

// MockLogCollectorMockRecorder is the mock recorder for MockLogCollector
type MockLogCollectorMockRecorder struct {
	mock *MockLogCollector
}

// NewMockLogCollector creates a new mock instance
func NewMockLogCollector(ctrl *gomock.Controller) *MockLogCollector {
	mock := &MockLogCollector{ctrl: ctrl}
	mock.recorder = &MockLogCollectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogCollector) EXPECT() *MockLogCollectorMockRecorder {
	return m.recorder
}

// Collect mocks base method
func (m *MockLogCollector) Collect(arg0 string, arg1 *opsmngr.LogCollectionJob) (*opsmngr.LogCollectionJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Collect", arg0, arg1)
	ret0, _ := ret[0].(*opsmngr.LogCollectionJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Collect indicates an expected call of Collect
func (mr *MockLogCollectorMockRecorder) Collect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collect", reflect.TypeOf((*MockLogCollector)(nil).Collect), arg0, arg1)
}

// MockLogJobLister is a mock of LogJobLister interface
type MockLogJobLister struct {
	ctrl     *gomock.Controller
	recorder *MockLogJobListerMockRecorder
}

// MockLogJobListerMockRecorder is the mock recorder for MockLogJobLister
type MockLogJobListerMockRecorder struct {
	mock *MockLogJobLister
}

// NewMockLogJobLister creates a new mock instance
func NewMockLogJobLister(ctrl *gomock.Controller) *MockLogJobLister {
	mock := &MockLogJobLister{ctrl: ctrl}
	mock.recorder = &MockLogJobListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogJobLister) EXPECT() *MockLogJobListerMockRecorder {
	return m.recorder
}

// LogCollectionJobs mocks base method
func (m *MockLogJobLister) LogCollectionJobs(arg0 string, arg1 *opsmngr.LogListOptions) (*opsmngr.LogCollectionJobs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogCollectionJobs", arg0, arg1)
	ret0, _ := ret[0].(*opsmngr.LogCollectionJobs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogCollectionJobs indicates an expected call of LogCollectionJobs
func (mr *MockLogJobListerMockRecorder) LogCollectionJobs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogCollectionJobs", reflect.TypeOf((*MockLogJobLister)(nil).LogCollectionJobs), arg0, arg1)
}

// MockLogJobDeleter is a mock of LogJobDeleter interface
type MockLogJobDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockLogJobDeleterMockRecorder
}

// MockLogJobDeleterMockRecorder is the mock recorder for MockLogJobDeleter
type MockLogJobDeleterMockRecorder struct {
	mock *MockLogJobDeleter
}

// NewMockLogJobDeleter creates a new mock instance
func NewMockLogJobDeleter(ctrl *gomock.Controller) *MockLogJobDeleter {
	mock := &MockLogJobDeleter{ctrl: ctrl}
	mock.recorder = &MockLogJobDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogJobDeleter) EXPECT() *MockLogJobDeleterMockRecorder {
	return m.recorder
}

// DeleteCollectionJob mocks base method
func (m *MockLogJobDeleter) DeleteCollectionJob(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollectionJob", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollectionJob indicates an expected call of DeleteCollectionJob
func (mr *MockLogJobDeleterMockRecorder) DeleteCollectionJob(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollectionJob", reflect.TypeOf((*MockLogJobDeleter)(nil).DeleteCollectionJob), arg0, arg1)
}