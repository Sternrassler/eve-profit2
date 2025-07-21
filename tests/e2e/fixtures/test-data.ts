/**
 * EVE Profit Calculator 2.0 - Test Data Factory
 * Universal Testing Guidelines konform
 */
import { faker } from '@faker-js/faker';

// EVE Online specific test data
export interface EveCharacterData {
  character_id: number;
  character_name: string;
  corporation_id: number;
  alliance_id?: number;
  security_status: number;
  location: {
    solar_system_id: number;
    station_id?: number;
    structure_id?: number;
  };
}

export interface EveItemData {
  type_id: number;
  name: string;
  group_id: number;
  category_id: number;
  volume: number;
  market_group_id?: number;
  published: boolean;
}

export interface MarketOrderData {
  order_id: number;
  type_id: number;
  location_id: number;
  is_buy_order: boolean;
  price: number;
  volume_remain: number;
  volume_total: number;
  min_volume: number;
  duration: number;
  issued: string;
  range: string;
}

// Test Data Generators
export function generateEveCharacter(overrides: Partial<EveCharacterData> = {}): EveCharacterData {
  const defaults: EveCharacterData = {
    character_id: faker.number.int({ min: 90000000, max: 98000000 }),
    character_name: faker.person.firstName() + ' ' + faker.person.lastName(),
    corporation_id: faker.number.int({ min: 1000000, max: 2000000 }),
    alliance_id: faker.number.int({ min: 99000000, max: 99999999 }),
    security_status: faker.number.float({ min: -10, max: 5, fractionDigits: 2 }),
    location: {
      solar_system_id: faker.number.int({ min: 30000001, max: 33000000 }),
      station_id: faker.number.int({ min: 60000000, max: 61000000 })
    }
  };
  
  return { ...defaults, ...overrides };
}

export function generateEveItem(overrides: Partial<EveItemData> = {}): EveItemData {
  const defaults: EveItemData = {
    type_id: faker.number.int({ min: 1, max: 50000 }),
    name: faker.commerce.productName(),
    group_id: faker.number.int({ min: 1, max: 1000 }),
    category_id: faker.number.int({ min: 1, max: 50 }),
    volume: faker.number.float({ min: 0.01, max: 1000, fractionDigits: 2 }),
    market_group_id: faker.number.int({ min: 1, max: 2000 }),
    published: true
  };
  
  return { ...defaults, ...overrides };
}

export function generateMarketOrder(overrides: Partial<MarketOrderData> = {}): MarketOrderData {
  const defaults: MarketOrderData = {
    order_id: faker.number.int({ min: 5000000000, max: 6000000000 }),
    type_id: faker.number.int({ min: 1, max: 50000 }),
    location_id: faker.number.int({ min: 60000000, max: 61000000 }),
    is_buy_order: faker.datatype.boolean(),
    price: faker.number.float({ min: 0.01, max: 1000000, fractionDigits: 2 }),
    volume_remain: faker.number.int({ min: 1, max: 1000000 }),
    volume_total: faker.number.int({ min: 1, max: 1000000 }),
    min_volume: 1,
    duration: faker.number.int({ min: 1, max: 90 }),
    issued: faker.date.recent().toISOString(),
    range: 'region'
  };
  
  return { ...defaults, ...overrides };
}

// Predefined EVE test data
export const eveTestData = {
  // Popular trading items
  items: {
    tritanium: generateEveItem({
      type_id: 34,
      name: 'Tritanium',
      group_id: 18,
      category_id: 4,
      volume: 0.01
    }),
    
    pyerite: generateEveItem({
      type_id: 35,
      name: 'Pyerite',
      group_id: 18,
      category_id: 4,
      volume: 0.01
    }),
    
    plex: generateEveItem({
      type_id: 11399,
      name: 'PLEX',
      group_id: 1875,
      category_id: 25,
      volume: 0.01
    })
  },
  
  // Test characters
  characters: {
    trader: generateEveCharacter({
      character_name: 'Test Trader',
      security_status: 5.0,
      location: {
        solar_system_id: 30000142, // Jita
        station_id: 60003760       // Jita IV - Moon 4 - Caldari Navy Assembly Plant
      }
    }),
    
    miner: generateEveCharacter({
      character_name: 'Test Miner',
      security_status: 2.5,
      location: {
        solar_system_id: 30000144  // New Caldari
      }
    })
  },
  
  // Trading hubs
  locations: {
    jita: {
      system_id: 30000142,
      station_id: 60003760,
      name: 'Jita IV - Moon 4 - Caldari Navy Assembly Plant'
    },
    amarr: {
      system_id: 30002187,
      station_id: 60008494,
      name: 'Amarr VIII (Oris) - Emperor Family Academy'
    },
    dodixie: {
      system_id: 30002659,
      station_id: 60011866,
      name: 'Dodixie IX - Moon 20 - Federation Navy Assembly Plant'
    }
  }
};

// Helper functions
export function createTestMarketData(typeId: number, locationId: number): {
  buyOrders: MarketOrderData[];
  sellOrders: MarketOrderData[];
} {
  const buyOrders = Array.from({ length: 5 }, () => 
    generateMarketOrder({
      type_id: typeId,
      location_id: locationId,
      is_buy_order: true,
      price: faker.number.float({ min: 100, max: 500, fractionDigits: 2 })
    })
  );
  
  const sellOrders = Array.from({ length: 5 }, () => 
    generateMarketOrder({
      type_id: typeId,
      location_id: locationId,
      is_buy_order: false,
      price: faker.number.float({ min: 520, max: 800, fractionDigits: 2 })
    })
  );
  
  return { buyOrders, sellOrders };
}

export function generateUniqueCharacter(): EveCharacterData {
  return generateEveCharacter({
    character_id: faker.number.int({ min: 90000000, max: 98000000 }),
    character_name: `Test_${faker.string.alphanumeric(8)}`
  });
}
