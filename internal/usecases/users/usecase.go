package users

import (
	"errors"
	"time"

	"rakhimjonshokirov/movie-crud-app/internal/entities"
	"rakhimjonshokirov/movie-crud-app/pkg/auth"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	Create(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *zap.Logger
}

func NewUserUseCase(log *zap.Logger, repo UserRepo) *UserUseCase {
	return &UserUseCase{log: log, repo: repo}
}

func (uc *UserUseCase) Register(email, username, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entities.User{
		Email:     email,
		Username:  username,
		Password:  string(hash),
		CreatedAt: ptrTime(time.Now()),
	}

	return uc.repo.Create(user)
}

func (uc *UserUseCase) Login(email, password string) (string, error) {
	user, err := uc.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	return auth.GenerateToken(user.ID)
}

func (uc *UserUseCase) ValidateToken(token string) (uint, error) {
	return auth.ValidateToken(token)
}

func ptrTime(t time.Time) *time.Time {
	return &t
}
