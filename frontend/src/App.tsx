// EVE Profit Calculator 2.0 - Main Application
// Based on Universal Development Guidelines - Clean Code + React Best Practices

import { useState, useEffect } from 'react';
import { healthService, ApiError } from './services';
import type { HealthResponse, EveItem } from './services';
import { ItemSearch } from './components/ItemSearch';
import './App.css';
import './components/ItemSearch.css';

function App() {
  const [healthStatus, setHealthStatus] = useState<HealthResponse | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [selectedItem, setSelectedItem] = useState<EveItem | null>(null);

  // Clean Code: Single Responsibility - Health Check
  useEffect(() => {
    checkBackendHealth();
  }, []);

  const checkBackendHealth = async (): Promise<void> => {
    try {
      setIsLoading(true);
      setError(null);
      
      const health = await healthService.getHealthStatus();
      setHealthStatus(health);
    } catch (err) {
      if (err instanceof ApiError) {
        setError(`Backend Error: ${err.message}`);
      } else {
        setError('Unknown error occurred');
      }
    } finally {
      setIsLoading(false);
    }
  };

  // Clean Code: Single Responsibility - Item Selection Handler
  const handleItemSelect = (item: EveItem): void => {
    setSelectedItem(item);
  };

  return (
    <div className="app">
      <header className="app-header">
        <h1>🚀 EVE Profit Calculator 2.0</h1>
        <p>Modern Trading Analysis for EVE Online</p>
      </header>

      <main className="app-main">
        {/* Backend Status Section */}
        <section className="status-section">
          <h2>Backend Status</h2>
          
          {isLoading && <p>Checking backend connection...</p>}
          
          {error && (
            <div className="error-message">
              <p>❌ {error}</p>
              <button onClick={checkBackendHealth}>Retry Connection</button>
            </div>
          )}
          
          {healthStatus && (
            <div className="success-message">
              <p>✅ Backend Connected</p>
              <ul>
                <li>Status: {healthStatus.status}</li>
                <li>Timestamp: {new Date(healthStatus.time).toLocaleString()}</li>
              </ul>
            </div>
          )}
        </section>

        {/* EVE Item Search Section */}
        <section className="search-section">
          <h2>🔍 EVE Item Search</h2>
          <ItemSearch onItemSelect={handleItemSelect} />
          
          {selectedItem && (
            <div className="selected-item">
              <h3>Selected Item</h3>
              <div className="item-details">
                <p><strong>Name:</strong> {selectedItem.type_name}</p>
                <p><strong>Type ID:</strong> {selectedItem.type_id}</p>
                <p><strong>Group ID:</strong> {selectedItem.group_id}</p>
                <p><strong>Volume:</strong> {selectedItem.volume} m³</p>
                {selectedItem.mass && <p><strong>Mass:</strong> {selectedItem.mass} kg</p>}
                {selectedItem.description && <p><strong>Description:</strong> {selectedItem.description}</p>}
              </div>
            </div>
          )}
        </section>

        {/* EVE Features Section */}
        <section className="features-section">
          <h2>EVE Trading Features</h2>
          <div className="feature-grid">
            <div className="feature-card">
              <h3>📊 Item Search</h3>
              <p>Search EVE Online items and view market data</p>
              <button disabled={!healthStatus}>
                {healthStatus ? '✅ Available' : '❌ Backend Required'}
              </button>
            </div>
            
            <div className="feature-card">
              <h3>💰 Profit Calculator</h3>
              <p>Calculate trading profits between stations</p>
              <button disabled>Coming Soon</button>
            </div>
            
            <div className="feature-card">
              <h3>👨‍🚀 Character Data</h3>
              <p>EVE SSO integration for character information</p>
              <button disabled>Coming Soon</button>
            </div>
          </div>
        </section>
      </main>

      <footer className="app-footer">
        <p>Built with React + TypeScript + Clean Code Principles</p>
      </footer>
    </div>
  );
}

export default App;
