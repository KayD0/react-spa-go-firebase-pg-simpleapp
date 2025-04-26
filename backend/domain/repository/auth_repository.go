package repository

import (
    "context" // コンテキストパッケージをインポート
)

// AuthRepository は認証操作のためのインターフェースを定義します
type AuthRepository interface {
    // VerifyToken は認証トークンを検証し、クレームを返します
    VerifyToken(ctx context.Context, token string) (map[string]interface{}, error)
}
