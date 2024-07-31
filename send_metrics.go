package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Data structure for the Datadog API request
type Point struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

type MetricSeries struct {
	Metric    string     `json:"metric"`
	Type      string     `json:"type"`
	Points    []Point    `json:"points"`
	Resources []Resource `json:"resources"`
	Tags      []string   `json:"tags"`
}

type Resource struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type DatadogRequest struct {
	Series []MetricSeries `json:"series"`
}

func main() {
	// Load the Trivy JSON report
	file, err := ioutil.ReadFile("trivy.json")
	if err != nil {
		log.Fatalf("Error reading Trivy JSON report: %v", err)
	}

	var trivyReport map[string]interface{}
	if err := json.Unmarshal(file, &trivyReport); err != nil {
		log.Fatalf("Error parsing Trivy JSON report: %v", err)
	}

	// Extract vulnerability counts
	highCount := 0
	criticalCount := 0
	for _, result := range trivyReport["Results"].([]interface{}) {
		vulnerabilities := result.(map[string]interface{})["Vulnerabilities"]
		if vulnerabilities != nil {
			for _, vuln := range vulnerabilities.([]interface{}) {
				severity := vuln.(map[string]interface{})["Severity"]
				switch severity {
				case "HIGH":
					highCount++
				case "CRITICAL":
					criticalCount++
				}
			}
		}
	}

	timestamp := time.Now().Unix()
	data := DatadogRequest{
		Series: []MetricSeries{
			{
				Metric: "trivy.vulnerabilities.high",
				Type:   "gauge",
				Points: []Point{{Timestamp: timestamp, Value: float64(highCount)}},
				Resources: []Resource{
					{Name: "example-host", Type: "host"},
				},
				Tags: []string{"severity:high", "source:trivy"},
			},
			{
				Metric: "trivy.vulnerabilities.critical",
				Type:   "gauge",
				Points: []Point{{Timestamp: timestamp, Value: float64(criticalCount)}},
				Resources: []Resource{
					{Name: "example-host", Type: "host"},
				},
				Tags: []string{"severity:critical", "source:trivy"},
			},
		},
	}

	// Send the data to Datadog
	sendDataToDatadog(data)
}

func sendDataToDatadog(data DatadogRequest) {
	url := "https://api.datadoghq.com/api/v1/series"
	apiKey := os.Getenv("DD_API_KEY")
	if apiKey == "" {
		log.Fatal("DD_API_KEY is not set")
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshaling data to JSON: %v", err)
	}

	// Debug: Print the JSON data being sent to Datadog
	fmt.Println("JSON Payload:", string(jsonData))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("DD-API-KEY", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending data to Datadog: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}
