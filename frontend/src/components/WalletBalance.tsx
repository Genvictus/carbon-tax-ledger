import React, { useEffect } from 'react';
import { Wallet, Plus } from 'lucide-react';
import { TopUpModal } from './TopUpModal';
import { PaymentHistory } from './PaymentHistory';
import { walletService } from '../services/wallet';
import { showErrorToast } from '../utils/toastUtils';

export function WalletBalance() {
  const [balance, setBalance] = React.useState(0); // Example initial balance
  const [isTopUpModalOpen, setIsTopUpModalOpen] = React.useState(false);

  useEffect(() => {
    const fetchBalance = async () => {
      try {
        const response = await walletService.getWallet();
        if (response?.success) {
          setBalance(response.data?.token);
        } else {
          showErrorToast(response?.error);
        }
      }
      catch (err) {
        showErrorToast("Failed to fetch wallet balance");
      }
    };

    fetchBalance();
  }, []);

  return (
    <>
      <div className="flex items-center gap-4">
        {/* <PaymentHistory /> */}
        <button
          onClick={() => setIsTopUpModalOpen(true)}
          className="flex items-center gap-1 bg-green-600 text-white px-3 py-2 rounded-lg hover:bg-green-700 transition-colors"
        >
          <Plus size={18} />
          Top Up
        </button>
        <div className="flex items-center bg-white rounded-lg px-4 py-2 shadow-sm">
          <Wallet className="text-green-600 mr-2" size={20} />
          <span className="font-semibold">{balance} Tokens</span>
        </div>
      </div>

      <TopUpModal
        isOpen={isTopUpModalOpen}
        onClose={() => setIsTopUpModalOpen(false)}
        currentBalance={balance}
      />
    </>
  );
}