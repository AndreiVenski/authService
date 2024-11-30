package models

import "github.com/google/uuid"

type User struct {
	UserID uuid.UUID `db:"user_id"`
	Name   string    `db:"name"`
	Email  string    `db:"email"`
}

type UserInfo struct {
	UserID uuid.UUID `json:"user_id"`
	IP     string    `json:"ip_addr"`
}
