// Mock Data Factory für EVE-spezifische Test Data
export const createMockEVEItem = (overrides = {}) => ({
  type_id: 34,
  type_name: 'Tritanium',
  group_id: 18,
  category_id: 4,
  published: 1,
  market_group_id: 1857,
  mass: 1.0,
  volume: 0.01,
  capacity: 0.0,
  ...overrides
})

export const createMockSearchResults = (count = 3) => {
  const baseItems = [
    { type_id: 34, type_name: 'Tritanium' },
    { type_id: 35, type_name: 'Pyerite' },
    { type_id: 36, type_name: 'Mexallon' }
  ]
  
  return baseItems.slice(0, count).map((item, index) => ({
    ...createMockEVEItem(item),
    type_id: item.type_id + index
  }))
}

export const createMockApiError = (status = 500, message = 'Internal Server Error') => ({
  response: {
    status,
    data: {
      error: message,
      success: false
    }
  },
  message,
  isAxiosError: true
})

export const createMockHealthResponse = (overrides = {}) => ({
  success: true,
  data: {
    status: 'healthy',
    timestamp: new Date().toISOString(),
    database: 'connected',
    esi: 'available',
    ...overrides
  }
})

// Helper function to wait for async operations
export const waitForApiCall = () => new Promise(resolve => setTimeout(resolve, 0))

// Mock user event helpers
export const mockUserInteraction = {
  search: async (searchTerm: string) => ({
    query: searchTerm,
    results: createMockSearchResults()
  }),
  
  clickItem: async (itemId: number) => ({
    selectedItem: createMockEVEItem({ type_id: itemId })
  })
}

// Test constants
export const TEST_CONSTANTS = {
  API_BASE_URL: 'http://localhost:9000/api/v1',
  MOCK_ITEM_ID: 34,
  MOCK_SEARCH_TERM: 'tritanium',
  DEBOUNCE_DELAY: 300
}
