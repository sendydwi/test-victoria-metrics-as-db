package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"golang.org/x/exp/rand"
)

func main() {

	// meter List Sunrise
	meter := []string{
		"215220049582",
		"215220049558",
		"215220049497",
		// "215220049495",
		// "215220049491",
		// "215220049244",
		// "215220048709",
		// "215220017925",
		// "215220014822",
		// "215220014723",
		// "215220014710",
		// "215220018443",
		// "215220018409",
		// "215220017902",
		// "215220017901",
		// "215220017900",
		// "215220017899",
		// "215220017898",
		// "215220017897",
		// "215220017848",
		// "215220017847",
		// "215220013399",
		// "215220013400",
		// "215220013398",
		// "215220013384",
		// "215220013383",
		// "215220013381",
		// "215220013379",
		// "215220011640",
		// "215220011639",
		// "215220013179",
		// "215220013178",
		// "215220013177",
		// "215220012029",
		// "215220011782",
		// "215220011779",
		// "215220011770",
		// "215220011769",
		// "215220011767",
		// "215220011731",
		// "215220011709",
		// "215220011708",
		// "215220011707",
		// "215220011706",
		// "215220049833",
		// "215220049132",
		// "215220048990",
		// "215220048927",
		// "215220025213",
		// "215220011670",
		// "215220011669",
		// "215220011662",
		// "215220011661",
		// "215220011660",
		// "215220011659",
		// "215220011658",
		// "215220011657",
		// "215220011644",
		// "215220011643",
		// "215221001485",
		// "215221001483",
		// "215220048688",
		// "215220048584",
		// "215220034774",
		// "215220027020",
		// "215220016185",
		// "215220014519",
		// "215220013583",
		// "215220012737",
		// "215220012736",
		// "215220012733",
		// "215220012732",
		// "215220012191",
		// "215220011514",
		// "215220010713",
		// "215220010710",
		// "215220010709",
		// "215220047797",
		// "215220022035",
		// "215220016679",
		// "215220013352",
		// "215220012441",
		// "215220012351",
		// "215220012349",
		// "215220010275",
		// "215220010266",
		// "215220019497",
		// "215220019222",
		// "215220019218",
		// "215220019207",
		// "215220019205",
		// "215220018766",
		// "215220018764",
		// "215220018501",
		// "215220018498",
		// "215220017843",
		// "215220017737",
		// "215220017506",
		// "215220017505",
		// "215220016300",
		// "215220016299",
		// "215220016298",
		// "215220016297",
		// "215220016296",
		// "215220016295",
		// "215220016262",
		// "215220016259",
		// "215220015835",
		// "215220014276",
		// "215220014256",
		// "215220013401",
		// "215220013361",
		// "215220013340",
		// "215220013309",
		// "215220013308",
		// "215220013307",
		// "215220013276",
		// "215220013275",
		// "215220013274",
		// "215220013270",
		// "215220013268",
		// "215220013266",
		// "215220013265",
		// "215220013258",
		// "215220013257",
		// "215220013256",
		// "215220013255",
		// "215220013254",
		// "215220013253",
		// "215220013196",
		// "215220013182",
		// "215220013181",
		// "215220013180",
		// "215220013176",
		// "215220013175",
		// "215220013156",
		// "215220013155",
		// "215220013154",
		// "215220013153",
		// "215220013152",
		// "215220012988",
		// "215220012986",
		// "215220012984",
		// "215220012868",
		// "215220012514",
		// "215220012510",
		// "215220012366",
		// "215220012358",
		// "215220012357",
		// "215220012356",
		// "215220012355",
		// "215220012354",
		// "215220012353",
		// "215220012089",
		// "215220012034",
		// "215220012032",
		// "215220012031",
		// "215220012030",
		// "215220011705",
		// "215220011686",
		// "215220011685",
		// "215220011684",
		// "215220011683",
		// "215220011682",
		// "215220011681",
		// "215220011679",
	}

	url := "http://165.22.242.191:4242/api/put"
	exampleData := getTemplateExamplePayload("loadprofile.json")
	exampleDataInstant := getTemplateExamplePayloadInstant("instant.json")
	exampleDataEob := getTemplateExamplePayloadEoB("eob.json")
	for _, v := range meter {
		unit := "51143"
		insert := "SNR"
		brandId := "Sunrise"
		projectId := "PROJECT-1"
		loadProfileDummyData(*exampleData, url, v, unit, insert, projectId, brandId)
		loadInstantDummyData(*exampleDataInstant, url, v, unit, insert, projectId, brandId)
		loadEobDummyData(*exampleDataEob, url, v, unit, insert, projectId, brandId)
	}

	// meter List wasion
	meterWasion := []string{
		"311126000347",
		"111125039014",
		// "111125039010",
		// "311126000193",
		// "111125033632",
		// "111125033637",
		// "111125039420",
		// "311126000171",
		// "111125034009",
		// "111125036647",
		// "111125036683",
		// "111125033640",
		// "111125037831",
		// "111125036694",
		// "111125036690",
	}

	exampleDataWasion := getTemplateExamplePayload("loadprofile.json")
	exampleDataInstantWasion := getTemplateExamplePayloadInstant("instant.json")
	exampleDataEobWasion := getTemplateExamplePayloadEoB("eob.json")
	for _, v := range meterWasion {
		unit := "53559"
		insert := "WSN"
		brandId := "Wasion"
		projectId := "PROJECT-1"
		loadProfileDummyData(*exampleDataWasion, url, v, unit, insert, projectId, brandId)
		loadInstantDummyData(*exampleDataInstantWasion, url, v, unit, insert, projectId, brandId)
		loadEobDummyData(*exampleDataEobWasion, url, v, unit, insert, projectId, brandId)
	}
}

