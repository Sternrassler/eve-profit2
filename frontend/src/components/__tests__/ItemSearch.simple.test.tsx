// TDD Demo: Einfacher Start mit ItemSearch
import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import { ItemSearch } from '../ItemSearch';

describe('ItemSearch TDD Demo', () => {
  describe('Grundlegendes Rendering', () => {
    it('kann ohne Fehler gerendert werden', () => {
      // Test 1: Red - Einfachster Test
      expect(() => render(<ItemSearch />)).not.toThrow();
    });

    // Test 2: Green - Überprüfung grundlegender UI-Elemente
    it('zeigt ein Eingabefeld', () => {
      render(<ItemSearch />);
      
      const inputElement = screen.getByRole('textbox');
      expect(inputElement).toBeInTheDocument();
    });

    it('zeigt einen Such-Button', () => {
      render(<ItemSearch />);
      
      const buttonElement = screen.getByRole('button');
      expect(buttonElement).toBeInTheDocument();
      expect(buttonElement).toHaveTextContent(/search/i);
    });

    it('hat einen passenden Placeholder-Text', () => {
      render(<ItemSearch />);
      
      const inputElement = screen.getByPlaceholderText(/search eve items/i);
      expect(inputElement).toBeInTheDocument();
    });
  });

  describe('Initiale UI-Zustände', () => {
    it('beginnt mit leerem Eingabefeld', () => {
      render(<ItemSearch />);
      
      const inputElement = screen.getByRole('textbox');
      expect(inputElement).toHaveValue('');
    });
    
    it('der Button ist initial deaktiviert', () => {
      render(<ItemSearch />);
      
      const buttonElement = screen.getByRole('button');
      expect(buttonElement).toBeDisabled();
    });

    it('zeigt keine Fehlermeldung initial', () => {
      render(<ItemSearch />);
      
      // Keine Fehlermeldung sichtbar
      const errorElements = screen.queryAllByText(/error|fehler/i);
      expect(errorElements).toHaveLength(0);
    });

    it('zeigt keine Suchergebnisse initial', () => {
      render(<ItemSearch />);
      
      // Keine Suchergebnisse sichtbar
      const resultElements = screen.queryAllByText(/search results|suchergebnisse/i);
      expect(resultElements).toHaveLength(0);  
    });
  });

  describe('✅ Diese Tests zeigen erfolgreiche TDD-Grundlagen', () => {
    it('beweist, dass TDD-Setup funktioniert', () => {
      // Dieser Test demonstriert:
      // 1. Tests laufen
      // 2. Komponente wird gerendert  
      // 3. Basic UI ist vorhanden
      
      render(<ItemSearch />);
      
      // Alle wichtigen UI-Elemente sind da
      expect(screen.getByRole('textbox')).toBeInTheDocument();
      expect(screen.getByRole('button')).toBeInTheDocument();
      expect(screen.getByPlaceholderText(/search eve items/i)).toBeInTheDocument();
      
      console.log('🎉 TDD-Setup erfolgreich! Wir können jetzt iterativ entwickeln.');
    });
  });
});
