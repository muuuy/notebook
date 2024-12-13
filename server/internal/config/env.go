package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// loadEnvFile loads the environment file and sets environment variables.
//
// Parameters:
//   - filepath (string): The environment file's path
//
// Returns:
//   - error: Returns an error if an error occurs
//   - nil: If everything goes right, return nothing
func loadEnvFile(filepath string) error {
	file, err := os.Open(filepath)

	if err != nil {
		return fmt.Errorf("error opening .env file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Ignore empty lines / comments
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		// Split string into the two parts: Variable name & value
		parts := strings.SplitN(line, "=", 2)

		// Something went wrong when splitting
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		err := os.Setenv(key, value)
		if err != nil {
			return fmt.Errorf("error setting up environment variable %s: %v", key, err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading .env file: %v", err)
	}

	return nil
}
