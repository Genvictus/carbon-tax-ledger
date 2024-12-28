import React, { useState } from 'react';
import { useAuth } from '../context/AuthContext';

interface LoginModalProps {
  isOpen: boolean;
}

export function LoginModal({ isOpen }: LoginModalProps) {
  const [mspID, setMspID] = useState('');
  const [cert, setCert] = useState<File | null>(null);
  const [key, setKey] = useState<File | null>(null);
  const [tlsCert, setTlsCert] = useState<File | null>(null);
  const { login } = useAuth();

  if (!isOpen) return null;

  const handleFileChange = (
    e: React.ChangeEvent<HTMLInputElement>,
    setFile: React.Dispatch<React.SetStateAction<File | null>>
  ) => {
    if (e.target.files) {
      setFile(e.target.files[0]);
    }
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    login(mspID, cert!, key!, tlsCert!);
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-white rounded-lg p-8 max-w-md w-full mx-4 relative">
        <h2 className="text-2xl font-bold mb-6 text-center">Welcome to Carbon Tax Portal</h2>

        <form onSubmit={handleSubmit} className="space-y-6">
          {/* MSP ID */}
          <div>
            <label htmlFor="mspID" className="block text-sm font-medium text-gray-700 mb-1">
              MSP ID
            </label>
            <input
              id="mspID"
              type="text"
              value={mspID}
              onChange={(e) => setMspID(e.target.value)}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-transparent"
              placeholder="Enter your MSP ID"
              required
            />
          </div>

          {/* Certificate File */}
          <div>
            <label htmlFor="cert" className="block text-sm font-medium text-gray-700 mb-1">
              Certificate (Cert)
            </label>
            <input
              id="cert"
              type="file"
              onChange={(e) => handleFileChange(e, setCert)}
              className="w-full border border-gray-300 rounded-lg px-4 py-2"
              required
            />
          </div>

          {/* Key File */}
          <div>
            <label htmlFor="key" className="block text-sm font-medium text-gray-700 mb-1">
              Private Key
            </label>
            <input
              id="key"
              type="file"
              onChange={(e) => handleFileChange(e, setKey)}
              className="w-full border border-gray-300 rounded-lg px-4 py-2"
              required
            />
          </div>

          {/* TLS Certificate File */}
          <div>
            <label htmlFor="tlsCert" className="block text-sm font-medium text-gray-700 mb-1">
              TLS Certificate
            </label>
            <input
              id="tlsCert"
              type="file"
              onChange={(e) => handleFileChange(e, setTlsCert)}
              className="w-full border border-gray-300 rounded-lg px-4 py-2"
              required
            />
          </div>

          {/* Submit Button */}
          <button
            type="submit"
            className="w-full bg-green-600 text-white py-2 rounded-lg hover:bg-green-700 transition-colors flex items-center justify-center gap-2"
          >
            Submit
          </button>
        </form>

        <p className="mt-4 text-center text-sm text-gray-600">
          Learn about how carbon tax works and make your contribution towards a sustainable future.
        </p>
      </div>
    </div>
  );
}