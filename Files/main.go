package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Directory is not Specified")
		return
	}

	root := os.Args[1]

	var format string
	fmt.Println("Enter file format:")
	fmt.Scanln(&format)

	fmt.Println("Size sorting options:")
	fmt.Println("1. Descending")
	fmt.Println("2. Ascending")

	var option int
	for {
		fmt.Println("Enter a sorting option:")
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Wrong option")
			continue
		}
		if option != 1 && option != 2 {
			fmt.Println("Wrong option")
			continue
		}
		break
	}

	fileGroups := make(map[int64][]string)

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if !d.IsDir() {
			ext := strings.TrimPrefix(filepath.Ext(path), ".")
			if format == "" || strings.EqualFold(ext, format) {
				info, err := d.Info()
				if err != nil {
					return err
				}
				size := info.Size()
				fileGroups[size] = append(fileGroups[size], path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	var sizes []int64
	for size := range fileGroups {
		sizes = append(sizes, size)
	}

	if option == 1 {
		sort.Slice(sizes, func(i, j int) bool {
			return sizes[i] > sizes[j]
		})
	} else {
		sort.Slice(sizes, func(i, j int) bool {
			return sizes[i] < sizes[j]
		})
	}

	fmt.Println()
	for _, size := range sizes {
		paths := fileGroups[size]
		fmt.Printf("%d bytes\n", size)
		for _, path := range paths {
			fmt.Println(path)
		}
		fmt.Println()
	}
}
