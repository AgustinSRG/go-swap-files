//go:build !unix

package swap_files

import "os"

func swapFilesInternal(firstFile string, secondFile string, swapFile string) error {
	var err error

	err = os.Rename(secondFile, swapFile)

	if err != nil {
		return err
	}

	err = os.Rename(firstFile, secondFile)

	if err != nil {
		return err
	}

	err = os.Rename(swapFile, firstFile)

	if err != nil {
		return err
	}

	return nil
}
