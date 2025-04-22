package repository

import (
	"context"

	"github.com/baseapp/backend/domain/repository"
	"github.com/baseapp/backend/infrastructure/auth"
)

// AuthRepositoryImpl implements the AuthRepository interface
type AuthRepositoryImpl struct {
	firebaseAuth auth.FirebaseAuth
}

// NewAuthRepository creates a new AuthRepositoryImpl
func NewAuthRepository(firebaseAuth auth.FirebaseAuth) repository.AuthRepository {
	return &AuthRepositoryImpl{
		firebaseAuth: firebaseAuth,
	}
}

// VerifyToken verifies an authentication token and returns the claims
func (r *AuthRepositoryImpl) VerifyToken(ctx context.Context, token string) (map[string]interface{}, error) {
	return r.firebaseAuth.VerifyIDToken(ctx, token)
}
