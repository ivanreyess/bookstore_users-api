package service

import (
	"github.com/ivanreyess/bookstore_users-api/domain/user"
	"github.com/ivanreyess/bookstore_users-api/utils/cryptoutils"
	"github.com/ivanreyess/bookstore_users-api/utils/dateutils"
	"github.com/ivanreyess/bookstore_users-api/utils/errors"
)

var (
	//UserService holds user services functions
	UserService userServiceInterface = &userService{}
)

type userService struct {
}

type userServiceInterface interface {
	CreateUser(user.User) (*user.User, *errors.RestErr)
	GetUser(int64) (*user.User, *errors.RestErr)
	UpdateUser(bool, user.User) (*user.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (user.Users, *errors.RestErr)
}

//CreateUser creates a new user reference
func (s *userService) CreateUser(u user.User) (*user.User, *errors.RestErr) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	u.DateCreated = dateutils.GetNowDBFormat()
	u.Status = user.StatusActive
	u.Password = cryptoutils.GetMD5(u.Password)
	if err := u.Save(); err != nil {
		return nil, err
	}
	return &u, nil
}

//GetUser return an user given its ID
func (s *userService) GetUser(userID int64) (*user.User, *errors.RestErr) {
	dao := &user.User{ID: userID}
	if err := dao.Get(); err != nil {
		return nil, err
	}
	return dao, nil
}

//UpdateUser updates user
func (s *userService) UpdateUser(isPartial bool, u user.User) (*user.User, *errors.RestErr) {
	current, err := s.GetUser(u.ID)
	if err != nil {
		return nil, err
	}
	if err := current.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if u.FirstName != "" {
			current.FirstName = u.FirstName
		}
		if u.LastName != "" {
			current.LastName = u.LastName
		}
		if u.Email != "" {
			current.Email = u.Email
		}

	} else {
		current.FirstName = u.FirstName
		current.LastName = u.LastName
		current.Email = u.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

//DeleteUser remove user
func (s *userService) DeleteUser(userID int64) *errors.RestErr {
	user := &user.User{ID: userID}
	return user.Delete()
}

//Search retrieve users given the status
func (s *userService) SearchUser(status string) (user.Users, *errors.RestErr) {
	dao := &user.User{}
	return dao.FindByStatus(status)
}