func loadProfileDummyData(template DataLoadProfile, url, meterId, unitId, insertBy, projectID, brandID string) {
	baseDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Now().UTC().Location())

	operation := template.MeterOperationType
	status := template.ErrorMessage

	for d := 0; d < 42; d += 3 {
		data := []ValidPayload{}
		for x := 0; x <= 2; x++ {
			day := baseDate.Add(time.Duration(d+x) * 24 * time.Hour)
			nod := 96
			increment := 1440 / nod
			for i := 1; i <= nod; i++ {

				dataDate := day.Add(time.Duration(increment) * time.Duration(i) * time.Minute)
				trxId := fmt.Sprintf("%s-%d", template.TrxID, dataDate.UnixMilli())

				dataType := reflect.TypeOf(template.MeterResponse)
				valueof := reflect.ValueOf(template.MeterResponse)
				if dataType.Kind() == reflect.Ptr {
					dataType = dataType.Elem()
				}
				for x := 0; x < dataType.NumField(); x++ {
					field := dataType.Field(x)
					jsonTag := field.Tag.Get("json")

					vmenabled := field.Tag.Get("vmdata")
					if vmenabled != "true" {
						continue
					}

					if valueof.Field(x).String() == "" {
						continue
					}
					value, err := strconv.ParseFloat(valueof.Field(x).String(), 64)
					if err != nil {
						fmt.Println(err.Error())
						continue
					}

					metricName := fmt.Sprintf("dummy_1_periodic_%s_%s", operation, jsonTag)
					format := ValidPayload{
						Metric:    metricName,
						Timestamp: dataDate.UnixMilli(),
						Value:     value * randomizeWithin50Percent(64),
						Tags: Tags{
							MeterID:   meterId,
							Status:    status,
							HesID:     insertBy,
							UnitId:    unitId,
							TrxID:     trxId,
							ProjectID: projectID,
							BrandID:   brandID,
						},
					}

					data = append(data, format)
				}

				metricName := fmt.Sprintf("dummy_1_periodic_%s_success", operation)
				format := ValidPayload{
					Metric:    metricName,
					Timestamp: dataDate.UnixMilli(),
					Value:     1,
					Tags: Tags{
						MeterID:   meterId,
						Status:    status,
						HesID:     insertBy,
						UnitId:    unitId,
						TrxID:     trxId,
						ProjectID: projectID,
						BrandID:   brandID,
					},
				}

				data = append(data, format)
			}
			fmt.Println("Work on ", day.String(), meterId)
		}

		time.Sleep(250 * time.Millisecond)

		payload, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
		}

		var req *http.Request
		for r := 0; r < 5; r++ {
			req, err = http.NewRequest("POST", url, bytes.NewBuffer(payload))
			if err != nil {
				fmt.Println("Error creating request:", err)
				time.Sleep(1 * time.Second)
			} else {
				continue
			}
		}

		if req == nil {
			fmt.Println("Error creating request")
			continue
		}

		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
		}
		defer resp.Body.Close()

		// Print response status
		fmt.Println("Response status:", resp.Status)
	}
}

