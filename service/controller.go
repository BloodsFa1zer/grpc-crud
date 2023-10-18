package service

import (
	"app4/database"
	_ "app4/proto"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	UserDB database.DbInterface
}

func NewUserService(UserDB database.DbInterface) *UserService {
	return &UserService{UserDB: UserDB}
}

func (us *UserService) Create(user database.User) (int64, error) {

	userID, err := us.UserDB.InsertUser(user)
	if err != nil {
		return 0, status.Errorf(codes.Aborted, "can`t insert user")
	}

	return userID, nil
}

func (us *UserService) Get(ID int64) (*database.User, error) {

	if ID == 0 {
		return nil, status.Errorf(codes.OutOfRange, "id cannot be 0")
	}
	user, err := us.UserDB.FindByID(ID)
	if err == sql.ErrNoRows {
		return nil, status.Errorf(codes.NotFound, "there is no user with that id ")
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can`t find user")
	}

	return user, nil
}

//func (uc *UseCase) Update(updateUser models.User) error {
//	var user models.User
//	var err error
//	// check if user exists
//	if user, err = uc.Get(string(updateUser.ID)); err != nil {
//		return err
//	}
//
//	// check if only name is going to change,
//	// as the email cannot be changed
//	if user.Email != updateUser.Email {
//		return errors.New("email cannot be changed")
//	}
//
//	err = uc.repo.Update(updateUser)
//	if err != nil {
//		// handle the error properly as the error might be something worth to debug
//	}
//
//	return nil
//}
//
//func (uc *UseCase) Delete(id string) error {
//	var err error
//	// check if user exists
//	if _, err = uc.Get(id); err != nil {
//		return err
//	}
//
//	err = uc.repo.Delete(id)
//	if err != nil {
//		// handle the error as it might be something worth to debug
//		return err
//	}
//
//	return nil
//}
