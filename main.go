package main


import "flag"
import "os"


func main() {

	// Declare the command line flags.
	//
	// NOTE that we need to run "flag.Parse()" before we will actually get the data from the command line.
	//
		count_ptr     := flag.Uint64("count",     1,    "The number of times to repeat the full pattern. I.e., the number of full iterations. (default: 1)")
		plus_ptr      := flag.Int64("plus",       0,    "How much of a partial iteration to do. Negative values permitted. (default: 0)")
		separator_ptr := flag.String("separator", "\n", "Used to \"glue\" pattern items. (default: \"\\n\")")
		format_ptr    := flag.String("format",    "",   "Used to format the output. (default: \"\")")

	// Parse the command line.
	//
	// This results in getting the command line flags we previously declared.
	// And also lets us get the (remaining) command line arguments (after the flags have been
	// removed) with "flag.Args()".
	//
		flag.Parse()

	// Decide what needs to get done.
	//
	// We decide that the output is going to STDOUT.
	//
	// We decide that the full pattern will repeat "count" times.
	//
	// We decide that a partial pattern will inclue "plus" mod (pattern length) items.
	//
	// We decide that the pattern comes from the command line arguments (after the command
	// line flags have been removed).
	//
	// We decide that the separator to "glue" the items of the pattern is "separator".
	//
		w          := os.Stdout
		count      := *count_ptr
		plus       := *plus_ptr
		pattern    := flag.Args()
		separator  := *separator_ptr

		var formatter Formatter
		if format := *format_ptr ; "" == format {
			formatter = NewIdentityFormatter()
		} else {
			formatter = NewCachedFormatter(NewPrintfStyleFormatter(format), len(pattern))
		}

	// Actually do what needs to get done.
	//               
	// I.e., repeat.
	//               
		repeat(w, count, plus, pattern, separator, formatter)
}
