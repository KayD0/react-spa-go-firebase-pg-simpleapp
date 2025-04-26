package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "net/http"
    "os"
)

// AuthResponse represents the response from the auth verification endpoint
type AuthResponse struct {
    Authenticated bool                   `json:"authenticated"`
    User          map[string]interface{} `json:"user"`
}

// ErrorResponse represents an error response from the API
type ErrorResponse struct {
    Error string `json:"error"`
}

// TestAuthVerification tests the auth verification endpoint
func TestAuthVerification(token string) {
    baseURL := os.Getenv("API_BASE_URL")
    if baseURL == "" {
        baseURL = "http://localhost:5000"
    }

    if token == "" {
        fmt.Println("トークンが提供されていません。未認証アクセスをテストします...")
        
        // Test without token
        resp, err := http.Post(fmt.Sprintf("%s/api/auth/verify", baseURL), "application/json", nil)
        if err != nil {
            fmt.Printf("エラー: %v\n", err)
            return
        }
        defer resp.Body.Close()

        fmt.Printf("ステータスコード: %d\n", resp.StatusCode)
        
        var errorResp ErrorResponse
        if err := json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
            fmt.Printf("レスポンスの解析エラー: %v\n", err)
            return
        }
        
        fmt.Printf("レスポンス: %s\n\n", errorResp.Error)

        // Test with invalid token
        req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/auth/verify", baseURL), nil)
        req.Header.Set("Authorization", "Bearer invalid_token")
        
        client := &http.Client{}
        resp, err = client.Do(req)
        if err != nil {
            fmt.Printf("無効なトークンでのエラー: %v\n", err)
            return
        }
        defer resp.Body.Close()

        fmt.Printf("無効なトークンでのステータスコード: %d\n", resp.StatusCode)
        
        if err := json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
            fmt.Printf("無効なトークンでのレスポンスの解析エラー: %v\n", err)
            return
        }
        
        fmt.Printf("無効なトークンでのレスポンス: %s\n", errorResp.Error)
    } else {
        fmt.Printf("提供されたトークンでテストしています...\n")
        
        // Test with provided token
        req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/auth/verify", baseURL), nil)
        req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
        
        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            fmt.Printf("エラー: %v\n", err)
            return
        }
        defer resp.Body.Close()

        fmt.Printf("ステータスコード: %d\n", resp.StatusCode)
        
        if resp.StatusCode == http.StatusOK {
            var authResp AuthResponse
            if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
                fmt.Printf("レスポンスの解析エラー: %v\n", err)
                return
            }
            
            userBytes, _ := json.MarshalIndent(authResp.User, "", "  ")
            fmt.Printf("認証成功: %v\n", authResp.Authenticated)
            fmt.Printf("ユーザー情報:\n%s\n", string(userBytes))
            
            // Test profile endpoint
            fmt.Println("\nプロフィールエンドポイントをテストしています...")
            req, _ = http.NewRequest("GET", fmt.Sprintf("%s/api/profile", baseURL), nil)
            req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
            
            resp, err = client.Do(req)
            if err != nil {
                fmt.Printf("プロフィール取得エラー: %v\n", err)
                return
            }
            defer resp.Body.Close()
            
            fmt.Printf("プロフィールステータスコード: %d\n", resp.StatusCode)
            
            var profileResp map[string]interface{}
            if err := json.NewDecoder(resp.Body).Decode(&profileResp); err != nil {
                fmt.Printf("プロフィールレスポンスの解析エラー: %v\n", err)
                return
            }
            
            profileBytes, _ := json.MarshalIndent(profileResp, "", "  ")
            fmt.Printf("プロフィールレスポンス:\n%s\n", string(profileBytes))
        } else {
            var errorResp ErrorResponse
            if err := json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
                fmt.Printf("エラーレスポンスの解析エラー: %v\n", err)
                return
            }
            
            fmt.Printf("認証エラー: %s\n", errorResp.Error)
        }
    }
}

func auth_test() {
    // Define command line flags
    tokenPtr := flag.String("token", "", "Firebase ID token for authentication")
    
    // Parse command line flags
    flag.Parse()
    
    // Run the test
    TestAuthVerification(*tokenPtr)
}
