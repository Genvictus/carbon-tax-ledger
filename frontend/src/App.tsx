import { TreePine, Factory, Wind } from 'lucide-react';
import { LoginModal } from './components/LoginModal';
import { WalletBalance } from './components/WalletBalance';
import { CarbonCalculator } from './components/CarbonCalculator';
import { useAuth } from './context/AuthContext';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

function App() {
  const { isAuthenticated, logout } = useAuth();

  return (
    <>
      <ToastContainer />
      <div className="min-h-screen bg-gray-50">
        <LoginModal isOpen={!isAuthenticated} />
        
        {isAuthenticated && (
          <>
            {/* Header with Wallet */}
            <div className="bg-gray-50 py-4 px-6 shadow-sm">
              <div className="container mx-auto flex justify-between items-center">
                <button
                  onClick={logout}
                  className="text-gray-600 hover:text-gray-800 transition-colors"
                >
                  Sign Out
                </button>
                <WalletBalance />
              </div>
            </div>

            {/* Hero Section */}
            <div className="bg-gradient-to-b from-green-600 to-green-800 text-white py-16">
              <div className="container mx-auto px-4 text-center">
                <h1 className="text-4xl md:text-5xl font-bold mb-6">
                  Understanding Carbon Tax
                </h1>
                <p className="text-xl max-w-2xl mx-auto mb-8">
                  A crucial step towards reducing greenhouse gas emissions and fighting climate change
                </p>
                
                {/* Carbon Calculator */}
                <CarbonCalculator />
              </div>
            </div>

            {/* Main Content */}
            <div className="container mx-auto px-4 py-12">
              {/* Info Cards */}
              <div className="grid md:grid-cols-3 gap-8 mb-12">
                <div className="bg-white p-6 rounded-lg shadow-md">
                  <div className="flex items-center mb-4">
                    <TreePine className="text-green-600 mr-2" size={24} />
                    <h3 className="text-xl font-semibold">Environmental Impact</h3>
                  </div>
                  <p className="text-gray-600">
                    Carbon tax helps reduce greenhouse gas emissions by encouraging businesses and
                    individuals to adopt cleaner technologies and practices.
                  </p>
                </div>

                <div className="bg-white p-6 rounded-lg shadow-md">
                  <div className="flex items-center mb-4">
                    <Factory className="text-green-600 mr-2" size={24} />
                    <h3 className="text-xl font-semibold">How It Works</h3>
                  </div>
                  <p className="text-gray-600">
                    Companies pay a fee based on their carbon emissions, incentivizing them to
                    reduce their carbon footprint and invest in clean energy.
                  </p>
                </div>

                <div className="bg-white p-6 rounded-lg shadow-md">
                  <div className="flex items-center mb-4">
                    <Wind className="text-green-600 mr-2" size={24} />
                    <h3 className="text-xl font-semibold">Clean Future</h3>
                  </div>
                  <p className="text-gray-600">
                    Revenue from carbon tax supports renewable energy projects, research, and
                    development of sustainable technologies.
                  </p>
                </div>
              </div>
            </div>
          </>
        )}
      </div>
    </>
  );
}

export default App;