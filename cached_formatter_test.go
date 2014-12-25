package main


import "testing"
import "fmt"


func TestCachedFormat(t *testing.T) {

	tests := []struct{
		actions []struct{s string ; was_in_cache bool}
	}{
		{
			actions: []struct{s string ; was_in_cache bool}{
				struct{s string ; was_in_cache bool}{
					s: "apple",
					was_in_cache: false,
				},
				struct{s string ; was_in_cache bool}{
					s: "banana",
					was_in_cache: false,
				},
				struct{s string ; was_in_cache bool}{
					s: "cherry",
					was_in_cache: false,
				},
				struct{s string ; was_in_cache bool}{
					s: "apple",
					was_in_cache: true,
				},
				struct{s string ; was_in_cache bool}{
					s: "banana",
					was_in_cache: true,
				},
				struct{s string ; was_in_cache bool}{
					s: "cherry",
					was_in_cache: true,
				},
			},
		},
		{
			actions: []struct{s string ; was_in_cache bool}{
				struct{s string ; was_in_cache bool}{
					s: "apple",
					was_in_cache: false,
				},
				struct{s string ; was_in_cache bool}{
					s: "apple",
					was_in_cache: true,
				},
				struct{s string ; was_in_cache bool}{
					s: "apple",
					was_in_cache: true,
				},
				struct{s string ; was_in_cache bool}{
					s: "apple",
					was_in_cache: true,
				},
				struct{s string ; was_in_cache bool}{
					s: "banana",
					was_in_cache: false,
				},
				struct{s string ; was_in_cache bool}{
					s: "apple",
					was_in_cache: true,
				},
				struct{s string ; was_in_cache bool}{
					s: "banana",
					was_in_cache: true,
				},
			},
		},
	}

	// Run each test.
	//
		for test_number, test := range tests {

			// Create new CachedFormatter.
				format := "<([%s])>"

				var formatter            Formatter = NewPrintfStyleFormatter(format)

				cached_formatter := NewCachedFormatter(formatter, 10)

			// Run each action, and make sure we get what is expected at each step.
				for action_number, action := range test.actions {

					actual_formatted, actual_was_in_cache := cached_formatter.CachedFormat(action.s)

					if expected_formatted := fmt.Sprintf(format, action.s) ; expected_formatted != actual_formatted {
						t.Errorf("Did not get what was expected for formatting for test #%d on action #%d.\nExpected: %q\nReceived: %q\n",  test_number, action_number, expected_formatted, actual_formatted)

					}
					if expected_was_in_cache := action.was_in_cache ; expected_was_in_cache != actual_was_in_cache {
						t.Errorf("Did not get what was expected for if it was in th cache for test #%d on action #%d.\nExpected: %q\nReceived: %q\n",  test_number, action_number, expected_was_in_cache, actual_was_in_cache)
					}
				}

		}

}
