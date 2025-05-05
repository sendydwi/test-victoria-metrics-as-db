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
		"215220049495",
		"215220049491",
		"215220049244",
		"215220048709",
		"215220017925",
		"215220014822",
		"215220014723",
		"215220014710",
		"215220018443",
		"215220018409",
		"215220017902",
		"215220017901",
		"215220017900",
		"215220017899",
		"215220017898",
		"215220017897",
		"215220017848",
		"215220017847",
		"215220013399",
		"215220013400",
		"215220013398",
		"215220013384",
		"215220013383",
		"215220013381",
		"215220013379",
		"215220011640",
		"215220011639",
		"215220013179",
		"215220013178",
		"215220013177",
		"215220012029",
		"215220011782",
		"215220011779",
		"215220011770",
		"215220011769",
		"215220011767",
		"215220011731",
		"215220011709",
		"215220011708",
		"215220011707",
		"215220011706",
		"215220049833",
		"215220049132",
		"215220048990",
		"215220048927",
		"215220025213",
		"215220011670",
		"215220011669",
		"215220011662",
		"215220011661",
		"215220011660",
		"215220011659",
		"215220011658",
		"215220011657",
		"215220011644",
		"215220011643",
		"215221001485",
		"215221001483",
		"215220048688",
		"215220048584",
		"215220034774",
		"215220027020",
		"215220016185",
		"215220014519",
		"215220013583",
		"215220012737",
		"215220012736",
		"215220012733",
		"215220012732",
		"215220012191",
		"215220011514",
		"215220010713",
		"215220010710",
		"215220010709",
		"215220047797",
		"215220022035",
		"215220016679",
		"215220013352",
		"215220012441",
		"215220012351",
		"215220012349",
		"215220010275",
		"215220010266",
		"215220019497",
		"215220019222",
		"215220019218",
		"215220019207",
		"215220019205",
		"215220018766",
		"215220018764",
		"215220018501",
		"215220018498",
		"215220017843",
		"215220017737",
		"215220017506",
		"215220017505",
		"215220016300",
		"215220016299",
		"215220016298",
		"215220016297",
		"215220016296",
		"215220016295",
		"215220016262",
		"215220016259",
		"215220015835",
		"215220014276",
		"215220014256",
		"215220013401",
		"215220013361",
		"215220013340",
		"215220013309",
		"215220013308",
		"215220013307",
		"215220013276",
		"215220013275",
		"215220013274",
		"215220013270",
		"215220013268",
		"215220013266",
		"215220013265",
		"215220013258",
		"215220013257",
		"215220013256",
		"215220013255",
		"215220013254",
		"215220013253",
		"215220013196",
		"215220013182",
		"215220013181",
		"215220013180",
		"215220013176",
		"215220013175",
		"215220013156",
		"215220013155",
		"215220013154",
		"215220013153",
		"215220013152",
		"215220012988",
		"215220012986",
		"215220012984",
		"215220012868",
		"215220012514",
		"215220012510",
		"215220012366",
		"215220012358",
		"215220012357",
		"215220012356",
		"215220012355",
		"215220012354",
		"215220012353",
		"215220012089",
		"215220012034",
		"215220012032",
		"215220012031",
		"215220012030",
		"215220011705",
		"215220011686",
		"215220011685",
		"215220011684",
		"215220011683",
		"215220011682",
		"215220011681",
		"215220011679",
	}

	url := "http://165.22.242.191:4242/api/put"

	exampleData := getTemplateExamplePayload("loadprofile.json")
	exampleDataInstant := getTemplateExamplePayloadInstant("instant.json")
	exampleDataEob := getTemplateExamplePayloadEoB("eob.json")
	for _, v := range meter {
		unit := "51143"
		insert := "SUNRISE"
		loadProfileDummyData(*exampleData, url, v, unit, insert)
		loadInstantDummyData(*exampleDataInstant, url, v, unit, insert)
		loadEobDummyData(*exampleDataEob, url, v, unit, insert)
	}

	// meter List wasion
	meterWasion := []string{
		"311126000347",
		"111125039014",
		"111125039010",
		"311126000193",
		"111125033632",
		"111125033637",
		"111125039420",
		"311126000171",
		"111125034009",
		"111125036647",
		"111125036683",
		"111125033640",
		"111125037831",
		"111125036694",
		"111125036690",
	}

	exampleDataWasion := getTemplateExamplePayload("loadprofile.json")
	exampleDataInstantWasion := getTemplateExamplePayloadInstant("instant.json")
	exampleDataEobWasion := getTemplateExamplePayloadEoB("eob.json")
	for _, v := range meterWasion {
		unit := "53559"
		insert := "WASIONHES"
		loadProfileDummyData(*exampleDataWasion, url, v, unit, insert)
		loadInstantDummyData(*exampleDataInstantWasion, url, v, unit, insert)
		loadEobDummyData(*exampleDataEobWasion, url, v, unit, insert)
	}
}

