package service

import (
	"errors"
	"fmt"
	"test/custom_errors"
	"test/database/model"
	"test/interface/auth"
	"test/interface/cache"
	"test/interface/hasher"
	"test/interface/repository"
	"test/interface/service"

	"github.com/jinzhu/gorm"
)

type userService struct {
	userRepository repository.UserRepository
	hasher         hasher.Hasher
	cacheClient    cache.CacheClient
	jwtWrapper     auth.JwtWrapper
}

var blackListCacheKey = "BLACKLIST:"

func (s *userService) Authenticate(username, password string) (*service.AuthorizationResponse, error) {
	if username == "" {
		return nil, errors.New("username argument should not be empty")
	}

	if password == "" {
		return nil, errors.New("password argument should not be empty")
	}

	user, err := s.userRepository.GetByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, custom_errors.NewUnauthorizedError("user not found")
		}

		return nil, err
	}

	if user == nil {
		return nil, errors.New(fmt.Sprint("user not found by username", username))
	}

	hash, err := s.hasher.Hash(password)
	if err != nil {
		return nil, err
	}

	if hash != user.PasswordHash {
		return nil, custom_errors.NewUnauthorizedError("not valid username or password passed")
	}

	token, err := s.jwtWrapper.GenerateToken(user.Username, user.ID)
	if err != nil {
		return nil, err
	}

	resp := &service.AuthorizationResponse{
		Token:     token,
		TokenType: "Bearer",
	}

	return resp, nil
}

func (s *userService) Register(username, password string) (*model.User, error) {
	if username == "" {
		return nil, errors.New("username argument should not be empty")
	}

	if password == "" {
		return nil, errors.New("password argument should not be empty")
	}

	hash, err := s.hasher.Hash(password)
	if err != nil {
		return nil, err
	}

	userFromDB, err := s.userRepository.GetByUsername(username)
	if err != nil {
		return nil, err
	} else if userFromDB != nil {
		return nil, custom_errors.NewEntityExistsError("user is already exists, cant register")
	}

	user := &model.User{Username: username, PasswordHash: hash}
	err = s.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (s *userService) ValidateToken(token string) (*auth.JwtClaim, error) {
	if token == "" {
		return nil, errors.New("token argument must not be empty")
	}

	blackListKey := fmt.Sprintf("%s%s", blackListCacheKey, token)
	_, err := s.cacheClient.GetString(blackListKey)
	// Если токен найден в черном списке, то токен невалидный
	if err == nil {
		return nil, custom_errors.NewUnauthorizedError("token is in blacklist")
	} else {
		// если это не ошибка отсутствия ключа в кеше, то это какая-та другая ошибка
		if _, ok := err.(*custom_errors.KeyNotFoundError); !ok {
			return nil, err
		}
	}

	claim, err := s.jwtWrapper.ValidateToken(token)
	if err != nil {
		return nil, custom_errors.NewUnauthorizedError(err.Error())
	}

	return claim, nil
}

func (s *userService) Logout(token string) error {
	if token == "" {
		return errors.New("token argument must not be empty")
	}

	blacklistKey := fmt.Sprintf("%s%s", blackListCacheKey, token)
	err := s.cacheClient.SetString(blacklistKey, "")
	return err
}

func (s *userService) Get(userID uint) (*model.User, error) {
	user, err := s.userRepository.Get(userID)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, custom_errors.NewEntityNotFoundError(fmt.Sprintf("user %d not found", userID))
	}

	return user, nil
}

func NewUserService(userRepository repository.UserRepository, hasher hasher.Hasher, jwtWrapper auth.JwtWrapper, cacheClient cache.CacheClient) *userService {
	if userRepository == nil {
		panic("userRepository must not be nil")
	}

	if hasher == nil {
		panic("hasher must not be nil")
	}

	if jwtWrapper == nil {
		panic("jwtWrapper must not be nil")
	}

	if cacheClient == nil {
		panic("cacheClient must not be nil")
	}

	return &userService{
		userRepository: userRepository,
		hasher:         hasher,
		jwtWrapper:     jwtWrapper,
		cacheClient:    cacheClient,
	}
}
