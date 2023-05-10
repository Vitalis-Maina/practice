package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Directory is not specified")
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

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if format == "" || filepath.Ext(path) == "."+format {
				size := info.Size()
				fileGroups[size] = append(fileGroups[size], path)
			}
		}
		return err
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