func loadInstantDummyData(template DataInstant, url, meterId, unitId, insertBy, projectID, brandID string) {
	baseDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Now().UTC().Location())

	operation := template.MeterOperationType
	status := template.ErrorMessage

	for d := 0; d < 42; d++ {
		day := baseDate.Add(time.Duration(d) * 24 * time.Hour)
		nod := 2
		increment := (24 * 60) / nod

		data := []ValidPayload{}
		for i := 1; i <= nod; i++ {
			dataDate := day.Add(time.Duration(increment) * time.Duration(i) * time.Minute)
			trxId := fmt.Sprintf("%s-%d", template.TrxID, dataDate.UnixMilli())

			dataType := reflect.TypeOf(template.MeterResponse)
			valueof := reflect.ValueOf(template.MeterResponse)
			if dataType.Kind() == reflect.Ptr {
				dataType = dataType.Elem()
			}
			for x := 0; x < dataType.NumField(); x++ {
				field := dataType.Field(x)
				jsonTag := field.Tag.Get("json")

				vmenabled := field.Tag.Get("vmdata")
				if vmenabled != "true" {
					continue
				}

				if valueof.Field(x).String() == "" {
					continue
				}
				value, err := strconv.ParseFloat(valueof.Field(x).String(), 64)
				if err != nil {
					fmt.Println(err.Error())
					continue
				}

				metricName := fmt.Sprintf("dummy_1_periodic_%s_%s", operation, jsonTag)
				format := ValidPayload{
					Metric:    metricName,
					Timestamp: dataDate.UnixMilli(),
					Value:     value * randomizeWithin50Percent(64),
					Tags: Tags{
						MeterID:   meterId,
						Status:    status,
						HesID:     insertBy,
						UnitId:    unitId,
						TrxID:     trxId,
						ProjectID: projectID,
						BrandID:   brandID,
					},
				}

				data = append(data, format)
			}

			metricName := fmt.Sprintf("dummy_1_periodic_%s_success", operation)
			format := ValidPayload{
				Metric:    metricName,
				Timestamp: dataDate.UnixMilli(),
				Value:     1,
				Tags: Tags{
					MeterID:   meterId,
					Status:    status,
					HesID:     insertBy,
					UnitId:    unitId,
					TrxID:     trxId,
					ProjectID: projectID,
					BrandID:   brandID,
				},
			}

			data = append(data, format)
		}

		time.Sleep(250 * time.Millisecond)

		payload, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
		}

		var req *http.Request
		for r := 0; r < 5; r++ {
			req, err = http.NewRequest("POST", url, bytes.NewBuffer(payload))
			if err != nil {
				fmt.Println("Error creating request:", err)
				time.Sleep(1 * time.Second)
			} else {
				continue
			}
		}

		if req == nil {
			fmt.Println("Error creating request")
			continue
		}

		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
		}
		defer resp.Body.Close()

		// Print response status
		fmt.Println("Response status:", resp.Status, day.String(), meterId)
	}
}

