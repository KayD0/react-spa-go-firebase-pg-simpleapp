package auth

import (
	"errors"
	"os"
	"strings"
)

// AuthType は認証サービスの種類を表す型です
type AuthType string

const (
	// Firebase 認証
	Firebase AuthType = "firebase"
	// モック認証（テスト用）
	Mock     AuthType = "mock"
	// 将来的に他の認証サービスを追加できます
	// Auth0    AuthType = "auth0"
	// Cognito  AuthType = "cognito"
	// など
)

// NewAuthService は指定された認証タイプに基づいて新しい AuthService を作成します
func NewAuthService(authType AuthType) (AuthService, error) {
	// 環境変数から認証タイプを取得（設定されていない場合はパラメータを使用）
	envAuthType := os.Getenv("AUTH_SERVICE_TYPE")
	if envAuthType != "" {
		authType = AuthType(strings.ToLower(envAuthType))
	}

	// 認証タイプに基づいて適切な認証サービスを作成
	switch authType {
	case Firebase:
		return NewFirebaseAuth()
	case Mock:
		return NewMockAuth()
	// 将来的に他の認証サービスを追加できます
	// case Auth0:
	//     return NewAuth0Auth()
	// case Cognito:
	//     return NewCognitoAuth()
	default:
		return nil, errors.New("サポートされていない認証タイプです: " + string(authType))
	}
}
