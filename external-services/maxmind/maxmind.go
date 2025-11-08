package maxmind

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"pebu-go-demo/internal/logger"
)

type Maxmind interface {
	GetGeoIpData(ctx context.Context, ipAddress string, bearerToken string) (*GeoIpResponse, error)
}

type maxmindImpl struct {
	logger logger.Logger
}

func New(logger logger.Logger) Maxmind {
	return maxmindImpl{
		logger: logger,
	}
}

func (m maxmindImpl) GetGeoIpData(ctx context.Context, ipAddress string, bearerToken string) (*GeoIpResponse, error) {
	// Define the IP address and URL
	url := fmt.Sprintf("https://geoip.maxmind.com/geoip/v2.1/city/%s?demo=1", ipAddress)

	// Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Add headers to the request
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "en-US,en;q=0.8")
	req.Header.Add("authorization", bearerToken)
	req.Header.Add("origin", "https://www.maxmind.com")
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("referer", "https://www.maxmind.com/")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("sec-gpc", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		m.logger.Error(ctx, "Error on maxmind response", map[string]any{
			"error": err.Error(),
		})
		return nil, err
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		m.logger.Error(ctx, "Error reading response body", map[string]any{
			"error": err.Error(),
		})
		return nil, err
	}

	m.logger.Info(ctx, "Maxmind returned result", map[string]any{
		"responseBody": string(body),
	})

	if resp.StatusCode != http.StatusOK {
		m.logger.Error(ctx, "Maxmind status did not return expected OK status", map[string]any{
			"status": resp.StatusCode,
		})
		// Unmarshal error response
		errorResponseBody := GeoIpErrorResponse{}
		err = json.Unmarshal(body, &errorResponseBody)
		if err != nil {
			m.logger.Error(ctx, "Error unmarshaling maxmind response", map[string]any{
				"error": err.Error(),
			})
			return nil, err
		}
		return nil, fmt.Errorf("expected %v but got %v. Error from maxmind: %v, %v",
			http.StatusOK, resp.StatusCode,
			errorResponseBody.Code, errorResponseBody.Error)
	}

	// Unmarshal success response
	responseBody := GeoIpResponse{}
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		m.logger.Error(ctx, "Error unmarshaling maxmind response", map[string]any{
			"error": err.Error(),
		})
		return nil, err
	}

	return &responseBody, nil
}
