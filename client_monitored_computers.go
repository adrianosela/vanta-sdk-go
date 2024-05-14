package vanta

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (v *vanta) ListMonitoredComputers(ctx context.Context) (*ListMonitoredComputersOutput, error) {
	tokenType, token := v.tokenStore.GetToken()
	if token == "" {
		return nil, errors.New("no auth token present")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/monitored-computers", v.baseURL), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build http request: %v", err)
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", tokenType, token))

	resp, err := v.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute http request: %v", err)
	}
	defer resp.Body.Close()

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read http response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 http response status code (%d), body: %s", resp.StatusCode, string(respBodyBytes))
	}

	var listMonitoredComputersOutput *ListMonitoredComputersOutput
	if err = json.Unmarshal(respBodyBytes, &listMonitoredComputersOutput); err != nil {
		return nil, fmt.Errorf("failed to JSON-decode response body: %v", err)
	}

	return listMonitoredComputersOutput, nil
}

func (v *vanta) GetMonitoredComputerByID(ctx context.Context, id string) (*MonitoredComputer, error) {
	tokenType, token := v.tokenStore.GetToken()
	if token == "" {
		return nil, errors.New("no auth token present")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/monitored-computers/%s", v.baseURL, id), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build http request: %v", err)
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", tokenType, token))

	resp, err := v.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute http request: %v", err)
	}
	defer resp.Body.Close()

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read http response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 http response status code (%d), body: %s", resp.StatusCode, string(respBodyBytes))
	}

	var monitoredComputer *MonitoredComputer
	if err = json.Unmarshal(respBodyBytes, &monitoredComputer); err != nil {
		return nil, fmt.Errorf("failed to JSON-decode response body: %v", err)
	}

	return monitoredComputer, nil
}
