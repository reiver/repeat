package main


// limit indirectly calculates the number of (logical) lines to be outputted. Its result is meant
// to be used at the limit to a for-loop.
//
// Consider the example where:
//
//	count   := 3
//	pattern := []string{"A", "b", "c", "d", "e"}
//
// If we only had to consider the "count" and "pattern" parameters, then we would get the following
// written onto "w":
//
//	A
//	b
//	c
//	d
//	e
//	A
//	b
//	c
//	d
//	e
//	A
//	b
//	c
//	d
//	e
//
// To make it clearer:
//
//	A	← Line #1  \
//	b	← Line #2  |
//	c	← Line #3  |- Iteration of the pattern #1
//	d	← Line #4  |
//	e	← Line #5  /
//	A	← Line #6  \
//	b	← Line #7  |
//	c	← Line #8  |- Iteration of the pattern #2
//	d	← Line #9  |
//	e	← Line #10 /
//	A	← Line #11 \
//	b	← Line #12 |
//	c	← Line #13 |- Iteration of the pattern #3
//	d	← Line #14 |
//	e	← Line #15 /
//
// Or to make it even clearer than that, by adding some spacing:
//
//	A	← Line #1  \
//	b	← Line #2  |
//	c	← Line #3  |- Iteration of the pattern #1
//	d	← Line #4  |
//	e	← Line #5  /
//
//	A	← Line #6  \
//	b	← Line #7  |
//	c	← Line #8  |- Iteration of the pattern #2
//	d	← Line #9  |
//	e	← Line #10 /
//
//	A	← Line #11 \
//	b	← Line #12 |
//	c	← Line #13 |- Iteration of the pattern #3
//	d	← Line #14 |
//	e	← Line #15 /
//
// If we only had the "count" and "pattern" parameter, then all we would need to do to
// calculate how many times the loop should iterate is:
//
//	limit := count * len(pattern)
//
// However, this is NOT sufficient for that calculation. It is more complex that this.
//
// It is more complex than this, because we also have to consider the "plus" parameter.
//
// The "plus" lets us add a "partial iteration" of the pattern.
//
// This can be seen in this following example.
//
// Consider the example where:
//
//	count   := 3
//	pattern := []string{"A", "b", "c"}
//	plus    := 2
//
// Then we would get the following written onto "w":
//
//	A
//	b
//	c
//	d
//	e
//	A
//	b
//	c
//	d
//	e
//	A
//	b
//	c
//	d
//	e
//	A
//	b
//
// To make it clearer:
//
//	A	← Line #1  \
//	b	← Line #2  |
//	c	← Line #3  |- Iteration of the pattern #1
//	d	← Line #4  |
//	e	← Line #5  /
//	A	← Line #6  \
//	b	← Line #7  |
//	c	← Line #8  |- Iteration of the pattern #2
//	d	← Line #9  |
//	e	← Line #10 /
//	A	← Line #11 \
//	b	← Line #12 |
//	c	← Line #13 |- Iteration of the pattern #3
//	d	← Line #14 |
//	e	← Line #15 /
//	A	← Line #16 \_ 
//	b	← Line #17 /  This is the partial iteration of the pattern
//
// Or to make it even clearer than that, by adding some spacing:
//
//	A	← Line #1  \
//	b	← Line #2  |
//	c	← Line #3  |- Iteration of the pattern #1
//	d	← Line #4  |
//	e	← Line #5  /
//
//	A	← Line #6  \
//	b	← Line #7  |
//	c	← Line #8  |- Iteration of the pattern #2
//	d	← Line #9  |
//	e	← Line #10 /
//
//	A	← Line #11 \
//	b	← Line #12 |
//	c	← Line #13 |- Iteration of the pattern #3
//	d	← Line #14 |
//	e	← Line #15 /
//
//	A	← Line #16 \_ 
//	b	← Line #17 /  This is the partial iteration of the pattern
//
// Hopefully you can see that the effect of "plus" being "2" is that the partial interation
// of the pattern consists of just the first 2 items of the pattern.
//
// Note that "plus" can also be negative!
//
// But what does a negative "plus" parameter mean?!
//
// What a negative "plus" parameter means is, you are counting from the end of the pattern.
//
// This can be seen in the following illustration:
//
//	 A	 b	 c	 d	 e
//	 1	 2	 3	 4	 5
//	-4	-3	-2	-1	
//
// Of course, if we had a pattern of a different length, things would line up
// differently. As in the following illustration:
//
//	 A	 b	 c	 d	 e	 f	 g
//	 1	 2	 3	 4	 5	 6	 7
//	-6	-5	-4	-3	-2	-1
//
// (Some may recognie this as being somewhat similar to modular arithmetic.
// Although it is important to not conflate it with modular arithmetic.
// In modular arithmetic 5 ≡ 0 (mod 5). However, with the "plus" parameter
// 5 and 0 are treated differently.)
//
// The effect of this can be seen in the following example:
//
//	count   := 3
//	pattern := []string{"A", "b", "c"}
//	plus    := -2
//
// Then we would get the following written onto "w":
//
//	A
//	b
//	c
//	d
//	e
//	A
//	b
//	c
//	d
//	e
//	A
//	b
//	c
//	d
//	e
//	A
//	b
//	c
//
// To make it clearer:
//
//	A	← Line #1  \
//	b	← Line #2  |
//	c	← Line #3  |- Iteration of the pattern #1
//	d	← Line #4  |
//	e	← Line #5  /
//	A	← Line #6  \
//	b	← Line #7  |
//	c	← Line #8  |- Iteration of the pattern #2
//	d	← Line #9  |
//	e	← Line #10 /
//	A	← Line #11 \
//	b	← Line #12 |
//	c	← Line #13 |- Iteration of the pattern #3
//	d	← Line #14 |
//	e	← Line #15 /
//	A	← Line #16 \ 
//	b	← Line #17 |- This is the partial iteration of the pattern
//	c	← Line #18 /
//
// Or to make it even clearer than that, by adding some spacing:
//
//	A	← Line #1  \
//	b	← Line #2  |
//	c	← Line #3  |- Iteration of the pattern #1
//	d	← Line #4  |
//	e	← Line #5  /
//
//	A	← Line #6  \
//	b	← Line #7  |
//	c	← Line #8  |- Iteration of the pattern #2
//	d	← Line #9  |
//	e	← Line #10 /
//
//	A	← Line #11 \
//	b	← Line #12 |
//	c	← Line #13 |- Iteration of the pattern #3
//	d	← Line #14 |
//	e	← Line #15 /
//
//	A	← Line #16 \ 
//	b	← Line #17 |- This is the partial iteration of the pattern
//	c	← Line #18 /
//
// As expected, a "plus" of -2 gave us 3 extra items for the partial iteration
// of the pattern when the length of the pattern equals 5.
//
func limit(count uint64, plus int64, pattern_length uint64) uint64 {

	// Make it so that the "plus" parameter does not go beyond one
	// the full length of the pattern. And also, so that it doesn't
	// do a the full length of the pattern either.
	//
	// NOTE that the way Golang does modular arithmetic, using the "%"
	// operator, negative values stay negative. (Which is good because
	// that is how we want it.
	//
		plus = plus % int64(pattern_length)

	// Calculate the limit.
		limit := count * pattern_length
		if 0 < plus {
			limit += uint64(plus)
		} else if 0 > plus {
			limit += pattern_length - uint64(-1*plus)
		}


	// Return.
		return limit
}
