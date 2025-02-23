package collection

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"errors"

	"github.com/devshansharma/luke/internal/utils"
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
	data, err := json.MarshalIndent(c, "", "    ")
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
	entries, err := utils.GetDirEntries(fileSuffix)
	if err != nil {
		return err
	}

	writer := writer.GetInstance()
	for index, entry := range entries {
		name := strings.TrimSuffix(entry.Name(), fileSuffix)
		writer.Response(fmt.Sprintf("%02d  %s", index+1, strings.ReplaceAll(name, "_", " ")))
	}

	return nil
}

func DeleteHandler(arg int) error {
	entries, err := utils.GetDirEntries(fileSuffix)
	if err != nil {
		return err
	}

	if arg < 1 || arg > len(entries) {
		return errors.New("invalid id")
	}

	name := strings.TrimSuffix(entries[arg-1].Name(), fileSuffix)
	name = strings.ReplaceAll(name, "_", " ")

	resp, err := utils.AskYesOrNo(fmt.Sprintf("Are you sure, you want to delete %q collection", name))
	if err != nil {
		return err
	}

	if resp {
		dirPath, err := utils.HomeDir()
		if err != nil {
			return err
		}

		fileName := fmt.Sprintf("%s/%s%s", dirPath, strings.ReplaceAll(name, " ", "_"), fileSuffix)

		err = os.Remove(fileName)
		if err != nil {
			return err
		}

		writer.GetInstance().Response(fmt.Sprintf("%q delete successfully", fileName))
	}

	return nil
}
