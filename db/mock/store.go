// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gufir/money-management/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	db "github.com/gufir/money-management/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateCategories mocks base method.
func (m *MockStore) CreateCategories(arg0 context.Context, arg1 db.CreateCategoriesParams) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategories", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCategories indicates an expected call of CreateCategories.
func (mr *MockStoreMockRecorder) CreateCategories(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategories", reflect.TypeOf((*MockStore)(nil).CreateCategories), arg0, arg1)
}

// CreateMonthlyReport mocks base method.
func (m *MockStore) CreateMonthlyReport(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMonthlyReport", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMonthlyReport indicates an expected call of CreateMonthlyReport.
func (mr *MockStoreMockRecorder) CreateMonthlyReport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMonthlyReport", reflect.TypeOf((*MockStore)(nil).CreateMonthlyReport), arg0, arg1)
}

// CreateReportUser mocks base method.
func (m *MockStore) CreateReportUser(arg0 context.Context, arg1 db.CreateReportUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReportUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateReportUser indicates an expected call of CreateReportUser.
func (mr *MockStoreMockRecorder) CreateReportUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReportUser", reflect.TypeOf((*MockStore)(nil).CreateReportUser), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 context.Context, arg1 db.CreateSessionParams) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// CreateTransaction mocks base method.
func (m *MockStore) CreateTransaction(arg0 context.Context, arg1 db.CreateTransactionParams) (db.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", arg0, arg1)
	ret0, _ := ret[0].(db.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockStoreMockRecorder) CreateTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockStore)(nil).CreateTransaction), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// CreateUserTx mocks base method.
