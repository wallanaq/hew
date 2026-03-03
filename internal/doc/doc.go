package doc

import (
	"regexp"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
)

var (
	indentRegex = regexp.MustCompile(`(?m)^`)
)

type normalizer struct {
	string
}

// heredoc normalizes the string using the heredoc library, removing common leading whitespace.
func (n normalizer) heredoc() normalizer {
	n.string = heredoc.Doc(n.string)
	return n
}

// indent adds two spaces of indentation to the beginning of each line in the string.
func (n normalizer) indent() normalizer {
	if len(strings.TrimSpace(n.string)) == 0 {
		return n
	}

	n.string = indentRegex.ReplaceAllLiteralString(n.string, "  ")
	return n
}

func (n normalizer) String() string {
	return n.string
}

// Description formats a string as a heredoc, which is suitable for
// the Long description of a cobra command. It un-indents the input string.
func Description(raw string) string {
	return normalizer{raw}.heredoc().string
}

// Example formats a string as a heredoc and then indents it with two spaces.
// This is suitable for the Example field of a cobra command, which is conventionally indented.
func Example(raw string) string {
	return normalizer{raw}.heredoc().indent().string
}
