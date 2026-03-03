package doc

import "testing"

func TestDescription(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name: "simple string",
			input: `
				This is a simple description.
				It has multiple lines.
			`,
			want: "This is a simple description.\nIt has multiple lines.\n",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Description(tt.input); got != tt.want {
				t.Errorf("Description() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestExample(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name: "simple example",
			input: `
				# A comment
				command --flag`,
			want: "  # A comment\n  command --flag",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "single line",
			input: `command`,
			want:  "  command",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Example(tt.input); got != tt.want {
				t.Errorf("Example() = %q, want %q", got, tt.want)
			}
		})
	}
}
