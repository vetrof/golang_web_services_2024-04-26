package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	return printDir(out, path, printFiles, "")
}

func printDir(out io.Writer, path string, printFiles bool, prefix string) error {
	// Получаем содержимое директории
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	// Сортировка (сначала идут файлы, потом папки)
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	// Фильтрация, если не нужно выводить файлы
	if !printFiles {
		var dirsOnly []os.DirEntry
		for _, entry := range entries {
			if entry.IsDir() {
				dirsOnly = append(dirsOnly, entry)
			}
		}
		entries = dirsOnly
	}

	// Обход элементов
	for i, entry := range entries {
		connector := "├───"
		if i == len(entries)-1 {
			connector = "└───"
		}

		// Путь к текущему элементу
		fullPath := filepath.Join(path, entry.Name())

		if entry.IsDir() {
			// Для директорий - печатаем имя и уходим в рекурсию
			fmt.Fprintf(out, "%s%s%s\n", prefix, connector, entry.Name())
			newPrefix := prefix + "│\t"
			if i == len(entries)-1 {
				newPrefix = prefix + "\t"
			}
			err := printDir(out, fullPath, printFiles, newPrefix)
			if err != nil {
				return err
			}
		} else {
			// Для файлов - добавляем информацию о размере
			info, err := entry.Info()
			if err != nil {
				return err
			}
			sizeStr := fmt.Sprintf("(%db)", info.Size())
			if info.Size() == 0 {
				sizeStr = "(empty)"
			}
			fmt.Fprintf(out, "%s%s%s %s\n", prefix, connector, entry.Name(), sizeStr)
		}
	}
	return nil
}
