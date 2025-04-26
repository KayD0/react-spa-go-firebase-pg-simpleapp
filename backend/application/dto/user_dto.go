package dto

import (
    "time"

    "github.com/baseapp/domain/entity"
)

// UserResponse はクライアントに送信されるユーザーデータを表す構造体です
type UserResponse struct {
    ID          uint   `json:"id"`               // ユーザーの一意の識別子
    FirebaseUID string `json:"firebase_uid"`     // Firebase UID
    DisplayName string `json:"display_name"`     // 表示名
    Bio         string `json:"bio"`              // 自己紹介
    Location    string `json:"location"`         // 所在地
    Website     string `json:"website"`          // ウェブサイト
    CreatedAt   string `json:"created_at"`      // 作成日時
    UpdatedAt   string `json:"updated_at"`      // 更新日時
}

// UserUpdateRequest はクライアントから受信するユーザー更新データを表す構造体です
type UserUpdateRequest struct {
    DisplayName string `json:"display_name"` // 表示名
    Bio         string `json:"bio"`          // 自己紹介
    Location    string `json:"location"`     // 所在地
    Website     string `json:"website"`      // ウェブサイト
}

// NewUserResponse は User エンティティから新しい UserResponse を作成する関数です
func NewUserResponse(user *entity.User) *UserResponse {
    return &UserResponse{
        ID:          user.ID,
        FirebaseUID: user.FirebaseUID,
        DisplayName: user.DisplayName,
        Bio:         user.Bio,
        Location:    user.Location,
        Website:     user.Website,
        CreatedAt:   user.CreatedAt.Format(time.RFC3339), // 作成日時をRFC3339形式にフォーマット
        UpdatedAt:   user.UpdatedAt.Format(time.RFC3339), // 更新日時をRFC3339形式にフォーマット
    }
}

// ToEntity は UserUpdateRequest を User エンティティに変換するメソッドです
func (r *UserUpdateRequest) ToEntity(existingUser *entity.User) *entity.User {
    // 既存のユーザーが nil の場合、新しいユーザーを作成
    if existingUser == nil {
        existingUser = &entity.User{}
    }
    
    // 各フィールドが空でない場合、既存のユーザーに値を設定
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
    
    // 更新日時を現在の時刻に設定
    existingUser.UpdatedAt = time.Now()
    return existingUser // 更新されたユーザーエンティティを返す
}
