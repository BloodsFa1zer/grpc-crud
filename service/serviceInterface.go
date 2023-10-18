package service

import "app4/database"

type ServiceInterface interface {
	Get(ID int64) (*database.User, error)
	Create(user database.User) (int64, error)
	//UpdateUser(ID int64, user database.User) (int64, error)
	//FindUsers() (*[]database.User, error)
	//DeleteUserByID(ID int64) error
	//	FindByNicknameToGetUserPassword(nickname string) (*User, error)
}
