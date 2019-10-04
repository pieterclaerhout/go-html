package html

// ElementOption defines the options for the elements
type ElementOption int8

const (
	// Void is the default (no options)
	Void ElementOption = 1 << iota

	// SelfClose defines the element as a self-closing element
	SelfClose

	// CSSElement defines the element as CSS element
	CSSElement

	// JSElement defines the element as a JS element
	JSElement

	// NoWhitespace defines that the element shouldn't add whitespace
	NoWhitespace
)
