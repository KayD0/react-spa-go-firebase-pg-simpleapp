package controller

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/baseapp/application/usecase"
)

// AuthController は認証関連のHTTPリクエストを処理する構造体です
type AuthController struct {
    authUseCase *usecase.AuthUseCase // 認証ユースケースのインスタンス
}

// NewAuthController は新しい AuthController を作成するコンストラクタです
func NewAuthController(authUseCase *usecase.AuthUseCase) *AuthController {
    return &AuthController{
        authUseCase: authUseCase, // 引数として渡されたユースケースをフィールドに設定
    }
}

// VerifyAuth は認証検証ルートを処理するメソッドです
func (c *AuthController) VerifyAuth(ctx *gin.Context) {
    // Authorization ヘッダーを取得
    authHeader := ctx.GetHeader("Authorization")
    if authHeader == "" {
        // ヘッダーが空の場合、401 Unauthorized を返す
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorizationヘッダーがありません"})
        return
    }

    // トークンを抽出（'Bearer ' プレフィックスを削除）
    idToken := authHeader
    if strings.HasPrefix(authHeader, "Bearer ") {
        idToken = strings.TrimPrefix(authHeader, "Bearer ")
    }

    // トークンを検証
    response, err := c.authUseCase.VerifyToken(ctx.Request.Context(), idToken)
    if err != nil {
        // 検証に失敗した場合、401 Unauthorized を返す
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    // 検証成功時、レスポンスを返す
    ctx.JSON(http.StatusOK, response)
}

// GetUserIDFromToken はコンテキスト内のトークンからユーザーIDを抽出するメソッドです
func (c *AuthController) GetUserIDFromToken(ctx *gin.Context) (string, error) {
    // Authorization ヘッダーを取得
    authHeader := ctx.GetHeader("Authorization")
    if authHeader == "" {
        return "", &AuthError{Message: "Authorizationヘッダーがありません"} // エラーを返す
    }

    // トークンを抽出（'Bearer ' プレフィックスを削除）
    idToken := authHeader
    if strings.HasPrefix(authHeader, "Bearer ") {
        idToken = strings.TrimPrefix(authHeader, "Bearer ")
    }

    // トークンを検証
    response, err := c.authUseCase.VerifyToken(ctx.Request.Context(), idToken)
    if err != nil {
        return "", &AuthError{Message: err.Error()} // エラーを返す
    }

    // ユーザーIDを抽出
    uid, err := c.authUseCase.GetUserIDFromClaims(response.User)
    if err != nil {
        return "", &AuthError{Message: err.Error()} // エラーを返す
    }

    return uid, nil // ユーザーIDを返す
}

// AuthError は認証エラーを表す構造体です
type AuthError struct {
    Message string // エラーメッセージ
}

// Error メソッドはエラーメッセージを返します
func (e *AuthError) Error() string {
    return e.Message
}