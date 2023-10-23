package service

import (
	"app4/database"
	_ "app4/proto"
	"database/sql"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	UserDB database.DbInterface
}

func NewUserService(UserDB database.DbInterface) *UserService {
	return &UserService{UserDB: UserDB}
}

// TODO: VALIDATION ON USER CREATION!!!!!!!!
func (us *UserService) Create(user database.User) (int64, error) {

	userID, err := us.UserDB.InsertUser(user)
	if err != nil {
		if err == errors.New("such nickName already exists") {
			return 0, status.Errorf(codes.AlreadyExists, err.Error())
		}
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

func (us *UserService) Update(userID int64, user database.User) (int64, error) {

	updatedUserID, err := us.UserDB.UpdateUser(userID, user)
	if err == sql.ErrNoRows {
		return 0, status.Errorf(codes.NotFound, "there is no user with that id ")
	} else if err == errors.New("such nickName already exists") {
		return 0, status.Errorf(codes.AlreadyExists, err.Error())
	} else if err != nil {
		return 0, status.Errorf(codes.DataLoss, "can`t update user")
	}

	//return updatedUserID, status.Errorf(codes.OK, "user successfully updated")
	return updatedUserID, nil
}

func (us *UserService) Delete(userID int64) error {

	err := us.UserDB.DeleteUserByID(userID)
	if err == sql.ErrNoRows {
		return status.Errorf(codes.NotFound, "there is no user with that ID")
	} else if err != nil {
		return status.Errorf(codes.DataLoss, "can`t delete user")
	}

	return nil
}
