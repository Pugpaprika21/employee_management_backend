package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/Pugpaprika21/db"
	"github.com/Pugpaprika21/router/department"
)

func TestDepartmentRouter(t *testing.T) {
	e := echo.New()
	dbConnection, err := db.New().UseDB()
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	departmentRouter := department.NewDepartmentRouter()
	departmentRouter.Register(e, dbConnection)

	testCases := []struct {
		name       string
		method     string
		path       string
		statusCode int
	}{
		{"CreateDepartment", http.MethodPost, "/api/v1/createDepartment", http.StatusCreated},
		{"GetAllDepartment", http.MethodGet, "/api/v1/getAllDepartment", http.StatusOK},
		{"GetDepartmentById", http.MethodGet, "/api/v1/getDepartmentById/1", http.StatusOK},
		{"UpdateDepartmentById", http.MethodPut, "/api/v1/updateDepartmentById/1", http.StatusOK},
		{"DeleteDepartmentById", http.MethodDelete, "/api/v1/deleteDepartmentById/1", http.StatusOK},
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
