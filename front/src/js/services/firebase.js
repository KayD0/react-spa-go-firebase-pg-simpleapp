/**
 * Firebase Service
 * 
 * This file is a compatibility layer that redirects to the new authentication service.
 * It maintains the same API as the original firebase.js file to ensure backward compatibility.
 * 
 * DEPRECATED: This file is maintained for backward compatibility only.
 * New code should use the auth-service.js file directly.
 */

import authService from './auth/auth-service';

// Initialize the authentication service
authService.initialize().catch(error => {
  console.error('Failed to initialize authentication service:', error);
});

// Authentication state observer for UI updates
let currentUser = null;
authService.onAuthStateChanged(user => {
  currentUser = user;
  
  // Dispatch custom event when auth state changes
  window.dispatchEvent(new CustomEvent('authStateChanged', { detail: { user } }));
  
  // Update UI based on auth state
  updateUI();
});

// Update UI based on authentication state
function updateUI() {
  const authLinks = document.getElementById('authLinks');
  const userInfo = document.getElementById('userInfo');
  
  if (!authLinks || !userInfo) return;
  
  if (currentUser) {
    // User is signed in
    authLinks.style.display = 'none';
    userInfo.style.display = 'block';
    userInfo.querySelector('.user-email').textContent = currentUser.email;
  } else {
    // User is signed out
    authLinks.style.display = 'block';
    userInfo.style.display = 'none';
  }
}

// Sign up with email and password
export async function signUp(email, password) {
  return authService.signUp(email, password);
}

// Sign in with email and password
export async function signIn(email, password) {
  return authService.signIn(email, password);
}

// Sign out
export async function signOut() {
  return authService.signOut();
}

// Get current user
export function getCurrentUser() {
  return authService.getCurrentUser();
}

// Check if user is authenticated
export function isAuthenticated() {
  return authService.isAuthenticated();
}

export default {
  signUp,
  signIn,
  signOut,
  getCurrentUser,
  isAuthenticated
};
