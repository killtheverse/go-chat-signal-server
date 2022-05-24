package services

import (
	"github.com/killtheverse/go-chat-signal-server/internal/core/ports"
)

type UserService struct {
    userRepository ports.UserRepository  
}

func NewUserService(repository ports.UserRepository) *UserService {
    return &UserService{
        userRepository: repository,
    }
}

func (s *UserService) Login(username string, password string) error {
    err := s.userRepository.Login(username, password)
    if err != nil {
        return err
    }
    return nil
}

func (s *UserService) Register(username string, password string, name string) error {
    err := s.userRepository.Register(username, password, name)
    if err != nil {
        return err
    }
    return nil
}
