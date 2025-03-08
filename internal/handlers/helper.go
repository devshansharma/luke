package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/devshansharma/luke/internal/models"
	"github.com/devshansharma/luke/pkg/utils"
)

type FileSuffix string

var (
	collectionSuffix  FileSuffix = "_collection.json"
	environmentSuffix FileSuffix = "_env.json"
)

// filterCollections to get collection list
func filter(entries []os.DirEntry, suffix FileSuffix) []os.DirEntry {
	dirEntries := make([]os.DirEntry, 0)

	for _, dir := range entries {
		if strings.HasSuffix(dir.Name(), string(suffix)) {
			dirEntries = append(dirEntries, dir)
		}
	}

	return dirEntries
}

func filterEnvironment(entries []os.DirEntry) []os.DirEntry {
	return filter(entries, environmentSuffix)
}

func filterCollection(entries []os.DirEntry) []os.DirEntry {
	return filter(entries, collectionSuffix)
}

// getCollectionName to get collection name
func getCollectionName(entry os.DirEntry) string {
	return strings.ReplaceAll(strings.TrimSuffix(entry.Name(), string(collectionSuffix)), "_", " ")
}

// getEnvName to get collection name
func getEnvName(entry os.DirEntry) string {
	return strings.ReplaceAll(strings.TrimSuffix(entry.Name(), string(environmentSuffix)), "_", " ")
}

func getCollectionFileName(name string) string {
	name = strings.ReplaceAll(name, " ", "_")
	return fmt.Sprintf("%s%s", name, collectionSuffix)
}

func getEnvFileName(name string) string {
	name = strings.ReplaceAll(name, " ", "_")
	return fmt.Sprintf("%s%s", name, environmentSuffix)
}

func writeToFile(obj models.Collection, fileName string) error {
	data, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}

	fileData := string(data) + "\n"
	err = os.WriteFile(fileName, []byte(fileData), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write to file: %s", err)
	}

	return nil
}

func getCollection(name string) (obj models.Collection, fileName string, err error) {
	dir, err := utils.GetConfigDir()
	if err != nil {
		return
	}

	fileName = fmt.Sprintf("%s/%s", dir, getCollectionFileName(name))

	_, err = os.Stat(fileName)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		err = fmt.Errorf("collection does not exist")
		return
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		err = fmt.Errorf("failed to get collection: %s", err.Error())
		return
	}

	err = json.Unmarshal(data, &obj)
	if err != nil {
		err = fmt.Errorf("failed to parse collection: %s", err.Error())
		return
	}

	return
}
