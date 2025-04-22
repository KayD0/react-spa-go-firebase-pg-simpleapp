package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// FirebaseAuth defines the interface for Firebase authentication
type FirebaseAuth interface {
	// VerifyIDToken verifies a Firebase ID token and returns the decoded token
	VerifyIDToken(ctx context.Context, idToken string) (map[string]interface{}, error)
}

// FirebaseAuthImpl implements the FirebaseAuth interface
type FirebaseAuthImpl struct {
	app       *firebase.App
	authClient *auth.Client
}

// AuthError represents an authentication error
type AuthError struct {
	Message string
}

// Error returns the error message
func (e *AuthError) Error() string {
	return e.Message
}

// NewFirebaseAuth creates a new FirebaseAuthImpl
func NewFirebaseAuth() (FirebaseAuth, error) {
	ctx := context.Background()
	var app *firebase.App
	var err error

	// Check if environment variables for service account are provided
	if os.Getenv("FIREBASE_PROJECT_ID") != "" {
		// Create a service account credential from environment variables
		serviceAccount := map[string]interface{}{
			"type":                        "service_account",
			"project_id":                  os.Getenv("FIREBASE_PROJECT_ID"),
			"private_key_id":              os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
			"private_key":                 strings.ReplaceAll(os.Getenv("FIREBASE_PRIVATE_KEY"), "\\n", "\n"),
			"client_email":                os.Getenv("FIREBASE_CLIENT_EMAIL"),
			"client_id":                   os.Getenv("FIREBASE_CLIENT_ID"),
			"auth_uri":                    os.Getenv("FIREBASE_AUTH_URI"),
			"token_uri":                   os.Getenv("FIREBASE_TOKEN_URI"),
			"auth_provider_x509_cert_url": os.Getenv("FIREBASE_AUTH_PROVIDER_X509_CERT_URL"),
			"client_x509_cert_url":        os.Getenv("FIREBASE_CLIENT_X509_CERT_URL"),
		}

		serviceAccountJSON, err := json.Marshal(serviceAccount)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal service account: %v", err)
		}

		// Initialize Firebase with service account
		opt := option.WithCredentialsJSON(serviceAccountJSON)
		app, err = firebase.NewApp(ctx, nil, opt)
		if err != nil {
			return nil, fmt.Errorf("error initializing Firebase with service account: %v", err)
		}
		log.Println("Firebase Admin SDK initialized with service account credentials")
	} else {
		// Initialize Firebase with application default credentials
		app, err = firebase.NewApp(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("error initializing Firebase with default credentials: %v", err)
		}
		log.Println("Firebase Admin SDK initialized with application default credentials")
	}

	// Get Auth client
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v", err)
	}

	return &FirebaseAuthImpl{
		app:       app,
		authClient: client,
	}, nil
}

// VerifyIDToken verifies a Firebase ID token and returns the decoded token
func (f *FirebaseAuthImpl) VerifyIDToken(ctx context.Context, idToken string) (map[string]interface{}, error) {
	// Verify the token
	token, err := f.authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("error verifying ID token: %v", err)
	}

	// Return the token claims
	return token.Claims, nil
}
