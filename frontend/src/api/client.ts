import axios from 'axios';

// The Nginx proxy routes /api/auth to the user-service
export const api = axios.create({
  baseURL: '/api',
});

// Interceptor to attach token
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token && config.headers) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});
