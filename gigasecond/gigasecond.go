// Build constraint.  It's like when you buy something new and you find a
// tag that says "Inspected by #12."  Nice to know it was inspected, but
// you don't need the tag.  Take it off.
// +build !example

// Package clause.
package gigasecond

// Constant declaration.
const TestVersion = ? // find the value in gigasecond_test.go

// API function.  It uses a type from the Go standard library.
func AddGigasecond(time.Time) time.Time

// API variable.  It needs a type at least.  A value would be nice too.
// See gigasecond_test.go.
var Birthday

// Nitpickers don't think much of stub comments.  Replace or remove.
