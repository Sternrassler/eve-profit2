// TDD Tests für ItemSearch Komponente - Angepasst an aktuelle Implementierung
import { describe, it, expect, beforeEach, vi } from 'vitest';
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { ItemSearch } from '../ItemSearch';
import { itemsService } from '../../services';
import type { EveItem } from '../../services';

// Mock the services with proper typing
vi.mock('../../services', () => ({
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

const mockItemsService = {
  searchItems: vi.fn(),
  getItemDetails: vi.fn(),
} as typeof itemsService & {
  searchItems: ReturnType<typeof vi.fn>;
  getItemDetails: ReturnType<typeof vi.fn>;
};

// Test constants
const TEST_CONSTANTS = {
  MOCK_SEARCH_TERM: 'tritanium'
};

// Mock data
const createMockEVEItem = (overrides: Partial<EveItem> = {}): EveItem => ({
  type_id: 34,
  type_name: 'Tritanium',
  group_id: 18,
  published: true,
  volume: 0.01,
  ...overrides
});

const createMockSearchResults = (): EveItem[] => [
  createMockEVEItem(),
  createMockEVEItem({ 
    type_id: 35, 
    type_name: 'Pyerite',
    volume: 0.013
  })
];

describe('ItemSearch Component', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  describe('Initial Rendering', () => {
    it('should render search components', () => {
      render(<ItemSearch />);

      expect(screen.getByRole('textbox')).toBeInTheDocument();
      expect(screen.getByPlaceholderText(/Search EVE items/i)).toBeInTheDocument();
      expect(screen.getByText(/🔍 Search/i)).toBeInTheDocument();
    });

    it('should have search button disabled initially', () => {
      render(<ItemSearch />);

      const searchButton = screen.getByText(/🔍 Search/i);
      expect(searchButton).toBeDisabled();
    });
  });

  describe('Search Button Interaction', () => {
    it('should enable search button when text is entered', async () => {
      const user = userEvent.setup();
      render(<ItemSearch />);

      const searchInput = screen.getByRole('textbox');
      const searchButton = screen.getByText(/🔍 Search/i);

      expect(searchButton).toBeDisabled();

      await user.type(searchInput, TEST_CONSTANTS.MOCK_SEARCH_TERM);

      expect(searchButton).not.toBeDisabled();
    });

    it('should call search API when search button is clicked', async () => {
      const user = userEvent.setup();
      const mockResults = createMockSearchResults();
      mockItemsService.searchItems.mockResolvedValue(mockResults);

      render(<ItemSearch />);
      
      const searchInput = screen.getByRole('textbox');
      const searchButton = screen.getByText(/🔍 Search/i);

      await user.type(searchInput, TEST_CONSTANTS.MOCK_SEARCH_TERM);
      await user.click(searchButton);

      expect(mockItemsService.searchItems).toHaveBeenCalledWith(TEST_CONSTANTS.MOCK_SEARCH_TERM);
    });
  });

  describe('Keyboard Navigation', () => {
    it('should trigger search on Enter key press', async () => {
      const user = userEvent.setup();
      const mockResults = createMockSearchResults();
      mockItemsService.searchItems.mockResolvedValue(mockResults);

      render(<ItemSearch />);
      
      const searchInput = screen.getByRole('textbox');

      await user.type(searchInput, TEST_CONSTANTS.MOCK_SEARCH_TERM);
      await user.keyboard('{Enter}');

      expect(mockItemsService.searchItems).toHaveBeenCalledWith(TEST_CONSTANTS.MOCK_SEARCH_TERM);
    });
  });

  describe('Search Results Display', () => {
    it('should display search results after successful search', async () => {
      const user = userEvent.setup();
      const mockResults = createMockSearchResults();
      mockItemsService.searchItems.mockResolvedValue(mockResults);

      render(<ItemSearch />);
      
      const searchInput = screen.getByRole('textbox');
      const searchButton = screen.getByText(/🔍 Search/i);

      await user.type(searchInput, TEST_CONSTANTS.MOCK_SEARCH_TERM);
      await user.click(searchButton);

      await waitFor(() => {
        expect(screen.getByText(/Search Results/i)).toBeInTheDocument();
        expect(screen.getByText(mockResults[0].type_name)).toBeInTheDocument();
        expect(screen.getByText(mockResults[1].type_name)).toBeInTheDocument();
      });
    });

    it('should display item details in results', async () => {
      const user = userEvent.setup();
      const mockResults = createMockSearchResults();
      mockItemsService.searchItems.mockResolvedValue(mockResults);

      render(<ItemSearch />);
      
      const searchInput = screen.getByRole('textbox');
      const searchButton = screen.getByText(/🔍 Search/i);

      await user.type(searchInput, TEST_CONSTANTS.MOCK_SEARCH_TERM);
      await user.click(searchButton);

      await waitFor(() => {
        expect(screen.getByText(`Type ID: ${mockResults[0].type_id}`)).toBeInTheDocument();
        expect(screen.getByText(`Volume: ${mockResults[0].volume} m³`)).toBeInTheDocument();
      });
    });
  });

  describe('Loading States', () => {
    it('should show loading state during search', async () => {
      const user = userEvent.setup();
      let resolveSearch: (value: EveItem[]) => void;
      const searchPromise = new Promise<EveItem[]>((resolve) => {
        resolveSearch = resolve;
      });
      mockItemsService.searchItems.mockReturnValue(searchPromise);

      render(<ItemSearch />);
      
      const searchInput = screen.getByRole('textbox');
      const searchButton = screen.getByText(/🔍 Search/i);

      await user.type(searchInput, TEST_CONSTANTS.MOCK_SEARCH_TERM);
      await user.click(searchButton);

      // Should show loading state
      expect(screen.getByText(/🔍 Searching.../i)).toBeInTheDocument();
      expect(searchInput).toBeDisabled();

      // Resolve the promise
      resolveSearch!(createMockSearchResults());

      await waitFor(() => {
        expect(screen.getByText(/🔍 Search/i)).toBeInTheDocument();
      });
    });
  });

  describe('Error Handling', () => {
    it('should display error message on search failure', async () => {
      const user = userEvent.setup();
      const errorMessage = 'Search failed';
      mockItemsService.searchItems.mockRejectedValue(new Error(errorMessage));

      render(<ItemSearch />);
      
      const searchInput = screen.getByRole('textbox');
      const searchButton = screen.getByText(/🔍 Search/i);

      await user.type(searchInput, TEST_CONSTANTS.MOCK_SEARCH_TERM);
      await user.click(searchButton);

      await waitFor(() => {
        expect(screen.getByText(/Unknown search error occurred/i)).toBeInTheDocument();
      });
    });

    it('should display error for empty search term', async () => {
      const user = userEvent.setup();
      render(<ItemSearch />);
      
      const searchInput = screen.getByRole('textbox');
      
      // Type some text then clear it
      await user.type(searchInput, 'test');
      await user.clear(searchInput);
      await user.keyboard('{Enter}');

      await waitFor(() => {
        expect(screen.getByText(/Please enter a search term/i)).toBeInTheDocument();
      });
    });

    it('should display no results message when search returns empty', async () => {
      const user = userEvent.setup();
      mockItemsService.searchItems.mockResolvedValue([]);

      render(<ItemSearch />);
      
      const searchInput = screen.getByRole('textbox');
      const searchButton = screen.getByText(/🔍 Search/i);

      await user.type(searchInput, 'nonexistent');
      await user.click(searchButton);

      await waitFor(() => {
        expect(screen.getByText(/No items found for "nonexistent"/i)).toBeInTheDocument();
      });
    });
  });

  describe('Item Selection', () => {
    it('should call onItemSelect when item is clicked', async () => {
      const onItemSelect = vi.fn();
      const user = userEvent.setup();
      const mockResults = createMockSearchResults();
      mockItemsService.searchItems.mockResolvedValue(mockResults);

      render(<ItemSearch onItemSelect={onItemSelect} />);
      
      const searchInput = screen.getByRole('textbox');
      const searchButton = screen.getByText(/🔍 Search/i);

      await user.type(searchInput, TEST_CONSTANTS.MOCK_SEARCH_TERM);
      await user.click(searchButton);

      await waitFor(() => {
        expect(screen.getByText(mockResults[0].type_name)).toBeInTheDocument();
      });

      const itemButton = screen.getByText(mockResults[0].type_name).closest('button');
      expect(itemButton).toBeInTheDocument();
      
      await user.click(itemButton!);

      expect(onItemSelect).toHaveBeenCalledWith(mockResults[0]);
    });
  });
});
