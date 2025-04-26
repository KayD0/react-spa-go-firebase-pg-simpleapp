package repository

import (
    "context" // コンテキストパッケージをインポート
    "errors"  // エラーパッケージをインポート
    "time"    // 時間パッケージをインポート

    "github.com/baseapp/domain/entity"     // ドメインエンティティをインポート
    "github.com/baseapp/domain/repository" // ドメインリポジトリをインポート
    "gorm.io/gorm" // GORM をインポート
)

// UserModel はデータベース内のユーザーを表します
type UserModel struct {
    ID          uint      `gorm:"primaryKey"` // 主キー
    FirebaseUID string    `gorm:"unique;not null;size:128"` // Firebase UID（ユニーク、NULL不可、サイズ128）
    DisplayName string    `gorm:"size:100"` // 表示名（サイズ100）
    Bio         string    `gorm:"type:text"` // 自己紹介（テキスト型）
    Location    string    `gorm:"size:100"` // 所在地（サイズ100）
    Website     string    `gorm:"size:255"` // ウェブサイト（サイズ255）
    CreatedAt   time.Time // 作成日時
    UpdatedAt   time.Time // 更新日時
}

// TableName は UserModel に対応するテーブル名を指定します
func (UserModel) TableName() string {
    return "user_profiles" // テーブル名を "user_profiles" に設定
}

// ToEntity は UserModel を User エンティティに変換します
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

// FromEntity は User エンティティを UserModel に変換します
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

// UserRepositoryImpl は UserRepository インターフェースを実装します
type UserRepositoryImpl struct {
    db *gorm.DB // GORM データベースインスタンス
}

// NewUserRepository は新しい UserRepositoryImpl を作成します
func NewUserRepository(db *gorm.DB) repository.UserRepository {
    return &UserRepositoryImpl{
        db: db, // データベースインスタンスを設定
    }
}

// FindByID は ID に基づいてユーザーを検索します
func (r *UserRepositoryImpl) FindByID(ctx context.Context, id uint) (*entity.User, error) {
    var model UserModel
    result := r.db.First(&model, id) // ID に基づいて最初のレコードを取得
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, errors.New("ユーザーが見つかりません") // ユーザーが見つからない場合のエラーメッセージ
        }
        return nil, result.Error // その他のエラーを返す
    }
    
    return model.ToEntity(), nil // ユーザーエンティティを返す
}

// FindByFirebaseUID は Firebase UID に基づいてユーザーを検索します
func (r *UserRepositoryImpl) FindByFirebaseUID(ctx context.Context, firebaseUID string) (*entity.User, error) {
    var model UserModel
    result := r.db.Where("firebase_uid = ?", firebaseUID).First(&model) // Firebase UID に基づいて最初のレコードを取得
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, errors.New("ユーザーが見つかりません") // ユーザーが見つからない場合のエラーメッセージ
        }
        return nil, result.Error // その他のエラーを返す
    }
    
    return model.ToEntity(), nil // ユーザーエンティティを返す
}

// Create は新しいユーザーを作成します
func (r *UserRepositoryImpl) Create(ctx context.Context, user *entity.User) error {
    var model UserModel
    model.FromEntity(user) // エンティティからモデルにデータをコピー
    
    result := r.db.Create(&model) // モデルをデータベースに作成
    if result.Error != nil {
        return result.Error // エラーが発生した場合はエラーを返す
    }
    
    // エンティティの ID を更新
    user.ID = model.ID
    return nil
}

// Update は既存のユーザーを更新します
func (r *UserRepositoryImpl) Update(ctx context.Context, user *entity.User) error {
    var model UserModel
    model.FromEntity(user) // エンティティからモデルにデータをコピー
    
    result := r.db.Save(&model) // モデルをデータベースに保存
    return result.Error // エラーを返す
}
