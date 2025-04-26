package usecase

import (
    "context"
    "errors"

    "github.com/baseapp/application/dto"
    "github.com/baseapp/domain/repository"
)

// AuthUseCase は認証関連のビジネスロジックを実装する構造体です
type AuthUseCase struct {
    authRepo repository.AuthRepository // 認証リポジトリのインターフェース
}

// NewAuthUseCase は新しい AuthUseCase を作成するコンストラクタです
func NewAuthUseCase(authRepo repository.AuthRepository) *AuthUseCase {
    return &AuthUseCase{
        authRepo: authRepo, // 引数として渡されたリポジトリをフィールドに設定
    }
}

// VerifyToken は認証トークンを検証するメソッドです
func (uc *AuthUseCase) VerifyToken(ctx context.Context, token string) (*dto.AuthResponse, error) {
    // リポジトリを使用してトークンを検証し、クレームを取得
    claims, err := uc.authRepo.VerifyToken(ctx, token)
    if err != nil {
        // トークンが無効な場合、エラーメッセージを返す
        return nil, errors.New("無効な認証トークンです: " + err.Error())
    }
    
    // クレームから新しい AuthResponse を作成して返す
    return dto.NewAuthResponse(claims), nil
}

// GetUserIDFromClaims はトークンのクレームからユーザーIDを抽出するメソッドです
func (uc *AuthUseCase) GetUserIDFromClaims(claims map[string]interface{}) (string, error) {
    // クレームから "user_id" を取得し、文字列型にアサーション
    uid, ok := claims["user_id"].(string)
    if !ok {
        // ユーザーIDが見つからない場合、エラーメッセージを返す
        return "", errors.New("ユーザーIDが見つかりません")
    }
    
    // ユーザーIDを返す
    return uid, nil
}