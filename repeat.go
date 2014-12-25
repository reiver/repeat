package main


import "io"


// repeat is where the actual work starts to be done in this program.
//
// Note that *what* is to be done is not decided in this func.
// *What* is to be done is communicated with this func via its parameters.
//
// The "w" parameter is an "io.Writer" that determines where the output from this func will go -- will be written.
// For example, this could be STDOUT (i.e., os.Stdout).
//
// The "count" parameter is a non-negative integer in the "uint64" range that specifies how many times to repeat
// full interations of the pattern, specified in the "pattern" parameter.
//
// The "plus" parameter is a possibly negtive integer in the "int64" range that species how much of a
// partial interation of the pattern to do.
//
// The "separator" parameter specifies the "glue" to put between items from the pattern when they are being
// outputted -- when they ae being wrttien -- using the "w" parameter. For example, if "separator" was "\n"
// then each item would be put on its own line.
//
// The "formatter" parameter is used to format the item before it is outputted -- before it is written -- using
// the "w" parameter.
//
// NOTE that at the end, this will output -- this will write -- an (extra) "\n" using the "w" parameter.
// NOTE that this "\n" has nothing to do with what is specified with the "separator" parameter.
//
func repeat(w io.Writer, count uint64, plus int64, pattern []string, separator string, formatter Formatter) {

	// Calculate the length of the pattern.
		pattern_length := uint64(len(pattern))

	// If the length of the patter is less than one (1), then just return without doing anything.
		if 1 > pattern_length {
			return
		}

	// Figure out how many times the loop should iterate.
		limit := limit(count, plus, pattern_length)

	// Repeat.
		first_iteration := true

		for i:=uint64(0); i<limit; i++ {

			// Output the separator.
				if ! first_iteration {
					_,_ = io.WriteString(w, separator)
				} else {
					first_iteration = false
				}

			// Calculate the index into the pattern.
				index_into_the_pattern := i % pattern_length

			// Get the item from the pattern that we are at in this iteration of the
			// loop (which corresponds to a line if the "separator" parameters equal
			// '\n').
				s := pattern[index_into_the_pattern]

			// Format the item.
				formatted := formatter.Format(s)

			// Output the formatted item.
				_,_ = io.WriteString(w, formatted)
		}

	// Regardless of what the separator is, at the end we output a '\n'.
	//
	// We consider this necessary for a text file to be well-formed.
	//
		_,_ = io.WriteString(w, "\n")
}
