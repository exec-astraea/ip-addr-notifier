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
