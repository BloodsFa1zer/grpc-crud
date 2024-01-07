package service

import "app4/database"

type UserServiceInterface interface {
	Create(user database.User) (int64, error)
	Read(ID int64) (*database.User, error)
	Update(ID int64, user database.User) (int64, error)
	Delete(ID int64) error
}
