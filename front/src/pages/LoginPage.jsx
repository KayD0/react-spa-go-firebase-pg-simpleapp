import React, { useState } from 'react';
import { Alert } from 'react-bootstrap';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import '../css/login.css';

const LoginPage = () => {
  const [formData, setFormData] = useState({
    username: '',
    password: ''
  });
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [error, setError] = useState('');
  const { signIn } = useAuth();
  const navigate = useNavigate();

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setIsSubmitting(true);

    try {
      // Using username as email for authentication
      const { username, password } = formData;
      // Treat username as email for Firebase authentication
      const email = username.includes('@') ? username : `${username}@example.com`;
      const { user, error: signInError } = await signIn(email, password);

      if (signInError) {
        setError(signInError);
      } else {
        // Redirect to home page on successful login
        navigate('/');
      }
    } catch (err) {
      console.error('Login error:', err);
      setError('ログイン処理中にエラーが発生しました。後でもう一度お試しください。');
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="login-page">
      <div className="ring">
        <i style={{ "--clr": "#00ff0a" }}></i>
        <i style={{ "--clr": "#ff0057" }}></i>
        <i style={{ "--clr": "#fffd44" }}></i>
        <div className="login">
          <h2>Login</h2>
          
          {error && (
            <div className="inputBx">
              <Alert dismissible onClose={() => setError('')}>
                {error}
              </Alert>
            </div>
          )}
          
          <form onSubmit={handleSubmit}>
            <div className="inputBx">
              <input
                type="text"
                name="username"
                placeholder="Username"
                value={formData.username}
                onChange={handleChange}
                required
              />
            </div>
            
            <div className="inputBx">
              <input
                type="password"
                name="password"
                placeholder="Password"
                value={formData.password}
                onChange={handleChange}
                required
              />
            </div>
            
            <div className="inputBx">
              <input
                type="submit"
                value={isSubmitting ? "Signing in..." : "Sign in"}
                disabled={isSubmitting}
              />
            </div>
          </form>
          
          <div className="links">
            <a href="#">Forget Password</a>
            <Link to="/register">Signup</Link>
          </div>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
