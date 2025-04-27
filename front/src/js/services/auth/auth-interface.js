/**
 * 認証インターフェース
 * 
 * このファイルは、すべての認証プロバイダーが実装しなければならないインターフェースを定義します。
 * これは認証実装の契約として機能し、アプリケーションがコアロジックを変更することなく
 * 異なる認証プロバイダー間を切り替えることを可能にします。
 */

/**
 * 抽象認証インターフェース
 * 
 * すべての認証プロバイダーは、アプリケーションと互換性を持つためにこれらのメソッドを実装する必要があります。
 */
export class AuthInterface {
    /**
     * 認証プロバイダーを初期化する
     * @returns {Promise<void>}
     */
    async initialize() {
        throw new Error('未実装のメソッド: initialize()');
    }

    /**
     * 新しいユーザーを登録する
     * @param {string} email - ユーザーのメールアドレス
     * @param {string} password - ユーザーのパスワード
     * @returns {Promise<{user: Object|null, error: string|null}>} - ユーザーデータまたはエラーを含む結果オブジェクト
     */
    async signUp(email, password) {
        throw new Error('未実装のメソッド: signUp()');
    }

    /**
     * 既存のユーザーでサインインする
     * @param {string} email - ユーザーのメールアドレス
     * @param {string} password - ユーザーのパスワード
     * @returns {Promise<{user: Object|null, error: string|null}>} - ユーザーデータまたはエラーを含む結果オブジェクト
     */
    async signIn(email, password) {
        throw new Error('未実装のメソッド: signIn()');
    }

    /**
     * 現在のユーザーをサインアウトする
     * @returns {Promise<{error: string|null}>} - エラーがあればそれを含む結果オブジェクト
     */
    async signOut() {
        throw new Error('未実装のメソッド: signOut()');
    }

    /**
     * 現在認証されているユーザーを取得する
     * @returns {Object|null} - 現在のユーザーオブジェクト、または認証されていない場合はnull
     */
    getCurrentUser() {
        throw new Error('未実装のメソッド: getCurrentUser()');
    }

    /**
     * 現在ユーザーが認証されているかどうかを確認する
     * @returns {boolean} - ユーザーが認証されていればtrue、そうでなければfalse
     */
    isAuthenticated() {
        throw new Error('未実装のメソッド: isAuthenticated()');
    }

    /**
     * 現在のユーザーの認証トークンを取得する
     * @param {boolean} forceRefresh - トークンの強制リフレッシュを行うかどうか
     * @returns {Promise<string|null>} - 認証トークン、または認証されていない場合はnull
     */
    async getAuthToken(forceRefresh = false) {
        throw new Error('未実装のメソッド: getAuthToken()');
    }

    /**
     * 認証状態の変更を購読する
     * @param {Function} callback - 認証状態が変更されたときに呼び出される関数
     * @returns {Function} - 購読解除関数
     */
    onAuthStateChanged(callback) {
        throw new Error('未実装のメソッド: onAuthStateChanged()');
    }
}

export default AuthInterface;
