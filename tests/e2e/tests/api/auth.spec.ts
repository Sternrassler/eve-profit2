/**
 * EVE Profit Calculator 2.0 - Auth API Tests
 * Authentication and SSO Configuration E2E Tests
 */
import { test, expect } from '@playwright/test';

test.describe('Authentication API Endpoints', () => {
  
  test('auth login endpoint should return configuration', async ({ page }) => {
    const response = await page.request.get('/api/v1/auth/login');
    expect(response.status()).toBe(200);
    
    const authData = await response.json();
    expect(authData).toHaveProperty('message');
    expect(authData).toHaveProperty('client_id');
    expect(authData).toHaveProperty('callback');
    expect(authData).toHaveProperty('scopes');
    expect(authData).toHaveProperty('status', 'configured');
    
    // Verify EVE SSO configuration
    expect(authData.message).toContain('EVE SSO');
    expect(Array.isArray(authData.scopes)).toBe(true);
    expect(authData.scopes.length).toBeGreaterThan(0);
  });
  
  test('auth login should include required EVE scopes', async ({ page }) => {
    const response = await page.request.get('/api/v1/auth/login');
    expect(response.status()).toBe(200);
    
    const authData = await response.json();
    const scopes = authData.scopes;
    
    // Verify essential EVE Online scopes are present
    expect(scopes).toContain('publicData');
    expect(scopes.some((scope: string) => scope.includes('location'))).toBe(true);
    expect(scopes.some((scope: string) => scope.includes('skills'))).toBe(true);
    expect(scopes.some((scope: string) => scope.includes('wallet'))).toBe(true);
  });
  
  test('auth login should respond quickly', async ({ page }) => {
    const startTime = Date.now();
    const response = await page.request.get('/api/v1/auth/login');
    const endTime = Date.now();
    
    expect(response.status()).toBe(200);
    expect(endTime - startTime).toBeLessThan(500); // Under 500ms for config endpoint
  });
});
