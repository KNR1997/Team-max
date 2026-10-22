package domain

import (
	"time"

	"todo/internal/shared/kernel"
)

type User struct {
	ID           kernel.UserID
	Email        string
	PasswordHash string
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

var (
	ErrInvalidCredentials = kernel.ErrUnauthorized
	ErrUserNotFound       = kernel.ErrNotFound
)
