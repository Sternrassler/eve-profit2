// EVE Profit Calculator 2.0 - Items API Service
// Based on Universal Development Guidelines - Single Responsibility

import { apiClient } from './ApiClient';
import type { EveItem, ItemSearchResponse } from './types';

export class ItemsService {
  // Clean Code: Meaningful Names & Single Responsibility
  public async getItemById(itemId: number): Promise<EveItem> {
    const response = await apiClient.get<{ success: boolean; data: EveItem }>(`/items/${itemId}`);
    return response.data;
  }

  public async searchItems(query: string): Promise<EveItem[]> {
    const response = await apiClient.get<ItemSearchResponse>(`/items/search?q=${encodeURIComponent(query)}`);
    return response.data;
  }

  // Helper method for EVE-specific item lookup
  public async findTritanium(): Promise<EveItem> {
    return this.getItemById(34); // Tritanium Type ID
  }
}

export const itemsService = new ItemsService();
