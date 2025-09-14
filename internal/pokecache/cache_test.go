package pokecache

import (
	"testing"
	"time"
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		expected []byte
	}{
		{
			name:     "data1",
			key:      "http://example.com",
			expected: []byte("Some sample data"),
		},
		{
			name:     "data2",
			key:      "https://api.example.com",
			expected: []byte("Some kind of api data"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := NewCache(5 * time.Second)
			cache.Add(tt.key, tt.expected)
			actual, ok := cache.Get(tt.key)

			if !ok {
				t.Error("Expected to find key in cache")
				return
			}

			if string(actual) != string(tt.expected) {
				t.Error("Expected to find same value in cache")
				return
			}
		})
	}
}
