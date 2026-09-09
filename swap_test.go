package swap_files

import (
	"os"
	"testing"
)

func TestSwap(t *testing.T) {
	// Prepare test files

	err := os.MkdirAll("./test-files", 0700)

	if err != nil {
		panic(err)
	}

	file_1 := "./test-files/file_1.txt"
	file_1_contents := "Contents of file 1"

	err = os.WriteFile(file_1, []byte(file_1_contents), 0600)

	if err != nil {
		panic(err)
	}

	file_2 := "./test-files/file_2.txt"
	file_2_contents := "Contents of file 2"

	err = os.WriteFile(file_2, []byte(file_2_contents), 0600)

	if err != nil {
		panic(err)
	}

	// Swap

	file_swap := "./test-files/file_swap.txt"

	err = SwapFiles(file_1, file_2, file_swap)

	if err != nil {
		panic(err)
	}

	// Check

	file_1_contents_new, err := os.ReadFile(file_1)

	if err != nil {
		panic(err)
	}

	if string(file_1_contents_new) != file_2_contents {
		t.Errorf("Expected %v contents to be '%v', but got '%v'", file_1, file_2_contents, file_1_contents_new)
	}

	file_2_contents_new, err := os.ReadFile(file_2)

	if err != nil {
		panic(err)
	}

	if string(file_2_contents_new) != file_1_contents {
		t.Errorf("Expected %v contents to be '%v', but got '%v'", file_2, file_1_contents, file_2_contents_new)
	}
}
