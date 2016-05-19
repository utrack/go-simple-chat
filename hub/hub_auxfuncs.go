package hub

import (
	"errors"
	"github.com/kennygrant/sanitize"
	"strings"
)

// NameChecker validates the users' nicknames.
// Should return transformed nickname and error if
// client should be refused to join.
type NameChecker func(string) (string, error)

// Sanitizer transforms the message's text.
type Sanitizer func(string) string

// DefaultNameChecker is a sample NameChecker which strips out HTML and
// leading/trailing spaces.
func DefaultNameChecker(name string) (string, error) {
	name = sanitize.HTML(name)
	name = strings.Trim(name, "\t\n ")

	if len(name) > 20 {
		return ``, errors.New("Nickname is too long")
	}
	if len(name) < 2 {
		return ``, errors.New("Nickname is too short")
	}
	return name, nil
}
