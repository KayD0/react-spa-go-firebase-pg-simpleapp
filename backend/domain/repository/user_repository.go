package repository

import (
	"context"

	"github.com/baseapp/backend/domain/entity"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	// FindByID finds a user by ID
	FindByID(ctx context.Context, id uint) (*entity.User, error)
	
	// FindByFirebaseUID finds a user by Firebase UID
	FindByFirebaseUID(ctx context.Context, firebaseUID string) (*entity.User, error)
	
	// Create creates a new user
	Create(ctx context.Context, user *entity.User) error
	
	// Update updates an existing user
	Update(ctx context.Context, user *entity.User) error
}
