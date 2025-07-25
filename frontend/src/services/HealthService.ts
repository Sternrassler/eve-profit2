// EVE Profit Calculator 2.0 - Health API Service
// Based on Universal Development Guidelines - Single Responsibility

import { apiClient } from './ApiClient';
import type { HealthResponse } from './types';

export class HealthService {
  // Clean Code: Meaningful Names
  public async getHealthStatus(): Promise<HealthResponse> {
    return apiClient.get<HealthResponse>('/health');
  }

  public async checkDatabaseConnection(): Promise<{ status: string }> {
    return apiClient.get<{ status: string }>('/sde/test');
  }

  public async checkEsiConnection(): Promise<{ status: string }> {
    return apiClient.get<{ status: string }>('/esi/test');
  }
}

export const healthService = new HealthService();
