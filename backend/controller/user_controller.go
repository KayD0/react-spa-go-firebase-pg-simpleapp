package controller

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/baseapp/application/dto"
    "github.com/baseapp/application/usecase"
)

// UserController はユーザー関連のHTTPリクエストを処理する構造体です
type UserController struct {
    userUseCase    *usecase.UserUseCase // ユーザー用ユースケースのインスタンス
    authController  *AuthController      // 認証コントローラのインスタンス
}

// NewUserController は新しい UserController を作成するコンストラクタです
func NewUserController(userUseCase *usecase.UserUseCase, authController *AuthController) *UserController {
    return &UserController{
        userUseCase: userUseCase,      // 引数として渡されたユーザー用ユースケースをフィールドに設定
        authController: authController,  // 引数として渡された認証コントローラをフィールドに設定
    }
}

// GetProfile はプロフィール取得ルートを処理するメソッドです
func (c *UserController) GetProfile(ctx *gin.Context) {
    // トークンからユーザーIDを取得
    firebaseUID, err := c.authController.GetUserIDFromToken(ctx)
    if err != nil {
        // ユーザーIDが取得できない場合、401 Unauthorized を返す
        ctx.JSON(http.StatusUnauthorized, gin.H{
            "success": false,
            "error":   err.Error(),
        })
        return
    }

    // プロフィールを取得
    profile, err := c.userUseCase.GetProfile(ctx.Request.Context(), firebaseUID)
    if err != nil {
        // プロフィール取得に失敗した場合、500 Internal Server Error を返す
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "error":   err.Error(),
        })
        return
    }

    // プロフィールを返す
    ctx.JSON(http.StatusOK, gin.H{
        "success": true,
        "profile": profile,
    })
}

// UpdateProfile はプロフィール更新ルートを処理するメソッドです
func (c *UserController) UpdateProfile(ctx *gin.Context) {
    // トークンからユーザーIDを取得
    firebaseUID, err := c.authController.GetUserIDFromToken(ctx)
    if err != nil {
        // ユーザーIDが取得できない場合、401 Unauthorized を返す
        ctx.JSON(http.StatusUnauthorized, gin.H{
            "success": false,
            "error":   err.Error(),
        })
        return
    }

    // リクエストボディをパース
    var requestBody dto.UserUpdateRequest
    if err := ctx.ShouldBindJSON(&requestBody); err != nil {
        // リクエストボディが無効な場合、400 Bad Request を返す
        ctx.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "無効なリクエストボディ: " + err.Error(),
        })
        return
    }

    // プロフィールを更新
    profile, err := c.userUseCase.UpdateProfile(ctx.Request.Context(), firebaseUID, &requestBody)
    if err != nil {
        // プロフィール更新に失敗した場合、500 Internal Server Error を返す
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "error":   err.Error(),
        })
        return
    }

    // 更新されたプロフィールを返す
    ctx.JSON(http.StatusOK, gin.H{
        "success":  true,
        "profile":  profile,
        "message":  "プロフィールが更新されました",
    })
}