func loadEobDummyData(template DataEndOfBilling, url, meterId, unitId, insertBy, projectID, brandID string) {
	operation := template.MeterOperationType
	status := template.ErrorMessage

	for d := 4; d <= 5; d++ {
		day := time.Date(2025, time.Month(d), 1, 0, 0, 0, 0, time.Now().UTC().Location())

		data := []ValidPayload{}
		trxId := fmt.Sprintf("%s-%d", template.TrxID, day.UnixMilli())

		dataType := reflect.TypeOf(template.MeterResponse)
		valueof := reflect.ValueOf(template.MeterResponse)
		if dataType.Kind() == reflect.Ptr {
			dataType = dataType.Elem()
		}
		for x := 0; x < dataType.NumField(); x++ {
			field := dataType.Field(x)
			jsonTag := field.Tag.Get("json")

			vmenabled := field.Tag.Get("vmdata")
			if vmenabled != "true" {
				continue
			}

			if valueof.Field(x).String() == "" {
				continue
			}
			value, err := strconv.ParseFloat(valueof.Field(x).String(), 64)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			metricName := fmt.Sprintf("dummy_1_periodic_%s_%s", operation, jsonTag)
			format := ValidPayload{
				Metric:    metricName,
				Timestamp: day.UnixMilli(),
				Value:     value * randomizeWithin50Percent(64),
				Tags: Tags{
					MeterID:   meterId,
					Status:    status,
					HesID:     insertBy,
					UnitId:    unitId,
					TrxID:     trxId,
					ProjectID: projectID,
					BrandID:   brandID,
				},
			}

			data = append(data, format)
		}

		metricName := fmt.Sprintf("dummy_1_periodic_%s_success", operation)
		format := ValidPayload{
			Metric:    metricName,
			Timestamp: day.UnixMilli(),
			Value:     1,
			Tags: Tags{
				MeterID:   meterId,
				Status:    status,
				HesID:     insertBy,
				UnitId:    unitId,
				TrxID:     trxId,
				ProjectID: projectID,
				BrandID:   brandID,
			},
		}

		data = append(data, format)

		time.Sleep(250 * time.Millisecond)

		payload, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
		}

		var req *http.Request
		for r := 0; r < 5; r++ {
			req, err = http.NewRequest("POST", url, bytes.NewBuffer(payload))
			if err != nil {
				fmt.Println("Error creating request:", err.Error())
				time.Sleep(1 * time.Second)
			} else {
				continue
			}
		}

		if req == nil {
			fmt.Println("Error creating request")
			continue
		}

		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
		}
		defer resp.Body.Close()

		// Print response status
		fmt.Println("Response status:", resp.Status, day.String(), meterId)
	}
}

func randomizeWithin50Percent(base float64) float64 {
	rand.Seed(uint64(time.Now().UnixNano())) // seed once per program
	factor := 1 + (rand.Float64()*1.0 - 0.5) // generates between 0.5 and 1.5
	return base * factor
}

func getTemplateExamplePayload(fileName string) *DataLoadProfile {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return nil
	}

	var examplePayload DataLoadProfile

	err = json.Unmarshal(data, &examplePayload)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v", err)
		return nil
	}

	return &examplePayload
}

func getTemplateExamplePayloadInstant(fileName string) *DataInstant {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return nil
	}

	var examplePayload DataInstant

	err = json.Unmarshal(data, &examplePayload)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v", err)
		return nil
	}

	return &examplePayload
}

func getTemplateExamplePayloadEoB(fileName string) *DataEndOfBilling {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return nil
	}

	var examplePayload DataEndOfBilling

	err = json.Unmarshal(data, &examplePayload)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v", err)
		return nil
	}

	return &examplePayload
}

type ValidPayload struct {
	Metric    string  `json:"metric"`
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
	Tags      Tags    `json:"tags"`
}

