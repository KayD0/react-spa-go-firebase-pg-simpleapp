package usecase

import (
    "context"
    "errors"

    "github.com/baseapp/application/dto"
    "github.com/baseapp/domain/repository"
)

// AuthUseCase は認証関連のビジネスロジックを実装する構造体です
type AuthUseCase struct {
    authRepo repository.AuthRepository
}

// NewAuthUseCase は新しい AuthUseCase を作成するコンストラクタです
func NewAuthUseCase(authRepo repository.AuthRepository) *AuthUseCase {
    return &AuthUseCase{
        authRepo: authRepo,
    }
}

// VerifyToken は認証トークンを検証するメソッドです
func (uc *AuthUseCase) VerifyToken(ctx context.Context, token string) (*dto.AuthResponse, error) {
    // リポジトリを使用してトークンを検証し、クレームを取得    
    claims, err := uc.authRepo.VerifyToken(ctx, token)
    if err != nil {
        return nil, errors.New("無効な認証トークンです: " + err.Error())
    }
    
    return dto.NewAuthResponse(claims), nil
}

// GetUserIDFromClaims extracts the user ID from token claims
func (uc *AuthUseCase) GetUserIDFromClaims(claims map[string]interface{}) (string, error) {
    uid, ok := claims["user_id"].(string)
    if !ok {
        return "", errors.New("ユーザーIDが見つかりません")
    }
    
    return uid, nil
}
