import { describe, it, expect, beforeEach, vi } from 'vitest'
import { render, screen, waitFor } from '../test/test-utils'
import App from '../App'
import { createMockHealthResponse, createMockApiError } from '../test/mocks'

// Mock the HealthService
vi.mock('../services/HealthService', () => ({
  HealthService: {
    getHealth: vi.fn()
  }
}))

// Mock the ItemsService
vi.mock('../services/ItemsService', () => ({
  ItemsService: {
    searchItems: vi.fn(),
    getItemDetails: vi.fn()
  }
}))

describe('App Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  describe('Initial Rendering', () => {
    it('should render the main application structure', () => {
      // Arrange & Act
      render(<App />)

      // Assert
      expect(screen.getByRole('main')).toBeInTheDocument()
      expect(screen.getByText(/EVE Profit Calculator/i)).toBeInTheDocument()
      expect(screen.getByText(/Trading-Optimierung/i)).toBeInTheDocument()
    })

    it('should render the header with correct title', () => {
      // Arrange & Act
      render(<App />)

      // Assert
      const header = screen.getByRole('banner')
      expect(header).toBeInTheDocument()
      
      const title = screen.getByRole('heading', { level: 1 })
      expect(title).toHaveTextContent('EVE Profit Calculator 2.0')
    })

    it('should render the ItemSearch component', () => {
      // Arrange & Act
      render(<App />)

      // Assert
      expect(screen.getByRole('search')).toBeInTheDocument()
      expect(screen.getByPlaceholderText(/EVE Item suchen/i)).toBeInTheDocument()
    })
  })

  describe('Health Status Integration', () => {
    it('should display backend connection status when healthy', async () => {
      // Arrange
      const { HealthService } = await import('../services/HealthService')
      const mockHealthService = HealthService as any
      mockHealthService.getHealth.mockResolvedValue(createMockHealthResponse())

      // Act
      render(<App />)

      // Assert
      await waitFor(() => {
        expect(screen.getByText(/Backend: Connected/i)).toBeInTheDocument()
      })

      expect(mockHealthService.getHealth).toHaveBeenCalledTimes(1)
    })

    it('should display connection error when health check fails', async () => {
      // Arrange
      const { HealthService } = await import('../services/HealthService')
      const mockHealthService = HealthService as any
      mockHealthService.getHealth.mockRejectedValue(createMockApiError(500, 'Connection failed'))

      // Act
      render(<App />)

      // Assert
      await waitFor(() => {
        expect(screen.getByText(/Backend: Disconnected/i)).toBeInTheDocument()
      })

      expect(mockHealthService.getHealth).toHaveBeenCalledTimes(1)
    })

    it('should show loading state initially', () => {
      // Arrange
      const { HealthService } = await import('../services/HealthService')
      const mockHealthService = HealthService as any
      // Mock that never resolves to test loading state
      mockHealthService.getHealth.mockImplementation(() => new Promise(() => {}))

      // Act
      render(<App />)

      // Assert
      expect(screen.getByText(/Backend: Checking/i)).toBeInTheDocument()
    })
  })

  describe('Responsive Design', () => {
    it('should have mobile-friendly CSS classes', () => {
      // Arrange & Act
      render(<App />)

      // Assert
      const appContainer = screen.getByRole('main')
      expect(appContainer).toHaveClass('app')
      
      // Check if container has responsive styling
      const computedStyle = window.getComputedStyle(appContainer)
      expect(computedStyle.minHeight).toBe('100vh')
    })
  })

  describe('Accessibility', () => {
    it('should have proper ARIA landmarks', () => {
      // Arrange & Act
      render(<App />)

      // Assert
      expect(screen.getByRole('banner')).toBeInTheDocument() // header
      expect(screen.getByRole('main')).toBeInTheDocument()   // main content
    })

    it('should have semantic HTML structure', () => {
      // Arrange & Act
      render(<App />)

      // Assert
      expect(screen.getByRole('heading', { level: 1 })).toBeInTheDocument()
      expect(screen.getByRole('search')).toBeInTheDocument()
    })
  })

  describe('Error Boundaries', () => {
    it('should handle component errors gracefully', async () => {
      // This would require an error boundary implementation
      // For now, we just ensure the component doesn't crash
      expect(() => render(<App />)).not.toThrow()
    })
  })
})
