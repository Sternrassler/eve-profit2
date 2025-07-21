/**
 * EVE Profit Calculator 2.0 - Items API Tests
 * Items Search and Details API E2E Tests
 */
import { test, expect } from '@playwright/test';

test.describe('Items API Endpoints', () => {
  
  test('should get item details by ID', async ({ page }) => {
    // Test Tritanium (TypeID: 34)
    const response = await page.request.get('/api/v1/items/34');
    expect(response.status()).toBe(200);
    
    const itemData = await response.json();
    expect(itemData).toHaveProperty('success', true);
    expect(itemData).toHaveProperty('data');
    expect(itemData.data).toHaveProperty('type_id', 34);
    expect(itemData.data).toHaveProperty('type_name', 'Tritanium');
  });
  
  test('should return 400 for invalid item ID', async ({ page }) => {
    const response = await page.request.get('/api/v1/items/invalid');
    expect(response.status()).toBe(400);
    
    const errorData = await response.json();
    expect(errorData).toHaveProperty('success', false);
    expect(errorData).toHaveProperty('error');
  });
  
  test('should return 404 for non-existent item', async ({ page }) => {
    const response = await page.request.get('/api/v1/items/999999');
    expect(response.status()).toBe(404);
    
    const errorData = await response.json();
    expect(errorData).toHaveProperty('success', false);
    expect(errorData).toHaveProperty('error', 'Item not found');
  });
  
  test('should search items by name', async ({ page }) => {
    const response = await page.request.get('/api/v1/items/search?q=Tritanium');
    expect(response.status()).toBe(200);
    
    const searchData = await response.json();
    expect(searchData).toHaveProperty('success', true);
    expect(searchData).toHaveProperty('data');
    expect(Array.isArray(searchData.data)).toBe(true);
    expect(searchData.data.length).toBeGreaterThan(0);
    
    // Verify search results contain relevant items
    const firstItem = searchData.data[0];
    expect(firstItem).toHaveProperty('type_id');
    expect(firstItem).toHaveProperty('type_name');
    expect(firstItem.type_name.toLowerCase()).toContain('tritanium');
  });
  
  test('should return 400 for empty search query', async ({ page }) => {
    const response = await page.request.get('/api/v1/items/search?q=');
    expect(response.status()).toBe(400);
    
    const errorData = await response.json();
    expect(errorData).toHaveProperty('success', false);
    expect(errorData).toHaveProperty('error', 'Search query is required');
  });
  
  test('should return empty results for non-matching search', async ({ page }) => {
    const response = await page.request.get('/api/v1/items/search?q=NonExistentItemXYZ123');
    expect(response.status()).toBe(200);
    
    const searchData = await response.json();
    expect(searchData).toHaveProperty('success', true);
    expect(searchData).toHaveProperty('data');
    // API returns null for empty results, not empty array
    expect(searchData.data === null || (Array.isArray(searchData.data) && searchData.data.length === 0)).toBe(true);
  });
});

test.describe('Items API Performance', () => {
  
  test('item details should respond quickly', async ({ page }) => {
    const startTime = Date.now();
    const response = await page.request.get('/api/v1/items/34');
    const endTime = Date.now();
    
    expect(response.status()).toBe(200);
    expect(endTime - startTime).toBeLessThan(1000); // Under 1 second
  });
  
  test('item search should respond quickly', async ({ page }) => {
    const startTime = Date.now();
    const response = await page.request.get('/api/v1/items/search?q=Tritanium');
    const endTime = Date.now();
    
    expect(response.status()).toBe(200);
    expect(endTime - startTime).toBeLessThan(2000); // Under 2 seconds (mehr Toleranz für DB-Suche)
  });
});
