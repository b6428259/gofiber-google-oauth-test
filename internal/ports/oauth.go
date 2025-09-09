package ports

import (
    "context"
    "gofiber-hex-google-oauth/internal/domain"
)

type OAuthClient interface {
    LoginURL(state string) string
    Exchange(ctx context.Context, code string) (domain.User, error)
}
