import React, { useEffect } from "react";
import { Leaf, CreditCard } from "lucide-react";
import { PaymentModal } from "./PaymentModal";
import { carbonService } from "../services/carbon";
import { walletService } from "../services/wallet";
import { showErrorToast } from "../utils/toastUtils";

export function CarbonCalculator() {
  const [carbonTokens, setCarbonTokens] = React.useState(0); // Example amount
  const [isPaymentModalOpen, setIsPaymentModalOpen] = React.useState(false);
  const [currentBalance, setCurrentBalance] = React.useState(1000);
  const [isLoading, setIsLoading] = React.useState(true);

  useEffect(() => {
    const fetchAll = async () => {
      try {
        const response = await walletService.getWallet();
        if (response?.success) {
          setCurrentBalance(response.data?.token);
        } else {
          showErrorToast(response?.error);
        }
      } catch (err) {
        showErrorToast("Failed to fetch wallet balance");
      }
      try {
        const response = await carbonService.getCarbon();
        if (response?.success) {
          setCarbonTokens(response.data?.token);
        } else {
          showErrorToast(response?.error);
        }
      } catch (err) {
        showErrorToast("Failed to fetch carbon tokens");
      }
      setIsLoading(false);
    };

    // Call fetchAll every 5 seconds
    const interval = setInterval(fetchAll, 5000);

    // Cleanup interval on component unmount
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="text-center mb-12">
      <div className="inline-block bg-white rounded-2xl shadow-lg p-8">
        <div className="flex items-center justify-center mb-4">
          <Leaf className="text-green-600 mr-2" size={28} />
          <h2 className="text-2xl font-bold text-gray-800">
            Required Carbon Tokens
          </h2>
        </div>
        <div className="text-5xl font-bold text-green-600 mb-2">
          {carbonTokens === 0 ? "-" : carbonTokens}
        </div>
        <p className="text-gray-600 mb-6">
          tokens to offset your carbon footprint
        </p>

        <button
          onClick={() => setIsPaymentModalOpen(true)}
          disabled={carbonTokens === 0}
          className={`flex items-center justify-center gap-2 text-white px-8 py-3 rounded-lg transition-colors w-full ${
            carbonTokens === 0
              ? "bg-gray-400 cursor-not-allowed"
              : "bg-green-600 hover:bg-green-700"
          }`}
        >
          <CreditCard size={20} />
          <span>Pay {carbonTokens === 0 ? "" : carbonTokens} Tokens</span>
        </button>
      </div>
      {!isLoading && (
        <PaymentModal
          isOpen={isPaymentModalOpen}
          onClose={() => setIsPaymentModalOpen(false)}
          amount={carbonTokens}
          currentBalance={currentBalance}
        />
      )}
    </div>
  );
}
