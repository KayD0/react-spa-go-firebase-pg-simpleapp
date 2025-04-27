/**
 * Firebase互換レイヤー
 * 
 * このファイルは、旧Firebase実装と新しい認証アーキテクチャの間の互換レイヤーを提供します。
 * 移行期間中に、アプリケーションが両方の実装を併用できるようにします。
 * 
 * 元のFirebase実装をインポートし、同じAPIをエクスポートしますが、
 * 内部的には新しい認証サービスを使用します。
 */

import * as originalFirebase from '../firebase';
import authService from './auth-service';

// 認証サービスを初期化
authService.initialize().catch(error => {
  console.error('認証サービスの初期化に失敗しました:', error);
});

/**
 * メールアドレスとパスワードでサインアップ
 * 新しい認証サービスを使用しますが、同じAPIを維持します。
 * 
 * @param {string} email - ユーザーのメールアドレス
 * @param {string} password - ユーザーのパスワード
 * @returns {Promise<{user: Object|null, error: string|null}>}
 */
export async function signUp(email, password) {
  // 新しい認証サービスを使用
  return authService.signUp(email, password);
}

/**
 * メールアドレスとパスワードでサインイン
 * 新しい認証サービスを使用しますが、同じAPIを維持します。
 * 
 * @param {string} email - ユーザーのメールアドレス
 * @param {string} password - ユーザーのパスワード
 * @returns {Promise<{user: Object|null, error: string|null}>}
 */
export async function signIn(email, password) {
  // 新しい認証サービスを使用
  return authService.signIn(email, password);
}

/**
 * サインアウト
 * 新しい認証サービスを使用しますが、同じAPIを維持します。
 * 
 * @returns {Promise<{error: string|null}>}
 */
export async function signOut() {
  // 新しい認証サービスを使用
  return authService.signOut();
}

/**
 * 現在のユーザーを取得
 * 新しい認証サービスを使用しますが、同じAPIを維持します。
 * 
 * @returns {Object|null}
 */
export function getCurrentUser() {
  // 新しい認証サービスを使用
  return authService.getCurrentUser();
}

/**
 * ユーザーが認証されているか確認
 * 新しい認証サービスを使用しますが、同じAPIを維持します。
 * 
 * @returns {boolean}
 */
export function isAuthenticated() {
  // 新しい認証サービスを使用
  return authService.isAuthenticated();
}

/**
 * 元のFirebase実装を取得
 * 必要に応じて元のFirebase実装にアクセスできます。
 * 
 * @returns {Object}
 */
export function getOriginalFirebase() {
  return originalFirebase;
}

export default {
  signUp,
  signIn,
  signOut,
  getCurrentUser,
  isAuthenticated,
  getOriginalFirebase
};
