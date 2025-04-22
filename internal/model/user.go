package model

import (
	"time"

	"github.com/google/uuid"
)

type UserStatus string

const (
	UserStatusPending  UserStatus = "pending"
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
	UserStatusSuspend  UserStatus = "suspend"
)

type Argon2Config struct {
	Salt      string `json:"salt"`
	Time      uint32 `json:"time"`
	Memory    uint32 `json:"memory"`
	Threads   uint8  `json:"threads"`
	KeyLength uint32 `json:"keyLen"`
}

type User struct {
	ID             int64        `json:"id"`
	UUID           uuid.UUID    `json:"uuid"`
	Username       string       `json:"username"`
	Email          string       `json:"email"`
	PhoneNumber    string       `json:"phone_number,omitempty"`
	PasswordHash   string       `json:"-"`
	PasswordConfig Argon2Config `json:"-"`
	Status         UserStatus   `json:"status"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	DeletedAt      *time.Time   `json:"deleted_at,omitempty"`
}
