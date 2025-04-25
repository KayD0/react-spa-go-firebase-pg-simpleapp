package usecase

import (
	"context"
	"errors"

	"github.com/baseapp/application/dto"
	"github.com/baseapp/domain/repository"
)

// AuthUseCase implements the authentication-related business logic
type AuthUseCase struct {
	authRepo repository.AuthRepository
}

// NewAuthUseCase creates a new AuthUseCase
func NewAuthUseCase(authRepo repository.AuthRepository) *AuthUseCase {
	return &AuthUseCase{
		authRepo: authRepo,
	}
}

// VerifyToken verifies an authentication token
func (uc *AuthUseCase) VerifyToken(ctx context.Context, token string) (*dto.AuthResponse, error) {
	claims, err := uc.authRepo.VerifyToken(ctx, token)
	if err != nil {
		return nil, errors.New("無効な認証トークンです: " + err.Error())
	}
	
	return dto.NewAuthResponse(claims), nil
}

// GetUserIDFromClaims extracts the user ID from token claims
func (uc *AuthUseCase) GetUserIDFromClaims(claims map[string]interface{}) (string, error) {
	uid, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("ユーザーIDが見つかりません")
	}
	
	return uid, nil
}
