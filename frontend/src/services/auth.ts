import { api } from '../config/api';

interface LoginResponse {
  success: boolean;
  message: string;
  error: string | null;
  data: {
    sessionID: string;
    mspID: string;
  } | null;
}

interface LoginRequest {
  mspID: string;
  cert: File;
  key: File;
  tlsCert: File;
}

export const authService = {
  async login(credentials: LoginRequest) {
    const { data } = await api.post<LoginResponse>('/login', credentials, {
      headers: { 'Content-Type': 'multipart/form-data' },
    });
    return data;
  },

  async logout() {
    await api.post('/logout', null);
  },
};