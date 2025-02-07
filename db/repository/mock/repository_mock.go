// Code generated by MockGen. DO NOT EDIT.
// Source: db/repository/repository.go

// Package mockrepo is a generated GoMock package.
package mockrepo

import (
	context "context"
	reflect "reflect"

	querier "github.com/gadhittana-01/book-go/db/repository"
	utils "github.com/gadhittana01/go-modules/utils"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	pgx "github.com/jackc/pgx/v5"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CheckBookExists mocks base method.
func (m *MockRepository) CheckBookExists(ctx context.Context, id uuid.UUID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckBookExists", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckBookExists indicates an expected call of CheckBookExists.
func (mr *MockRepositoryMockRecorder) CheckBookExists(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckBookExists", reflect.TypeOf((*MockRepository)(nil).CheckBookExists), ctx, id)
}

// CheckEmailExists mocks base method.
func (m *MockRepository) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckEmailExists", ctx, email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckEmailExists indicates an expected call of CheckEmailExists.
func (mr *MockRepositoryMockRecorder) CheckEmailExists(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckEmailExists", reflect.TypeOf((*MockRepository)(nil).CheckEmailExists), ctx, email)
}

// CheckOrderExists mocks base method.
func (m *MockRepository) CheckOrderExists(ctx context.Context, id uuid.UUID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckOrderExists", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckOrderExists indicates an expected call of CheckOrderExists.
func (mr *MockRepositoryMockRecorder) CheckOrderExists(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckOrderExists", reflect.TypeOf((*MockRepository)(nil).CheckOrderExists), ctx, id)
}

// CreateBook mocks base method.
func (m *MockRepository) CreateBook(ctx context.Context, arg querier.CreateBookParams) (querier.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", ctx, arg)
	ret0, _ := ret[0].(querier.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockRepositoryMockRecorder) CreateBook(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockRepository)(nil).CreateBook), ctx, arg)
}

// CreateOrder mocks base method.
func (m *MockRepository) CreateOrder(ctx context.Context, arg querier.CreateOrderParams) (querier.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", ctx, arg)
	ret0, _ := ret[0].(querier.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockRepositoryMockRecorder) CreateOrder(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockRepository)(nil).CreateOrder), ctx, arg)
}

// CreateOrderDetail mocks base method.
func (m *MockRepository) CreateOrderDetail(ctx context.Context, arg querier.CreateOrderDetailParams) (querier.OrderDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderDetail", ctx, arg)
	ret0, _ := ret[0].(querier.OrderDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrderDetail indicates an expected call of CreateOrderDetail.
func (mr *MockRepositoryMockRecorder) CreateOrderDetail(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderDetail", reflect.TypeOf((*MockRepository)(nil).CreateOrderDetail), ctx, arg)
}

// CreateUser mocks base method.
func (m *MockRepository) CreateUser(ctx context.Context, arg querier.CreateUserParams) (querier.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, arg)
	ret0, _ := ret[0].(querier.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockRepositoryMockRecorder) CreateUser(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockRepository)(nil).CreateUser), ctx, arg)
}

