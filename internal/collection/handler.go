package collection

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"errors"

	"github.com/devshansharma/luke/pkg/writer"
	"github.com/google/uuid"
)

const fileSuffix = "_collection.json"

func AddHandler(c Collection) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dirPath := fmt.Sprintf("%s/.luke", homeDir)
	fileName := fmt.Sprintf("%s/%s%s", dirPath, strings.ReplaceAll(c.Name, " ", "_"), fileSuffix)

	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return err
	}

	_, err = os.Stat(fileName)
	if err == nil {
		return errors.New("collection already exists")
	}

	if !os.IsNotExist(err) {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	c.ID = uuid.NewString()
	data, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func ListHandler() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	dirPath := fmt.Sprintf("%s/.luke", homeDir)

	dirEntry, err := os.ReadDir(dirPath)
	if err != nil {
		return err
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

	writer := writer.GetInstance()
	for index, entry := range filteredList {
		name := strings.TrimSuffix(entry.Name(), fileSuffix)
		writer.Response(fmt.Sprintf("%02d  %s", index+1, strings.ReplaceAll(name, "_", " ")))
	}

	return nil
}
