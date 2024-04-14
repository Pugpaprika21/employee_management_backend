package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/Pugpaprika21/db"
	"github.com/Pugpaprika21/router/position"
)

func TestPositionRouter(t *testing.T) {
	e := echo.New()
	dbConnection, err := db.New().UseDB()
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	positionRouter := position.NewPositionRouter()
	positionRouter.Register(e, dbConnection)

	testCases := []struct {
		name       string
		method     string
		path       string
		statusCode int
	}{
		{"CreatePosition", http.MethodPost, "/api/v1/createPosition", http.StatusCreated},
		{"GetAllPosition", http.MethodGet, "/api/v1/getAllPosition", http.StatusOK},
		{"GetPositionById", http.MethodGet, "/api/v1/getPositionById/1", http.StatusOK},
		{"UpdatePositionById", http.MethodPut, "/api/v1/updatePositionById/1", http.StatusOK},
		{"DeletePositionById", http.MethodDelete, "/api/v1/deletePositionById/1", http.StatusOK},
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
