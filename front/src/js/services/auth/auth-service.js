/**
 * 認証サービス
 * 
 * このサービスは認証操作のための統一されたインターフェースを提供します。
 * 設定に基づいて適切な認証プロバイダーを作成するために、authファクトリを使用します。
 */

// 認証プロバイダーを作成するためのファクトリ関数とプロバイダータイプをインポート
import { createAuthProvider, AUTH_PROVIDERS } from './auth-factory';

// 環境変数から認証プロバイダーのタイプを取得、またはデフォルトを使用
const AUTH_PROVIDER_TYPE = import.meta.env.VITE_AUTH_PROVIDER || AUTH_PROVIDERS.FIREBASE;

class AuthService {
    constructor() {
        this.provider = null;
        this.initialized = false;
        this.initializationPromise = null;
    }

    /**
     * 認証サービスを初期化する
     * @param {Object} config - 認証プロバイダーの設定
     * @returns {Promise<void>}
     */
    async initialize(config = {}) {
        if (this.initialized) {
            return Promise.resolve();
        }

        if (this.initializationPromise) {
            return this.initializationPromise;
        }

        this.initializationPromise = new Promise(async (resolve, reject) => {
            try {
                // 認証プロバイダーを作成
                this.provider = createAuthProvider(AUTH_PROVIDER_TYPE, config);
                
                // プロバイダーを初期化
                await this.provider.initialize();
                
                this.initialized = true;
                resolve();
            } catch (error) {
                console.error('認証サービスの初期化に失敗しました:', error);
                reject(error);
            }
        });

        return this.initializationPromise;
    }

    /**
     * 操作を実行する前にサービスが初期化されていることを確認
     * @private
     */
    _ensureInitialized() {
        if (!this.initialized) {
            throw new Error('認証サービスが初期化されていません。まずinitialize()を呼び出してください。');
        }
    }

    /**
     * 新しいユーザーを登録する
     * @param {string} email - ユーザーのメールアドレス
     * @param {string} password - ユーザーのパスワード
     * @returns {Promise<{user: Object|null, error: string|null}>}
     */
    async signUp(email, password) {
        this._ensureInitialized();
        return this.provider.signUp(email, password);
    }

    /**
     * 既存のユーザーでサインインする
     * @param {string} email - ユーザーのメールアドレス
     * @param {string} password - ユーザーのパスワード
     * @returns {Promise<{user: Object|null, error: string|null}>}
     */
    async signIn(email, password) {
        this._ensureInitialized();
        return this.provider.signIn(email, password);
    }

    /**
     * 現在のユーザーをサインアウトする
     * @returns {Promise<{error: string|null}>}
     */
    async signOut() {
        this._ensureInitialized();
        return this.provider.signOut();
    }

    /**
     * 現在認証されているユーザーを取得する
     * @returns {Object|null}
     */
    getCurrentUser() {
        if (!this.initialized) return null;
        return this.provider.getCurrentUser();
    }

    /**
     * 現在ユーザーが認証されているかどうかを確認する
     * @returns {boolean}
     */
    isAuthenticated() {
        if (!this.initialized) return false;
        return this.provider.isAuthenticated();
    }

    /**
     * 現在のユーザーの認証トークンを取得する
     * @param {boolean} forceRefresh - トークンの強制リフレッシュを行うかどうか
     * @returns {Promise<string|null>}
     */
    async getAuthToken(forceRefresh = false) {
        this._ensureInitialized();
        return this.provider.getAuthToken(forceRefresh);
    }

    /**
     * 認証状態の変更を購読する
     * @param {Function} callback - 認証状態が変更されたときに呼び出される関数
     * @returns {Function} - 購読解除関数
     */
    onAuthStateChanged(callback) {
        this._ensureInitialized();
        return this.provider.onAuthStateChanged(callback);
    }

    /**
     * 現在の認証プロバイダーのタイプを取得する
     * @returns {string}
     */
    getProviderType() {
        return AUTH_PROVIDER_TYPE;
    }
}

// シングルトンインスタンスを作成
const authService = new AuthService();

export default authService;
