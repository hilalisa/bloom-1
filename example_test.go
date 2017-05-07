package bloom_test

import (
	"fmt"
	"github.com/yourbasic/bloom"
)

// Create and use a Bloom filter.
func Example() {
	// Create a Bloom filter with room for 10000 elements
	// at a false-positives rate less than 0.5 percent.
	blacklist := bloom.New(10000, 200)

	// Add a string to the filter.
	url := "https://rascal.com"
	blacklist.Add(url)

	// Check for membership.
	if blacklist.Likely(url) {
		fmt.Println(url, "seems to be shady")
	}
	// Output: https://rascal.com seems to be shady
}
