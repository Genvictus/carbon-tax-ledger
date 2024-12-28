import React from 'react';
import { History } from 'lucide-react';
import { PaymentHistoryModal } from './PaymentHistoryModal';

export function PaymentHistory() {
  const [isHistoryModalOpen, setIsHistoryModalOpen] = React.useState(false);

  return (
    <>
      <button
        onClick={() => setIsHistoryModalOpen(true)}
        className="flex items-center gap-1 bg-gray-600 text-white px-3 py-2 rounded-lg hover:bg-gray-700 transition-colors"
      >
        <History size={18} />
        History
      </button>

      <PaymentHistoryModal
        isOpen={isHistoryModalOpen}
        onClose={() => setIsHistoryModalOpen(false)}
      />
    </>
  );
}