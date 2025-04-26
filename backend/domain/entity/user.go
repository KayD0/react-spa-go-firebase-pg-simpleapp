package entity

import (
    "time" // 時間関連のパッケージをインポート
)

// User はドメイン内のユーザーエンティティを表します
type User struct {
    ID          uint      // ユーザーの一意な識別子
    FirebaseUID string    // Firebase におけるユーザーの UID
    DisplayName string    // ユーザーの表示名
    Bio         string    // ユーザーの自己紹介
    Location    string    // ユーザーの所在地
    Website     string    // ユーザーのウェブサイトURL
    CreatedAt   time.Time // ユーザーが作成された日時
    UpdatedAt   time.Time // ユーザーが最後に更新された日時
}
