// basename removes directory components and a .suffix
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c

package main

import (
	"fmt"
)

func main() {
	s := "a/b/c.d.go"
	sr := basename(s)
	fmt.Printf("hahah: %s", sr)
}

func basename(s string) string {

	// Discard last '/' and everyting before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Preserve everyting before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}
