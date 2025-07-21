/**
 * EVE Profit Calculator 2.0 - API Health Tests
 * Backend API Health Checks
 */
import { test, expect } from '@playwright/test';
import { ApiHealthPage } from '../../pages/api-health-page';

test.describe('API Health Endpoints', () => {
  let healthPage: ApiHealthPage;
  
  test.beforeEach(async ({ page }) => {
    healthPage = new ApiHealthPage(page);
  });
  
  test('health endpoint should return healthy status', async () => {
    // Act & Assert
    await healthPage.expectHealthResponse();
  });
  
  test('health endpoint should confirm database connection', async () => {
    // Act & Assert
    await healthPage.expectDatabaseConnection();
  });
  
  test('health endpoint should confirm EVE API connectivity', async () => {
    // Act & Assert
    await healthPage.expectEveApiConnection();
  });
  
  test('health endpoint should respond quickly', async () => {
    // Act & Assert - Expect response within 100ms
    await healthPage.expectResponseTime(100);
  });
  
  test('api root endpoint should be accessible', async ({ page }) => {
    const response = await page.request.get('/');
    expect(response.status()).toBe(200);
    
    const apiData = await response.json();
    expect(apiData).toHaveProperty('name');
    expect(apiData).toHaveProperty('version');
    expect(apiData).toHaveProperty('status');
    expect(apiData.status).toBe('running');
  });
});

test.describe('API Performance', () => {
  let healthPage: ApiHealthPage;
  
  test.beforeEach(async ({ page }) => {
    healthPage = new ApiHealthPage(page);
  });
  
  test('health checks should be consistently fast', async () => {
    // Arrange - Run multiple health checks
    const runs = 5;
    const results: number[] = [];
    
    // Act - Measure response times
    for (let i = 0; i < runs; i++) {
      const startTime = Date.now();
      await healthPage.expectHealthResponse();
      const responseTime = Date.now() - startTime;
      results.push(responseTime);
    }
    
    // Assert - All responses should be fast
    const averageTime = results.reduce((a, b) => a + b, 0) / results.length;
    expect(averageTime).toBeLessThan(50); // Average under 50ms
    
    const maxTime = Math.max(...results);
    expect(maxTime).toBeLessThan(200); // No response over 200ms
  });
});
