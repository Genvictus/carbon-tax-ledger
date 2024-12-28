import { X, ArrowUpRight, ArrowDownRight } from 'lucide-react';

interface PaymentHistoryModalProps {
  isOpen: boolean;
  onClose: () => void;
}

// Example transaction data
const transactions = [
  { id: 1, type: 'payment', amount: 245, date: '2024-03-10', description: 'Carbon Tax Payment' },
  { id: 2, type: 'topup', amount: 500, date: '2024-03-09', description: 'Wallet Top Up' },
  { id: 3, type: 'payment', amount: 180, date: '2024-03-08', description: 'Carbon Tax Payment' },
  { id: 4, type: 'topup', amount: 1000, date: '2024-03-07', description: 'Wallet Top Up' },
];

export function PaymentHistoryModal({ isOpen, onClose }: PaymentHistoryModalProps) {
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-white rounded-lg p-8 max-w-2xl w-full mx-4 relative max-h-[80vh] overflow-hidden flex flex-col">
        <button
          onClick={onClose}
          className="absolute top-4 right-4 text-gray-500 hover:text-gray-700"
        >
          <X size={24} />
        </button>
        
        <h2 className="text-2xl font-bold mb-6">Payment History</h2>
        
        <div className="overflow-y-auto flex-1">
          <div className="space-y-4">
            {transactions.map((transaction) => (
              <div
                key={transaction.id}
                className="flex items-center justify-between p-4 bg-gray-50 rounded-lg"
              >
                <div className="flex items-center gap-3">
                  {transaction.type === 'payment' ? (
                    <ArrowUpRight className="text-red-500" size={20} />
                  ) : (
                    <ArrowDownRight className="text-green-500" size={20} />
                  )}
                  <div>
                    <p className="font-medium">{transaction.description}</p>
                    <p className="text-sm text-gray-500">{transaction.date}</p>
                  </div>
                </div>
                <span className={`font-semibold ${
                  transaction.type === 'payment' ? 'text-red-500' : 'text-green-500'
                }`}>
                  {transaction.type === 'payment' ? '-' : '+'}
                  {transaction.amount} Tokens
                </span>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}