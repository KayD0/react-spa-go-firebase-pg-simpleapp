/**
 * Firebase 認証プロバイダー
 * 
 * このファイルは Firebase 認証を使用して AuthInterface を実装します。
 * Firebase 固有の認証ロジックをカプセル化しています。
 */

import { initializeApp } from 'firebase/app';
import { 
    getAuth, 
    createUserWithEmailAndPassword, 
    signInWithEmailAndPassword, 
    signOut as firebaseSignOut,
    onAuthStateChanged,
    browserLocalPersistence,
    setPersistence
} from 'firebase/auth';
import { AuthInterface } from './auth-interface';

export class FirebaseAuthProvider extends AuthInterface {
    constructor(config) {
        super();
        this.config = config;
        this.app = null;
        this.auth = null;
        this.currentUser = null;
    }

    /**
     * Firebase 認証の初期化
     * @returns {Promise<void>}
     */
    async initialize() {
        try {
            // Firebase アプリの初期化
            this.app = initializeApp(this.config);
            this.auth = getAuth(this.app);
            
            // 永続性を有効にして、ユーザーがログイン状態を維持できるようにする
            await setPersistence(this.auth, browserLocalPersistence);
            console.log('Firebase の永続性がローカルに設定されました');
            
            // 認証状態リスナーの設定
            this.unsubscribeAuthState = onAuthStateChanged(this.auth, (user) => {
                this.currentUser = user;
                
                // 認証状態が変化したときにカスタムイベントをディスパッチ
                window.dispatchEvent(new CustomEvent('authStateChanged', { detail: { user } }));
            });
            
            return Promise.resolve();
        } catch (error) {
            console.error('Firebase 初期化エラー:', error);
            
            // 開発用のフォールバックを提供
            if (!this.app) this.app = initializeApp({ projectId: 'demo-project' });
            if (!this.auth) this.auth = getAuth(this.app);
            
            return Promise.reject(error);
        }
    }

    /**
     * メールアドレスとパスワードでサインアップ
     * @param {string} email - ユーザーのメールアドレス
     * @param {string} password - ユーザーのパスワード
     * @returns {Promise<{user: Object|null, error: string|null}>}
     */
    async signUp(email, password) {
        try {
            const userCredential = await createUserWithEmailAndPassword(this.auth, email, password);
            return { user: userCredential.user, error: null };
        } catch (error) {
            console.error("サインアップエラー:", error);
            return { user: null, error: error.message };
        }
    }

    /**
     * メールアドレスとパスワードでサインイン
     * @param {string} email - ユーザーのメールアドレス
     * @param {string} password - ユーザーのパスワード
     * @returns {Promise<{user: Object|null, error: string|null}>}
     */
    async signIn(email, password) {
        try {
            const userCredential = await signInWithEmailAndPassword(this.auth, email, password);
            return { user: userCredential.user, error: null };
        } catch (error) {
            console.error("サインインエラー:", error);
            return { user: null, error: error.message };
        }
    }

    /**
     * 現在のユーザーをサインアウト
     * @returns {Promise<{error: string|null}>}
     */
    async signOut() {
        try {
            await firebaseSignOut(this.auth);
            return { error: null };
        } catch (error) {
            console.error("サインアウトエラー:", error);
            return { error: error.message };
        }
    }

    /**
     * 現在認証されているユーザーを取得
     * @returns {Object|null}
     */
    getCurrentUser() {
        return this.currentUser;
    }

    /**
     * 現在ユーザーが認証されているか確認
     * @returns {boolean}
     */
    isAuthenticated() {
        return !!this.currentUser;
    }

    /**
     * 現在のユーザーの認証トークンを取得
     * @param {boolean} forceRefresh - トークンを強制的にリフレッシュするかどうか
     * @returns {Promise<string|null>}
     */
    async getAuthToken(forceRefresh = false) {
        if (!this.currentUser) {
            console.warn('現在サインインしているユーザーがいません');
            return null;
        }
        
        try {
            const token = await this.currentUser.getIdToken(forceRefresh);
            return token;
        } catch (error) {
            console.error('認証トークン取得エラー:', error);
            return null;
        }
    }

    /**
     * 認証状態の変化を購読
     * @param {Function} callback - 認証状態が変化したときに呼び出される関数
     * @returns {Function} - 購読解除関数
     */
    onAuthStateChanged(callback) {
        return onAuthStateChanged(this.auth, callback);
    }

    /**
     * プロバイダーが不要になったときにリソースをクリーンアップ
     */
    cleanup() {
        if (this.unsubscribeAuthState) {
            this.unsubscribeAuthState();
        }
    }
}

export default FirebaseAuthProvider;
