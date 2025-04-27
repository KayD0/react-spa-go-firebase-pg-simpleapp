package repository

import (
	"context" // コンテキストパッケージをインポート

	"github.com/baseapp/domain/authentication"
	"github.com/baseapp/domain/repository" // ドメインリポジトリをインポート
	// 認証インフラストラクチャをインポート
)

// AuthRepositoryImpl は AuthRepository インターフェースを実装します
type AuthRepositoryImpl struct {
    authService authentication.AuthService // 認証サービスインスタンス
}

// NewAuthRepository は新しい AuthRepositoryImpl を作成します
func NewAuthRepository(authService authentication.AuthService) repository.AuthRepository {
    return &AuthRepositoryImpl{
        authService: authService, // 認証サービスインスタンスを設定
    }
}

// VerifyToken は認証トークンを検証し、クレームを返します
func (r *AuthRepositoryImpl) VerifyToken(ctx context.Context, token string) (map[string]interface{}, error) {
    // 認証サービスインスタンスを使用してトークンを検証
    return r.authService.VerifyToken(ctx, token)
}
