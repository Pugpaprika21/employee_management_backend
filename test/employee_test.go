package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/Pugpaprika21/db"
	"github.com/Pugpaprika21/router/employee"
)

func TestEmployeeRouter(t *testing.T) {
	e := echo.New()
	dbConnection, err := db.New().UseDB()
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	employeeRouter := employee.NewEmployeeRouter()
	employeeRouter.Register(e, dbConnection)

	testCases := []struct {
		name       string
		method     string
		path       string
		statusCode int
	}{
		{"CreateEmployee", http.MethodPost, "/api/v1/createEmployee", http.StatusCreated},
		{"GetAllEmployee", http.MethodGet, "/api/v1/getAllEmployee", http.StatusOK},
		{"GetEmployeeById", http.MethodGet, "/api/v1/getEmployeeById/1", http.StatusOK},
		{"UpdateEmployeeById", http.MethodPut, "/api/v1/updateEmployeeById/1", http.StatusOK},
		{"DeleteEmployeeById", http.MethodDelete, "/api/v1/deleteEmployeeById/1", http.StatusOK},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
			assert.Equal(t, tc.statusCode, rec.Code)
		})
	}
}
