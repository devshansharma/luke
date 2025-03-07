package utils

import (
	"errors"
	"fmt"
	"os"
)

// GetConfigDir to get luke configuration directory
// also used for the purpose of storing collections and environments.
func GetConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home dir:%s", err)
	}

	lukeDir := fmt.Sprintf("%s/.luke", homeDir)

	_, err = os.Stat(lukeDir)

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("failed to get config directory for luke: %s", err)
	}

	if err != nil && errors.Is(err, os.ErrNotExist) {
		makeErr := os.MkdirAll(lukeDir, 0755)
		if makeErr != nil {
			return "", makeErr
		}
	}

	return lukeDir, nil
}
