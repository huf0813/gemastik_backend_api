package infrastructure

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

type DriverSupabase struct {
	AnonKey    string `json:"anon_key"`
	ServiceKey string `json:"service_key"`
	Address    string `json:"address"`
}

func NewDriverSupabase() (DriverSupabase, error) {
	if err := godotenv.Load(".env"); err != nil {
		return DriverSupabase{}, nil
	}

	return DriverSupabase{
		Address:    os.Getenv("SUPABASE_ADDRESS"),
		AnonKey:    os.Getenv("SUPABASE_ANON_KEY"),
		ServiceKey: os.Getenv("SUPABASE_SERVICE_KEY"),
	}, nil
}

func (d *DriverSupabase) RequestFormula(table, httpMethod string, jsonByte []byte) (*http.Request, error) {
	var request *http.Request
	url := fmt.Sprintf("%s/%s", d.Address, table)

	if httpMethod == http.MethodPost || httpMethod == http.MethodPatch {
		requestWithPayload, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(jsonByte))
		if err != nil {
			return nil, err
		}
		requestWithPayload.Header.Set("Content-Type", "application/json")
		requestWithPayload.Header.Set("Prefer", "return=representation")
		request = requestWithPayload
	} else {
		requestWithoutPayload, err := http.NewRequest(httpMethod, url, nil)
		if err != nil {
			return nil, err
		}
		request = requestWithoutPayload
	}

	request.Header.Set("apikey", d.ServiceKey)
	request.Header.Set("Authorization", fmt.Sprintf("%s %s", "Bearer", d.ServiceKey))

	return request, nil
}

func (d *DriverSupabase) ExecuteRequestFormula(request *http.Request) (*http.Response, error) {
	client := &http.Client{}
	result, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
