import { describe, it, expect } from 'vitest';
import { isValidName, formatGreeting } from './utils';

describe('isValidName', () => {
  it('returns true for non-empty strings', () => {
    expect(isValidName('John')).toBe(true);
    expect(isValidName('Jane Doe')).toBe(true);
    expect(isValidName('a')).toBe(true);
  });

  it('returns false for empty strings', () => {
    expect(isValidName('')).toBe(false);
  });

  it('returns false for whitespace-only strings', () => {
    expect(isValidName('   ')).toBe(false);
    expect(isValidName('\t')).toBe(false);
    expect(isValidName('\n')).toBe(false);
  });

  it('trims whitespace before validation', () => {
    expect(isValidName('  John  ')).toBe(true);
  });
});

describe('formatGreeting', () => {
  it('formats greeting with provided name', () => {
    expect(formatGreeting('World')).toBe('Hello World!');
    expect(formatGreeting('User')).toBe('Hello User!');
  });

  it('handles empty string', () => {
    expect(formatGreeting('')).toBe('Hello !');
  });

  it('preserves special characters', () => {
    expect(formatGreeting('User123')).toBe('Hello User123!');
    expect(formatGreeting("O'Brien")).toBe("Hello O'Brien!");
  });
});

// NOTE: DOM-dependent functions (createElement, handleError) require jsdom environment
// To test those, update vitest.config.ts to use environment: 'jsdom' and add @testing-library/dom
