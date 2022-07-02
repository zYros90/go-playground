package examples

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type EmployeesRequest struct {
	EmployeeName   string `json:"name"`
	EmployeeSalary int    `json:"salary"`
	EmployeeAge    int    `json:"age"`
}

func PostExampleEmployees(url string) (map[string]interface{}, error) {
	employee := &EmployeesRequest{
		EmployeeName:   "felix",
		EmployeeSalary: 1000,
		EmployeeAge:    30,
	}

	data, err := json.Marshal(employee)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error in request, got status code %d, msg: %s", resp.StatusCode, string(data))
	}
	employeeResp := make(map[string]interface{})
	err = json.Unmarshal(data, &employeeResp)
	if err != nil {
		return nil, err
	}

	return employeeResp, nil
}
