package models

import "github.com/google/uuid"

type UserInfo struct {
	UserID uuid.UUID
	IP     string
}
