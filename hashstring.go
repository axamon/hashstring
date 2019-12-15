package hashstring

import (
	"crypto/md5"
	"fmt"
)

// Md5Sum returns the md5sum in exadecimal format of the string s.
func Md5Sum(s string) string {
	h := md5.New()

	// new line is added at the end to obtain same output of
	// common md5sum cli tools.
	h.Write([]byte(s + "\n"))
	hashedpass := h.Sum(nil)

	// output is in exadecimal format.
	return fmt.Sprintf("%x", hashedpass)
}
