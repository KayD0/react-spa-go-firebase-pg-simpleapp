/**
 * 認証ファクトリー
 * 
 * このファイルは認証プロバイダーを作成するためのファクトリーを提供します。
 * アプリケーションが異なる認証プロバイダー間を簡単に切り替えられるようにします。
 */

import FirebaseAuthProvider from './firebase-auth-provider';
import MockAuthProvider from './mock-auth-provider';

// 認証プロバイダーの種類
export const AUTH_PROVIDERS = {
  FIREBASE: 'firebase',
  MOCK: 'mock',
  // 実装され次第、ここに他のプロバイダーを追加
  // COGNITO: 'cognito',
  // AUTH0: 'auth0',
  // CUSTOM: 'custom',
};

// デフォルトのプロバイダータイプ
const DEFAULT_PROVIDER = AUTH_PROVIDERS.FIREBASE;

/**
 * 指定されたタイプに基づいて認証プロバイダーを作成
 * 
 * @param {string} type - 作成する認証プロバイダーのタイプ
 * @param {Object} config - 認証プロバイダーの設定
 * @returns {AuthInterface} - 指定された認証プロバイダーのインスタンス
 * @throws {Error} - 指定されたプロバイダータイプがサポートされていない場合
 */
export function createAuthProvider(type = DEFAULT_PROVIDER, config = {}) {
  switch (type) {
    case AUTH_PROVIDERS.FIREBASE:
      // Firebaseの設定が提供されていない場合は環境変数を使用
      const firebaseConfig = config.firebase || {
        apiKey: import.meta.env.VITE_FIREBASE_API_KEY,
        authDomain: import.meta.env.VITE_FIREBASE_AUTH_DOMAIN,
        projectId: import.meta.env.VITE_FIREBASE_PROJECT_ID,
        storageBucket: import.meta.env.VITE_FIREBASE_STORAGE_BUCKET,
        messagingSenderId: import.meta.env.VITE_FIREBASE_MESSAGING_SENDER_ID,
        appId: import.meta.env.VITE_FIREBASE_APP_ID
      };
      
      // Firebaseの設定が正しく読み込まれているか確認
      if (!firebaseConfig.apiKey || firebaseConfig.apiKey === 'YOUR_API_KEY') {
        console.error('Firebaseの設定が不足しているか無効です。.envファイルが正しく設定されていることを確認してください。');
      }
      
      return new FirebaseAuthProvider(firebaseConfig);
    
    case AUTH_PROVIDERS.MOCK:
      return new MockAuthProvider(config.mock);
      
    // 他のプロバイダーが実装され次第、ケースを追加
    // case AUTH_PROVIDERS.COGNITO:
    //   return new CognitoAuthProvider(config.cognito);
    // case AUTH_PROVIDERS.AUTH0:
    //   return new Auth0AuthProvider(config.auth0);
    // case AUTH_PROVIDERS.CUSTOM:
    //   return new CustomAuthProvider(config.custom);
    
    default:
      throw new Error(`サポートされていない認証プロバイダータイプです: ${type}`);
  }
}

export default {
  createAuthProvider,
  AUTH_PROVIDERS
};
