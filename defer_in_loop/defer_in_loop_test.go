package defer_in_loop_test

import (
	"fmt"
	"os"
	"testing"
)

var (
	tmpFileList = []string{"a.txt", "b.txt", "c.txt"}
)

func closeAndRemoveFile(file *os.File) (err error) {
	if err = file.Close(); err != nil {
		return err
	}
	fmt.Println("file", file.Name(), "closed!")
	if err = os.Remove(file.Name()); err != nil {
		return err
	}
	fmt.Println("file", file.Name(), "removed!")
	return nil
}

/*
Outputs:

loop ends.
file c.txt closed!
file c.txt removed!
file b.txt closed!
file b.txt removed!
file a.txt closed!
file a.txt removed!
*/
func TestDeferInLoop(t *testing.T) {
	for i := range tmpFileList {
		file, err := os.Create(tmpFileList[i])
		if err != nil {
			t.Fail()
		}
		// Possible resource leak, 'defer' is called in the 'for' loop
		// All files will be removed after loop ends
		defer func() {
			if err := closeAndRemoveFile(file); err != nil {
				fmt.Println("closeAndRemoveFile failed, get error:", err.Error())
			}
		}()
	}
	fmt.Println("loop ends.")
}

/*
Outputs:

file a.txt closed!
file a.txt removed!
file b.txt closed!
file b.txt removed!
file c.txt closed!
file c.txt removed!
loop ends.
*/
func TestDeferInLoopWrappedInFunction(t *testing.T) {
	for i := range tmpFileList {
		func() {
			file, err := os.Create(tmpFileList[i])
			if err != nil {
				t.Fail()
			}
			// No resource leak
			// Every file will be removed at the end of each iteration
			defer func() {
				if err := closeAndRemoveFile(file); err != nil {
					fmt.Println("closeAndRemoveFile failed, get error:", err.Error())
				}
			}()
		}()
	}
	fmt.Println("loop ends.")
}
