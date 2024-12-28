import axios from 'axios';

// Create an axios instance with default config
export const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:5000/api',
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add a request interceptor for auth
api.interceptors.request.use(
  (config) => {
    const mspID = localStorage.getItem('mspID');
    const sessionID = localStorage.getItem('sessionID');
    if (mspID && sessionID) {
      config.headers['msp-id'] = mspID;
      config.headers['session-id'] = sessionID;
  }
  return config;
});

// Add a response interceptor for auth
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 400) {
      localStorage.removeItem('mspID');
      localStorage.removeItem('sessionID');
      window.location.reload();
    }
    return Promise.reject(error);
  },
);