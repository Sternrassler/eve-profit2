// Type definitions for jest-dom matchers in Vitest
import type { TestingLibraryMatchers } from '@testing-library/jest-dom/matchers'

declare module 'vitest' {
  interface Assertion<T = unknown> extends jest.Matchers<void>, TestingLibraryMatchers<T, void> {}
  interface AsymmetricMatchersContaining extends TestingLibraryMatchers<unknown, void> {
    // This interface extends TestingLibraryMatchers for asymmetric matcher support
  }
}
