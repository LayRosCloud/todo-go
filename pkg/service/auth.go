package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/layroscloud/todo-go/entity"
	"github.com/layroscloud/todo-go/pkg/repository"
)

const (
	Salt      = "jhgbhfdjkfjeiwgbryuheha"
	Signature = "dsnjfsanujifsnaujid"
	TokenTtl  = 12 * time.Hour
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a AuthService) CreateUser(user entity.User) (int, error) {
	user.Password = GenerateHashPassword(user.Password)
	return a.repo.CreateUser(user)
}

func (a AuthService) GenerateToken(username string, password string) (string, error) {
	passwordHash := GenerateHashPassword(password)
	user, err := a.repo.FindByUsernameAndPassword(username, passwordHash)

	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTtl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	token, _ := jwtToken.SignedString([]byte(Signature))
	return token, nil
}

func (a AuthService) ParseToken(accessToken string) (int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signature method")
		}
		return []byte(Signature), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("error! not received claims data")
	}

	return claims.UserId, nil
}

func GenerateHashPassword(password string) string {
	hash := sha1.New()
	bytesPassword := []byte(password)
	hash.Write(bytesPassword)
	return fmt.Sprintf("%x", hash.Sum([]byte(Salt)))
}
