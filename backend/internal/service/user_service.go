package service

import (
	"github.com/Sanandrew123/FileSystem/internal/model"
	"github.com/Sanandrew123/FileSystem/internal/repository"
)

func Register(username, password string) (*model.User, error) {
	return repository.CreateUser(username, password)
}