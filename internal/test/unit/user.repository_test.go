package unit

import (
	"evolvecredit-test/internal/user"
	"log"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/suite"
)


type UserRepositoryUnitTestSuite struct {
	suite.Suite
	userRepository *user.UserRepository
}

func TestUserRepositoryUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryUnitTestSuite))
}

func (s *UserRepositoryUnitTestSuite) SetupTest() {
	ur, err := user.NewUserRepository()
	if err != nil {
		log.Fatal("Error occurred in user repository setup")
	}
	s.userRepository = ur
	s.userRepository.DeleteAll() 
}

func (s *UserRepositoryUnitTestSuite) TestPlaceholder() {
	s.True(true)
}

func (s *UserRepositoryUnitTestSuite) TestUserRepositoryCreation() {
	userRepository, err := user.NewUserRepository()

	s.NoError(err)
	s.Empty(userRepository)
}

func (s *UserRepositoryUnitTestSuite) TestCountOneUser() {
	user :=  user.User{}
	faker.FakeData(&user)

	s.userRepository.Save(&user)

	expected, err := s.userRepository.Size()

	s.NoError(err)
	s.Equal(1, expected)
}

func (s *UserRepositoryUnitTestSuite) TestCountMoreThanOneUser() {
	user := user.User{}
	faker.FakeData(&user)

	s.userRepository.Save(&user)

	faker.FakeData(&user)
	s.userRepository.Save(&user)

	expected, err := s.userRepository.Size()

	s.NoError(err)
	s.Equal(2, expected)
}

func (s *UserRepositoryUnitTestSuite) TestSaveNewUser() {
	userRepository, _ := user.NewUserRepository()
	user := user.User{}
	_  = faker.FakeData(&user) 

	response, err := userRepository.Save(&user)

	s.NoError(err)
	s.True(response)
}

func (s *UserRepositoryUnitTestSuite) TestDeleteAllUsers() {
	userRepository, _ := user.NewUserRepository()

	response, err := userRepository.DeleteAll()

	s.NoError(err)
	s.True(response)
}
