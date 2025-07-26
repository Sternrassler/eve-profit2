import React from 'react'

// Test Wrapper Component für React Testing
export const TestWrapper = ({ children }: { children: React.ReactNode }) => {
  return (
    <div data-testid="test-wrapper">
      {children}
    </div>
  )
}
