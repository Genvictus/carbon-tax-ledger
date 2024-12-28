import React, { createContext, useContext, useState } from 'react';
import { authService } from '../services/auth';
import { showErrorToast } from '../utils/toastUtils';

interface AuthContextType {
  login: (mspID: string, cert: File, key: File, tlsCert: File) => Promise<void>;
  logout: () => Promise<void>;
  isAuthenticated: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(!!localStorage.getItem('mspID') && !!localStorage.getItem('sessionID'));

  const login = async (mspID: string, cert: File, key: File, tlsCert: File) => {
    try {
      const response = await authService.login({ mspID, cert, key, tlsCert });
      if (response?.success) {
        localStorage.setItem('mspID', response.data?.mspID!);
        localStorage.setItem('sessionID', response.data?.sessionID!);
        setIsAuthenticated(true);
      } else {
        showErrorToast(response?.error || 'Failed to login');
      }
    } catch (err) {
      showErrorToast('Failed to login');
    }
  };

  const logout = async () => {
    try {
      await authService.logout();
    } catch (err) {
      // do nothing!
    } finally {
      localStorage.removeItem('mspID');
      localStorage.removeItem('sessionID');
      setIsAuthenticated(false);
    }
  };

  return (
    <AuthContext.Provider value={{ login, logout, isAuthenticated }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}