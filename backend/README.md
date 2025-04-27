# ユーザープロファイルAPIのGoバックエンド

これは、元のPython実装を置き換える形でのユーザープロファイルAPIバックエンドのGo実装です。

## 機能

- 柔軟な認証サービスアーキテクチャ（Firebase、モックなど）
- ユーザープロフィール管理
- PostgreSQLデータベース統合
- RESTful APIエンドポイント

## 前提条件

- Go 1.21以上
- PostgreSQLデータベース
- Firebaseプロジェクト

## 認証サービス

このアプリケーションは、異なる認証プロバイダーを簡単に切り替えられる柔軟な認証アーキテクチャを採用しています。現在サポートされている認証サービス：

- **Firebase認証** - デフォルトの認証プロバイダー
- **モック認証** - テスト環境用の簡易認証

認証サービスは環境変数 `AUTH_SERVICE_TYPE` で設定できます：
- `firebase` - Firebase認証を使用（デフォルト）
- `mock` - モック認証を使用（テスト用）

新しい認証サービスを追加するには：
1. `AuthService` インターフェースを実装する新しいサービスを作成
2. `auth_factory.go` に新しいサービスタイプを追加
3. 環境変数で新しいサービスを選択

## 環境変数

`.env.example`ファイルを`.env`にコピーし、値を入力してください：

```bash
cp .env.example .env
```

## インストール

1. 依存関係をインストールします：

```bash
go mod download
```

2. アプリケーションをビルドします：

```bash
go build -o app
```

## アプリケーションの実行

```bash
./app
```

または、Goを直接使用して実行します：

```bash
go run main.go
```

## APIエンドポイント

- `GET /` - API情報
- `POST /api/auth/verify` - 認証トークンの検証
- `GET /api/profile` - ユーザープロフィールの取得
- `PUT /api/profile` - ユーザープロフィールの更新
