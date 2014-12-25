package main


import "fmt"
import "strings"


// PrintfStyleFormatter purpose is to provide printf style formatting.
//
// It implements the "Formatter" interface with its implementation of
// "Format(string) string".
//
// Note that it does not do any caching itself, but the user of this
// can put it in a "CachedFormatter" to have caching.
//
type PrintfStyleFormatter struct {
	format string
}


// sanitizeFormat takes the format (usually specified by the user) and sanitizes it,
// so that it can be used with things like "fmt.Sprintf".
func sanitizeFormat(format string) string {

	processed_format := strings.Replace(strings.Replace(format, "%", "%%", -1), "%%s", "%s", -1)
	
	return processed_format
}


// NewPrintfStyleFormatter creates a new PrintfStyleFormatter.
func NewPrintfStyleFormatter(format string) (*PrintfStyleFormatter) {

	me := PrintfStyleFormatter{
		format: sanitizeFormat(format),
	}

	return &me
}


// Format implements the the "Format(string) string" method from the "Formatter" interface.
func (me *PrintfStyleFormatter) Format(s string) string {
	formatted := fmt.Sprintf(me.format, s)

	return formatted
}
