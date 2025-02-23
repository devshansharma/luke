package collection

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"errors"

	"github.com/google/uuid"
)

func Handler(cfg Config) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dirPath := fmt.Sprintf("%s/.luke", homeDir)
	fileName := fmt.Sprintf("%s/%s_collection.json", dirPath, strings.ReplaceAll(cfg.Name, " ", "_"))

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

	cfg.ID = uuid.NewString()
	data, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
