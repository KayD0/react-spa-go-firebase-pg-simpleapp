/**
 * モック認証プロバイダー
 * 
 * このファイルは、モック認証プロバイダーを使用してAuthInterfaceを実装します。
 * 開発、テスト、または他の認証プロバイダーを実装するためのテンプレートとして使用できます。
 */

import { AuthInterface } from './auth-interface';

export class MockAuthProvider extends AuthInterface {
  constructor(config = {}) {
    super();
    this.config = config;
    this.currentUser = null;
    this.listeners = [];
    this.mockUsers = config.mockUsers || [
      { email: 'test@example.com', password: 'password123' }
    ];
  }

  /**
   * モック認証プロバイダーを初期化する
   * @returns {Promise<void>}
   */
  async initialize() {
    console.log('モック認証プロバイダーが初期化されました');
    
    // localStorageに保存されたユーザーがいるか確認
    const storedUser = localStorage.getItem('mock_auth_user');
    if (storedUser) {
      try {
        this.currentUser = JSON.parse(storedUser);
        // 認証されたユーザーをリスナーに通知
        this._notifyListeners(this.currentUser);
      } catch (error) {
        console.error('保存されたユーザーの解析エラー:', error);
        localStorage.removeItem('mock_auth_user');
      }
    }
    
    return Promise.resolve();
  }

  /**
   * メールアドレスとパスワードでサインアップ
   * @param {string} email - ユーザーのメールアドレス
   * @param {string} password - ユーザーのパスワード
   * @returns {Promise<{user: Object|null, error: string|null}>}
   */
  async signUp(email, password) {
    // ユーザーが既に存在するか確認
    const existingUser = this.mockUsers.find(user => user.email === email);
    if (existingUser) {
      return { user: null, error: 'このメールアドレスのユーザーは既に存在します' };
    }
    
    // 新しいユーザーを作成
    const newUser = {
      uid: `mock-${Date.now()}`,
      email,
      emailVerified: false,
      displayName: email.split('@')[0],
      createdAt: new Date().toISOString()
    };
    
    // モックユーザーに追加
    this.mockUsers.push({ email, password });
    
    // 現在のユーザーとして設定
    this.currentUser = newUser;
    localStorage.setItem('mock_auth_user', JSON.stringify(newUser));
    
    // リスナーに通知
    this._notifyListeners(newUser);
    
    return { user: newUser, error: null };
  }

  /**
   * メールアドレスとパスワードでサインイン
   * @param {string} email - ユーザーのメールアドレス
   * @param {string} password - ユーザーのパスワード
   * @returns {Promise<{user: Object|null, error: string|null}>}
   */
  async signIn(email, password) {
    // ユーザーを検索
    const user = this.mockUsers.find(
      user => user.email === email && user.password === password
    );
    
    if (!user) {
      return { user: null, error: 'メールアドレスまたはパスワードが無効です' };
    }
    
    // ユーザーオブジェクトを作成（パスワードを除く）
    const authUser = {
      uid: `mock-${email.split('@')[0]}`,
      email,
      emailVerified: true,
      displayName: email.split('@')[0],
      lastLoginAt: new Date().toISOString()
    };
    
    // 現在のユーザーとして設定
    this.currentUser = authUser;
    localStorage.setItem('mock_auth_user', JSON.stringify(authUser));
    
    // リスナーに通知
    this._notifyListeners(authUser);
    
    return { user: authUser, error: null };
  }

  /**
   * 現在のユーザーをサインアウト
   * @returns {Promise<{error: string|null}>}
   */
  async signOut() {
    this.currentUser = null;
    localStorage.removeItem('mock_auth_user');
    
    // リスナーに通知
    this._notifyListeners(null);
    
    return { error: null };
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
   * @param {boolean} forceRefresh - トークンの強制リフレッシュを行うかどうか
   * @returns {Promise<string|null>}
   */
  async getAuthToken(forceRefresh = false) {
    if (!this.currentUser) {
      return null;
    }
    
    // モックトークンを生成
    const payload = {
      sub: this.currentUser.uid,
      email: this.currentUser.email,
      name: this.currentUser.displayName,
      iat: Math.floor(Date.now() / 1000),
      exp: Math.floor(Date.now() / 1000) + 3600 // 1時間の有効期限
    };
    
    // シンプルなモックトークンを作成（実際のJWTではなく、デモ用）
    const mockToken = btoa(JSON.stringify(payload));
    return mockToken;
  }

  /**
   * 認証状態の変更を購読
   * @param {Function} callback - 認証状態が変更されたときに呼び出される関数
   * @returns {Function} - 購読解除関数
   */
  onAuthStateChanged(callback) {
    this.listeners.push(callback);
    
    // 現在の状態を即座に呼び出す
    callback(this.currentUser);
    
    // 購読解除関数を返す
    return () => {
      this.listeners = this.listeners.filter(listener => listener !== callback);
    };
  }

  /**
   * 認証状態の変更をすべてのリスナーに通知
   * @private
   * @param {Object|null} user - 現在のユーザー、またはサインアウト時はnull
   */
  _notifyListeners(user) {
    this.listeners.forEach(listener => {
      try {
        listener(user);
      } catch (error) {
        console.error('認証状態変更リスナーのエラー:', error);
      }
    });
  }

  /**
   * プロバイダーが不要になったときにリソースをクリーンアップ
   */
  cleanup() {
    this.listeners = [];
  }
}

export default MockAuthProvider;
