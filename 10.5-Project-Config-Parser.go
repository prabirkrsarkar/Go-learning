package main

import (
	"bufio"
	"fmt"
	"strings"
)

func parseConfig(content string) (map[string]string, error) {

	config := make(map[string]string)

	// Create a new scanner to read the content line by line
	// bufio.NewScanner accepts a io.Reader interface type
	// type Reader interface {
	//     Read(p []byte) (n int, err error)
	// }
	// Anything that implements Read can be used as a data source for bufio.NewScanner

	scanner := bufio.NewScanner(strings.NewReader(content))
	lineNo := 0

	// Iterate through each line of the configuration content
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		// Skip empty lines and comments
		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "#") {
			continue
		}

		// Split the line into key and value
		parts := strings.SplitN(trimmedLine, "=", 2)

		// If the string is not in key=value format discard it
		if len(parts) != 2 || !strings.Contains(trimmedLine, "=") {
			fmt.Printf("Invalid configuration at line %d: %s\n", lineNo, line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove surrounding quotes from the value if present
		if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
			value = strings.Trim(value, `"`)
		}

		config[key] = value
	}

	// After the scanning is complete, check for any errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading configuration: %v", err)
	}

	return config, nil
}

func main() {

	// This is a multiline string of configuration settings
	envFileContent := `
#Application Configurations
APP_NAME="optra"
APP_VERSION="1.0.2"
DEBUG_MODE=true
#Database Configurations
DB_HOST="localhost"
DB_PORT = 5432
DB_USER="user"
DB_PASSWORD="p@ass$word WITH SP@ACES!"
#API Configurations
API_KEY=
API_SECRET="your_api_secret"
`
	config, err := parseConfig(envFileContent)
	if err != nil {
		fmt.Printf("Error parsing configuration: %v\n", err)
		return
	}

	// Print the parsed configuration
	for key, value := range config {
		fmt.Printf("%s: %s\n", key, value)
	}
}
