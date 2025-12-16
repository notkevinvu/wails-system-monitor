import './style.css';
import './app.css';

import logo from './assets/images/logo-universal.png';
import {Greet, GetItems} from '../wailsjs/go/main/App';

// Build the UI using safe DOM methods
const app = document.querySelector('#app');

const logoImg = document.createElement('img');
logoImg.id = 'logo';
logoImg.className = 'logo';
logoImg.src = logo;
app.appendChild(logoImg);

const resultDiv = document.createElement('div');
resultDiv.className = 'result';
resultDiv.id = 'result';
resultDiv.textContent = 'Please enter your name below';
app.appendChild(resultDiv);

const inputBox = document.createElement('div');
inputBox.className = 'input-box';
inputBox.id = 'input';

const nameInput = document.createElement('input');
nameInput.className = 'input';
nameInput.id = 'name';
nameInput.type = 'text';
nameInput.autocomplete = 'off';
inputBox.appendChild(nameInput);

const greetBtn = document.createElement('button');
greetBtn.className = 'btn';
greetBtn.textContent = 'Greet';
greetBtn.addEventListener('click', greet);
inputBox.appendChild(greetBtn);

app.appendChild(inputBox);

// Items section
const itemsSection = document.createElement('div');
itemsSection.style.marginTop = '20px';

const itemsBtn = document.createElement('button');
itemsBtn.className = 'btn';
itemsBtn.textContent = 'Get Items from Go';
itemsBtn.addEventListener('click', fetchItems);
itemsSection.appendChild(itemsBtn);

const itemsList = document.createElement('div');
itemsList.id = 'items-list';
itemsList.style.marginTop = '10px';
itemsSection.appendChild(itemsList);

app.appendChild(itemsSection);

nameInput.focus();

// Greet function - calls Go backend
function greet() {
    const name = nameInput.value.trim();
    if (name === "") return;

    Greet(name)
        .then((result) => {
            resultDiv.textContent = result;
        })
        .catch((err) => {
            console.error(err);
            resultDiv.textContent = 'Error: Could not greet. Please try again.';
        });
}

// FetchItems function - gets list from Go backend
function fetchItems() {
    GetItems()
        .then((items) => {
            // Clear existing items
            itemsList.textContent = '';

            // Build list using safe DOM methods
            const ul = document.createElement('ul');
            items.forEach(item => {
                const li = document.createElement('li');
                li.textContent = item;
                ul.appendChild(li);
            });
            itemsList.appendChild(ul);
        })
        .catch((err) => {
            console.error(err);
            itemsList.textContent = 'Error: Could not fetch items. Please try again.';
        });
}
