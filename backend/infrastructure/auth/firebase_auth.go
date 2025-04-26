package auth

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "strings"

    firebase "firebase.google.com/go/v4" // Firebase SDK をインポート
    "firebase.google.com/go/v4/auth"       // Firebase 認証をインポート
    "google.golang.org/api/option"         // Google API オプションをインポート
)

// FirebaseAuth は Firebase 認証のインターフェースを定義します
type FirebaseAuth interface {
    // VerifyIDToken は Firebase ID トークンを検証し、デコードされたトークンを返します
    VerifyIDToken(ctx context.Context, idToken string) (map[string]interface{}, error)
}

// FirebaseAuthImpl は FirebaseAuth インターフェースを実装します
type FirebaseAuthImpl struct {
    app        *firebase.App // Firebase アプリのインスタンス
    authClient *auth.Client  // Firebase 認証クライアント
}

// AuthError は認証エラーを表します
type AuthError struct {
    Message string // エラーメッセージ
}

// Error はエラーメッセージを返します
func (e *AuthError) Error() string {
    return e.Message
}

// NewFirebaseAuth は新しい FirebaseAuthImpl を作成します
func NewFirebaseAuth() (FirebaseAuth, error) {
    ctx := context.Background()
    var app *firebase.App
    var err error

    // 環境変数にサービスアカウントの情報が提供されているか確認
    if os.Getenv("FIREBASE_PROJECT_ID") != "" {
        // 環境変数からサービスアカウントの資格情報を作成
        serviceAccount := map[string]interface{}{
            "type":                        "service_account",
            "project_id":                  os.Getenv("FIREBASE_PROJECT_ID"),
            "private_key_id":              os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
            "private_key":                 strings.ReplaceAll(os.Getenv("FIREBASE_PRIVATE_KEY"), "\\n", "\n"),
            "client_email":                os.Getenv("FIREBASE_CLIENT_EMAIL"),
            "client_id":                   os.Getenv("FIREBASE_CLIENT_ID"),
            "auth_uri":                    os.Getenv("FIREBASE_AUTH_URI"),
            "token_uri":                   os.Getenv("FIREBASE_TOKEN_URI"),
            "auth_provider_x509_cert_url": os.Getenv("FIREBASE_AUTH_PROVIDER_X509_CERT_URL"),
            "client_x509_cert_url":        os.Getenv("FIREBASE_CLIENT_X509_CERT_URL"),
        }

        // サービスアカウントを JSON 形式にマシュアル
        serviceAccountJSON, err := json.Marshal(serviceAccount)
        if err != nil {
            return nil, fmt.Errorf("サービスアカウントのマシュアルに失敗しました: %v", err)
        }

        // サービスアカウントで Firebase を初期化
        opt := option.WithCredentialsJSON(serviceAccountJSON)
        app, err = firebase.NewApp(ctx, nil, opt)
        if err != nil {
            return nil, fmt.Errorf("サービスアカウントで Firebase を初期化中にエラーが発生しました: %v", err)
        }
        log.Println("サービスアカウント資格情報で Firebase Admin SDK が初期化されました")
    } else {
        // アプリケーションのデフォルト資格情報で Firebase を初期化
        app, err = firebase.NewApp(ctx, nil)
        if err != nil {
            return nil, fmt.Errorf("デフォルト資格情報で Firebase を初期化中にエラーが発生しました: %v", err)
        }
        log.Println("アプリケーションのデフォルト資格情報で Firebase Admin SDK が初期化されました")
    }

    // Auth クライアントを取得
    client, err := app.Auth(ctx)
    if err != nil {
        return nil, fmt.Errorf("Auth クライアントの取得中にエラーが発生しました: %v", err)
    }

    return &FirebaseAuthImpl{
        app:       app,
        authClient: client,
    }, nil
}

// VerifyIDToken は Firebase ID トークンを検証し、デコードされたトークンを返します
func (f *FirebaseAuthImpl) VerifyIDToken(ctx context.Context, idToken string) (map[string]interface{}, error) {
    // トークンを検証
    token, err := f.authClient.VerifyIDToken(ctx, idToken)
    if err != nil {
        return nil, fmt.Errorf("ID トークンの検証中にエラーが発生しました: %v", err)
    }

    // トークンのクレームを返す
    return token.Claims, nil
}
