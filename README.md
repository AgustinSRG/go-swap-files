# Library for swapping files (Go)

This is a simple library for **Go** to swap 2 files, ensuring cross-platform compatibility while ensuring maximum efficiency and atomicity if possible.

 - In Unix operating systems, the swap is atomic using the [renameat2 syscall](https://lwn.net/Articles/569134/).
 - For other operating systems, the swap is done using an intermediary swap file that must be specified. 

[Documentation](https://pkg.go.dev/github.com/AgustinSRG/go-swap-files)

## Usage

Import the module into your project:

```sh
go get github.com/AgustinSRG/go-swap-files
```

Example usage:

```go

package main

import (
	// Import the module
	swap_files "github.com/AgustinSRG/go-child-process-manager"
)

func main() {
    err := swap_files.SwapFiles("file.original", "file.copy", "file.swap")

    if err != nil {
		panic(err)
	}
}
```

## Testing

In order to run the tests, type:

```sh
go test -v
```
