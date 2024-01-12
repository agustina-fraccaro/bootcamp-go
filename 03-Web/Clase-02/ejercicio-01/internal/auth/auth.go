package auth

import "errors"

var (
	ErrAuthTokenInternal = errors.New("authenticator: internal error")

	ErrAuthTokenInvalid = errors.New("authenticator: token invalid")

	ErrAuthTokenNotFound = errors.New("authenticator: token not found")

	ErrAuthTokenExpired = errors.New("authenticator: token expired")
)

type AuthToken interface {
	Auth(token string) (err error)
}
