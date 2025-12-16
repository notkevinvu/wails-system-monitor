/**
 * Utility functions for the frontend application.
 * These are separated for easier testing.
 */

/**
 * Helper function to create and configure DOM elements with type safety.
 */
export function createElement<K extends keyof HTMLElementTagNameMap>(
  tag: K,
  attrs: Partial<HTMLElementTagNameMap[K]> & { parent?: HTMLElement }
): HTMLElementTagNameMap[K] {
  const el = document.createElement(tag);
  const { parent, ...rest } = attrs;
  Object.assign(el, rest);
  parent?.appendChild(el);
  return el;
}

/**
 * Helper function for consistent error handling and user feedback.
 */
export function handleError(err: unknown, target: HTMLElement, action: string): void {
  console.error(err);
  target.textContent = `Error: Could not ${action}. Please try again.`;
}

/**
 * Validates that a name input is non-empty after trimming.
 */
export function isValidName(name: string): boolean {
  return name.trim().length > 0;
}

/**
 * Formats a greeting message.
 */
export function formatGreeting(name: string): string {
  return `Hello ${name}!`;
}
