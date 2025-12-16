package main

import (
	"fmt"
	"os"
	"flag"
)

const (
	col_reset = "\033[0m"
	col_blue = "\033[34m"
)

type Path struct {
	d []os.DirEntry
	e string
}

func main() {
	all := flag.Bool("a", false, "a bool")
	flag.Parse()

	paths := flag.Args()
	if len(paths) == 0 {
		paths = append(paths, ".")
	}

	results := make(map[string]Path)
	
	for _, path := range paths {
		files, err := os.ReadDir(path)
		if err != nil {
			results[path] = Path{e: fmt.Sprintf("gols: cannot access '%s': No such file or directory\n", path)}
			continue
		}

		results[path] = Path{d: files}
	}

	num_paths := len(paths)

	for i, path := range paths {
		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Printf("gols: cannot access '%s': No such file or directory\n", path)
			continue
		}

		if i > 0 {
			fmt.Printf("\n")
		}

		if num_paths > 1 {
			fmt.Printf("%v:\n", path)
		}

		for _, f := range files {
			file_name := f.Name()
			if file_name[0] == '.' && !*all {
				continue
			}

			if f.IsDir() {
				fmt.Printf("%s%s%s  ", col_blue, file_name, col_reset)
				continue
			}

			fmt.Printf("%v\t", file_name)
		}
		fmt.Printf("\n")
	}
}
