package services

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	ent "github.com/naumovrus/weather-api/internal/entities"
	passwordhasher "github.com/naumovrus/weather-api/internal/pkg/passwordhash"
	"github.com/naumovrus/weather-api/internal/repository"
)

const (
	tokenTTL  = time.Hour * 12
	signedKey = ("grj#zjaAJzj$%askj4551##sa")
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user ent.User) (int, error) {
	user.Password = passwordhasher.GeneratePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, passwordhasher.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
	},
		user.Id,
	})
	return token.SignedString([]byte(signedKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signedKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type")
	}
	return claims.UserId, nil
}