// FindBook mocks base method.
func (m *MockRepository) FindBook(ctx context.Context, arg querier.FindBookParams) ([]querier.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBook", ctx, arg)
	ret0, _ := ret[0].([]querier.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBook indicates an expected call of FindBook.
func (mr *MockRepositoryMockRecorder) FindBook(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBook", reflect.TypeOf((*MockRepository)(nil).FindBook), ctx, arg)
}

// FindBookByID mocks base method.
func (m *MockRepository) FindBookByID(ctx context.Context, id uuid.UUID) (querier.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBookByID", ctx, id)
	ret0, _ := ret[0].(querier.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBookByID indicates an expected call of FindBookByID.
func (mr *MockRepositoryMockRecorder) FindBookByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBookByID", reflect.TypeOf((*MockRepository)(nil).FindBookByID), ctx, id)
}

// FindOrderByID mocks base method.
func (m *MockRepository) FindOrderByID(ctx context.Context, arg querier.FindOrderByIDParams) (querier.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrderByID", ctx, arg)
	ret0, _ := ret[0].(querier.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrderByID indicates an expected call of FindOrderByID.
func (mr *MockRepositoryMockRecorder) FindOrderByID(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrderByID", reflect.TypeOf((*MockRepository)(nil).FindOrderByID), ctx, arg)
}

// FindOrderByUserID mocks base method.
func (m *MockRepository) FindOrderByUserID(ctx context.Context, arg querier.FindOrderByUserIDParams) ([]querier.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrderByUserID", ctx, arg)
	ret0, _ := ret[0].([]querier.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrderByUserID indicates an expected call of FindOrderByUserID.
func (mr *MockRepositoryMockRecorder) FindOrderByUserID(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrderByUserID", reflect.TypeOf((*MockRepository)(nil).FindOrderByUserID), ctx, arg)
}

// FindOrderDetailByOrderID mocks base method.
func (m *MockRepository) FindOrderDetailByOrderID(ctx context.Context, arg querier.FindOrderDetailByOrderIDParams) ([]querier.FindOrderDetailByOrderIDRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrderDetailByOrderID", ctx, arg)
	ret0, _ := ret[0].([]querier.FindOrderDetailByOrderIDRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrderDetailByOrderID indicates an expected call of FindOrderDetailByOrderID.
func (mr *MockRepositoryMockRecorder) FindOrderDetailByOrderID(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrderDetailByOrderID", reflect.TypeOf((*MockRepository)(nil).FindOrderDetailByOrderID), ctx, arg)
}

// FindUserByEmail mocks base method.
func (m *MockRepository) FindUserByEmail(ctx context.Context, email string) (querier.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", ctx, email)
	ret0, _ := ret[0].(querier.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockRepositoryMockRecorder) FindUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockRepository)(nil).FindUserByEmail), ctx, email)
}

// GetBookCount mocks base method.
func (m *MockRepository) GetBookCount(ctx context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookCount", ctx)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookCount indicates an expected call of GetBookCount.
func (mr *MockRepositoryMockRecorder) GetBookCount(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookCount", reflect.TypeOf((*MockRepository)(nil).GetBookCount), ctx)
}

// GetBookPurchasedByUserID mocks base method.
func (m *MockRepository) GetBookPurchasedByUserID(ctx context.Context, userID uuid.UUID) ([]querier.GetBookPurchasedByUserIDRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookPurchasedByUserID", ctx, userID)
	ret0, _ := ret[0].([]querier.GetBookPurchasedByUserIDRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookPurchasedByUserID indicates an expected call of GetBookPurchasedByUserID.
func (mr *MockRepositoryMockRecorder) GetBookPurchasedByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookPurchasedByUserID", reflect.TypeOf((*MockRepository)(nil).GetBookPurchasedByUserID), ctx, userID)
}

// GetDB mocks base method.
func (m *MockRepository) GetDB() utils.PGXPool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDB")
	ret0, _ := ret[0].(utils.PGXPool)
	return ret0
}

// GetDB indicates an expected call of GetDB.
func (mr *MockRepositoryMockRecorder) GetDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDB", reflect.TypeOf((*MockRepository)(nil).GetDB))
}

// GetOrderCountByUserId mocks base method.
func (m *MockRepository) GetOrderCountByUserId(ctx context.Context, userID uuid.UUID) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderCountByUserId", ctx, userID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderCountByUserId indicates an expected call of GetOrderCountByUserId.
func (mr *MockRepositoryMockRecorder) GetOrderCountByUserId(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderCountByUserId", reflect.TypeOf((*MockRepository)(nil).GetOrderCountByUserId), ctx, userID)
}

// UpdateOrderByID mocks base method.
func (m *MockRepository) UpdateOrderByID(ctx context.Context, arg querier.UpdateOrderByIDParams) (querier.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderByID", ctx, arg)
	ret0, _ := ret[0].(querier.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrderByID indicates an expected call of UpdateOrderByID.
func (mr *MockRepositoryMockRecorder) UpdateOrderByID(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderByID", reflect.TypeOf((*MockRepository)(nil).UpdateOrderByID), ctx, arg)
}

// WithTx mocks base method.
func (m *MockRepository) WithTx(tx pgx.Tx) querier.Querier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTx", tx)
	ret0, _ := ret[0].(querier.Querier)
	return ret0
}

// WithTx indicates an expected call of WithTx.
func (mr *MockRepositoryMockRecorder) WithTx(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTx", reflect.TypeOf((*MockRepository)(nil).WithTx), tx)
}
