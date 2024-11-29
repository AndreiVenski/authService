package models

import "github.com/google/uuid"

type User struct {
	UserID uuid.UUID
	Name   string
	Email  string
}

type UserInfo struct {
	UserID uuid.UUID
	IP     string
}
