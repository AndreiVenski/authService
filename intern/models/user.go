package models

import "github.com/google/uuid"

type UserInfo struct {
	userID uuid.UUID
	ip     int
}

type Tokens struct {
}
