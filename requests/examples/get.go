package examples

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type EmployeesResponse struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    []EmployeeData `json:"data"`
}

type EmployeeData struct {
	ID             int    `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary int    `json:"employee_salary"`
	EmployeeAge    int    `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}

func GetExampleEmployees(url string) (*EmployeesResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error in request, got status code %d, msg: %s", resp.StatusCode, string(data))
	}
	employees := new(EmployeesResponse)
	err = json.Unmarshal(data, employees)
	if err != nil {
		return nil, err
	}

	return employees, nil
}
