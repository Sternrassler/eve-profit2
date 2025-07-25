// EVE Profit Calculator 2.0 - API Client Configuration
// Based on Universal Development Guidelines - Clean Code + Error Handling

import axios from 'axios';

// Backend API Configuration
const API_BASE_URL = 'http://localhost:9000';
const API_VERSION = 'v1';

// API Client Instance with Clean Code principles
export class ApiClient {
  private readonly client = axios.create({
    baseURL: `${API_BASE_URL}/api/${API_VERSION}`,
    timeout: 10000, // 10 second timeout
    headers: {
      'Content-Type': 'application/json',
      'Accept': 'application/json',
    },
  });

  constructor() {
    // Request interceptor for logging (Development Guidelines)
    this.client.interceptors.request.use(
      (config) => {
        console.log(`[API] ${config.method?.toUpperCase()} ${config.url}`);
        return config;
      },
      (error) => {
        console.error('[API] Request error:', error);
        return Promise.reject(new Error(error));
      }
    );

    // Response interceptor for error handling
    this.client.interceptors.response.use(
      (response) => {
        console.log(`[API] ${response.status} ${response.config.url}`);
        return response;
      },
      (error) => {
        console.error('[API] Response error:', error.response?.status, error.message);
        return Promise.reject(this.handleApiError(error));
      }
    );
  }

  // Clean Code: Single Responsibility - Error handling
  private handleApiError(error: any): ApiError {
    if (error.response) {
      // Server responded with error status
      return new ApiError(
        error.response.status,
        error.response.data as string || error.message,
        'SERVER_ERROR'
      );
    } else if (error.request) {
      // Network error
      return new ApiError(
        0,
        'Network error - Backend server not reachable',
        'NETWORK_ERROR'
      );
    } else {
      // Request setup error
      return new ApiError(
        0,
        error.message,
        'REQUEST_ERROR'
      );
    }
  }

  // Public API method
  public async get<T>(endpoint: string): Promise<T> {
    const response = await this.client.get<T>(endpoint);
    return response.data;
  }

  public async post<T>(endpoint: string, data?: unknown): Promise<T> {
    const response = await this.client.post<T>(endpoint, data);
    return response.data;
  }
}

// Custom Error Class for API errors
export class ApiError extends Error {
  public readonly status: number;
  public readonly message: string;
  public readonly type: 'SERVER_ERROR' | 'NETWORK_ERROR' | 'REQUEST_ERROR';

  constructor(
    status: number,
    message: string,
    type: 'SERVER_ERROR' | 'NETWORK_ERROR' | 'REQUEST_ERROR'
  ) {
    super(message);
    this.name = 'ApiError';
    this.status = status;
    this.message = message;
    this.type = type;
  }
}

// Singleton instance following Clean Code principles
export const apiClient = new ApiClient();
