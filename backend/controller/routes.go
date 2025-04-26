package controller

import (
    "github.com/gin-gonic/gin"
)

// Router はすべてのルートを処理する構造体です
type Router struct {
    engine         *gin.Engine       // Gin エンジンのインスタンス
    authController *AuthController   // 認証コントローラのインスタンス
    userController *UserController   // ユーザコントローラのインスタンス
    mainController *MainController   // メインコントローラのインスタンス
}

// NewRouter は新しい Router を作成するコンストラクタです
func NewRouter(
    engine *gin.Engine,
    authController *AuthController,
    userController *UserController,
    mainController *MainController,
) *Router {
    return &Router{
        engine:         engine,         // 引数として渡された Gin エンジンをフィールドに設定
        authController: authController, // 認証コントローラをフィールドに設定
        userController: userController, // ユーザコントローラをフィールドに設定
        mainController: mainController, // メインコントローラをフィールドに設定
    }
}

// SetupRoutes はすべてのルートを設定するメソッドです
func (r *Router) SetupRoutes() {
    // メインルート
    r.engine.GET("/", r.mainController.Index)

    // 認証ルート
    authGroup := r.engine.Group("/api/auth")
    {
        // 認証検証のための POST ルートを設定
        authGroup.POST("/verify", r.AuthMiddleware(), r.authController.VerifyAuth)
    }

    // プロフィールルート
    profileGroup := r.engine.Group("/api")
    {
        // プロフィール取得のための GET ルートを設定
        profileGroup.GET("/profile", r.AuthMiddleware(), r.userController.GetProfile)
        // プロフィール更新のための PUT ルートを設定
        profileGroup.PUT("/profile", r.AuthMiddleware(), r.userController.UpdateProfile)
    }
}

// AuthMiddleware はリクエストに有効な認証トークンがあるかをチェックするミドルウェアです
func (r *Router) AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Authorization ヘッダーを取得
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            // ヘッダーが空の場合、401 Unauthorized を返す
            c.JSON(401, gin.H{"error": "Authorizationヘッダーがありません"})
            c.Abort() // リクエストの処理を中止
            return
        }

        c.Next() // 次のハンドラに処理を渡す
    }
}
