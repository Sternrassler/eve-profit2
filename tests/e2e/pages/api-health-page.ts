/**
 * EVE Profit Calculator 2.0 - API Health Page Object
 * Für Backend API Testing
 */
import { expect, Page } from '@playwright/test';
import { BasePage } from './base-page';

export class ApiHealthPage extends BasePage {
  constructor(page: Page) {
    super(page);
  }
  
  // Navigation
  async navigateToHealth(): Promise<void> {
    await this.goto('/api/v1/health');
    await this.waitForPageLoad();
  }
  
  async navigateToHealthz(): Promise<void> {
    await this.goto('/healthz');
    await this.waitForPageLoad();
  }
  
  // API Health Checks
  async expectHealthResponse(): Promise<void> {
    const response = await this.page.request.get('/api/v1/health');
    expect(response.status()).toBe(200);
    
    const healthData = await response.json();
    expect(healthData).toHaveProperty('status');
    expect(healthData.status).toBe('healthy');
  }

  async expectDatabaseConnection(): Promise<void> {
    // Test SDE database connection via dedicated endpoint
    const response = await this.page.request.get('/api/v1/sde/test');
    expect(response.status()).toBe(200);
    
    const sdeData = await response.json();
    expect(sdeData).toHaveProperty('sde_status');
    expect(sdeData.sde_status).toBe('connected');
  }

  async expectEveApiConnection(): Promise<void> {
    // Test EVE ESI API connection via dedicated endpoint  
    const response = await this.page.request.get('/api/v1/esi/test');
    expect(response.status()).toBe(200);
    
    const esiData = await response.json();
    expect(esiData).toHaveProperty('esi_status');
    expect(esiData.esi_status).toBe('connected');
  }

  async expectResponseTime(maxMs: number = 1000): Promise<void> {
    const startTime = Date.now();
    const response = await this.page.request.get('/api/v1/health');
    const responseTime = Date.now() - startTime;
    
    expect(response.status()).toBe(200);
    expect(responseTime).toBeLessThan(maxMs);
  }
}
