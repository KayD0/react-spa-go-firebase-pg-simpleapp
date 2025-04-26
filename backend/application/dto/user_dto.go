package dto

import (
    "time"

    "github.com/baseapp/domain/entity"
)

// UserResponse represents the user data that will be sent to clients
type UserResponse struct {
    ID          uint   `json:"id"`
    FirebaseUID string `json:"firebase_uid"`
    DisplayName string `json:"display_name"`
    Bio         string `json:"bio"`
    Location    string `json:"location"`
    Website     string `json:"website"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
}

// UserUpdateRequest represents the data received from clients to update a user
type UserUpdateRequest struct {
    DisplayName string `json:"display_name"`
    Bio         string `json:"bio"`
    Location    string `json:"location"`
    Website     string `json:"website"`
}

// NewUserResponse creates a new UserResponse from a User entity
func NewUserResponse(user *entity.User) *UserResponse {
    return &UserResponse{
        ID:          user.ID,
        FirebaseUID: user.FirebaseUID,
        DisplayName: user.DisplayName,
        Bio:         user.Bio,
        Location:    user.Location,
        Website:     user.Website,
        CreatedAt:   user.CreatedAt.Format(time.RFC3339),
        UpdatedAt:   user.UpdatedAt.Format(time.RFC3339),
    }
}

// ToEntity converts a UserUpdateRequest to a User entity
func (r *UserUpdateRequest) ToEntity(existingUser *entity.User) *entity.User {
    if existingUser == nil {
        existingUser = &entity.User{}
    }
    
    if r.DisplayName != "" {
        existingUser.DisplayName = r.DisplayName
    }
    if r.Bio != "" {
        existingUser.Bio = r.Bio
    }
    if r.Location != "" {
        existingUser.Location = r.Location
    }
    if r.Website != "" {
        existingUser.Website = r.Website
    }
    
    existingUser.UpdatedAt = time.Now()
    return existingUser
}
