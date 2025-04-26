package repository

import (
	"context"
	"errors"
	"time"

	"github.com/baseapp/domain/entity"
	"github.com/baseapp/domain/repository"
	"gorm.io/gorm"
)

// UserModel represents a user in the database
type UserModel struct {
	ID          uint      `gorm:"primaryKey"`
	FirebaseUID string    `gorm:"unique;not null;size:128"`
	DisplayName string    `gorm:"size:100"`
	Bio         string    `gorm:"type:text"`
	Location    string    `gorm:"size:100"`
	Website     string    `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// TableName specifies the table name for the UserModel
func (UserModel) TableName() string {
	return "user_profiles"
}

// ToEntity converts a UserModel to a User entity
func (m *UserModel) ToEntity() *entity.User {
	return &entity.User{
		ID:          m.ID,
		FirebaseUID: m.FirebaseUID,
		DisplayName: m.DisplayName,
		Bio:         m.Bio,
		Location:    m.Location,
		Website:     m.Website,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

// FromEntity converts a User entity to a UserModel
func (m *UserModel) FromEntity(user *entity.User) {
	m.ID = user.ID
	m.FirebaseUID = user.FirebaseUID
	m.DisplayName = user.DisplayName
	m.Bio = user.Bio
	m.Location = user.Location
	m.Website = user.Website
	m.CreatedAt = user.CreatedAt
	m.UpdatedAt = user.UpdatedAt
}

// UserRepositoryImpl implements the UserRepository interface
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepositoryImpl
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

// FindByID finds a user by ID
func (r *UserRepositoryImpl) FindByID(ctx context.Context, id uint) (*entity.User, error) {
	var model UserModel
	result := r.db.First(&model, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("ユーザーが見つかりません")
		}
		return nil, result.Error
	}
	
	return model.ToEntity(), nil
}

// FindByFirebaseUID finds a user by Firebase UID
func (r *UserRepositoryImpl) FindByFirebaseUID(ctx context.Context, firebaseUID string) (*entity.User, error) {
	var model UserModel
	result := r.db.Where("firebase_uid = ?", firebaseUID).First(&model)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("ユーザーが見つかりません")
		}
		return nil, result.Error
	}
	
	return model.ToEntity(), nil
}

// Create creates a new user
func (r *UserRepositoryImpl) Create(ctx context.Context, user *entity.User) error {
	var model UserModel
	model.FromEntity(user)
	
	result := r.db.Create(&model)
	if result.Error != nil {
		return result.Error
	}
	
	// Update the ID in the entity
	user.ID = model.ID
	return nil
}

// Update updates an existing user
func (r *UserRepositoryImpl) Update(ctx context.Context, user *entity.User) error {
	var model UserModel
	model.FromEntity(user)
	
	result := r.db.Save(&model)
	return result.Error
}
