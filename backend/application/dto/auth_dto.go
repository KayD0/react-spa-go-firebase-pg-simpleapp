package dto

// AuthResponse は認証レスポンスデータを表す構造体です
type AuthResponse struct {
    Authenticated bool                   `json:"authenticated"` // 認証成功フラグ
    User          map[string]interface{} `json:"user"`         // ユーザー情報を含むマップ
}

// NewAuthResponse は新しい AuthResponse を作成するコンストラクタです
func NewAuthResponse(claims map[string]interface{}) *AuthResponse {
    return &AuthResponse{
        Authenticated: true, // 認証が成功したことを示す
        User: map[string]interface{}{
            "uid":            claims["uid"],            // ユーザーID
            "user_id":        claims["user_id"],        // ユーザーの一意の識別子
            "email":          claims["email"],          // ユーザーのメールアドレス
            "email_verified": claims["email_verified"], // メールアドレスの確認状況
            "auth_time":      claims["auth_time"],      // 認証時刻
        },
    }
}
