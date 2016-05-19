package hub

// NameChecker validates the users' nicknames.
// Should return transformed nickname and error if
// client should be refused to join.
type NameChecker func(string) (string, error)

// Sanitizer transforms the message's text.
type Sanitizer func(string) string
