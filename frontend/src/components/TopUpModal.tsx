import React from 'react';
import { X, CreditCard, Wallet } from 'lucide-react';
import { showErrorToast } from '../utils/toastUtils';
import { walletService } from '../services/wallet';

interface TopUpModalProps {
  isOpen: boolean;
  onClose: () => void;
  currentBalance: number;
}

export function TopUpModal({ isOpen, onClose, currentBalance }: TopUpModalProps) {
  const [amount, setAmount] = React.useState(100);
  
  if (!isOpen) return null;

  const handleTopUp = async () => {
    try {
      const response = await walletService.topup({ amount });
      if (response?.success) {
        window.location.reload();
      } else {
        showErrorToast(response?.error || 'Failed to top up');
      }
    } catch (err) {
      showErrorToast('Failed to top up');
    }
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-white rounded-lg p-8 max-w-md w-full mx-4 relative">
        <button
          onClick={onClose}
          className="absolute top-4 right-4 text-gray-500 hover:text-gray-700"
        >
          <X size={24} />
        </button>
        
        <h2 className="text-2xl font-bold mb-6">Top Up Carbon Tokens</h2>
        
        <div className="bg-gray-50 p-4 rounded-lg mb-6">
          <div className="flex items-center justify-between mb-2">
            <span className="text-gray-600">Current Balance</span>
            <div className="flex items-center">
              <Wallet className="text-green-600 mr-2" size={16} />
              <span className="font-semibold">{currentBalance} Tokens</span>
            </div>
          </div>
        </div>

        <div className="mb-6">
          <label className="block text-sm font-medium text-gray-700 mb-2">
            Amount to Top Up
          </label>
          <input
            type="number"
            value={amount}
            onChange={(e) => setAmount(Number(e.target.value))}
            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-transparent"
            min="1"
          />
        </div>

        <button
          onClick={handleTopUp}
          className="w-full flex items-center justify-center gap-2 bg-green-600 text-white py-3 rounded-lg hover:bg-green-700 transition-colors"
        >
          <CreditCard size={20} />
          <span>Pay {amount} Tokens</span>
        </button>
      </div>
    </div>
  );
}