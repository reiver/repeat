package main


// Formatter is an interface that defines a single method: "Format(string) string".
//
// The purposes of this interface is to create an abstraction for the "--format"
// command line flag. (The "--format" command line flag is delared in the file
// "main.go"; which includes help text.)
//
// Also related to this is "IdentityFormatter" defined in "identity_formatter.go"
// and "PrintfStyleFormatter" defined in "printf_style_formatter.go".
type Formatter interface {
	Format(string) string
}
