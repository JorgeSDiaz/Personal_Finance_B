package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JorgeSDiaz/Personal_Finance_B/internal/user/handler"
	"github.com/JorgeSDiaz/Personal_Finance_B/internal/user/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func setupRouter(h *handler.UserHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/users/registry", h.Registry)
	r.POST("/users/login", h.LogIn)
	r.GET("/users", h.Users)
	return r
}

func TestHandler_Registry(t *testing.T) {
	cases := []struct {
		name           string
		body           string
		mockService    *mockService
		expectedStatus int
	}{
		{
			name:           "Test Handler - Registry - 200 OK",
			body:           dummyRawRegistryRequest,
			mockService:    buildMockRegistryService(&dummyUserRegistryRequest, dummyUserResponse, nil),
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Test Handler - Registry - 400 BadRequest - invalid JSON",
			body:           `{"email":123,"password":true}`,
			mockService:    new(mockService),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Test Handler - Registry - 400 BadRequest - service error (duplicated)",
			body:           dummyRawRegistryRequest,
			mockService:    buildMockRegistryService(&dummyUserRegistryRequest, model.UserResponse{}, errAlready),
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			h := handler.NewHandler(tc.mockService)
			r := setupRouter(h)

			req := httptest.NewRequest(http.MethodPost, "/users/registry", bytes.NewReader([]byte(tc.body)))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			require.Equal(t, tc.expectedStatus, w.Code, w.Body.String())
		})
	}
}

func TestUserHandler_LogIn_Table(t *testing.T) {
	cases := []struct {
		name           string
		body           string
		mockService    *mockService
		expectedStatus int
	}{
		{
			name:           "200 OK",
			body:           dummyRawLogInRequest,
			mockService:    buildMockLogInService(&dummyUserLogInRequest, dummyUserResponse, nil),
			expectedStatus: http.StatusOK,
		},
		{
			name:           "400 BadRequest - invalid JSON",
			body:           `{"email":false,"password":123}`,
			mockService:    new(mockService),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "400 BadRequest - service error",
			body:           dummyRawLogInRequest,
			mockService:    buildMockLogInService(&dummyUserLogInRequest, model.UserResponse{}, errInvalid),
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			h := handler.NewHandler(tc.mockService)
			r := setupRouter(h)

			req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewReader([]byte(tc.body)))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			require.Equal(t, tc.expectedStatus, w.Code, w.Body.String())
		})
	}
}

func TestUserHandler_Users(t *testing.T) {
	ms := buildMockAllUsersService(dummyUsersResponse)

	h := handler.NewHandler(ms)
	r := setupRouter(h)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, w.Body.String())

	var got []model.UserResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &got))
	require.Len(t, got, 2)

	ms.AssertExpectations(t)
}
