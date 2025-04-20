import React, { useState } from 'react';
import { Alert } from 'react-bootstrap';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import '../css/register.css';

const RegisterPage = () => {
  const [formData, setFormData] = useState({
    email: '',
    password: '',
    confirmPassword: ''
  });
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [error, setError] = useState('');
  const { signUp } = useAuth();
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

    // Validate passwords match
    if (formData.password !== formData.confirmPassword) {
      setError('パスワードが一致しません');
      return;
    }

    // Validate password strength
    if (formData.password.length < 6) {
      setError('パスワードは6文字以上である必要があります');
      return;
    }

    setIsSubmitting(true);

    try {
      const { email, password } = formData;
      const { user, error: signUpError } = await signUp(email, password);

      if (signUpError) {
        setError(signUpError);
      } else {
        // Redirect to home page on successful registration
        navigate('/');
      }
    } catch (err) {
      console.error('Registration error:', err);
      setError('登録処理中にエラーが発生しました。後でもう一度お試しください。');
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="register-page">
      <div className="ring">
        <i style={{ "--clr": "#00ff0a" }}></i>
        <i style={{ "--clr": "#ff0057" }}></i>
        <i style={{ "--clr": "#fffd44" }}></i>
        <div className="register">
          <h2>Signup</h2>
          
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
                type="email"
                name="email"
                placeholder="Email"
                value={formData.email}
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
              <div className="text-muted">
                Password must be at least 6 characters
              </div>
            </div>
            
            <div className="inputBx">
              <input
                type="password"
                name="confirmPassword"
                placeholder="Confirm Password"
                value={formData.confirmPassword}
                onChange={handleChange}
                required
              />
            </div>
            
            <div className="inputBx">
              <input
                type="submit"
                value={isSubmitting ? "Processing..." : "Sign up"}
                disabled={isSubmitting}
              />
            </div>
          </form>
          
          <div className="links">
            <Link to="/login">Already have an account? Login</Link>
          </div>
        </div>
      </div>
    </div>
  );
};

export default RegisterPage;
