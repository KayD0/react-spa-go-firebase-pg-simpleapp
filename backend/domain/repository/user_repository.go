package repository

import (
    "context" // コンテキストパッケージをインポート

    "github.com/baseapp/domain/entity" // ドメインエンティティをインポート
)

// UserRepository はユーザーデータ操作のためのインターフェースを定義します
type UserRepository interface {
    // FindByID は ID に基づいてユーザーを検索します
    FindByID(ctx context.Context, id uint) (*entity.User, error)
    
    // FindByFirebaseUID は Firebase UID に基づいてユーザーを検索します
    FindByFirebaseUID(ctx context.Context, firebaseUID string) (*entity.User, error)
    
    // Create は新しいユーザーを作成します
    Create(ctx context.Context, user *entity.User) error
    
    // Update は既存のユーザーを更新します
    Update(ctx context.Context, user *entity.User) error
}
