import { describe, it, expect, beforeEach, vi } from 'vitest'
import { render, screen } from '@testing-library/react'
import App from '../App'

// Mock the services mit korrekter Struktur
vi.mock('../services', () => ({
  healthService: {
    getHealth: vi.fn(),
  },
  itemsService: {
    searchItems: vi.fn(),
    getItemDetails: vi.fn(),
  },
  ApiError: class ApiError extends Error {
    constructor(message: string) {
      super(message);
      this.name = 'ApiError';
    }
  }
}));

describe('App Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  describe('Initial Rendering', () => {
    it('should render the main application structure', () => {
      // Act
      render(<App />)

      // Assert - Check for main app components
      expect(screen.getByText(/EVE Profit Calculator/i)).toBeInTheDocument()
    })

    it('should render the ItemSearch component', () => {
      // Act
      render(<App />)

      // Assert - Check that ItemSearch is rendered
      expect(screen.getByRole('textbox')).toBeInTheDocument()
      expect(screen.getByPlaceholderText(/Search EVE items/i)).toBeInTheDocument()
    })

    it('should render with proper CSS classes', () => {
      // Act
      const { container } = render(<App />)

      // Assert - Check for proper structure
      expect(container.firstChild).toHaveClass('app')
    })
  })

  describe('Component Integration', () => {
    it('should integrate ItemSearch and other components properly', () => {
      // Act
      render(<App />)

      // Assert - Check that main components are present
      expect(screen.getByRole('textbox')).toBeInTheDocument()
      
      // Check for any header/title
      const headingElements = screen.getAllByRole('heading')
      expect(headingElements.length).toBeGreaterThan(0)
    })
  })

  describe('Application State', () => {
    it('should initialize without errors', () => {
      // This test ensures the app can render without throwing
      expect(() => render(<App />)).not.toThrow()
    })
  })
});
