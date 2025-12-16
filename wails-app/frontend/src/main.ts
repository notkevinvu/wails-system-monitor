import './style.css';
import './app.css';

import logo from './assets/images/logo-universal.png';
import { Greet, GetItems } from '../wailsjs/go/main/App';

// Helper function to create and configure DOM elements
function createElement<K extends keyof HTMLElementTagNameMap>(
  tag: K,
  attrs: Partial<HTMLElementTagNameMap[K]> & { parent?: HTMLElement }
): HTMLElementTagNameMap[K] {
  const el = document.createElement(tag);
  const { parent, ...rest } = attrs;
  Object.assign(el, rest);
  parent?.appendChild(el);
  return el;
}

// Helper function for consistent error handling
function handleError(err: unknown, target: HTMLElement, action: string): void {
  console.error(err);
  target.textContent = `Error: Could not ${action}. Please try again.`;
}

// Initialize app container with null check
const app = document.querySelector<HTMLDivElement>('#app');
if (!app) {
  throw new Error('App container element not found');
}

// Build UI using helper function
const logoImg = createElement('img', { id: 'logo', className: 'logo', src: logo, parent: app });

const resultDiv = createElement('div', {
  className: 'result',
  id: 'result',
  textContent: 'Please enter your name below',
  parent: app
});

const inputBox = createElement('div', { className: 'input-box', id: 'input', parent: app });

const nameInput = createElement('input', {
  className: 'input',
  id: 'name',
  type: 'text',
  autocomplete: 'off',
  parent: inputBox
});

const greetBtn = createElement('button', {
  className: 'btn',
  textContent: 'Greet',
  parent: inputBox
});
greetBtn.addEventListener('click', greet);

// Items section
const itemsSection = createElement('div', { className: 'items-section', parent: app });

const itemsBtn = createElement('button', {
  className: 'btn',
  textContent: 'Get Items from Go',
  parent: itemsSection
});
itemsBtn.addEventListener('click', fetchItems);

const itemsList = createElement('div', { className: 'items-list', parent: itemsSection });

nameInput.focus();

// Greet function - calls Go backend
async function greet(): Promise<void> {
  const name = nameInput.value.trim();
  if (name === '') return;

  try {
    const result = await Greet(name);
    resultDiv.textContent = result;
  } catch (err) {
    handleError(err, resultDiv, 'greet');
  }
}

// FetchItems function - gets list from Go backend
async function fetchItems(): Promise<void> {
  try {
    const items = await GetItems();
    itemsList.textContent = '';

    const ul = document.createElement('ul');
    items.forEach((item) => {
      const li = document.createElement('li');
      li.textContent = item;
      ul.appendChild(li);
    });
    itemsList.appendChild(ul);
  } catch (err) {
    handleError(err, itemsList, 'fetch items');
  }
}
