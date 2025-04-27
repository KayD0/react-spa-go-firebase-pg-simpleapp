# 認証システム移行ガイド

このドキュメントは、Firebaseの直接認証実装から新しい柔軟な認証アーキテクチャへの移行に関するガイダンスを提供します。

## 概要

アプリケーションは、異なる認証プロバイダー間で簡単に切り替えが可能な柔軟な認証アーキテクチャを使用するようにリファクタリングされました。これにより、コアアプリケーションロジックを変更することなく、将来的にFirebase認証を他の認証方法に置き換えることが可能になります。

## 移行戦略

移行は以下の手順に従って段階的に行うことができます：

1. **互換レイヤーを使用する**: `firebase-compat.js`ファイルは、元のFirebase実装と同じAPIを維持しつつ、新しい認証サービスを内部で使用する互換レイヤーを提供します。これにより、既存のコードを壊すことなくスムーズに移行できます。

2. **インポートを更新する**: コンポーネントやサービス内のインポートを、元のFirebase実装から新しい認証サービスに徐々に更新します。

3. **Firebaseの直接依存を削除する**: すべてのコンポーネントとサービスが新しい認証サービスを使用するようになったら、Firebase認証への直接依存を削除します。

## ステップバイステップ移行

### ステップ 1: 互換レイヤーを使用する

`../firebase.js`から直接インポートする代わりに、`./auth/firebase-compat.js`からインポートします：

```javascript
// 変更前
import { signIn, signUp, signOut, getCurrentUser, isAuthenticated } from '../firebase';

// 変更後
import { signIn, signUp, signOut, getCurrentUser, isAuthenticated } from './auth/firebase-compat';
```

この変更はアプリケーションの他の部分に対して透過的であり、互換レイヤーが同じAPIを維持します。

### ステップ 2: 新しい認証サービスに更新する

完全に新しい認証アーキテクチャに移行する準備ができたら、インポートを認証サービスに直接更新します：

```javascript
// 変更前
import { signIn, signUp, signOut, getCurrentUser, isAuthenticated } from './auth/firebase-compat';

// 変更後
import authService from './auth/auth-service';

// authServiceメソッドを使用
const { user, error } = await authService.signIn(email, password);
const isLoggedIn = authService.isAuthenticated();
```

### ステップ 3: 認証プロバイダーを切り替える

異なる認証プロバイダーに切り替えるには、`.env`ファイル内の`VITE_AUTH_PROVIDER`環境変数を更新します：

```
# .envファイル
VITE_AUTH_PROVIDER=mock  # オプション: firebase, mock
```

プロバイダーを切り替えるためにコードの変更は必要ありません。認証サービスは設定されたプロバイダーを自動的に使用します。

## 利用可能な認証プロバイダー

### Firebase認証プロバイダー

Firebase認証プロバイダーは、認証操作にFirebase認証を使用します。環境変数にFirebaseの設定を指定する必要があります。

```
VITE_AUTH_PROVIDER=firebase
```

### モック認証プロバイダー

モック認証プロバイダーは、メモリとlocalStorageにユーザーデータを保存するシンプルな実装です。開発やテストに便利です。

```
VITE_AUTH_PROVIDER=mock
```

## 新しい認証プロバイダーの追加

新しい認証プロバイダーを追加するには、以下の手順に従います：

1. プロバイダー用の新しいファイルを作成します（例: `my-auth-provider.js`）。
2. プロバイダー固有のロジックで`AuthInterface`クラスを実装します。
3. `auth-factory.js`内の`AUTH_PROVIDERS`オブジェクトにプロバイダーを追加します。
4. `auth-factory.js`内の`createAuthProvider`関数にプロバイダーを作成するケースを追加します。
5. `.env.example`ファイルを更新し、プロバイダーをオプションとして追加します。

新しい認証プロバイダーの追加に関する詳細は、[README.md](./README.md)ファイルを参照してください。

## ベストプラクティス

1. 常にプロバイダーを直接使用するのではなく、`AuthService`を使用してください。
2. 認証サービスを使用する前に初期化してください。
3. 認証エラーを適切に処理してください。
4. 不要になった認証状態リスナーをクリーンアップしてください。
5. 環境変数を使用して設定を行ってください。
