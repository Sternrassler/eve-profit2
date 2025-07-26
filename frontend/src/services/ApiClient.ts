// EVE Profit Calculator 2.0 - API Client Configuration
// Based on Universal Development Guidelines - Clean Code + Error Handling

import axios from 'axios';

// Backend API Configuration
const API_BASE_URL = 'http://localhost:9000';
const API_VERSION = 'v1';

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
  private handleApiError(error: unknown): ApiError {
    const axiosError = error as { 
      response?: { status: number; data?: string }; 
      message?: string; 
      request?: unknown 
    };
    
    if (axiosError.response) {
      // Server responded with error status
      return new ApiError(
        axiosError.response.status,
        axiosError.response.data as string || axiosError.message || 'Unknown error',
        'SERVER_ERROR'
      );
    } else if (axiosError.request) {
      // Request was made but no response received
      return new ApiError(
        0,
        'No response from server',
        'NETWORK_ERROR'
      );
    } else {
      // Something else happened
      return new ApiError(
        0,
        axiosError.message || 'Request setup error',
        'REQUEST_ERROR'
      );
    }
  }

  // Public API methods
  public async get<T>(endpoint: string): Promise<T> {
    const response = await this.client.get<T>(endpoint);
    return response.data;
  }

  public async post<T>(endpoint: string, data?: unknown): Promise<T> {
    const response = await this.client.post<T>(endpoint, data);
    return response.data;
  }
}

// Singleton instance
export const apiClient = new ApiClient();
