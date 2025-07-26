import type { ReactElement } from 'react'
import { render } from '@testing-library/react'
import type { RenderOptions } from '@testing-library/react'
import { TestWrapper } from './TestWrapper'

// Custom render function for testing React components
const customRender = (
  ui: ReactElement,
  options?: Omit<RenderOptions, 'wrapper'>
) => render(ui, { wrapper: TestWrapper, ...options })

// Re-export everything from testing library
export { customRender as render }
export { screen, waitFor, fireEvent, within } from '@testing-library/react'
export { default as userEvent } from '@testing-library/user-event'

// Export test utilities and mocks
export {
  createMockEVEItem,
  createMockSearchResults,
  createMockApiError,
  createMockHealthResponse,
  waitForApiCall,
  mockUserInteraction,
  TEST_CONSTANTS
} from './mocks'
