package services

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/baseapp/models" // モデルをインポート
    "gorm.io/driver/postgres"   // PostgreSQL ドライバをインポート
    "gorm.io/gorm"              // GORM をインポート
    "gorm.io/gorm/logger"       // GORM ロガーをインポート
)

var (
    // DB はグローバルなデータベースインスタンス
    DB *gorm.DB
)

// InitDB はデータベース接続を初期化します
func InitDB() error {
    // 環境変数からデータベース接続情報を取得
    dbHost := os.Getenv("DB_HOST")
    if dbHost == "" {
        dbHost = "localhost" // デフォルトのホスト名
    }

    dbPort := os.Getenv("DB_PORT")
    if dbPort == "" {
        dbPort = "5432" // デフォルトのポート番号
    }

    dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        dbName = "youtubeapp" // デフォルトのデータベース名
    }

    dbUser := os.Getenv("DB_USER")
    if dbUser == "" {
        dbUser = "postgres" // デフォルトのユーザー名
    }

    dbPassword := os.Getenv("DB_PASSWORD")
    if dbPassword == "" {
        dbPassword = "postgres" // デフォルトのパスワード
    }

    // DSN（データソース名）文字列を作成
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)

    // GORM ロガーを設定
    newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags), // ログ出力先を標準出力に設定
        logger.Config{
            SlowThreshold:             time.Second, // 遅いクエリの閾値
            LogLevel:                  logger.Info,  // ログレベルを情報に設定
            IgnoreRecordNotFoundError: true,         // レコードが見つからないエラーを無視
            Colorful:                  true,         // カラフルなログ出力
        },
    )

    // データベースに接続
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: newLogger, // 設定したロガーを使用
    })
    if err != nil {
        return fmt.Errorf("データベースへの接続に失敗しました: %v", err)
    }

    // スキーマを自動マイグレーション
    err = DB.AutoMigrate(&models.UserProfile{})
    if err != nil {
        return fmt.Errorf("データベーススキーマのマイグレーションに失敗しました: %v", err)
    }

    log.Println("データベースに接続し、マイグレーションが成功しました")
    return nil
}
