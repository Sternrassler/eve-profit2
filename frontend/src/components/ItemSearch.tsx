// EVE Profit Calculator 2.0 - Item Search Component
// Based on Universal Development Guidelines - Single Responsibility + Clean Code

import { useState } from 'react';
import { itemsService, ApiError } from '../services';
import type { EveItem } from '../services';

interface ItemSearchProps {
  onItemSelect?: (item: EveItem) => void;
}

export function ItemSearch({ onItemSelect }: Readonly<ItemSearchProps>) {
  const [query, setQuery] = useState<string>('');
  const [results, setResults] = useState<EveItem[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  // Clean Code: Single Responsibility - Search Logic
  const handleSearch = async () => {
    if (!query.trim()) {
      setError('Please enter a search term');
      return;
    }

    try {
      setIsLoading(true);
      setError(null);
      
      const items: EveItem[] = await itemsService.searchItems(query);
      setResults(items);
      
      if (items.length === 0) {
        setError(`No items found for "${query}"`);
      }
    } catch (err) {
      if (err instanceof ApiError) {
        setError(`Search failed: ${err.message}`);
      } else {
        setError('Unknown search error occurred');
      }
      setResults([]);
    } finally {
      setIsLoading(false);
    }
  };

  // Clean Code: Meaningful Names
  const handleItemClick = (item: EveItem): void => {
    if (onItemSelect) {
      onItemSelect(item);
    }
  };

  const handleKeyPress = (event: React.KeyboardEvent): void => {
    if (event.key === 'Enter') {
      handleSearch();
    }
  };

  return (
    <div className="item-search">
      <div className="search-input-group">
        <input
          type="text"
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          onKeyDown={handleKeyPress}
          placeholder="Search EVE items (e.g., Tritanium, Veldspar)"
          className="search-input"
          disabled={isLoading}
        />
        <button 
          onClick={handleSearch}
          disabled={isLoading || !query.trim()}
          className="search-button"
        >
          {isLoading ? '🔍 Searching...' : '🔍 Search'}
        </button>
      </div>

      {error && (
        <div className="search-error">
          ❌ {error}
        </div>
      )}

      {results.length > 0 && (
        <div className="search-results">
          <h3>Search Results ({results.length} items)</h3>
          <div className="items-grid">
            {results.map((item) => (
              <button
                key={item.type_id}
                className="item-card"
                onClick={() => handleItemClick(item)}
                type="button"
              >
                <h4>{item.type_name}</h4>
                <p>Type ID: {item.type_id}</p>
                <p>Group: {item.group_id}</p>
                <p>Volume: {item.volume} m³</p>
                {item.published ? (
                  <span className="published">📈 Published</span>
                ) : (
                  <span className="unpublished">� Database Item</span>
                )}
              </button>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}
