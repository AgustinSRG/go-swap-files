//go:build unix

package swap_files

import "golang.org/x/sys/unix"

func swapFilesInternal(firstFile string, secondFile string, _ string) error {
	return unix.Renameat2(
		unix.AT_FDCWD,
		firstFile,
		unix.AT_FDCWD,
		secondFile,
		unix.RENAME_EXCHANGE,
	)
}
