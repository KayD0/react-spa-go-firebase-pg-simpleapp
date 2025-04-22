package persistence

import (
	"time"

	"gorm.io/gorm"
)

// UserProfileModel represents a user profile in the database
type UserProfileModel struct {
	ID          uint      `gorm:"primaryKey"`
	FirebaseUID string    `gorm:"unique;not null;size:128"`
	DisplayName string    `gorm:"size:100"`
	Bio         string    `gorm:"type:text"`
	Location    string    `gorm:"size:100"`
	Website     string    `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// TableName specifies the table name for the UserProfileModel
func (UserProfileModel) TableName() string {
	return "user_profiles"
}

// BeforeCreate is a GORM hook that sets the CreatedAt and UpdatedAt fields
func (u *UserProfileModel) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
	return nil
}

// BeforeUpdate is a GORM hook that sets the UpdatedAt field
func (u *UserProfileModel) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
