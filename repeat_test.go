package main


import "testing"
import "bytes"


func TestRepeat(t *testing.T) {

	tests := []struct{
		count       uint64
		plus         int64
		pattern   []string
		separator   string
		formatter   Formatter

		expected    string
	}{
		{
			count:     uint64(2),
			plus:      int64(-10),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-9),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\nA\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-8),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\nA\nb\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-7),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\nA\nb\nc\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-6),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\nA\nb\nc\nd\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-5),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-4),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\nA\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-3),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\nA\nb\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-2),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\nA\nb\nc\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-1),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\nA\nb\nc\nd\n",
		},
		{
			count:     uint64(2),
			plus:      int64(0),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\n",
		},
		{
			count:     uint64(2),
			plus:      int64(1),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\nA\n",
		},
		{
			count:     uint64(2),
			plus:      int64(2),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\nA\nb\n",
		},
		{
			count:     uint64(2),
			plus:      int64(3),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\nA\nb\nc\n",
		},
		{
			count:     uint64(2),
			plus:      int64(4),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\nA\nb\nc\nd\n",
		},
		{
			count:     uint64(2),
			plus:      int64(5),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\n",
		},
		{
			count:     uint64(2),
			plus:      int64(6),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\nA\n",
		},
		{
			count:     uint64(2),
			plus:      int64(7),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\nA\nb\n",
		},
		{
			count:     uint64(2),
			plus:      int64(8),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\nA\nb\nc\n",
		},
		{
			count:     uint64(2),
			plus:      int64(9),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\nA\nb\nc\nd\n",
		},
		{
			count:     uint64(2),
			plus:      int64(10),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "A\nb\nc\nd\ne\nA\nb\nc\nd\ne\n",
		},


		{
			count:     uint64(2),
			plus:      int64(1),
			pattern:   []string{"1", "0", "0", "0"},
			separator: "\n",
			formatter: NewIdentityFormatter(),

			expected:  "1\n0\n0\n0\n1\n0\n0\n0\n1\n",
		},


		{
			count:     uint64(2),
			plus:      int64(-10),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-9),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\tA\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-8),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\tA\tb\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-7),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\tA\tb\tc\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-6),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\tA\tb\tc\td\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-5),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-4),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\tA\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-3),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\tA\tb\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-2),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\tA\tb\tc\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-1),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\tA\tb\tc\td\n",
		},
		{
			count:     uint64(2),
			plus:      int64(0),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\n",
		},
		{
			count:     uint64(2),
			plus:      int64(1),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\tA\n",
		},
		{
			count:     uint64(2),
			plus:      int64(2),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\tA\tb\n",
		},
		{
			count:     uint64(2),
			plus:      int64(3),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\tA\tb\tc\n",
		},
		{
			count:     uint64(2),
			plus:      int64(4),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\tA\tb\tc\td\n",
		},
		{
			count:     uint64(2),
			plus:      int64(5),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\n",
		},
		{
			count:     uint64(2),
			plus:      int64(6),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\tA\n",
		},
		{
			count:     uint64(2),
			plus:      int64(7),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\tA\tb\n",
		},
		{
			count:     uint64(2),
			plus:      int64(8),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\tA\tb\tc\n",
		},
		{
			count:     uint64(2),
			plus:      int64(9),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\tA\tb\tc\td\n",
		},
		{
			count:     uint64(2),
			plus:      int64(10),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: "\t",
			formatter: NewIdentityFormatter(),

			expected:  "A\tb\tc\td\te\tA\tb\tc\td\te\n",
		},


		{
			count:     uint64(2),
			plus:      int64(1),
			pattern:   []string{"apple", "banana", "cherry"},
			separator: ",",
			formatter: NewIdentityFormatter(),

			expected:  "apple,banana,cherry,apple,banana,cherry,apple\n",
		},


		{
			count:     uint64(2),
			plus:      int64(-10),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-9),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e join A\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-8),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e join A join b\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-7),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e join A join b join c\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-6),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e join A join b join c join d\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-5),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-4),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e join A\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-3),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e join A join b\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-2),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e join A join b join c\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-1),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e join A join b join c join d\n",
		},
		{
			count:     uint64(2),
			plus:      int64(0),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e\n",
		},
		{
			count:     uint64(2),
			plus:      int64(1),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e join A\n",
		},
		{
			count:     uint64(2),
			plus:      int64(2),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e join A join b\n",
		},
		{
			count:     uint64(2),
			plus:      int64(3),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e join A join b join c\n",
		},
		{
			count:     uint64(2),
			plus:      int64(4),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e join A join b join c join d\n",
		},
		{
			count:     uint64(2),
			plus:      int64(5),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e\n",
		},
		{
			count:     uint64(2),
			plus:      int64(6),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e join A\n",
		},
		{
			count:     uint64(2),
			plus:      int64(7),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e join A join b\n",
		},
		{
			count:     uint64(2),
			plus:      int64(8),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e join A join b join c\n",
		},
		{
			count:     uint64(2),
			plus:      int64(9),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e join A join b join c join d\n",
		},
		{
			count:     uint64(2),
			plus:      int64(10),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " join ",
			formatter: NewIdentityFormatter(),

			expected:  "A join b join c join d join e join A join b join c join d join e\n",
		},


		{
			count:     uint64(1),
			plus:      int64(0),
			pattern:   []string{"yek", "do", "se", "chahâr", "panj", "šeš", "haft", "hašt", "noh", "dah"},
			separator: ", ",
			formatter: NewIdentityFormatter(),

			expected:  "yek, do, se, chahâr, panj, šeš, haft, hašt, noh, dah\n",
		},


		{
			count:     uint64(2),
			plus:      int64(-10),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " ++ ",
			formatter: NewPrintfStyleFormatter("f(x = %s)"),

			expected:  "f(x = A) ++ f(x = b) ++ f(x = c) ++ f(x = d) ++ f(x = e) ++ f(x = A) ++ f(x = b) ++ f(x = c) ++ f(x = d) ++ f(x = e)\n",
		},
		{
			count:     uint64(2),
			plus:      int64(-9),
			pattern:   []string{"A", "b", "c", "d", "e"},
			separator: " ++ ",
			formatter: NewPrintfStyleFormatter("f(x = %s)"),

			expected:  "f(x = A) ++ f(x = b) ++ f(x = c) ++ f(x = d) ++ f(x = e) ++ f(x = A) ++ f(x = b) ++ f(x = c) ++ f(x = d) ++ f(x = e) ++ f(x = A)\n",
		},
	}


	// Run each test.
	//
		for test_number, test := range tests {

			// Create the writer that we can use to capture the output.
				var w bytes.Buffer

			// Invoke repeat().
				repeat(&w, test.count, test.plus, test.pattern, test.separator, test.formatter)

			// Check that repeat() outputted what we expected.
				if actual := w.String() ; test.expected != actual {
					t.Errorf("Did not get what was expected for test #%d.\nExpected: %q\nReceived: %q\n", test_number, test.expected, actual)
				}
		}
}
