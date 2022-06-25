package user

import (
	"errors"
	"time"
)

type UserService struct {
	ErrDataNotFound error
	userRepository  *UserRepository
}

func NewUserService(userRepository *UserRepository) (*UserService, error) {
	return &UserService{userRepository: userRepository, ErrDataNotFound: errors.New("data not found")}, nil
}

func (s *UserService) List(page int, limit int) ([]*User, error) {
	users, err := s.userRepository.Get(page, limit)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) ListTotalCount() (int, error) {
	count, err := s.userRepository.GetCount()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *UserService) GetBtwDatesTotalCount(startDate time.Time, endDate time.Time, page, limit int) (int, error) {
	count, err := s.userRepository.SearchBtwDatesCount(startDate, endDate, page, limit)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *UserService) GetByEmail(email string) (*User, error) {
	user, err := s.userRepository.GetByEmail(email)
	if err != nil {
		if err == s.userRepository.ErrEmptyResult {
			return nil, s.ErrDataNotFound
		}
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetBtwDates(startDate time.Time, endDate time.Time, page, limit int) ([]*User, error) {
	users, err := s.userRepository.SearchBtwDates(startDate, endDate, page, limit)
	if err != nil {
		return nil, err
	}

	return users, err
}
