package main


// IdentityFormatter is something that can be used as a "Formatter", because
// it implements the "Format(string) string" method, but conceptually does nothing
// to it; just like the concept of an "identity" in mathematics.
//
// In other words "IdentityFormatter.Format(s)" always just return the value of "s".
//
// This is useful for the case when you don't want to add any formatting, and just use
// it as is.
//
// You may want to do this if the user does not specify the "--format" flag or the user
// does specify the "--format" flag but makes it an empty string, as in '--format=""'.
type IdentityFormatter struct {}


// NewIdentityFormatter creates a new IdentityFormatter.
func NewIdentityFormatter() (*IdentityFormatter) {
	me := IdentityFormatter{}

	return &me
}


// Format implements the the "Format(string) string" method from the "Formatter" interface.
func (me *IdentityFormatter) Format(s string) string {
	return s
}
