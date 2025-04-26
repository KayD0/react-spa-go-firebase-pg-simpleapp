package persistence

import (
    "time"

    "gorm.io/gorm" // GORM をインポート
)

// UserProfileModel はデータベース内のユーザープロフィールを表します
type UserProfileModel struct {
    ID          uint      `gorm:"primaryKey"` // 主キー
    FirebaseUID string    `gorm:"unique;not null;size:128"` // Firebase UID（ユニーク、NULL不可、サイズ128）
    DisplayName string    `gorm:"size:100"` // 表示名（サイズ100）
    Bio         string    `gorm:"type:text"` // 自己紹介（テキスト型）
    Location    string    `gorm:"size:100"` // 所在地（サイズ100）
    Website     string    `gorm:"size:255"` // ウェブサイト（サイズ255）
    CreatedAt   time.Time // 作成日時
    UpdatedAt   time.Time // 更新日時
}

// TableName は UserProfileModel に対応するテーブル名を指定します
func (UserProfileModel) TableName() string {
    return "user_profiles" // テーブル名を "user_profiles" に設定
}

// BeforeCreate は GORM フックで、CreatedAt と UpdatedAt フィールドを設定します
func (u *UserProfileModel) BeforeCreate(tx *gorm.DB) error {
    now := time.Now() // 現在の時刻を取得
    u.CreatedAt = now // CreatedAt フィールドに現在の時刻を設定
    u.UpdatedAt = now // UpdatedAt フィールドにも現在の時刻を設定
    return nil
}

// BeforeUpdate は GORM フックで、UpdatedAt フィールドを設定します
func (u *UserProfileModel) BeforeUpdate(tx *gorm.DB) error {
    u.UpdatedAt = time.Now() // UpdatedAt フィールドに現在の時刻を設定
    return nil
}
