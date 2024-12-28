import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@': '/src', // Allows importing with '@/components/...' style
    },
  },
  optimizeDeps: {
    exclude: ['lucide-react'],
  },
  server: {
    watch: {
      usePolling: true, // Fixes file system watcher issues
    },
  },
});
