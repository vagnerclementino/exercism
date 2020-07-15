package hamming

import "testing"

func TestHamming(t *testing.T) {
	for _, tc := range testCases {
		got, err := Distance(tc.s1, tc.s2)
		if tc.expectError {
			// check if err is of error type
			var _ error = err

			// we expect error
			if err == nil {
				t.Fatalf("Distance(%q, %q); expected error, got nil.",
					tc.s1, tc.s2)
			}
		} else {
			// we do not expect error
			if err != nil {
				t.Fatalf("Distance(%q, %q) returned unexpected error: %v",
					tc.s1, tc.s2, err)
			}
			if got != tc.want {
				t.Fatalf("Distance(%q, %q) = %d, want %d.",
					tc.s1, tc.s2, got, tc.want)
			}

		}
	}
}

func BenchmarkHamming(b *testing.B) {
	// bench combined time to run through all test cases
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			// ignoring errors and results because we're just timing function execution
			_, _ = Distance(tc.s1, tc.s2)
		}
	}
}

func TestDivideInChunks(t *testing.T) {
	for _, tc := range divideInChunksTestCases {
		got, err := divideInChunks(tc.textToDivide, tc.chunkSize)
		if tc.expectedError != nil {
			// check if err is of error type
			var _ error = err

			// we expect error
			if err == nil {
				t.Fatalf(" %q: divideInChunks(%s, %d); expected error, got nil.",
					tc.title, tc.textToDivide, tc.chunkSize)
			}
		} else {
			// we do not expect error
			if err != nil {
				t.Fatalf("%q: divideInChunks(%s, %d) returned unexpected error: %v",
					tc.title, tc.textToDivide, tc.chunkSize, err)
			}
			if !compareChunk(got, tc.expectedChunks) {
				t.Fatalf("%q : divideInChunks(%s, %d) = %v, want %v.", tc.title,
					tc.textToDivide, tc.chunkSize, got, tc.expectedChunks)
			}
		}
	}
}

func compareChunk(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
