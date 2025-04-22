package dto

// AuthResponse represents the authentication response data
type AuthResponse struct {
	Authenticated bool                   `json:"authenticated"`
	User          map[string]interface{} `json:"user"`
}

// NewAuthResponse creates a new AuthResponse
func NewAuthResponse(claims map[string]interface{}) *AuthResponse {
	return &AuthResponse{
		Authenticated: true,
		User: map[string]interface{}{
			"uid":            claims["uid"],
			"email":          claims["email"],
			"email_verified": claims["email_verified"],
			"auth_time":      claims["auth_time"],
		},
	}
}
