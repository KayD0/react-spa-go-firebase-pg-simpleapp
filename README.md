# React SPA Go Postgres アプリケーション

Firebase認証を使用したReact SPAとGoバックエンドのウェブアプリケーション

## 目次

- [概要](#概要)
- [主な機能](#主な機能)
- [技術スタック](#技術スタック)
- [システムアーキテクチャ](#システムアーキテクチャ)
- [セットアップ手順](#セットアップ手順)
  - [前提条件](#前提条件)
  - [APIキーの取得](#apiキーの取得)
  - [バックエンドのセットアップ](#バックエンドのセットアップ)
  - [フロントエンドのセットアップ](#フロントエンドのセットアップ)
- [開発ガイド](#開発ガイド)
- [使用方法](#使用方法)
- [プロジェクト構造](#プロジェクト構造)
- [トラブルシューティング](#トラブルシューティング)
- [ライセンス](#ライセンス)

## 概要

このアプリケーションは、ユーザーがFirebase認証でログインし、プロフィール情報を管理できるシンプルなウェブアプリケーションです。バックエンドはGoで実装され、フロントエンドはReact（Vite）とBootstrapで構築されています。データベースにはPostgreSQLを使用しています。

## 主な機能

- **ユーザー認証**: Firebase Authenticationを使用したユーザー登録・ログイン機能
- **認証トークン検証**: バックエンドでのFirebase IDトークン検証
- **レスポンシブUI**: モバイルデバイスにも対応したユーザーインターフェース
- **プロフィール管理**: ユーザープロフィール情報の管理機能

## 技術スタック

### フロントエンド

- **React**: UIライブラリ
- **Vite**: モジュールバンドラー・開発サーバー
- **Bootstrap 5**: UIコンポーネントとレスポンシブデザイン
- **Firebase SDK**: 認証とIDトークン管理
- **SCSS**: スタイリングの拡張機能

### バックエンド

- **Go**: サーバーサイド言語
- **PostgreSQL**: データベース
- **Firebase Admin SDK**: IDトークン検証

### クラウドサービス

- **Firebase Authentication**: ユーザー認証

## システムアーキテクチャ

このアプリケーションは以下のコンポーネントで構成されています：

1. **フロントエンド（React SPA）**: ユーザーインターフェースを提供し、Firebase SDKを使用して認証を処理します。
2. **バックエンド（Go API）**: フロントエンドからのリクエストを処理し、データベースとの通信を行います。
3. **データベース（PostgreSQL）**: ユーザープロフィールなどのデータを永続化します。
4. **Firebase Authentication**: ユーザー認証サービスを提供します。

## セットアップ手順

### 前提条件

- Go 1.16以上
- Node.js 14以上
- npm または yarn
- PostgreSQL
- Google Cloudアカウント
- Firebaseプロジェクト

#### Firebase設定

1. [Firebase Console](https://console.firebase.google.com/)にアクセス
2. プロジェクトを作成または選択
3. Authentication機能を有効化し、メール/パスワード認証を設定
4. プロジェクト設定からWebアプリを追加
5. 提供されるFirebase設定オブジェクトをコピー
6. プロジェクト設定 > サービスアカウントから新しい秘密鍵を生成

### バックエンドのセットアップ

1. リポジトリをクローン:
   ```bash
   git clone <repository-url>
   cd <repository-name>
   ```

2. バックエンドディレクトリに移動:
   ```bash
   cd backend
   ```

3. 依存関係をインストール:
   ```bash
   go mod download
   ```

4. `.env.example`をコピーして`.env`ファイルを作成:
   ```bash
   cp .env.example .env
   ```

5. `.env`ファイルを編集し、必要な環境変数を設定:
   ```
   # データベース設定
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=yourpassword
   DB_NAME=yourdbname

   # CORS設定
   CORS_ORIGIN=http://localhost:5173

   # Firebase設定
   FIREBASE_PROJECT_ID=あなたのfirebaseプロジェクトID
   FIREBASE_PRIVATE_KEY_ID=あなたのprivate_key_id
   FIREBASE_PRIVATE_KEY=あなたのprivate_key
   FIREBASE_CLIENT_EMAIL=あなたのclient_email
   FIREBASE_CLIENT_ID=あなたのclient_id
   FIREBASE_AUTH_URI=https://accounts.google.com/o/oauth2/auth
   FIREBASE_TOKEN_URI=https://oauth2.googleapis.com/token
   FIREBASE_AUTH_PROVIDER_X509_CERT_URL=https://www.googleapis.com/oauth2/v1/certs
   FIREBASE_CLIENT_X509_CERT_URL=あなたのclient_x509_cert_url
   ```

6. バックエンドサーバーを起動:
   ```bash
   # Windows
   setup.bat
   # または
   go run main.go

   # macOS/Linux
   ./setup.sh
   # または
   go run main.go
   ```
   サーバーは`http://localhost:8080`で実行されます。

### フロントエンドのセットアップ

1. フロントエンドディレクトリに移動:
   ```bash
   cd front
   ```

2. 依存関係をインストール:
   ```bash
   npm install
   # または
   yarn
   ```

3. `.env.example`をコピーして`.env`ファイルを作成:
   ```bash
   cp .env.example .env
   ```

4. `.env`ファイルを編集し、Firebase設定を追加:
   ```
   VITE_FIREBASE_API_KEY=あなたのapiKey
   VITE_FIREBASE_AUTH_DOMAIN=あなたのauthDomain
   VITE_FIREBASE_PROJECT_ID=あなたのprojectId
   VITE_FIREBASE_STORAGE_BUCKET=あなたのstorageBucket
   VITE_FIREBASE_MESSAGING_SENDER_ID=あなたのmessagingSenderId
   VITE_FIREBASE_APP_ID=あなたのappId
   
   VITE_API_BASE_URL=http://localhost:8080
   ```

5. 開発サーバーを起動:
   ```bash
   npm run dev
   # または
   yarn dev
   ```
   フロントエンドは`http://localhost:5173`で実行されます。

## 開発ガイド

### バックエンド開発

- **コントローラーの追加**:
  1. `backend/controllers/`に新しいコントローラーファイルを作成
  2. 必要なルートとハンドラーを実装
  3. `controllers/routes.go`にルートを登録

- **サービスの追加**:
  1. `backend/services/`に新しいサービスファイルを作成
  2. 必要な関数とメソッドを実装
  3. コントローラーからサービスをインポートして使用

- **モデルの追加**:
  1. `backend/models/`に新しいモデルファイルを作成
  2. 必要な構造体とメソッドを実装

- **テスト**:
  1. `test_api.go`と`test_auth.go`を参考にテストを作成
  2. `go test`を使用してテストを実行

### フロントエンド開発

- **コンポーネントの追加**:
  1. `front/src/components/`に新しいコンポーネントファイルを作成
  2. 必要なJSXとロジックを実装
  3. 他のコンポーネントやページからインポートして使用

- **ページの追加**:
  1. `front/src/pages/`に新しいページファイルを作成
  2. App.jsxにルートを追加

- **サービスの追加**:
  1. `front/src/services/`または`front/src/js/services/`に新しいサービスファイルを作成
  2. APIとの通信ロジックを実装

- **スタイルの追加**:
  1. `front/src/css/`にスタイルを追加
  2. SCSSを使用する場合は`main.scss`に追加

## 使用方法

1. アプリケーションにアクセス: `http://localhost:5173`
2. ログインまたは新規ユーザー登録
3. ホームページでアプリケーションの機能を利用
4. プロフィールページでユーザー情報を管理
5. 認証テストページで認証状態とトークンを確認可能

## プロジェクト構造

```
プロジェクトルート/
├── backend/                  # バックエンドアプリケーション
│   ├── controllers/          # コントローラー
│   │   ├── auth_controller.go
│   │   ├── main_controller.go
│   │   ├── profile_controller.go
│   │   └── routes.go
│   ├── middleware/           # ミドルウェア
│   │   └── auth_middleware.go
│   ├── models/               # データモデル
│   │   └── user_profile.go
│   ├── services/             # サービス
│   │   ├── auth_service.go
│   │   └── db_service.go
│   ├── main.go               # メインGoアプリケーション
│   ├── go.mod                # Goモジュール定義
│   ├── go.sum                # Goモジュール依存関係
│   ├── setup.bat             # Windowsセットアップスクリプト
│   ├── setup.sh              # Unix/Linuxセットアップスクリプト
│   ├── test_api.go           # APIテスト
│   ├── test_auth.go          # 認証テスト
│   └── .env.example          # 環境変数のサンプル
│
├── front/                    # フロントエンドアプリケーション
│   ├── src/
│   │   ├── components/       # 再利用可能なコンポーネント
│   │   │   └── NavBar.jsx
│   │   ├── contexts/         # Reactコンテキスト
│   │   │   └── AuthContext.jsx
│   │   ├── css/              # スタイルシート
│   │   │   ├── main.scss     # メインSCSSファイル
│   │   │   └── style.css     # カスタムスタイル
│   │   ├── js/               # JavaScriptファイル
│   │   │   ├── services/     # APIサービス
│   │   │   │   ├── auth-api-service.js
│   │   │   │   ├── firebase.js
│   │   │   │   └── profile-api-service.js
│   │   │   ├── utils/        # ユーティリティ関数
│   │   │   │   └── auth-test.js
│   │   │   ├── bootstrap.js  # Bootstrap初期化
│   │   │   └── main.js       # メインJavaScriptファイル
│   │   ├── pages/            # ページコンポーネント
│   │   │   ├── AboutPage.jsx
│   │   │   ├── AuthTestPage.jsx
│   │   │   ├── ContactPage.jsx
│   │   │   ├── HomePage.jsx
│   │   │   ├── LoginPage.jsx
│   │   │   ├── ProfilePage.jsx
│   │   │   └── RegisterPage.jsx
│   │   ├── services/         # サービス
│   │   │   └── api.js
│   │   ├── App.jsx           # メインReactコンポーネント
│   │   └── main.jsx          # Reactエントリーポイント
│   ├── public/               # 静的ファイル
│   │   └── _redirects        # Netlify用リダイレクト設定
│   ├── etc/                  # その他の設定ファイル
│   ├── index.html            # メインHTMLファイル
│   ├── package.json          # npm依存関係
│   ├── package-lock.json     # npm依存関係ロック
│   ├── vite.config.js        # Vite設定
│   └── .env.example          # 環境変数のサンプル
│
└── LICENSE                   # ライセンスファイル
```

## トラブルシューティング

### Firebase認証エラー

- **auth/configuration-not-found エラー**:
  1. `.env`ファイルが正しく設定されているか確認
  2. Firebase Consoleで認証機能が有効になっているか確認
  3. ブラウザのキャッシュとCookieをクリア
  4. 開発サーバーを再起動

- **トークン検証エラー**:
  1. バックエンドとフロントエンドのFirebase設定が同じプロジェクトを指しているか確認
  2. Firebase Admin SDKの秘密鍵が正しく設定されているか確認

### API接続エラー

- **CORS エラー**:
  1. バックエンドの`CORS_ORIGIN`設定がフロントエンドのURLと一致しているか確認
  2. フロントエンドの`VITE_API_BASE_URL`がバックエンドのURLと一致しているか確認

### データベース接続エラー

- **PostgreSQL接続エラー**:
  1. PostgreSQLサービスが実行されているか確認
  2. データベース接続情報（ホスト、ポート、ユーザー名、パスワード、データベース名）が正しいか確認
  3. データベースユーザーに適切な権限があるか確認

## ライセンス

このプロジェクトはMITライセンスの下で公開されています。詳細は[LICENSE](LICENSE)ファイルを参照してください。
