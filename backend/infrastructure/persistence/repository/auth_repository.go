package repository

import (
    "context" // コンテキストパッケージをインポート

    "github.com/baseapp/domain/repository" // ドメインリポジトリをインポート
    "github.com/baseapp/infrastructure/auth" // 認証インフラストラクチャをインポート
)

// AuthRepositoryImpl は AuthRepository インターフェースを実装します
type AuthRepositoryImpl struct {
    firebaseAuth auth.FirebaseAuth // Firebase 認証インスタンス
}

// NewAuthRepository は新しい AuthRepositoryImpl を作成します
func NewAuthRepository(firebaseAuth auth.FirebaseAuth) repository.AuthRepository {
    return &AuthRepositoryImpl{
        firebaseAuth: firebaseAuth, // Firebase 認証インスタンスを設定
    }
}

// VerifyToken は認証トークンを検証し、クレームを返します
func (r *AuthRepositoryImpl) VerifyToken(ctx context.Context, token string) (map[string]interface{}, error) {
    // Firebase 認証インスタンスを使用してトークンを検証
    return r.firebaseAuth.VerifyIDToken(ctx, token)
}
