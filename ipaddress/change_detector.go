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
)

type ChangeDetectionResult struct {
	Changed   bool
	CurrentIP string
}

func DetectChange() (ChangeDetectionResult, error) {
	currentIP, err := FetchPublicIP()
	if err != nil {
		return ChangeDetectionResult{}, fmt.Errorf("unable to get current IP: %w", err)
	}

	lastIP, err := LoadLastIP()
	if err != nil {
		return ChangeDetectionResult{}, fmt.Errorf("unable to load last IP: %w", err)
	}

	if currentIP == lastIP {
		return ChangeDetectionResult{Changed: false, CurrentIP: currentIP}, nil
	}

	err = UpdateLastIP(currentIP)
	if err != nil {
		return ChangeDetectionResult{}, fmt.Errorf("unable to update last IP: %w", err)
	}

	return ChangeDetectionResult{Changed: true, CurrentIP: currentIP}, nil
}
