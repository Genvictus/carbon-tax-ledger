import { useState } from 'react';
import { X, CreditCard, Wallet } from 'lucide-react';
import { carbonService } from '../services/carbon';
import { showErrorToast } from '../utils/toastUtils';

interface PaymentModalProps {
  isOpen: boolean;
  onClose: () => void;
  amount: number;
  currentBalance: number;
}

export function PaymentModal({ isOpen, onClose, amount: suggestedAmount, currentBalance }: PaymentModalProps) {
  const [amount, setAmount] = useState(suggestedAmount);

  if (!isOpen) return null;

  const handlePayment = async () => {
    try {
      const response = await carbonService.payCarbonTax({ amount });
      if (response?.success) {
        window.location.reload();
      } else {
        showErrorToast(response?.error || 'Failed to make payment');
      }
    } catch (err) {
      showErrorToast('Failed to make payment');
    }
  };

  const insufficientFunds = currentBalance < amount;

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-white rounded-lg p-8 max-w-md w-full mx-4 relative">
        <button
          onClick={onClose}
          className="absolute top-4 right-4 text-gray-500 hover:text-gray-700"
        >
          <X size={24} />
        </button>
        
        <h2 className="text-2xl font-bold mb-6">Confirm Payment</h2>
        
        <div className="bg-gray-50 p-4 rounded-lg mb-6">
          <div className="flex items-center justify-between mb-4">
            <span className="text-gray-600">Current Balance</span>
            <div className="flex items-center">
              <Wallet className="text-green-600 mr-2" size={16} />
              <span className="font-semibold text-black">{currentBalance} Tokens</span>
            </div>
          </div>
          
          <div className="space-y-2">
            <label htmlFor="amount" className="block text-sm font-medium text-gray-600">
              Payment Amount
            </label>
            <div className="relative">
              <input
                id="amount"
                type="number"
                value={amount}
                onChange={(e) => setAmount(Number(e.target.value))}
                min="1"
                max={currentBalance}
                className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-transparent text-black"
              />
              <span className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-500">
                Tokens
              </span>
            </div>
          </div>
        </div>

        {insufficientFunds && (
          <p className="text-red-500 mb-4">
            Insufficient funds. Please top up your wallet or enter a smaller amount.
          </p>
        )}

        <button
          onClick={handlePayment}
          disabled={insufficientFunds || amount <= 0}
          className={`w-full flex items-center justify-center gap-2 py-3 rounded-lg transition-colors ${
            insufficientFunds || amount <= 0
              ? 'bg-gray-400 cursor-not-allowed'
              : 'bg-green-600 hover:bg-green-700'
          } text-white`}
        >
          <CreditCard size={20} />
          <span>Pay {amount} Tokens</span>
        </button>
      </div>
    </div>
  );
}