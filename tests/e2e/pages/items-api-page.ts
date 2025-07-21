/**
 * EVE Profit Calculator 2.0 - Items API Page Object
 * Für EVE Items und Market Data Testing
 */
import { expect, Page } from '@playwright/test';
import { BasePage } from './base-page';

export class ItemsApiPage extends BasePage {
  constructor(page: Page) {
    super(page);
  }
  
  // Items API Tests
  async expectItemsListResponse(): Promise<void> {
    const response = await this.page.request.get('/api/items');
    expect(response.status()).toBe(200);
    
    const itemsData = await response.json();
    expect(Array.isArray(itemsData)).toBeTruthy();
    
    if (itemsData.length > 0) {
      const firstItem = itemsData[0];
      expect(firstItem).toHaveProperty('type_id');
      expect(firstItem).toHaveProperty('name');
      expect(firstItem).toHaveProperty('group_id');
    }
  }
  
  async expectSpecificItem(typeId: number): Promise<any> {
    const response = await this.page.request.get(`/api/items/${typeId}`);
    expect(response.status()).toBe(200);
    
    const itemData = await response.json();
    expect(itemData).toHaveProperty('type_id', typeId);
    expect(itemData).toHaveProperty('name');
    expect(itemData).toHaveProperty('group_id');
    
    return itemData;
  }
  
  async expectItemNotFound(typeId: number): Promise<void> {
    const response = await this.page.request.get(`/api/items/${typeId}`);
    expect(response.status()).toBe(404);
  }
  
  // Search functionality
  async searchItems(searchTerm: string): Promise<any[]> {
    const response = await this.page.request.get(`/api/items/search?q=${encodeURIComponent(searchTerm)}`);
    expect(response.status()).toBe(200);
    
    const searchResults = await response.json();
    expect(Array.isArray(searchResults)).toBeTruthy();
    
    // Verify search results contain the search term
    if (searchResults.length > 0) {
      const hasMatchingResult = searchResults.some(item => 
        item.name.toLowerCase().includes(searchTerm.toLowerCase())
      );
      expect(hasMatchingResult).toBeTruthy();
    }
    
    return searchResults;
  }
  
  // EVE-specific item tests
  async expectTritaniumItem(): Promise<void> {
    const tritaniumTypeId = 34; // Tritanium in EVE Online
    const tritanium = await this.expectSpecificItem(tritaniumTypeId);
    
    expect(tritanium.name).toBe('Tritanium');
    expect(tritanium.group_id).toBe(18); // Mineral group
  }
  
  async expectPopularTradingItems(): Promise<void> {
    const popularItems = [
      34,   // Tritanium
      35,   // Pyerite
      36,   // Mexallon
      37,   // Isogen
      11399 // PLEX
    ];
    
    for (const typeId of popularItems) {
      await this.expectSpecificItem(typeId);
    }
  }
  
  // Performance tests
  async expectFastItemsResponse(maxMs: number = 500): Promise<void> {
    const startTime = Date.now();
    await this.expectItemsListResponse();
    const responseTime = Date.now() - startTime;
    
    expect(responseTime).toBeLessThan(maxMs);
  }
  
  async expectFastItemSearch(searchTerm: string, maxMs: number = 300): Promise<void> {
    const startTime = Date.now();
    await this.searchItems(searchTerm);
    const responseTime = Date.now() - startTime;
    
    expect(responseTime).toBeLessThan(maxMs);
  }
}
