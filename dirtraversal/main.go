package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSize := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSize)
		}
		close(fileSize)
	}()

	printDiskUsage(fileSize)
}

func walkDir(dir string, fileSizes chan int64) {
	entries := dirents(dir)
	for _, entry := range entries {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileInfo, err := entry.Info()
			if err != nil {
				panic(err)
			}
			fileSizes <- fileInfo.Size()
		}
	}
}

func dirents(dir string) []fs.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	return entries
}

func printDiskUsage(fileSize chan int64) {
	var nfiles, nbytes int64

	for nbyte := range fileSize {
		nfiles++
		nbytes += nbyte
	}
	fmt.Printf("The Number of Files: %d\n", nfiles)
	fmt.Printf("Total Size: %.1f GB\n", float64(nbytes/1e9))
}
