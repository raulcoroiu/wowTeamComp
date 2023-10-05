import React, { useState } from 'react';
import axios from 'axios'; // Import Axios

export const Register = () => {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleRegister = async (e) => {
    e.preventDefault();

    try {
      const response = await axios.post('http://localhost:8080/register', {
        username: username,
        email: email,
        password: password,
      });

      if (response.status === 201) {
        // Registration successful
        console.log('User registered successfully');
        // You can redirect the user to the login page or perform other actions
      } else {
        // Handle other status codes or error responses
      }
    } catch (error) {
      console.error('Registration error:', error);
      // Handle registration error, display an error message to the user, etc.
    }
  };

  return (
    <div className="register-container">
      <h2>Register</h2>
      <form onSubmit={handleRegister}>
        <label>Username:</label>
        <input
          type="text"
          required
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
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
        <button type="submit">Register</button>
      </form>
    </div>
  );
};

