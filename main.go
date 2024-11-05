package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/rand"
)

func main() {
	// suggested_metric_name `{sr/odr}_{readingType}_{fieldName}`
	// suggested_tag `meter_id={}, status={parent}, insert_by={fieldname}, trx_id={transaction_id}, request_date={}`
	// suggested_value `field_value` in float64
	// suggested_timestampt `read_date` in timestamptz

	// 4KiB

	url := "http://localhost:4242/api/put"

	// generateRandomSingleData(url)
	// ReadSingleData()

	exampleData := getTemplateExamplePayload()
	generateRandomDataFromTemplate(*exampleData, url)
}

func getTemplateExamplePayload() *Data {
	data, err := os.ReadFile("example.json")
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return nil
	}

	var examplePayload Data

	err = json.Unmarshal(data, &examplePayload)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v", err)
		return nil
	}

	return &examplePayload
}

func generateRandomDataFromTemplate(template Data, url string) {
	meterId := template.MeterID
	operation := template.MeterOperationType
	status := template.ErrorCode
	insertBy := template.MeterResponse.InsertBy
	trxId := template.TrxID
	dataType := reflect.TypeOf(template.MeterResponse)
	valueof := reflect.ValueOf(template.MeterResponse)
	if dataType.Kind() == reflect.Ptr {
		dataType = dataType.Elem()
	}

	fmt.Println(valueof)
	for i := 1; i <= 10000; i++ {
		time.Sleep(5 * time.Second)

		data := []ValidPayload{}
		for i := 0; i < dataType.NumField(); i++ {
			field := dataType.Field(i)
			jsonTag := field.Tag.Get("json")

			vmenabled := field.Tag.Get("vmdata")
			if vmenabled != "true" {
				continue
			}

			value, err := strconv.ParseFloat(valueof.Field(i).String(), 64)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			metricName := fmt.Sprintf("odr_%s_%s", operation, jsonTag)
			format := ValidPayload{
				Metric: metricName,
				Value:  value,
				Tags: Tags{
					MeterID:  meterId,
					Status:   status,
					InsertBy: insertBy,
					TrxID:    trxId,
				},
			}

			data = append(data, format)
		}

		payload, err := json.Marshal(data)
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
		fmt.Println(data, "Response status:", resp.Status)
	}
}

func generateRandomSingleData(url string) {
	maxGoroutines := 2
	guard := make(chan int, maxGoroutines)
	var wg sync.WaitGroup

	time := time.Now().UTC().UnixNano()
	randomizer := rand.New(rand.NewSource(uint64(time)))
	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		guard <- 1
		go func() {
			defer func() {
				<-guard
				wg.Done()
			}()

			var status string

			if randomizer.Intn(100) < 20 {
				status = "Failed"
			} else {
				status = "Success"
			}

			format := ValidPayload{
				Metric: "end_of_billing",
				Value:  float64(randomizer.Intn(100)),
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

		}()
	}

	wg.Wait()
}

func ReadSingleData() {
	queryUrl := "http://localhost:8428/prometheus/api/v1/query_range"

	countFailed := `count(end_of_billing{status="Failed"}[30d])`
	countSuccess := `count(end_of_billing{status="Success"}[30d])`

	resFailed, err := http.Get(fmt.Sprintf("%s?query=%s", queryUrl, countFailed))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	bodyFailed := new(Response)
	err = json.NewDecoder(resFailed.Body).Decode(bodyFailed)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	resSuccess, err := http.Get(fmt.Sprintf("%s?query=%s", queryUrl, countSuccess))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	bodySuccess := new(Response)
	err = json.NewDecoder(resSuccess.Body).Decode(bodySuccess)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	countSuccessRes, _ := strconv.Atoi(bodySuccess.Data.Result[0].Values[0][1].(string))
	countFailedRes, _ := strconv.Atoi(bodyFailed.Data.Result[0].Values[0][1].(string))
	total := countSuccessRes + countFailedRes

	fmt.Printf("Total success : %.2f%% with distribution : %v success / %v total", (float64(countSuccessRes)/float64(total))*100, countSuccessRes, total)
}

type ValidPayload struct {
	Metric string  `json:"metric"`
	Value  float64 `json:"value"`
	Tags   Tags    `json:"tags"`
}

type Tags struct {
	MeterID  string `json:"meter_id"`
	Status   string `json:"status"`
	InsertBy string `json:"insert_by"`
	TrxID    string `json:"trx_id"`
}

type Response struct {
	Status string           `json:"status"`
	Data   ResponseVictoria `json:"data"`
}

type ResponseVictoria struct {
	ResultType string                 `json:"resultType"`
	Result     []MetrixResultVictoria `json:"result"`
}

type MetrixResultVictoria struct {
	Metric interface{} `json:"metric"`
	Values [][]any     `json:"values"`
}

type MeterResponse struct {
	ActivePowerExport string `json:"active_power_export"`
	ActivePowerImport string `json:"active_power_import"`
	AlarmRegister     string `json:"alarm_register"`
	Clock             string `json:"clock"`
	CurrentL1         string `json:"current_l1" vmdata:"true"`
	CurrentL2         string `json:"current_l2"`
	CurrentL3         string `json:"current_l3"`
	InsertBy          string `json:"insert_by"`
	KvarhBilling      string `json:"kvarh_billing"`
	KvarhExportTotal  string `json:"kvarh_export_total"`
	KvarhImportTotal  string `json:"kvarh_import_total"`
	KwhExportL1       string `json:"kwh_export_l1"`
	KwhExportL2       string `json:"kwh_export_l2"`
	KwhExportL3       string `json:"kwh_export_l3"`
	KwhExportTotal    string `json:"kwh_export_total" vmdata:"true"`
	KwhImportL1       string `json:"kwh_import_l1"`
	KwhImportL2       string `json:"kwh_import_l2"`
	KwhImportL3       string `json:"kwh_import_l3"`
	KwhImportTotal    string `json:"kwh_import_total"`
	MeterID           string `json:"meter_id"`
	PowerFactor       string `json:"power_factor"`
	ServerTime        string `json:"server_time"`
	Status            string `json:"status"`
	VoltageL1         string `json:"voltage_l1" vmdata:"true"`
	VoltageL2         string `json:"voltage_l2"`
	VoltageL3         string `json:"voltage_l3"`
}

type Data struct {
	CommType           string        `json:"comm_type"`
	DcuID              string        `json:"dcu_id"`
	ErrorCode          string        `json:"error_code"`
	ErrorMessage       string        `json:"error_message"`
	HesID              string        `json:"hes_id"`
	MeterID            string        `json:"meter_id"`
	MeterOperationType string        `json:"meter_operation_type"`
	MeterResponse      MeterResponse `json:"meter_response"`
	ReadDate           string        `json:"read_date"`
	RequestDate        string        `json:"request_date"`
	TrxID              string        `json:"trx_id"`
}
