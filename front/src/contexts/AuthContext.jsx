import React, { createContext, useContext, useState, useEffect } from 'react';
import authService from '../js/services/auth/auth-service';

// Create the auth context
const AuthContext = createContext();

// Custom hook to use the auth context
export const useAuth = () => {
  return useContext(AuthContext);
};

// Auth provider component
export const AuthProvider = ({ children }) => {
  const [currentUser, setCurrentUser] = useState(null);
  const [loading, setLoading] = useState(true);
  const [initialized, setInitialized] = useState(false);

  // Initialize the authentication service
  useEffect(() => {
    const initAuth = async () => {
      try {
        await authService.initialize();
        setInitialized(true);
      } catch (error) {
        console.error("Error initializing auth service:", error);
        setInitialized(true); // Still mark as initialized to prevent infinite loading
      }
    };

    initAuth();
  }, []);

  // Sign up with email and password
  const signUp = async (email, password) => {
    return authService.signUp(email, password);
  };

  // Sign in with email and password
  const signIn = async (email, password) => {
    return authService.signIn(email, password);
  };

  // Sign out
  const signOut = async () => {
    return authService.signOut();
  };

  // Listen for auth state changes
  useEffect(() => {
    if (!initialized) return;

    const unsubscribe = authService.onAuthStateChanged((user) => {
      setCurrentUser(user);
      setLoading(false);
    });

    // Cleanup subscription on unmount
    return unsubscribe;
  }, [initialized]);

  // Context value
  const value = {
    currentUser,
    isAuthenticated: !!currentUser,
    signUp,
    signIn,
    signOut,
    authProviderType: initialized ? authService.getProviderType() : null,
  };

  return (
    <AuthContext.Provider value={value}>
      {!loading && children}
    </AuthContext.Provider>
  );
};
