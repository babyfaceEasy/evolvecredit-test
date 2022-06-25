package user

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/uptrace/bun"
)

var users = make([]*User, 0)

type UserRepo interface {
	Size() (int, error)
	DeleteAll() (bool, error)
	Save(user *User) (bool, error)
	Get() ([]*User, error)
	GetByEmail(email string) (*User, error)
	SearchBtwDates(startDate string, endDate string) ([]*User, error)
	SearchByEmailAndBtwDates(email string, startDate, endDate string) ([]*User, error)
}

type UserRepository struct {
	ErrEmptyResult error
	db             bun.IDB
}

func NewUserRepository(db bun.IDB) (*UserRepository, error) {
	return &UserRepository{db: db, ErrEmptyResult: errors.New("row(s) not found")}, nil
	/*
		userFakeData := User{}
		faker.FakeData(&userFakeData)
		users = append(users, &userFakeData)

		faker.FakeData(&userFakeData)
		users = append(users, &userFakeData)

		faker.FakeData(&userFakeData)
		users = append(users, &userFakeData)

		faker.FakeData(&userFakeData)
		users = append(users, &userFakeData)
	*/

	/*
		var id []int
		var userData *User
		for i := 0; i < 10; i++ {
			id, _ = faker.RandomInt(1, 200, 1)
			userData = NewUser(
			id[0],
			faker.Email(),
			faker.FirstName(),
			faker.LastName(),
			time.Now())

			users = append(users, userData)
		}

		return &UserRepository{}, nil
	*/

}

func (r *UserRepository) Size() (int, error) {
	return len(users), nil
}

func (r *UserRepository) DeleteAll() (bool, error) {
	users = make([]*User, 0)
	return true, nil
}

func (r *UserRepository) Save(user *User) (bool, error) {
	users = append(users, user)
	return true, nil
}

func (r *UserRepository) GetCount() (int, error) {
	ctx := context.Background()

	count, err := r.db.NewSelect().Model((*User)(nil)).Count(ctx)
	if err != nil {
		log.Printf("error occurred while trying to get all users count: %v", err)
		return 0, err
	}

	return count, nil
}

func (r *UserRepository) SearchBtwDatesCount(startDate time.Time, endDate time.Time, page, limit int) (int, error) {
	ctx := context.Background()

	count, err := r.db.NewSelect().Model((*User)(nil)).Where("dob >= ? and dob <= ?", startDate, endDate).Count(ctx)
	if err != nil {
		log.Printf("error occurred while trying to get all users btw dates count: %v", err)
		return 0, err
	}
	return count, nil
}

func (r *UserRepository) Get(page int, limit int) ([]*User, error) {

	var users []*User
	var userModels []UserModel

	log.Println((page - 1) * limit)

	ctx := context.Background()
	_, err := r.db.NewSelect().
		Model(&userModels).
		Order("id ASC").
		Limit(limit).
		Offset((page - 1) * limit).
		ScanAndCount(ctx)
	if err != nil {
		log.Printf("error occurred while trying to get all users: %v", err)
		return nil, err
	}

	for _, userModel := range userModels {
		user := NewUser(
			userModel.ID,
			userModel.Email,
			userModel.FirstName,
			userModel.LastName,
			userModel.DOB)

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) GetByEmail(email string) (*User, error) {
	userModel := new(UserModel)
	ctx := context.Background()

	err := r.db.NewSelect().Model(userModel).Where("email = ?", email).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, r.ErrEmptyResult
		}
		log.Printf("error occurred while trying to search by email: %v", err)
		return nil, err
	}

	user := NewUser(
		userModel.ID,
		userModel.Email,
		userModel.FirstName,
		userModel.LastName,
		userModel.DOB)
	return user, nil
}

func (r *UserRepository) SearchBtwDates(startDate time.Time, endDate time.Time, page, limit int) ([]*User, error) {
	var userModels []UserModel
	ctx := context.Background()

	_, err := r.db.NewSelect().
		Model(&userModels).
		Where("dob >= ? and dob <= ?", startDate, endDate).
		Order("id ASC").
		Limit(limit).
		Offset((page - 1) * limit).
		ScanAndCount(ctx)
	if err != nil {
		log.Printf("error occurred while trying to get all users btw dates: %v", err)
		return nil, err
	}

	users := make([]*User, 0)
	for _, userModel := range userModels {
		user := NewUser(
			userModel.ID,
			userModel.Email,
			userModel.FirstName,
			userModel.LastName,
			userModel.DOB)

		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) SearchByEmailAndBtwDates(email string, startDate, endDate string) ([]*User, error) {
	return nil, nil
}
