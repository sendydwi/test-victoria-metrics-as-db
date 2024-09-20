package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"math/rand"

	"time"

	"github.com/google/uuid"
)

type ValidPayload struct {
	Metric string  `json:"metric"`
	Value  float32 `json:"value"`
	Tags   Tags    `json:"tags"`
}

type Tags struct {
	MeterID string `json:"meter_id"`
	Status  string `json:"Status"`
}

func main() {
	url := "http://localhost:4242/api/put"

	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	for i := 1; i < 1000000; i++ {
		var status string

		if randomizer.Intn(100) < 20 {
			status = "Failed"
		} else {
			status = "Success"
		}

		format := ValidPayload{
			Metric: "metric_meter_1",
			Value:  float32(randomizer.Intn(100)),
			Tags: Tags{
				MeterID: uuid.New().String(),
				Status:  status,
			},
		}

		payload, err := json.Marshal(format)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()

		// Print response status
		fmt.Println(i, "Response status:", resp.Status)
	}

	queryUrl := "http://localhost:8428/prometheus/api/v1/query_range"

	countFailed := `count(metric_meter_1{Status="Failed"}[30d])`
	countSuccess := `count(metric_meter_1{Status="Success"}[30d])`

	resFailed, err := http.Get(fmt.Sprintf("%s?query=%s", queryUrl, countFailed))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	bodyFailed, err := io.ReadAll(resFailed.Body)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	fmt.Println(string(bodyFailed))

	resSuccess, err := http.Get(fmt.Sprintf("%s?query=%s", queryUrl, countSuccess))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	bodySuccess, err := io.ReadAll(resSuccess.Body)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	fmt.Println(string(bodySuccess))
}