func (m *MockStore) CreateUserTx(arg0 context.Context, arg1 db.CreateUserTxParams) (db.CreateUserTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserTx", arg0, arg1)
	ret0, _ := ret[0].(db.CreateUserTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserTx indicates an expected call of CreateUserTx.
func (mr *MockStoreMockRecorder) CreateUserTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserTx", reflect.TypeOf((*MockStore)(nil).CreateUserTx), arg0, arg1)
}

// GetAllCategories mocks base method.
func (m *MockStore) GetAllCategories(arg0 context.Context) ([]db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCategories", arg0)
	ret0, _ := ret[0].([]db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCategories indicates an expected call of GetAllCategories.
func (mr *MockStoreMockRecorder) GetAllCategories(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCategories", reflect.TypeOf((*MockStore)(nil).GetAllCategories), arg0)
}

// GetCategoryById mocks base method.
func (m *MockStore) GetCategoryById(arg0 context.Context, arg1 uuid.UUID) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryById", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoryById indicates an expected call of GetCategoryById.
func (mr *MockStoreMockRecorder) GetCategoryById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryById", reflect.TypeOf((*MockStore)(nil).GetCategoryById), arg0, arg1)
}

// GetCategoryByName mocks base method.
func (m *MockStore) GetCategoryByName(arg0 context.Context, arg1 string) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryByName", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoryByName indicates an expected call of GetCategoryByName.
func (mr *MockStoreMockRecorder) GetCategoryByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryByName", reflect.TypeOf((*MockStore)(nil).GetCategoryByName), arg0, arg1)
}

// GetDetailsReportByUser mocks base method.
func (m *MockStore) GetDetailsReportByUser(arg0 context.Context, arg1 uuid.UUID) ([]db.GetDetailsReportByUserRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetailsReportByUser", arg0, arg1)
	ret0, _ := ret[0].([]db.GetDetailsReportByUserRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetailsReportByUser indicates an expected call of GetDetailsReportByUser.
func (mr *MockStoreMockRecorder) GetDetailsReportByUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetailsReportByUser", reflect.TypeOf((*MockStore)(nil).GetDetailsReportByUser), arg0, arg1)
}

// GetReportByCategory mocks base method.
func (m *MockStore) GetReportByCategory(arg0 context.Context) ([]db.GetReportByCategoryRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReportByCategory", arg0)
	ret0, _ := ret[0].([]db.GetReportByCategoryRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReportByCategory indicates an expected call of GetReportByCategory.
func (mr *MockStoreMockRecorder) GetReportByCategory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReportByCategory", reflect.TypeOf((*MockStore)(nil).GetReportByCategory), arg0)
}

// GetReportByDate mocks base method.
func (m *MockStore) GetReportByDate(arg0 context.Context, arg1 db.GetReportByDateParams) ([]db.GetReportByDateRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReportByDate", arg0, arg1)
	ret0, _ := ret[0].([]db.GetReportByDateRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReportByDate indicates an expected call of GetReportByDate.
func (mr *MockStoreMockRecorder) GetReportByDate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReportByDate", reflect.TypeOf((*MockStore)(nil).GetReportByDate), arg0, arg1)
}

// GetReportByUser mocks base method.
func (m *MockStore) GetReportByUser(arg0 context.Context, arg1 uuid.UUID) (db.GetReportByUserRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReportByUser", arg0, arg1)
	ret0, _ := ret[0].(db.GetReportByUserRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReportByUser indicates an expected call of GetReportByUser.
func (mr *MockStoreMockRecorder) GetReportByUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReportByUser", reflect.TypeOf((*MockStore)(nil).GetReportByUser), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(arg0 context.Context, arg1 uuid.UUID) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetTransaction mocks base method.
func (m *MockStore) GetTransaction(arg0 context.Context, arg1 uuid.UUID) (db.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransaction", arg0, arg1)
	ret0, _ := ret[0].(db.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransaction indicates an expected call of GetTransaction.
func (mr *MockStoreMockRecorder) GetTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockStore)(nil).GetTransaction), arg0, arg1)
}

// GetTransactionByType mocks base method.
func (m *MockStore) GetTransactionByType(arg0 context.Context, arg1 db.GetTransactionByTypeParams) ([]db.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByType", arg0, arg1)
	ret0, _ := ret[0].([]db.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionByType indicates an expected call of GetTransactionByType.
func (mr *MockStoreMockRecorder) GetTransactionByType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByType", reflect.TypeOf((*MockStore)(nil).GetTransactionByType), arg0, arg1)
}

// GetTransactionByuserId mocks base method.
func (m *MockStore) GetTransactionByuserId(arg0 context.Context, arg1 uuid.UUID) ([]db.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByuserId", arg0, arg1)
	ret0, _ := ret[0].([]db.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionByuserId indicates an expected call of GetTransactionByuserId.
func (mr *MockStoreMockRecorder) GetTransactionByuserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByuserId", reflect.TypeOf((*MockStore)(nil).GetTransactionByuserId), arg0, arg1)
}

// GetUserByEmail mocks base method.
func (m *MockStore) GetUserByEmail(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockStoreMockRecorder) GetUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockStore)(nil).GetUserByEmail), arg0, arg1)
}

// GetUserById mocks base method.
func (m *MockStore) GetUserById(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockStoreMockRecorder) GetUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockStore)(nil).GetUserById), arg0, arg1)
}

// GetUserByUsername mocks base method.
func (m *MockStore) GetUserByUsername(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockStoreMockRecorder) GetUserByUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockStore)(nil).GetUserByUsername), arg0, arg1)
}

// UpdateCategories mocks base method.
func (m *MockStore) UpdateCategories(arg0 context.Context, arg1 db.UpdateCategoriesParams) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategories", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCategories indicates an expected call of UpdateCategories.
func (mr *MockStoreMockRecorder) UpdateCategories(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategories", reflect.TypeOf((*MockStore)(nil).UpdateCategories), arg0, arg1)
}

// UpdateTransaction mocks base method.
func (m *MockStore) UpdateTransaction(arg0 context.Context, arg1 db.UpdateTransactionParams) (db.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransaction", arg0, arg1)
	ret0, _ := ret[0].(db.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTransaction indicates an expected call of UpdateTransaction.
func (mr *MockStoreMockRecorder) UpdateTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransaction", reflect.TypeOf((*MockStore)(nil).UpdateTransaction), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}
