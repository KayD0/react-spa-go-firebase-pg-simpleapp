package usecase

import (
    "context"
    "errors"
    "time"

    "github.com/baseapp/application/dto"
    "github.com/baseapp/domain/entity"
    "github.com/baseapp/domain/repository"
)

// UserUseCase はユーザー関連のビジネスロジックを実装する構造体です
type UserUseCase struct {
    userRepo repository.UserRepository // ユーザーリポジトリのインスタンス
}

// NewUserUseCase は新しい UserUseCase を作成するコンストラクタです
func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
    return &UserUseCase{
        userRepo: userRepo, // 引数として渡されたユーザーリポジトリをフィールドに設定
    }
}

// GetProfile は Firebase UID に基づいてユーザープロフィールを取得するメソッドです
func (uc *UserUseCase) GetProfile(ctx context.Context, firebaseUID string) (*dto.UserResponse, error) {
    // Firebase UID でユーザーを検索
    user, err := uc.userRepo.FindByFirebaseUID(ctx, firebaseUID)
    if err != nil {
        // ユーザーが見つからない場合、新しいユーザーを作成
        newUser := &entity.User{
            FirebaseUID: firebaseUID,
            CreatedAt:   time.Now(), // 作成日時を設定
            UpdatedAt:   time.Now(), // 更新日時を設定
        }
        
        // 新しいユーザーをリポジトリに保存
        if err := uc.userRepo.Create(ctx, newUser); err != nil {
            return nil, errors.New("プロフィールの作成に失敗しました: " + err.Error())
        }
        
        // 新しいユーザーのレスポンスを返す
        return dto.NewUserResponse(newUser), nil
    }
    
    // 既存のユーザーのレスポンスを返す
    return dto.NewUserResponse(user), nil
}

// UpdateProfile はユーザープロフィールを更新するメソッドです
func (uc *UserUseCase) UpdateProfile(ctx context.Context, firebaseUID string, updateReq *dto.UserUpdateRequest) (*dto.UserResponse, error) {
    // Firebase UID でユーザーを検索
    user, err := uc.userRepo.FindByFirebaseUID(ctx, firebaseUID)
    if err != nil {
        // ユーザーが見つからない場合、新しいユーザーを作成
        newUser := &entity.User{
            FirebaseUID: firebaseUID,
            DisplayName: updateReq.DisplayName, // 表示名を設定
            Bio:         updateReq.Bio,         // 自己紹介を設定
            Location:    updateReq.Location,    // 所在地を設定
            Website:     updateReq.Website,     // ウェブサイトを設定
            CreatedAt:   time.Now(),            // 作成日時を設定
            UpdatedAt:   time.Now(),            // 更新日時を設定
        }
        
        // 新しいユーザーをリポジトリに保存
        if err := uc.userRepo.Create(ctx, newUser); err != nil {
            return nil, errors.New("プロフィールの作成に失敗しました: " + err.Error())
        }
        
        // 新しいユーザーのレスポンスを返す
        return dto.NewUserResponse(newUser), nil
    }
    
    // 既存のユーザーを更新
    updatedUser := updateReq.ToEntity(user)
    
    // ユーザーをリポジトリで更新
    if err := uc.userRepo.Update(ctx, updatedUser); err != nil {
        return nil, errors.New("プロフィールの更新に失敗しました: " + err.Error())
    }
    
    // 更新されたユーザーのレスポンスを返す
    return dto.NewUserResponse(updatedUser), nil
}
