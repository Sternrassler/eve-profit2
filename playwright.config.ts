import { defineConfig, devices } from '@playwright/test';

/**
 * Read environment variables from file.
 * https://github.com/motdotla/dotenv
 */
// import dotenv from 'dotenv';
// import path from 'path';
// dotenv.config({ path: path.resolve(__dirname, '.env') });

/**
 * EVE Profit Calculator 2.0 - Playwright E2E Configuration
 * Following Universal Testing Guidelines for EVE-specific testing
 */
export default defineConfig({
  // Test-Verzeichnis gemäß Universal Testing Guidelines
  testDir: './tests/e2e',
  
  // Parallel Testing für bessere Performance
  fullyParallel: true,
  
  // CI-spezifische Einstellungen
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 2 : 0,
  workers: process.env.CI ? 1 : undefined,
  
  // Reporter für verschiedene Outputs
  reporter: [
    ['html'],
    ['json', { outputFile: 'test-results/results.json' }],
    ['junit', { outputFile: 'test-results/results.xml' }]
  ],
  
  // Global Test Settings
  use: {
    // EVE Profit Calculator Backend URL
    baseURL: process.env.BASE_URL || 'http://localhost:9000',
    
    // Screenshots bei Fehlern für Debugging
    screenshot: 'only-on-failure',
    
    // Video bei Fehlern für detaillierte Analyse
    video: 'retain-on-failure',
    
    // Trace für Debugging
    trace: 'on-first-retry',
    
    // Timeout-Einstellungen für EVE API calls
    navigationTimeout: 30000,
    actionTimeout: 10000,
    
    // User Agent für EVE Online ESI API
    userAgent: 'EVE-Profit-Calculator/2.0 (Playwright E2E Tests)'
  },

  // Browser-Konfiguration für Multi-Browser Testing
  projects: [
    {
      name: 'chromium',
      use: { ...devices['Desktop Chrome'] },
    },

    {
      name: 'firefox',
      use: { ...devices['Desktop Firefox'] },
    },

    {
      name: 'webkit',
      use: { ...devices['Desktop Safari'] },
    },
    
    // Mobile Testing für responsive Design
    {
      name: 'Mobile Chrome',
      use: { ...devices['Pixel 5'] },
    },
    
    {
      name: 'Mobile Safari',
      use: { ...devices['iPhone 12'] },
    },
  ],

  // EVE Profit Calculator Backend Development Server
  webServer: {
    command: 'cd backend && go run cmd/server/main.go',
    url: 'http://localhost:9000',
    reuseExistingServer: !process.env.CI,
    timeout: 120 * 1000, // 2 minutes for Go compilation
  },
});
