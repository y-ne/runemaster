package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
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
	Salt      string `json:"salt" db:"salt"`
	Time      uint32 `json:"time" db:"time"`
	Memory    uint32 `json:"memory" db:"memory"`
	Threads   uint8  `json:"threads" db:"threads"`
	KeyLength uint32 `json:"keyLen" db:"keyLen"`
}

func (a Argon2Config) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Argon2Config) Scan(value any) error {
	bytes, ok := value.([]byte)

	if !ok {
		return fmt.Errorf("failed to convert password_config to bytes")
	}

	return json.Unmarshal(bytes, a)
}

type User struct {
	ID             int64        `json:"id" db:"id"`
	UUID           uuid.UUID    `json:"uuid" db:"user_uuid"`
	Username       string       `json:"username" db:"username"`
	Email          string       `json:"email" db:"email"`
	PhoneNumber    string       `json:"phone_number,omitempty" db:"phone_number"`
	PasswordHash   string       `json:"-" db:"password_hash"`
	PasswordConfig Argon2Config `json:"-" db:"password_config"`
	Status         UserStatus   `json:"status" db:"status"`
	CreatedAt      time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time   `json:"deleted_at,omitempty" db:"deleted_at"`
}
