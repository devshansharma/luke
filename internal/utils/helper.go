package utils

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func HomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dirPath := fmt.Sprintf("%s/.luke", homeDir)

	return dirPath, nil
}

func AskYesOrNo(prompt string) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt + " (y/n): ")

		input, err := reader.ReadString('\n')
		if err != nil {
			return false, err
		}

		input = strings.ToLower(strings.TrimSpace(input))
		if input == "y" || input == "yes" {
			return true, nil
		} else if input == "n" || input == "no" {
			return false, nil
		}

		fmt.Println("invalid input. Please enter 'y' or 'n'.")
	}
}

// GetDirEntries to get dir entries
func GetDirEntries(fileSuffix string) ([]os.DirEntry, error) {
	dirPath, err := HomeDir()
	if err != nil {
		return nil, err
	}

	dirEntry, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	sort.Slice(dirEntry, func(i, j int) bool {
		return dirEntry[i].Name() < dirEntry[j].Name()
	})

	filteredList := make([]os.DirEntry, 0)

	for _, entry := range dirEntry {
		if strings.HasSuffix(entry.Name(), fileSuffix) {
			filteredList = append(filteredList, entry)
		}
	}

	return filteredList, nil
}
