package services

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"time"
     
	"user-service/models"
	"user-service/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(email, password string) (*models.User, error)
	Login(email, password string) (string, error)
	RequestPasswordReset(email string) error
	ResetPassword(token, newPassword string) error
}

type userService struct {
	repo        repository.UserRepository
	jwtSecret   string
	redisClient *redis.Client
}

func NewUserService(repo repository.UserRepository, jwtSecret string, redisClient *redis.Client) UserService {
	return &userService{repo, jwtSecret, redisClient}
}

func (s *userService) Register(email, password string) (*models.User, error) {
	_, err := s.repo.FindByEmail(email)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT with token_version
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":       user.ID,
		"token_version": user.TokenVersion,
		"exp":           time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *userService) RequestPasswordReset(email string) error {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		// Do not leak if user exists or not, but return no error outwardly if needed,
		// though returning nil is fine for consistent API response.
		return nil
	}

	// Generate 32-byte secure token
	tokenBytes := scratchBytes(32)
	_, err = rand.Read(tokenBytes)
	if err != nil {
		return err
	}
	resetToken := hex.EncodeToString(tokenBytes)

	// Hash the token
	hash := sha256.Sum256([]byte(resetToken))
	hashedToken := hex.EncodeToString(hash[:])
	
	expr := time.Now().Add(15 * time.Minute)

	user.ResetTokenHash = &hashedToken
	user.ResetTokenExpiresAt = &expr

	// Update DB
	err = s.repo.UpdateUser(user)
	if err != nil {
		return err
	}

	// In a real app we'd send an email. For now, log the URL:
	log.Printf("Password Reset URL: http://localhost:5174/reset-password?token=%s\n", resetToken)
	return nil
}

func scratchBytes(n int) []byte {
	return make([]byte, n)
}

func (s *userService) ResetPassword(token, newPassword string) error {
	// Hash incoming token
	hash := sha256.Sum256([]byte(token))
	hashedToken := hex.EncodeToString(hash[:])

	user, err := s.repo.FindByResetToken(hashedToken)
	if err != nil {
		return errors.New("invalid or expired token")
	}

	if user.ResetTokenExpiresAt == nil || user.ResetTokenExpiresAt.Before(time.Now()) {
		return errors.New("invalid or expired token")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	user.TokenVersion++ // Increment version to invalidate old JWTs
	user.ResetTokenHash = nil
	user.ResetTokenExpiresAt = nil

	if err := s.repo.UpdateUser(user); err != nil {
		return err
	}

	// Push new token version to Redis so task-service (AuthMiddleware) knows instantly
	ctx := context.Background()
	redisKey := fmt.Sprintf("user_version:%d", user.ID)
	s.redisClient.Set(ctx, redisKey, user.TokenVersion, time.Hour*72)

	return nil
}