func loadProfileDummyData(template Data, url, meterId, unitId, insertBy string) {
	baseDate := time.Date(2025, 2, 1, 0, 0, 0, 0, time.Local)

	operation := template.MeterOperationType
	status := template.ErrorMessage

	for d := 1; d <= 90; d += 3 {
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
							MeterID:  meterId,
							Status:   status,
							InsertBy: insertBy,
							UnitId:   unitId,
							TrxID:    trxId,
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
						MeterID:  meterId,
						Status:   status,
						InsertBy: insertBy,
						UnitId:   unitId,
						TrxID:    trxId,
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

func loadInstantDummyData(template DataInstant, url, meterId, unitId, insertBy string) {
	baseDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)

	operation := template.MeterOperationType
	status := template.ErrorMessage

	for d := 1; d <= 30; d++ {
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
						MeterID:  meterId,
						Status:   status,
						InsertBy: insertBy,
						UnitId:   unitId,
						TrxID:    trxId,
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
					MeterID:  meterId,
					Status:   status,
					InsertBy: insertBy,
					UnitId:   unitId,
					TrxID:    trxId,
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
			}
			if err == nil {
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

func loadEobDummyData(template DataEndOfBilling, url, meterId, unitId, insertBy string) {
	operation := template.MeterOperationType
	status := template.ErrorMessage

	for d := 4; d <= 5; d++ {
		day := time.Date(2025, time.Month(d), 5, 0, 0, 0, 0, time.Local)

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
					MeterID:  meterId,
					Status:   status,
					InsertBy: insertBy,
					UnitId:   unitId,
					TrxID:    trxId,
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
				MeterID:  meterId,
				Status:   status,
				InsertBy: insertBy,
				UnitId:   unitId,
				TrxID:    trxId,
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
				fmt.Println("Error creating request:", err)
				time.Sleep(1 * time.Second)
			}
			if err == nil {
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

func getTemplateExamplePayload(fileName string) *Data {
	data, err := os.ReadFile(fileName)
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
	MeterID  string `json:"meter_id"`
	Status   string `json:"status"`
	InsertBy string `json:"insert_by"`
	UnitId   string `json:"unit_id"`
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
	ActivePowerExport string `json:"active_power_export" vmdata:"true"`
	ActivePowerImport string `json:"active_power_import" vmdata:"true"`
	AlarmRegister     string `json:"alarm_register`
	Clock             string `json:"clock`
	CurrentL1         string `json:"current_l1" vmdata:"true"`
	CurrentL2         string `json:"current_l2" vmdata:"true"`
	CurrentL3         string `json:"current_l3" vmdata:"true"`
	InsertBy          string `json:"insert_by`
	KvarhBilling      string `json:"kvarh_billing" vmdata:"true"`
	KvarhExportTotal  string `json:"kvarh_export_total" vmdata:"true"`
	KvarhImportTotal  string `json:"kvarh_import_total" vmdata:"true"`
	KwhExportL1       string `json:"kwh_export_l1" vmdata:"true"`
	KwhExportL2       string `json:"kwh_export_l2" vmdata:"true"`
	KwhExportL3       string `json:"kwh_export_l3" vmdata:"true"`
	KwhExportTotal    string `json:"kwh_export_total" vmdata:"true"`
	KwhImportL1       string `json:"kwh_import_l1" vmdata:"true"`
	KwhImportL2       string `json:"kwh_import_l2" vmdata:"true"`
	KwhImportL3       string `json:"kwh_import_l3" vmdata:"true"`
	KwhImportTotal    string `json:"kwh_import_total" vmdata:"true"`
	MeterID           string `json:"meter_id"`
	PowerFactor       string `json:"power_factor"`
	ServerTime        string `json:"server_time"`
	Status            string `json:"status"`
	VoltageL1         string `json:"voltage_l1" vmdata:"true"`
	VoltageL2         string `json:"voltage_l2" vmdata:"true"`
	VoltageL3         string `json:"voltage_l3" vmdata:"true"`
}

type Data struct {
	CommType           string        `json:"comm_type"`
	DcuID              string        `json:"dcu_id"`
	ErrorCode          int32         `json:"error_code"`
	ErrorMessage       string        `json:"error_message"`
	HesID              string        `json:"hes_id"`
	MeterID            string        `json:"meter_id"`
	MeterOperationType string        `json:"meter_operation_type"`
	MeterResponse      MeterResponse `json:"meter_response"`
	ReadDate           string        `json:"read_date"`
	RequestDate        string        `json:"request_date"`
	TrxID              string        `json:"trx_id"`
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
}

type Instant struct {
	InsertBy         string `json:"insert_by"`
	ServerTime       string `json:"server_time"`
	ActivePowerTotal string `json:"active_power_total" vmdata:"true"`
	CommType         string `json:"comm_type"`
	CurrentL1        string `json:"current_l1" vmdata:"true"`
	KvarhExportTotal string `json:"kvarh_export_total" vmdata:"true"`
	KvarhImportTotal string `json:"kvarh_import_total" vmdata:"true"`
	KwhExportTotal   string `json:"kwh_export_total" vmdata:"true"`
	KwhImportTotal   string `json:"kwh_import_total" vmdata:"true"`
	LqiPlc           string `json:"lqi_plc"`
	LqiRf            string `json:"lqi_rf"`
}

type EndOfBilling struct {
	BatteryCapacity  string `json:"battery_capacity" vmdata:"true"`
	InsertBy         string `json:"insert_by"`
	KvaMaxR1         string `json:"kva_max_r1" vmdata:"true"`
	KvaMaxR2         string `json:"kva_max_r2" vmdata:"true"`
	KvaMaxR3         string `json:"kva_max_r3" vmdata:"true"`
	KvaMaxTotal      string `json:"kva_max_total" vmdata:"true"`
	KvarhAbsR1       string `json:"kvarh_abs_r1" vmdata:"true"`
	KvarhAbsR2       string `json:"kvarh_abs_r2" vmdata:"true"`
	KvarhAbsR3       string `json:"kvarh_abs_r3" vmdata:"true"`
	KvarhAbsTotal    string `json:"kvarh_abs_total" vmdata:"true"`
	KvarhExportL1    string `json:"kvarh_export_l1" vmdata:"true"`
	KvarhExportL2    string `json:"kvarh_export_l2" vmdata:"true"`
	KvarhExportL3    string `json:"kvarh_export_l3" vmdata:"true"`
	KvarhExportTotal string `json:"kvarh_export_total" vmdata:"true"`
	KvarhImportL1    string `json:"kvarh_import_l1" vmdata:"true"`
	KvarhImportL2    string `json:"kvarh_import_l2" vmdata:"true"`
	KvarhImportL3    string `json:"kvarh_import_l3" vmdata:"true"`
	KvarhImportTotal string `json:"kvarh_import_total" vmdata:"true"`
	KwhAbsR1         string `json:"kwh_abs_r1" vmdata:"true"`
	KwhAbsR2         string `json:"kwh_abs_r2" vmdata:"true"`
	KwhAbsR3         string `json:"kwh_abs_r3" vmdata:"true"`
	KwhAbsTotal      string `json:"kwh_abs_total" vmdata:"true"`
	KwhExportL1      string `json:"kwh_export_l1" vmdata:"true"`
	KwhExportL2      string `json:"kwh_export_l2" vmdata:"true"`
	KwhExportL3      string `json:"kwh_export_l3" vmdata:"true"`
	KwhExportTotal   string `json:"kwh_export_total" vmdata:"true"`
	KwhImportL1      string `json:"kwh_import_l1" vmdata:"true"`
	KwhImportL2      string `json:"kwh_import_l2" vmdata:"true"`
	KwhImportL3      string `json:"kwh_import_l3" vmdata:"true"`
	KwhImportTotal   string `json:"kwh_import_total" vmdata:"true"`
	MeterCode        string `json:"meter_code"`

	PowerOffTotal string `json:"power_off_total"`
	TamperTotal   string `json:"tamper_total"`
	ReadTime      string `json:"read_date"`
	ServerTime    string `json:"server_time"`

	TimeKvaR1    string `json:"time_kva_r1"`
	TimeKvaR2    string `json:"time_kva_r2"`
	TimeKvaR3    string `json:"time_kva_r3"`
	TimeKvaTotal string `json:"time_kva_total"`
}
