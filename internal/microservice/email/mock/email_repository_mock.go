// Code generated by MockGen. DO NOT EDIT.
// Source: ./iemail_repo.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	domain_models "mail/internal/microservice/models/domain_models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEmailRepository is a mock of EmailRepository interface.
type MockEmailRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEmailRepositoryMockRecorder
}

// MockEmailRepositoryMockRecorder is the mock recorder for MockEmailRepository.
type MockEmailRepositoryMockRecorder struct {
	mock *MockEmailRepository
}

// NewMockEmailRepository creates a new mock instance.
func NewMockEmailRepository(ctrl *gomock.Controller) *MockEmailRepository {
	mock := &MockEmailRepository{ctrl: ctrl}
	mock.recorder = &MockEmailRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailRepository) EXPECT() *MockEmailRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockEmailRepository) Add(emailModelCore *domain_models.Email, ctx context.Context) (uint64, *domain_models.Email, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", emailModelCore, ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(*domain_models.Email)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Add indicates an expected call of Add.
func (mr *MockEmailRepositoryMockRecorder) Add(emailModelCore, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockEmailRepository)(nil).Add), emailModelCore, ctx)
}

// AddAttachment mocks base method.
func (m *MockEmailRepository) AddAttachment(emailID, fileID uint64, ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAttachment", emailID, fileID, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAttachment indicates an expected call of AddAttachment.
func (mr *MockEmailRepositoryMockRecorder) AddAttachment(emailID, fileID, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAttachment", reflect.TypeOf((*MockEmailRepository)(nil).AddAttachment), emailID, fileID, ctx)
}

// AddFile mocks base method.
func (m *MockEmailRepository) AddFile(fileID, fileType, fileName, fileSize string, ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFile", fileID, fileType, fileName, fileSize, ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddFile indicates an expected call of AddFile.
func (mr *MockEmailRepositoryMockRecorder) AddFile(fileID, fileType, fileName, fileSize, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFile", reflect.TypeOf((*MockEmailRepository)(nil).AddFile), fileID, fileType, fileName, fileSize, ctx)
}

// AddProfileEmail mocks base method.
func (m *MockEmailRepository) AddProfileEmail(email_id uint64, sender, recipient string, ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProfileEmail", email_id, sender, recipient, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProfileEmail indicates an expected call of AddProfileEmail.
func (mr *MockEmailRepositoryMockRecorder) AddProfileEmail(email_id, sender, recipient, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProfileEmail", reflect.TypeOf((*MockEmailRepository)(nil).AddProfileEmail), email_id, sender, recipient, ctx)
}

// AddProfileEmailMyself mocks base method.
func (m *MockEmailRepository) AddProfileEmailMyself(email_id uint64, login string, ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProfileEmailMyself", email_id, login, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProfileEmailMyself indicates an expected call of AddProfileEmailMyself.
func (mr *MockEmailRepositoryMockRecorder) AddProfileEmailMyself(email_id, login, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProfileEmailMyself", reflect.TypeOf((*MockEmailRepository)(nil).AddProfileEmailMyself), email_id, login, ctx)
}

// Delete mocks base method.
func (m *MockEmailRepository) Delete(id uint64, login string, ctx context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id, login, ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockEmailRepositoryMockRecorder) Delete(id, login, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEmailRepository)(nil).Delete), id, login, ctx)
}

// DeleteFileByID mocks base method.
func (m *MockEmailRepository) DeleteFileByID(fileID uint64, ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFileByID", fileID, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFileByID indicates an expected call of DeleteFileByID.
func (mr *MockEmailRepositoryMockRecorder) DeleteFileByID(fileID, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFileByID", reflect.TypeOf((*MockEmailRepository)(nil).DeleteFileByID), fileID, ctx)
}

// FindEmail mocks base method.
func (m *MockEmailRepository) FindEmail(login string, ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEmail", login, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindEmail indicates an expected call of FindEmail.
func (mr *MockEmailRepositoryMockRecorder) FindEmail(login, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEmail", reflect.TypeOf((*MockEmailRepository)(nil).FindEmail), login, ctx)
}

// GetAllDraft mocks base method.
func (m *MockEmailRepository) GetAllDraft(login string, offset, limit int64, ctx context.Context) ([]*domain_models.Email, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDraft", login, offset, limit, ctx)
	ret0, _ := ret[0].([]*domain_models.Email)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDraft indicates an expected call of GetAllDraft.
func (mr *MockEmailRepositoryMockRecorder) GetAllDraft(login, offset, limit, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDraft", reflect.TypeOf((*MockEmailRepository)(nil).GetAllDraft), login, offset, limit, ctx)
}

// GetAllIncoming mocks base method.
func (m *MockEmailRepository) GetAllIncoming(login string, offset, limit int64, ctx context.Context) ([]*domain_models.Email, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllIncoming", login, offset, limit, ctx)
	ret0, _ := ret[0].([]*domain_models.Email)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllIncoming indicates an expected call of GetAllIncoming.
func (mr *MockEmailRepositoryMockRecorder) GetAllIncoming(login, offset, limit, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllIncoming", reflect.TypeOf((*MockEmailRepository)(nil).GetAllIncoming), login, offset, limit, ctx)
}

// GetAllSent mocks base method.
func (m *MockEmailRepository) GetAllSent(login string, offset, limit int64, ctx context.Context) ([]*domain_models.Email, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSent", login, offset, limit, ctx)
	ret0, _ := ret[0].([]*domain_models.Email)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSent indicates an expected call of GetAllSent.
func (mr *MockEmailRepositoryMockRecorder) GetAllSent(login, offset, limit, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSent", reflect.TypeOf((*MockEmailRepository)(nil).GetAllSent), login, offset, limit, ctx)
}

// GetAllSpam mocks base method.
func (m *MockEmailRepository) GetAllSpam(login string, offset, limit int64, ctx context.Context) ([]*domain_models.Email, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSpam", login, offset, limit, ctx)
	ret0, _ := ret[0].([]*domain_models.Email)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSpam indicates an expected call of GetAllSpam.
func (mr *MockEmailRepositoryMockRecorder) GetAllSpam(login, offset, limit, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSpam", reflect.TypeOf((*MockEmailRepository)(nil).GetAllSpam), login, offset, limit, ctx)
}

// GetAvatarFileIDByLogin mocks base method.
func (m *MockEmailRepository) GetAvatarFileIDByLogin(login string, ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvatarFileIDByLogin", login, ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvatarFileIDByLogin indicates an expected call of GetAvatarFileIDByLogin.
func (mr *MockEmailRepositoryMockRecorder) GetAvatarFileIDByLogin(login, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvatarFileIDByLogin", reflect.TypeOf((*MockEmailRepository)(nil).GetAvatarFileIDByLogin), login, ctx)
}

// GetByID mocks base method.
func (m *MockEmailRepository) GetByID(id uint64, login string, ctx context.Context) (*domain_models.Email, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id, login, ctx)
	ret0, _ := ret[0].(*domain_models.Email)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockEmailRepositoryMockRecorder) GetByID(id, login, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockEmailRepository)(nil).GetByID), id, login, ctx)
}

// GetFileByID mocks base method.
func (m *MockEmailRepository) GetFileByID(id uint64, ctx context.Context) (*domain_models.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileByID", id, ctx)
	ret0, _ := ret[0].(*domain_models.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileByID indicates an expected call of GetFileByID.
func (mr *MockEmailRepositoryMockRecorder) GetFileByID(id, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileByID", reflect.TypeOf((*MockEmailRepository)(nil).GetFileByID), id, ctx)
}

// GetFilesByEmailID mocks base method.
func (m *MockEmailRepository) GetFilesByEmailID(emailID uint64, ctx context.Context) ([]*domain_models.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilesByEmailID", emailID, ctx)
	ret0, _ := ret[0].([]*domain_models.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilesByEmailID indicates an expected call of GetFilesByEmailID.
func (mr *MockEmailRepositoryMockRecorder) GetFilesByEmailID(emailID, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilesByEmailID", reflect.TypeOf((*MockEmailRepository)(nil).GetFilesByEmailID), emailID, ctx)
}

// Update mocks base method.
func (m *MockEmailRepository) Update(newEmail *domain_models.Email, ctx context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", newEmail, ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockEmailRepositoryMockRecorder) Update(newEmail, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEmailRepository)(nil).Update), newEmail, ctx)
}

// UpdateFileByID mocks base method.
func (m *MockEmailRepository) UpdateFileByID(fileID uint64, newFileID, newFileType, newFileName, newFileSize string, ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFileByID", fileID, newFileID, newFileType, newFileName, newFileSize, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFileByID indicates an expected call of UpdateFileByID.
func (mr *MockEmailRepositoryMockRecorder) UpdateFileByID(fileID, newFileID, newFileType, newFileName, newFileSize, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFileByID", reflect.TypeOf((*MockEmailRepository)(nil).UpdateFileByID), fileID, newFileID, newFileType, newFileName, newFileSize, ctx)
}