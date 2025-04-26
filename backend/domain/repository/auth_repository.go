package repository

import (
    "context"
)

// AuthRepository defines the interface for authentication operations
type AuthRepository interface {
    // VerifyToken verifies an authentication token and returns the claims
    VerifyToken(ctx context.Context, token string) (map[string]interface{}, error)
}
