package main

import (
	"testing"
)

func TestGreet(t *testing.T) {
	app := NewApp()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "greets World",
			input:    "World",
			expected: "Hello World, It's show time!",
		},
		{
			name:     "greets empty string",
			input:    "",
			expected: "Hello , It's show time!",
		},
		{
			name:     "greets with special characters",
			input:    "User123",
			expected: "Hello User123, It's show time!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := app.Greet(tt.input)
			if result != tt.expected {
				t.Errorf("Greet(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGetItems(t *testing.T) {
	app := NewApp()
	items := app.GetItems()

	// Verify we get the expected number of items
	expectedCount := 5
	if len(items) != expectedCount {
		t.Errorf("GetItems() returned %d items, want %d", len(items), expectedCount)
	}

	// Verify specific items exist
	expectedItems := []string{"Apple", "Banana", "Cherry", "Date", "Elderberry"}
	for i, expected := range expectedItems {
		if items[i] != expected {
			t.Errorf("GetItems()[%d] = %q, want %q", i, items[i], expected)
		}
	}
}

func TestNewApp(t *testing.T) {
	app := NewApp()
	if app == nil {
		t.Error("NewApp() returned nil")
	}
}
