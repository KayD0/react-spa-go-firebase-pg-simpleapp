package usecase

import (
    "context"
    "errors"
    "time"

    "github.com/baseapp/application/dto"
    "github.com/baseapp/domain/entity"
    "github.com/baseapp/domain/repository"
)

// UserUseCase implements the user-related business logic
type UserUseCase struct {
    userRepo repository.UserRepository
}

// NewUserUseCase creates a new UserUseCase
func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
    return &UserUseCase{
        userRepo: userRepo,
    }
}

// GetProfile gets a user profile by Firebase UID
func (uc *UserUseCase) GetProfile(ctx context.Context, firebaseUID string) (*dto.UserResponse, error) {
    user, err := uc.userRepo.FindByFirebaseUID(ctx, firebaseUID)
    if err != nil {
        // If user not found, create a new one
        newUser := &entity.User{
            FirebaseUID: firebaseUID,
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        }
        
        if err := uc.userRepo.Create(ctx, newUser); err != nil {
            return nil, errors.New("プロフィールの作成に失敗しました: " + err.Error())
        }
        
        return dto.NewUserResponse(newUser), nil
    }
    
    return dto.NewUserResponse(user), nil
}

// UpdateProfile updates a user profile
func (uc *UserUseCase) UpdateProfile(ctx context.Context, firebaseUID string, updateReq *dto.UserUpdateRequest) (*dto.UserResponse, error) {
    user, err := uc.userRepo.FindByFirebaseUID(ctx, firebaseUID)
    if err != nil {
        // If user not found, create a new one
        newUser := &entity.User{
            FirebaseUID: firebaseUID,
            DisplayName: updateReq.DisplayName,
            Bio:         updateReq.Bio,
            Location:    updateReq.Location,
            Website:     updateReq.Website,
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        }
        
        if err := uc.userRepo.Create(ctx, newUser); err != nil {
            return nil, errors.New("プロフィールの作成に失敗しました: " + err.Error())
        }
        
        return dto.NewUserResponse(newUser), nil
    }
    
    // Update the user
    updatedUser := updateReq.ToEntity(user)
    
    if err := uc.userRepo.Update(ctx, updatedUser); err != nil {
        return nil, errors.New("プロフィールの更新に失敗しました: " + err.Error())
    }
    
    return dto.NewUserResponse(updatedUser), nil
}
