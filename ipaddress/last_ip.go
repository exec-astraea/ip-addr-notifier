// Copyright (C) 2025 Alisa Frelia
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

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
