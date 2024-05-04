// Add a test for every unique use case in every project array is used in
// so that changes to the source will break the test before being pushed
// and breaking the project they are used in
package array_test

import (
	"bytes"
	"log"
	"os"
)

func cap(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
