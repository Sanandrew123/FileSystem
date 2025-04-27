package repository

import (
	"errors"
	"sync"

	"github.com/Sanandrew123/FileSystem/internal/model"
)

var (
	users      = make(map[string]*model.User)
	usersMutex sync.RWMutex
	userIDSeq  int64 = 1
)

func CreateUser(username, password string) (*model.User, error) {
	usersMutex.Lock()
	defer usersMutex.Unlock()

	if _, exists := users[username]; exists {
		return nil, ErrUserExists
	}
	user := &model.User{
		ID:       userIDSeq,
		Username: username,
		Password: password,
	}
	users[username] = user
	userIDSeq++
	return user, nil
}

func GetUserByUsername(username string) (*model.User, error) {
	usersMutex.RLock()
	defer usersMutex.RLock()

	user, exists := users[username]
	if !exists {
		return nil, ErrUserNotFound
	}
	return user, nil
}

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
)
