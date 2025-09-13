package repl

import "testing"

func TestCleanInput(t *testing.T) {
    tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "multiple-spaces",
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "multiple-words",
			input:    "Bulbasaur Charmander Squirtle PIKACHU",
			expected: []string{"Bulbasaur", "Charmander", "Squirtle", "PIKACHU"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := cleanInput(tt.input)
			if len(actual) != len(tt.expected) {
				t.Errorf("Expected length %d, Actual length %d", len(tt.expected), len(actual))
			} else {
				for i := range actual {
					actualWord := actual[i]
					expectedWord := tt.expected[i]
					
					if expectedWord != actualWord {
						t.Errorf("At index %d, expected '%s', actual '%s'", i, expectedWord, actualWord)
					}
				}
			}
		})
	}
}