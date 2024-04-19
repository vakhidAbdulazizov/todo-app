package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/vakhidAbdulazizov/todo-app/models"
	"github.com/vakhidAbdulazizov/todo-app/pkg/repository"
	"math/rand"
	"strings"
	"time"
)

const (
	salt      = "ytequetyutdcbzmncbmxzhdsjhqwkqjkweh"
	tokenTTl  = 12 * time.Hour
	signinkey = "dqj$%$51bdhdaak"
)

type AuthService struct {
	repo repository.Authorization
}

type jwtTokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"userId"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generateHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) ForgotPassword(email string) error {
	confirmKey := randKey()
	return s.repo.ForgotPassword(email, generateHash(confirmKey), confirmKey)
}

func (s *AuthService) GenerateToken(username string, password string) (string, error) {
	user, err := s.repo.GetUser(username, generateHash(password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signinkey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &jwtTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signinkey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*jwtTokenClaims)

	if !ok {
		return 0, errors.New("token claims are not type *jwtTokenClaims")
	}

	return claims.UserId, nil
}

func (s *AuthService) RestorePassword(email string, confirmKey string, password string) error {
	return s.repo.RestorePassword(email, generateHash(confirmKey), generateHash(password))
}
func generateHash(chars string) string {
	hash := sha1.New()

	hash.Write([]byte(chars))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func randKey() string {
	rand.Seed(time.Now().UnixNano())

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "012343567899")
	length := 6
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
