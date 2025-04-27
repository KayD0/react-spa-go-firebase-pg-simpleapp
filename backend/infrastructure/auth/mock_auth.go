package auth

import (
	"context"
	"errors"
	"time"
)

// MockAuthImpl は AuthService インターフェースを実装するモックの認証サービスです
// テストや開発環境で使用できます
type MockAuthImpl struct {
	// モックの設定やテストデータを保持するフィールドを追加できます
	validTokens map[string]map[string]interface{}
}

// NewMockAuth は新しい MockAuthImpl を作成します
func NewMockAuth() (AuthService, error) {
	// テスト用のトークンとクレームを初期化
	validTokens := make(map[string]map[string]interface{})
	
	// テスト用のトークンを追加
	validTokens["test-token"] = map[string]interface{}{
		"user_id":    "mock-user-123",
		"email":      "test@example.com",
		"name":       "Test User",
		"exp":        time.Now().Add(1 * time.Hour).Unix(),
		"auth_time":  time.Now().Unix(),
		"iat":        time.Now().Unix(),
		"iss":        "mock-auth-service",
		"sub":        "mock-user-123",
	}
	
	return &MockAuthImpl{
		validTokens: validTokens,
	}, nil
}

// VerifyToken はトークンを検証し、クレームを返します
// AuthService インターフェースの実装
func (m *MockAuthImpl) VerifyToken(ctx context.Context, idToken string) (map[string]interface{}, error) {
	// モックの検証ロジック
	claims, exists := m.validTokens[idToken]
	if !exists {
		return nil, errors.New("無効なトークンです")
	}
	
	// トークンの有効期限をチェック
	exp, ok := claims["exp"].(int64)
	if !ok || exp < time.Now().Unix() {
		return nil, errors.New("トークンの有効期限が切れています")
	}
	
	return claims, nil
}

// RegisterToken はテスト用のトークンを登録するヘルパーメソッドです
func (m *MockAuthImpl) RegisterToken(token string, claims map[string]interface{}) {
	m.validTokens[token] = claims
}
