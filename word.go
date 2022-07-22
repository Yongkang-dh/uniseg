package uniseg

import "unicode/utf8"

// firstWord returns the first word (as a slice of bytes) found in the given
// byte slice, according to the rules of Unicode Standard Annex #29 (Word
// Boundaries). This function can be called continuously to extract all words
// from a byte slice, as follows:
//
//   state := -1
//   for len(b) > 0 {
//       word, b, state = firstWord(b, state)
//   }
//
// If you don't know the current state, for example when calling the function
// for the first time, you must pass -1. For consecutive calls, pass the state
// returned by the previous call.
//
// The "rest" slice is the subslice of the original byte slice "b" starting
// after the last byte of the identified word. If the length of the "rest" slice
// is 0, the entire byte slice "b" has been processed.
//
// Given an empty byte slice "b", the function returns nil values.
//
// For the time being, this function is private because its signature might
// still change.
func firstWord(b []byte, state int) (cluster, rest []byte, newState int) {
	// An empty byte slice returns nothing.
	if len(b) == 0 {
		return
	}

	// Extract the first rune.
	r, length := utf8.DecodeRune(b)
	if len(b) <= length { // If we're already past the end, there is nothing else to parse.
		return b, nil, grAny
	}

	// If we don't know the state, determine it now.
	if state < 0 {
		state, _ = transitionWordBreakState(state, r, b[length:], "")
	}

	// Transition until we find a boundary.
	var boundary bool
	for {
		r, l := utf8.DecodeRune(b[length:])
		state, boundary = transitionWordBreakState(state, r, b[length+l:], "")

		if boundary {
			return b[:length], b[length:], state
		}

		length += l
		if len(b) <= length {
			return b, nil, grAny
		}
	}
}

// firstWordInString is like firstWord() but its input and outputs are a string.
func firstWordInString(str string, state int) (cluster, rest string, newState int) {
	// An empty byte slice returns nothing.
	if len(str) == 0 {
		return
	}

	// Extract the first rune.
	r, length := utf8.DecodeRuneInString(str)
	if len(str) <= length { // If we're already past the end, there is nothing else to parse.
		return str, "", grAny
	}

	// If we don't know the state, determine it now.
	if state < 0 {
		state, _ = transitionWordBreakState(state, r, nil, str[length:])
	}

	// Transition until we find a boundary.
	var boundary bool
	for {
		r, l := utf8.DecodeRuneInString(str[length:])
		state, boundary = transitionWordBreakState(state, r, nil, str[length+l:])

		if boundary {
			return str[:length], str[length:], state
		}

		length += l
		if len(str) <= length {
			return str, "", grAny
		}
	}
}