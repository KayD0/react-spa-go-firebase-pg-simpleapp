package models

import (
    "time"

    "gorm.io/gorm"
)

// UserProfile represents a user profile in the database
type UserProfile struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    FirebaseUID string    `gorm:"unique;not null;size:128" json:"firebase_uid"`
    DisplayName string    `gorm:"size:100" json:"display_name"`
    Bio         string    `gorm:"type:text" json:"bio"`
    Location    string    `gorm:"size:100" json:"location"`
    Website     string    `gorm:"size:255" json:"website"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

// TableName specifies the table name for the UserProfile model
func (UserProfile) TableName() string {
    return "user_profiles"
}

// BeforeCreate is a GORM hook that sets the CreatedAt and UpdatedAt fields
func (u *UserProfile) BeforeCreate(tx *gorm.DB) error {
    now := time.Now()
    u.CreatedAt = now
    u.UpdatedAt = now
    return nil
}

// BeforeUpdate is a GORM hook that sets the UpdatedAt field
func (u *UserProfile) BeforeUpdate(tx *gorm.DB) error {
    u.UpdatedAt = time.Now()
    return nil
}

// ToMap converts the UserProfile to a map
func (u *UserProfile) ToMap() map[string]interface{} {
    return map[string]interface{}{
        "id":           u.ID,
        "firebase_uid": u.FirebaseUID,
        "display_name": u.DisplayName,
        "bio":          u.Bio,
        "location":     u.Location,
        "website":      u.Website,
        "created_at":   u.CreatedAt.Format(time.RFC3339),
        "updated_at":   u.UpdatedAt.Format(time.RFC3339),
    }
}
