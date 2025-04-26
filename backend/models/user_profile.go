package models

import (
    "time"

    "gorm.io/gorm"
)

// UserProfile はデータベース内のユーザープロフィールを表す構造体です
type UserProfile struct {
    ID          uint      `gorm:"primaryKey" json:"id"`               // ユーザーの一意の識別子
    FirebaseUID string    `gorm:"unique;not null;size:128" json:"firebase_uid"` // Firebase UID（ユニークかつ非NULL）
    DisplayName string    `gorm:"size:100" json:"display_name"`       // 表示名
    Bio         string    `gorm:"type:text" json:"bio"`                // 自己紹介
    Location    string    `gorm:"size:100" json:"location"`            // 所在地
    Website     string    `gorm:"size:255" json:"website"`             // ウェブサイト
    CreatedAt   time.Time `json:"created_at"`                          // 作成日時
    UpdatedAt   time.Time `json:"updated_at"`                          // 更新日時
}

// TableName は UserProfile モデルのテーブル名を指定します
func (UserProfile) TableName() string {
    return "user_profiles" // テーブル名を "user_profiles" に設定
}

// BeforeCreate は GORM フックで、CreatedAt と UpdatedAt フィールドを設定します
func (u *UserProfile) BeforeCreate(tx *gorm.DB) error {
    now := time.Now() // 現在の時刻を取得
    u.CreatedAt = now // CreatedAt に現在の時刻を設定
    u.UpdatedAt = now // UpdatedAt に現在の時刻を設定
    return nil
}

// BeforeUpdate は GORM フックで、UpdatedAt フィールドを設定します
func (u *UserProfile) BeforeUpdate(tx *gorm.DB) error {
    u.UpdatedAt = time.Now() // UpdatedAt に現在の時刻を設定
    return nil
}

// ToMap は UserProfile をマップに変換します
func (u *UserProfile) ToMap() map[string]interface{} {
    return map[string]interface{}{
        "id":           u.ID,
        "firebase_uid": u.FirebaseUID,
        "display_name": u.DisplayName,
        "bio":          u.Bio,
        "location":     u.Location,
        "website":      u.Website,
        "created_at":   u.CreatedAt.Format(time.RFC3339), // 作成日時をRFC3339形式にフォーマット
        "updated_at":   u.UpdatedAt.Format(time.RFC3339), // 更新日時をRFC3339形式にフォーマット
    }
}
