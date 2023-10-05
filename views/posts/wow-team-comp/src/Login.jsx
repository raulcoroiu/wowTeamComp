import React, { useState } from 'react';
import axios from 'axios'; // Import Axios

export const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = async (e) => {
    e.preventDefault();

    try {
      const response = await axios.post('http://localhost:8080/login', {
        email: email,
        password: password,
      });

      if (response.status === 200) {
        // Authentication successful
        console.log(response.data.token); // Access the JWT token
        // You can save the token to local storage or a cookie for future requests
        // Redirect the user to a protected route or perform other actions
      } else {
        // Handle other status codes or error responses
      }
    } catch (error) {
      console.error('Login error:', error);
      // Handle login error, display an error message to the user, etc.
    }
  };

  return (
    <div className="login-container">
      <h2>Login</h2>
      <form onSubmit={handleLogin}>
        <label>Email:</label>
        <input
          type="email"
          required
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <label>Password:</label>
        <input
          type="password"
          required
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <button type="submit">Login</button>
      </form>
    </div>
  );
};

