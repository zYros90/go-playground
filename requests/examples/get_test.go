package examples

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetExampleEmployees(t *testing.T) {
	resp := &EmployeesResponse{
		Status:  "success",
		Message: "abc",
		Data: []EmployeeData{
			{
				ID:             1,
				EmployeeName:   "felix",
				EmployeeSalary: 500,
				EmployeeAge:    31,
				ProfileImage:   "",
			},
		},
	}
	pattern := "/api/employees"
	srv := HttpMock(pattern, http.StatusOK, resp)
	defer srv.Close()
	url := srv.URL + pattern

	patternErr := "/api/employees/err"
	srvErr := HttpMock(patternErr, http.StatusBadRequest, resp)
	defer srvErr.Close()
	urlErr := srvErr.URL + patternErr

	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    *EmployeesResponse
		wantErr bool
	}{
		{
			"get 200",
			args{url: url},
			resp,
			false,
		},
		{
			"get 400",
			args{url: urlErr},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetExampleEmployees(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetExampleEmployees() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExampleEmployees() = %v, want %v", got, tt.want)
			}
		})
	}
}
