/**
 * EVE Profit Calculator 2.0 - Base Page Object
 * Universal Testing Guidelines konform
 */
import { Page, expect } from '@playwright/test';

export abstract class BasePage {
  readonly page: Page;
  
  constructor(page: Page) {
    this.page = page;
  }
  
  // Navigation methods
  async goto(path: string): Promise<void> {
    await this.page.goto(path);
  }
  
  async waitForPageLoad(): Promise<void> {
    await this.page.waitForLoadState('domcontentloaded');
  }
  
  // Common interaction methods
  async clickElement(selector: string): Promise<void> {
    await this.page.click(selector);
  }
  
  async fillInput(selector: string, text: string): Promise<void> {
    await this.page.fill(selector, text);
  }
  
  async getText(selector: string): Promise<string> {
    return await this.page.textContent(selector) || '';
  }
  
  // Wait helpers
  async waitForElement(selector: string): Promise<void> {
    await this.page.waitForSelector(selector);
  }
  
  async waitForUrl(urlPattern: string | RegExp): Promise<void> {
    await this.page.waitForURL(urlPattern);
  }
  
  // EVE-specific helpers
  async waitForEveApiResponse(): Promise<void> {
    // Warte auf EVE ESI API responses
    await this.page.waitForResponse(response => 
      response.url().includes('esi.evetech.net') && response.status() === 200
    );
  }
  
  async waitForMarketDataLoad(): Promise<void> {
    // Warte bis Market Data geladen ist
    await this.page.waitForFunction(() => {
      return document.querySelector('[data-testid="market-data-loaded"]') !== null;
    });
  }
  
  // Screenshot helper
  async takeScreenshot(name: string): Promise<void> {
    await this.page.screenshot({ 
      path: `test-results/screenshots/${name}.png`,
      fullPage: true 
    });
  }
  
  // Error handling
  async expectNoErrors(): Promise<void> {
    const errorElements = this.page.locator('[data-testid="error-message"]');
    await expect(errorElements).toHaveCount(0);
  }
  
  async expectLoadingComplete(): Promise<void> {
    const loadingElements = this.page.locator('[data-testid="loading"]');
    await expect(loadingElements).toBeHidden();
  }
}
