package app

import (
    "context"

    "gofiber-hex-google-oauth/internal/domain"
    "gofiber-hex-google-oauth/internal/ports"
)

type AuthService struct {
    oauth ports.OAuthClient
}

func NewAuthService(o ports.OAuthClient) *AuthService {
    return &AuthService{oauth: o}
}

func (s *AuthService) OAuthLoginURL(state string) string {
    return s.oauth.LoginURL(state)
}

func (s *AuthService) LoginWithGoogle(ctx context.Context, code string) (domain.User, error) {
    return s.oauth.Exchange(ctx, code)
}
