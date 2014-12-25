package main


import "testing"


func TestSanitizeFormat(t *testing.T) {

	tests := []struct{
		format  string
		expected string
	}{
		{
			format:   "apple",
			expected: "apple",
		},
		{
			format:   "banana",
			expected: "banana",
		},
		{
			format:   "cherry",
			expected: "cherry",
		},


		{
			format:   "",
			expected: "",
		},


		{
			format:   "%",
			expected: "%%",
		},


		{
			format:   "%s",
			expected: "%s",
		},

		{
			format:   "%d %s %f",
			expected: "%%d %s %%f",
		},


		{
			format:   "The string is: %s. And it is 15% of it.",
			expected: "The string is: %s. And it is 15%% of it.",
		},
	}


	// Run each test.
	//
		for test_number, test := range tests {

			// Invoke sanitizeFormat().
				actual := sanitizeFormat(test.format)

			// Check that repeat() outputted what we expected.
				if test.expected != actual {
					t.Errorf("Did not get what was expected for test #%d.\nExpected: %q\nReceived: %q\n", test_number, test.expected, actual)
				}
		}
}
