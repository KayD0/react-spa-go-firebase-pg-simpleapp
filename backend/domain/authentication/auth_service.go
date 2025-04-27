package authentication

import (
	"context"
)

// AuthService は認証サービスのインターフェースを定義します
// 様々な認証プロバイダー（Firebase、Auth0など）で実装できます
type AuthService interface {
	// VerifyToken は認証トークンを検証し、デコードされたトークンを返します
	VerifyToken(ctx context.Context, idToken string) (map[string]interface{}, error)
}
