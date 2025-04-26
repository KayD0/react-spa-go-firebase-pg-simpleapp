package controller

import (
    "net/http" // HTTP関連のパッケージをインポート

    "github.com/gin-gonic/gin" // Ginフレームワークをインポート
)

// MainController は一般的なHTTPリクエストを処理するためのコントローラです
type MainController struct {
}

// NewMainController は新しい MainController を作成します
func NewMainController() *MainController {
    return &MainController{} // MainController の新しいインスタンスを返す
}

// Index はインデックスルートを処理します
func (c *MainController) Index(ctx *gin.Context) {
    // JSON形式でレスポンスを返す
    ctx.JSON(http.StatusOK, gin.H{
        "message": "ユーザープロフィールAPIが実行中です", // APIの稼働メッセージ
        "endpoints": gin.H{ // 利用可能なエンドポイントのリスト
            "auth_verify": "/api/auth/verify (Authorizationヘッダーを持つPOST)", // 認証トークン検証エンドポイント
            "profile_get": "/api/profile/ (Authorizationヘッダーを持つGET)", // プロフィール取得エンドポイント
            "profile_update": "/api/profile/ (JSONボディとAuthorizationヘッダーを持つPUT)", // プロフィール更新エンドポイント
        },
    })
}
