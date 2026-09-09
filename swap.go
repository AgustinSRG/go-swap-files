package swap_files

// Swaps 2 files.
//
// Parameters:
//   - firstFile: Path to the first file
//   - secondFile: Path to the second file
//   - swapFile: Path to an auxiliary file to use if necessary
//
// In Unix-based operating systems, the swap is atomic. The swap file is ignored.
//
// For other operating systems, the swap is divided in 3 operations:
//  1. rename(secondFile, swapFile)
//  2. rename(firstFile, secondFile)
//  3. rename(swapFile, firstFile)
//
// Due to its non-atomic nature in non-Unix operating systems,
// make sure to add a fail-safe in case the firstFile does not exists
// but swapFile does, indicating the swap was interrupted.
//
// Return value: An error (if the swap fails), or nil (success)
func SwapFiles(firstFile string, secondFile string, swapFile string) error {
	return swapFilesInternal(firstFile, secondFile, swapFile)
}
