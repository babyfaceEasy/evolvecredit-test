package unit

import (
	"evolvecredit-test/internal/user"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var mockUserRepo *MockUserRepository
var userService user.UserService

type UserServiceUnitTestSuite struct {
	suite.Suite
}

func TestUserServiceUnit(t *testing.T) {
	suite.Run(t, new(UserServiceUnitTestSuite))
}

func (s *UserServiceUnitTestSuite) SetupTest() {
	mockUserRepo = new(MockUserRepository)
	//userService, _ = user.NewUserService(mockUserRepo)
}

// mocks
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Size() (int, error) {
	args := m.Called()

	return args.Int(0), nil
}

func (m *MockUserRepository) DeleteAll() (bool, error) {
	args := m.Called()

	return args.Bool(0), nil
}

func (m *MockUserRepository) Save(user *user.User) (bool, error) {
	args := m.Called(user)

	return args.Bool(0), nil
}