type Tags struct {
	MeterID   string `json:"meter_id"`
	Status    string `json:"status"`
	HesID     string `json:"hes_id"`
	BrandID   string `json:"brand_id"`
	UnitId    string `json:"unit_id"`
	TrxID     string `json:"trx_id"`
	ProjectID string `json:"project_id"`
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

type LoadProfile struct {
	MeterId                string    `json:"meter_id"`
	Clock                  time.Time `json:"clock"` // Use time.Time if it's an ISO date
	Status                 string    `json:"status"`
	AlarmRegister          string    `json:"alarm_register"`
	VoltageL1              string    `json:"voltage_l1" vmdata:"true"`
	VoltageL2              string    `json:"voltage_l2" vmdata:"true"`
	VoltageL3              string    `json:"voltage_l3" vmdata:"true"`
	CurrentL1              string    `json:"current_l1" vmdata:"true"`
	CurrentL2              string    `json:"current_l2" vmdata:"true"`
	CurrentL3              string    `json:"current_l3" vmdata:"true"`
	WhExportL1             string    `json:"wh_export_l1" vmdata:"true"`
	WhExportL2             string    `json:"wh_export_l2" vmdata:"true"`
	WhExportL3             string    `json:"wh_export_l3" vmdata:"true"`
	WhExportTotal          string    `json:"wh_export_total" vmdata:"true"`
	WhImportL1             string    `json:"wh_import_l1" vmdata:"true"`
	WhImportL2             string    `json:"wh_import_l2" vmdata:"true"`
	WhImportL3             string    `json:"wh_import_l3" vmdata:"true"`
	WhImportTotal          string    `json:"wh_import_total" vmdata:"true"`
	ActivePowerExportTotal string    `json:"active_power_export_total" vmdata:"true"`
	ActivePowerImportTotal string    `json:"active_power_import_total" vmdata:"true"`
	VarhExportTotal        string    `json:"varh_export_total" vmdata:"true"`
	VarhImportTotal        string    `json:"varh_import_total" vmdata:"true"`
	VarhTotal              string    `json:"varh_total" vmdata:"true"`
	PowerFactor            string    `json:"power_factor" vmdata:"true"`
}

type DataLoadProfile struct {
	CommType           string      `json:"comm_type"`
	DcuID              string      `json:"dcu_id"`
	ErrorCode          int32       `json:"error_code"`
	ErrorMessage       string      `json:"error_message"`
	HesID              string      `json:"hes_id"`
	MeterID            string      `json:"meter_id"`
	MeterOperationType string      `json:"meter_operation_type"`
	MeterResponse      LoadProfile `json:"meter_response"`
	ReadDate           string      `json:"read_date"`
	RequestDate        string      `json:"request_date"`
	TrxID              string      `json:"trx_id"`
	ProjectId          string      `json:"project_id"`
}

type DataInstant struct {
	CommType           string  `json:"comm_type"`
	DcuID              string  `json:"dcu_id"`
	ErrorCode          int32   `json:"error_code"`
	ErrorMessage       string  `json:"error_message"`
	HesID              string  `json:"hes_id"`
	MeterID            string  `json:"meter_id"`
	MeterOperationType string  `json:"meter_operation_type"`
	MeterResponse      Instant `json:"meter_response"`
	ReadDate           string  `json:"read_date"`
	RequestDate        string  `json:"request_date"`
	TrxID              string  `json:"trx_id"`
}

type DataEndOfBilling struct {
	CommType           string       `json:"comm_type"`
	DcuID              string       `json:"dcu_id"`
	ErrorCode          int32        `json:"error_code"`
	ErrorMessage       string       `json:"error_message"`
	HesID              string       `json:"hes_id"`
	MeterID            string       `json:"meter_id"`
	MeterOperationType string       `json:"meter_operation_type"`
	MeterResponse      EndOfBilling `json:"meter_response"`
	ReadDate           string       `json:"read_date"`
	RequestDate        string       `json:"request_date"`
	TrxID              string       `json:"trx_id"`
	ProjectId          string       `json:"project_id"`
}

type Instant struct {
	MeterId                string    `json:"meter_id"`
	Clock                  time.Time `json:"clock"` // assuming ISO8601 format
	BatteryCapacity        string    `json:"battery_cappacity" vmdata:"true"`
	Frequency              string    `json:"frequency" vmdata:"true"`
	VoltageL1              string    `json:"voltage_l1" vmdata:"true"`
	VoltageL2              string    `json:"voltage_l2" vmdata:"true"`
	VoltageL3              string    `json:"voltage_l3" vmdata:"true"`
	CurrentL1              string    `json:"current_l1" vmdata:"true"`
	CurrentL2              string    `json:"current_l2" vmdata:"true"`
	CurrentL3              string    `json:"current_l3" vmdata:"true"`
	CurrentN               string    `json:"current_n" vmdata:"true"`
	WhExportL1             string    `json:"wh_export_l1" vmdata:"true"`
	WhExportL2             string    `json:"wh_export_l2" vmdata:"true"`
	WhExportL3             string    `json:"wh_export_l3" vmdata:"true"`
	WhExportTotal          string    `json:"wh_export_total" vmdata:"true"`
	WhImportL1             string    `json:"wh_import_l1" vmdata:"true"`
	WhImportL2             string    `json:"wh_import_l2" vmdata:"true"`
	WhImportL3             string    `json:"wh_import_l3" vmdata:"true"`
	WhImportTotal          string    `json:"wh_import_total" vmdata:"true"`
	ActivePowerExportL1    string    `json:"active_power_export_l1" vmdata:"true"`
	ActivePowerExportL2    string    `json:"active_power_export_l2" vmdata:"true"`
	ActivePowerExportL3    string    `json:"active_power_export_l3" vmdata:"true"`
	ActivePowerExportTotal string    `json:"active_power_export_total" vmdata:"true"`
	ActivePowerImportL1    string    `json:"active_power_import_l1" vmdata:"true"`
	ActivePowerImportL2    string    `json:"active_power_import_l2" vmdata:"true"`
	ActivePowerImportL3    string    `json:"active_power_import_l3" vmdata:"true"`
	ActivePowerImportTotal string    `json:"active_power_import_total" vmdata:"true"`
	PowerFactor            string    `json:"power_factor" vmdata:"true"`
	PowerFactorL1          string    `json:"power_factor_l1" vmdata:"true"`
	PowerFactorL2          string    `json:"power_factor_l2" vmdata:"true"`
	PowerFactorL3          string    `json:"power_factor_l3" vmdata:"true"`
	ReactivePowerExport    string    `json:"reactive_power_export" vmdata:"true"`
	ReactivePowerExportL1  string    `json:"reactive_power_export_l1" vmdata:"true"`
	ReactivePowerExportL2  string    `json:"reactive_power_export_l2" vmdata:"true"`
	ReactivePowerExportL3  string    `json:"reactive_power_export_l3" vmdata:"true"`
	ReactivePowerImport    string    `json:"reactive_power_import" vmdata:"true"`
	ReactivePowerImportL1  string    `json:"reactive_power_import_l1" vmdata:"true"`
	ReactivePowerImportL2  string    `json:"reactive_power_import_l2" vmdata:"true"`
	ReactivePowerImportL3  string    `json:"reactive_power_import_l3" vmdata:"true"`
	VaExportL1             string    `json:"va_export_l1" vmdata:"true"`
	VaExportL2             string    `json:"va_export_l2" vmdata:"true"`
	VaExportL3             string    `json:"va_export_l3" vmdata:"true"`
	VaExportTotal          string    `json:"va_export_total" vmdata:"true"`
	VaImportL1             string    `json:"va_import_l1" vmdata:"true"`
	VaImportL2             string    `json:"va_import_l2" vmdata:"true"`
	VaImportL3             string    `json:"va_import_l3" vmdata:"true"`
	VaImportTotal          string    `json:"va_import_total" vmdata:"true"`
	VarhBilling            string    `json:"varh_billing" vmdata:"true"`
	VarhExportL1           string    `json:"varh_export_l1" vmdata:"true"`
	VarhExportL2           string    `json:"varh_export_l2" vmdata:"true"`
	VarhExportL3           string    `json:"varh_export_l3" vmdata:"true"`
	VarhExportTotal        string    `json:"varh_export_total" vmdata:"true"`
	VarhImportL1           string    `json:"varh_import_l1" vmdata:"true"`
	VarhImportL2           string    `json:"varh_import_l2" vmdata:"true"`
	VarhImportL3           string    `json:"varh_import_l3" vmdata:"true"`
	VarhImportTotal        string    `json:"varh_import_total" vmdata:"true"`
	TddCurrent             string    `json:"tdd_current" vmdata:"true"`
	ThdCurrent             string    `json:"thd_current" vmdata:"true"`
	VoltageAngleL1         string    `json:"voltage_angle_l1" vmdata:"true"`
	VoltageAngleL2         string    `json:"voltage_angle_l2" vmdata:"true"`
	VoltageAngleL3         string    `json:"voltage_angle_l3" vmdata:"true"`
	CurrentAngleL1         string    `json:"current_angle_l1" vmdata:"true"`
	CurrentAngleL2         string    `json:"current_angle_l2" vmdata:"true"`
	CurrentAngleL3         string    `json:"current_angle_l3" vmdata:"true"`
}

type EndOfBilling struct {
	MeterId         string    `json:"meter_id"`
	Clock           time.Time `json:"clock"`
	BatteryCapacity string    `json:"battery_cappacity" vmdata:"true"`
	MeterOffTotal   string    `json:"meter_off_total" vmdata:"true"`
	TamperTotal     string    `json:"tamper_total" vmdata:"true"`

	WhExportR1    string `json:"wh_export_r1" vmdata:"true"`
	WhExportR2    string `json:"wh_export_r2" vmdata:"true"`
	WhExportR3    string `json:"wh_export_r3" vmdata:"true"`
	WhExportR4    string `json:"wh_export_r4" vmdata:"true"`
	WhExportR5    string `json:"wh_export_r5" vmdata:"true"`
	WhExportTotal string `json:"wh_export_total" vmdata:"true"`

	WhImportR1    string `json:"wh_import_r1" vmdata:"true"`
	WhImportR2    string `json:"wh_import_r2" vmdata:"true"`
	WhImportR3    string `json:"wh_import_r3" vmdata:"true"`
	WhImportR4    string `json:"wh_import_r4" vmdata:"true"`
	WhImportR5    string `json:"wh_import_r5" vmdata:"true"`
	WhImportTotal string `json:"wh_import_total" vmdata:"true"`

	VaMaxR1    string `json:"va_max_r1" vmdata:"true"`
	VaMaxR2    string `json:"va_max_r2" vmdata:"true"`
	VaMaxR3    string `json:"va_max_r3" vmdata:"true"`
	VaMaxR4    string `json:"va_max_r4" vmdata:"true"`
	VaMaxR5    string `json:"va_max_r5" vmdata:"true"`
	VaMaxTotal string `json:"va_max_total" vmdata:"true"`

	TimeVaMax   time.Time `json:"time_va_max"`
	TimeVaMaxR1 time.Time `json:"time_va_max_r1"`
	TimeVaMaxR2 time.Time `json:"time_va_max_r2"`
	TimeVaMaxR3 time.Time `json:"time_va_max_r3"`
	TimeVaMaxR4 time.Time `json:"time_va_max_r4"`
	TimeVaMaxR5 time.Time `json:"time_va_max_r5"`

	VahExportR1    string `json:"vah_export_r1" vmdata:"true"`
	VahExportR2    string `json:"vah_export_r2" vmdata:"true"`
	VahExportR3    string `json:"vah_export_r3" vmdata:"true"`
	VahExportR4    string `json:"vah_export_r4" vmdata:"true"`
	VahExportR5    string `json:"vah_export_r5" vmdata:"true"`
	VahExportTotal string `json:"vah_export_total" vmdata:"true"`

	VahImportR1    string `json:"vah_import_r1" vmdata:"true"`
	VahImportR2    string `json:"vah_import_r2" vmdata:"true"`
	VahImportR3    string `json:"vah_import_r3" vmdata:"true"`
	VahImportR4    string `json:"vah_import_r4" vmdata:"true"`
	VahImportR5    string `json:"vah_import_r5" vmdata:"true"`
	VahImportTotal string `json:"vah_import_total" vmdata:"true"`

	VarhExportR1    string `json:"varh_export_r1" vmdata:"true"`
	VarhExportR2    string `json:"varh_export_r2" vmdata:"true"`
	VarhExportR3    string `json:"varh_export_r3" vmdata:"true"`
	VarhExportR4    string `json:"varh_export_r4" vmdata:"true"`
	VarhExportR5    string `json:"varh_export_r5" vmdata:"true"`
	VarhExportTotal string `json:"varh_export_total" vmdata:"true"`

	VarhImportR1    string `json:"varh_import_r1" vmdata:"true"`
	VarhImportR2    string `json:"varh_import_r2" vmdata:"true"`
	VarhImportR3    string `json:"varh_import_r3" vmdata:"true"`
	VarhImportR4    string `json:"varh_import_r4" vmdata:"true"`
	VarhImportR5    string `json:"varh_import_r5" vmdata:"true"`
	VarhImportTotal string `json:"varh_import_total" vmdata:"true"`

	VarhTotal string `json:"varh_total" vmdata:"true"`
}
