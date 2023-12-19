package tempfiles

import (
	"fmt"
	"os"
)

// WorkWithTemp will give some basic patterns for working
// with temporary files and directories
func WorkWithTemp() error {
	t, err := os.MkdirTemp("", "tmp")
	if err != nil {
		return err
	}

	defer os.RemoveAll(t)
	// the directory must exist to create the tempfile
	// created. t is an *os.File object.
	tf, err := os.MkdirTemp(t, "tmp")
	if err != nil {
		return err
	}

	fmt.Println(tf)
	return nil
}
