package middleware

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/baseapp/services"
)

// AuthRequired はリクエストに有効な Firebase ID トークンが含まれているかをチェックするミドルウェアです
func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Authorization ヘッダーを取得
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            // ヘッダーが存在しない場合、401 Unauthorized を返す
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorizationヘッダーがありません"})
            c.Abort() // リクエストの処理を中止
            return
        }

        // トークンを抽出（'Bearer ' プレフィックスを削除）
        idToken := authHeader
        if strings.HasPrefix(authHeader, "Bearer ") {
            idToken = strings.TrimPrefix(authHeader, "Bearer ")
        }

        // トークンを検証
        token, err := services.VerifyIDToken(idToken)
        if err != nil {
            // トークンが無効な場合、401 Unauthorized を返す
            c.JSON(http.StatusUnauthorized, gin.H{"error": "無効な認証トークンです: " + err.Error()})
            c.Abort() // リクエストの処理を中止
            return
        }

        // トークンをコンテキストに保存
        c.Set("user", token)
        c.Next() // 次のハンドラーを呼び出す
    }
}

// GetUserIDFromToken はコンテキスト内のトークンからユーザーIDを抽出します
func GetUserIDFromToken(c *gin.Context) (string, error) {
    // コンテキストからユーザーを取得
    user, exists := c.Get("user")
    if !exists {
        return "", &services.AuthError{Message: "認証されていません。AuthRequiredミドルウェアを使用してください。"}
    }

    // ユーザーIDを抽出
    token := user.(map[string]interface{})

    uid, ok := token["user_id"].(string)
    if !ok {
        return "", &services.AuthError{Message: "ユーザーIDが見つかりません"}
    }

    return uid, nil // ユーザーIDを返す
}
