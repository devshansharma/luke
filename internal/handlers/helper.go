package handlers

import (
	"fmt"
	"os"
	"strings"
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
