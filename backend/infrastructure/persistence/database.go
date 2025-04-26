package persistence

import (
    "fmt"
    "log"
    "os"
    "time"

    "gorm.io/driver/postgres" // PostgreSQL ドライバをインポート
    "gorm.io/gorm"            // GORM をインポート
    "gorm.io/gorm/logger"     // GORM ロガーをインポート
)

// Database はデータベースへのアクセスを提供します
type Database struct {
    DB *gorm.DB // GORM データベースインスタンス
}

// NewDatabase は新しい Database を作成します
func NewDatabase() (*Database, error) {
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
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: newLogger, // 設定したロガーを使用
    })
    if err != nil {
        return nil, fmt.Errorf("データベースへの接続に失敗しました: %v", err)
    }

    return &Database{
        DB: db, // データベースインスタンスを返す
    }, nil
}

// AutoMigrate は指定されたモデルの自動マイグレーションを実行します
func (d *Database) AutoMigrate(models ...interface{}) error {
    return d.DB.AutoMigrate(models...) // GORM の AutoMigrate メソッドを呼び出す
}
