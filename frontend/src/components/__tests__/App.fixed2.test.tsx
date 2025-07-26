import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen } from '@testing-library/react'
import App from '../../App'

// Mock the services
vi.mock('../../services', () => ({
  healthService: {
    getHealth: vi.fn(),
  },
  itemsService: {
    searchItems: vi.fn(),
  },
  ApiError: class ApiError extends Error {
    constructor(message: string) {
      super(message);
      this.name = 'ApiError';
    }
  }
}))

describe('App Component - Fixed', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  describe('Initial Rendering', () => {
    it('should render the main application structure', () => {
      render(<App />)
      
      expect(screen.getByText(/EVE Profit Calculator 2.0/i)).toBeInTheDocument()
    })

    it('should render the ItemSearch component', () => {
      render(<App />)
      
      expect(screen.getByPlaceholderText(/Search EVE items/i)).toBeInTheDocument()
    })

    it('should render with proper CSS classes', () => {
      const { container } = render(<App />)
      
      expect(container.firstChild).toHaveClass('app')
    })
  })

  describe('Component Integration', () => {
    it('should integrate ItemSearch and other components properly', () => {
      render(<App />)
      
      // Check that main components are present
      expect(screen.getByText(/EVE Profit Calculator 2.0/i)).toBeInTheDocument()
      expect(screen.getByPlaceholderText(/Search EVE items/i)).toBeInTheDocument()
    })
  })

  describe('Application State', () => {
    it('should initialize without errors', () => {
      expect(() => render(<App />)).not.toThrow()
    })
  })
})
