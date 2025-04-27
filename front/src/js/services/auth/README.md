# 認証アーキテクチャ

このディレクトリには、アプリケーションの認証アーキテクチャが含まれています。このアーキテクチャは柔軟かつ拡張可能に設計されており、異なる認証プロバイダー間で簡単に切り替えることができます。

## 概要

認証アーキテクチャはプロバイダーパターンに従っており、異なる認証プロバイダーが共通のインターフェースを実装します。これにより、アプリケーションのコアロジックを変更することなく、異なる認証プロバイダーを使用できます。

### 主なコンポーネント

1. **AuthInterface**: すべての認証プロバイダーが実装すべきインターフェースを定義する抽象クラス。
2. **AuthFactory**: 設定に基づいて適切な認証プロバイダーを作成するファクトリー。
3. **AuthService**: 認証操作のための統一インターフェースを提供するサービス。
4. **Providers**: 異なる認証プロバイダーのための AuthInterface の具体的な実装。

## 使用方法

### 認証サービスの使用

`AuthService` は認証操作のメインエントリーポイントです。これはシングルトンインスタンスであり、アプリケーション全体でインポートして使用できます。

```javascript
import authService from './auth/auth-service';

// 認証サービスを初期化
await authService.initialize();

// 新しいユーザーを登録
const { user, error } = await authService.signUp(email, password);

// 既存のユーザーでサインイン
const { user, error } = await authService.signIn(email, password);

// サインアウト
await authService.signOut();

// 現在のユーザーを取得
const currentUser = authService.getCurrentUser();

// ユーザーが認証されているか確認
const isAuthenticated = authService.isAuthenticated();

// 認証トークンを取得
const token = await authService.getAuthToken();

// 認証状態の変更を購読
const unsubscribe = authService.onAuthStateChanged((user) => {
    console.log('認証状態が変更されました:', user);
});

// 必要なくなったら購読を解除
unsubscribe();
```

### 認証プロバイダーの設定

認証プロバイダーは環境変数を使用して設定されます。`VITE_AUTH_PROVIDER` 環境変数が使用するプロバイダーを決定します。

```
# .env ファイル
VITE_AUTH_PROVIDER=firebase  # オプション: firebase, mock
```

## 新しい認証プロバイダーの追加

新しい認証プロバイダーを追加するには、以下の手順を実行します：

1. プロバイダー用の新しいファイルを作成（例: `my-auth-provider.js`）。
2. プロバイダー固有のロジックで `AuthInterface` クラスを実装。
3. `auth-factory.js` の `AUTH_PROVIDERS` オブジェクトにプロバイダーを追加。
4. `auth-factory.js` の `createAuthProvider` 関数にプロバイダーを作成するケースを追加。
5. `.env.example` ファイルを更新してプロバイダーをオプションとして追加。

### 例: 新しいプロバイダーの追加

```javascript
// my-auth-provider.js
import { AuthInterface } from './auth-interface';

export class MyAuthProvider extends AuthInterface {
    constructor(config) {
        super();
        this.config = config;
        // プロバイダーを初期化
    }

    async initialize() {
        // 認証プロバイダーを初期化
        return Promise.resolve();
    }

    async signUp(email, password) {
        // サインアップロジックを実装
        return { user: null, error: null };
    }

    async signIn(email, password) {
        // サインインロジックを実装
        return { user: null, error: null };
    }

    async signOut() {
        // サインアウトロジックを実装
        return { error: null };
    }

    getCurrentUser() {
        // 現在のユーザー取得ロジックを実装
        return null;
    }

    isAuthenticated() {
        // 認証済みかどうかのロジックを実装
        return false;
    }

    async getAuthToken(forceRefresh = false) {
        // 認証トークン取得ロジックを実装
        return null;
    }

    onAuthStateChanged(callback) {
        // 認証状態変更リスナーを実装
        return () => {};
    }

    cleanup() {
        // クリーンアップロジックを実装
    }
}

export default MyAuthProvider;
```

次に `auth-factory.js` ファイルを更新します：

```javascript
// auth-factory.js
import MyAuthProvider from './my-auth-provider';

// 認証プロバイダータイプ
export const AUTH_PROVIDERS = {
    FIREBASE: 'firebase',
    MOCK: 'mock',
    MY_AUTH: 'my-auth',
    // ...
};

// createAuthProvider 関数内
case AUTH_PROVIDERS.MY_AUTH:
    return new MyAuthProvider(config.myAuth);
```

そして `.env.example` ファイルを更新します：

```
# 認証設定
VITE_AUTH_PROVIDER=firebase  # オプション: firebase, mock, my-auth
```

## 利用可能なプロバイダー

### Firebase 認証プロバイダー

Firebase 認証プロバイダーは Firebase Authentication を使用して認証操作を行います。環境変数に Firebase 設定を設定する必要があります。

### モック認証プロバイダー

モック認証プロバイダーは、ユーザーデータをメモリと localStorage に保存するシンプルな実装です。開発やテストに便利です。

## ベストプラクティス

1. プロバイダーを直接使用せず、常に `AuthService` を使用してください。
2. 認証サービスを使用する前に初期化してください。
3. 認証エラーを適切に処理してください。
4. 認証状態リスナーを不要になったらクリーンアップしてください。
5. 設定には環境変数を使用してください。
