package service

import "app4/database"

type UserServiceInterface interface {
	Get(ID int64) (*database.User, error)
	Create(user database.User) (int64, error)
	Update(ID int64, user database.User) (int64, error)
	Delete(ID int64) error
	//FindUsers() (*[]database.User, error)
}
