package httpErrors

import "github.com/pkg/errors"

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrRefreshTokenNotFound  = errors.New("refresh token not found")
	ErrRefreshTokenExpires   = errors.New("token expires")
	ErrRefreshTokenIncorrect = errors.New("token is incorrect")
)
