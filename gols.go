package main

import (
	"fmt"
	"os"
	"strings"
	"golang.org/x/term"
)

const (
	col_reset = "\033[0m"
	col_blue = "\033[34m"
)


func main() {
	args := os.Args[1:]
	paths := make([]string, len(os.Args))
	
	all := false

	paths_count := 0
	for _, arg := range args {
		if arg[0] == '-' {
			if arg == "-a" || arg == "--all" {
				all = true
			}
			continue
		}

		paths[paths_count] = arg
		paths_count++
	}

	if paths_count == 0 {
		paths[paths_count] = "."
	}

	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	for i, path := range paths {
		if path == "" {
			continue
		}

		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Printf("gols: cannot access '%s': No such file or directory\n", path)
			continue
		}

		files, max_len := filterFiles(files, all)
		column_num := width / (max_len + 2)

		column_spaces := getSpaces(column_num, files)

		if i > 0 {
			fmt.Printf("\n")
		}

		if paths_count > 1 {
			fmt.Printf("%v:\n", path)
		}

		printFiles(files, column_spaces, column_num)
		fmt.Printf("\n")
	}
}


func filterFiles(sd []os.DirEntry, all bool) ([]os.DirEntry, int) {
	max_len, count := 0, 0
	filtered := make([]os.DirEntry, len(sd))

	for _, d := range sd {
		entry := d.Name()
		if entry[0] == '.' && !all {
			continue
		}

		filtered[count] = d
		count++

		if entry_len := len(entry); entry_len > max_len {
			max_len = entry_len
		}
	}

	return filtered[:count], max_len
}


func getSpaces(columns int, dir_entries []os.DirEntry) []int {
	column_spaces := make([]int, columns)

	for i, entry := range dir_entries {
		idx := i % columns
		if len_entry := len(entry.Name()); len_entry > column_spaces[idx] {
			column_spaces[idx] = len_entry
		}
	}

	return column_spaces
}


func printFiles(files []os.DirEntry, spaces []int, columns int) {
	for i, file := range files {
		idx := i % columns
		if idx == 0 && i != 0 {
			fmt.Print("\n")
		}

		name := file.Name()
		space := spaces[idx] - len(name) + 2
		indent := strings.Repeat(" ", space)

		if file.IsDir() {
			fmt.Printf("%s%s%s%s", col_blue, name, col_reset, indent)
			continue
		}
		fmt.Printf("%s%s", name, indent)
	}
}
