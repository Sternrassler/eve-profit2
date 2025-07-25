// EVE Profit Calculator 2.0 - Type Definitions
// Based on Universal Development Guidelines - Clean Code + Type Safety

// EVE Item Types (Backend SDE Models)
export interface EveItem {
  type_id: number;
  type_name: string;
  group_id: number;
  volume: number;
  published?: boolean;
  mass?: number;
  description?: string;
}

// API Response Types (matching backend structure)
export interface ApiResponse<T> {
  success: boolean;
  data: T;
  message?: string;
}

export interface HealthResponse {
  status: string;
  time: string;
}

export interface ItemSearchResponse {
  success: boolean;
  data: EveItem[];
}
