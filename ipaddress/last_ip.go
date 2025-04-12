package ipaddress

import (
	"fmt"
	"io"
	"os"
)

const lastIPFilePath = "last_ip.txt"

func UpdateLastIP(ip string) error {
	file, err := os.Create(lastIPFilePath)
	if err != nil {
		return fmt.Errorf("unable to create file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(ip)
	if err != nil {
		return fmt.Errorf("unable to write to file: %w", err)
	}

	return nil
}

func LoadLastIP() (string, error) {
	file, err := os.Open(lastIPFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil // File does not exist, treat as no previous IP
		}
		return "", fmt.Errorf("unable to open file: %w", err)
	}
	defer file.Close()

	ip, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("unable to read file: %w", err)
	}

	return string(ip), nil
}
