package ipaddress

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IPResponse struct {
	IP string `json:"ip"`
}

func FetchPublicIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		return "", fmt.Errorf("unable to fetch public IP: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("unable to read response body: %w", err)
	}

	var ipResp IPResponse
	err = json.Unmarshal(body, &ipResp)
	if err != nil {
		return "", fmt.Errorf("unable to parse response JSON: %w", err)
	}

	return ipResp.IP, nil
}
