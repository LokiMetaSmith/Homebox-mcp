package main

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGroup(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/groups", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"123","name":"Test Group","currency":"USD"}`))
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	_, output, err := getGroup(context.Background(), nil, GetGroupInput{})
	assert.NoError(t, err)
	assert.Equal(t, "123", output.ID)
	assert.Equal(t, "Test Group", output.Name)
	assert.Equal(t, "USD", output.Currency)
}

func TestUpdateGroup(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/groups", r.URL.Path)
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"123","name":"Updated Group","currency":"EUR"}`))
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	_, output, err := updateGroup(context.Background(), nil, UpdateGroupInput{Name: "Updated Group", Currency: "EUR"})
	assert.NoError(t, err)
	assert.Equal(t, "123", output.ID)
	assert.Equal(t, "Updated Group", output.Name)
	assert.Equal(t, "EUR", output.Currency)
}

func TestGetGroupStatistics(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/groups/statistics", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"totalItems":42,"totalUsers":3}`))
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	_, output, err := getGroupStatistics(context.Background(), nil, GetGroupStatisticsInput{})
	assert.NoError(t, err)
	assert.Equal(t, 42, output.TotalItems)
	assert.Equal(t, 3, output.TotalUsers)
}

func TestGetAssetLabel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/labelmaker/assets/123", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("image data"))
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	_, output, err := getAssetLabel(context.Background(), nil, GetAssetLabelInput{ID: "123"})
	assert.NoError(t, err)
	assert.Equal(t, base64.StdEncoding.EncodeToString([]byte("image data")), output.Image)
}

func TestGetItemLabel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/labelmaker/item/456", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("image data"))
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	_, output, err := getItemLabel(context.Background(), nil, GetItemLabelInput{ID: "456"})
	assert.NoError(t, err)
	assert.Equal(t, base64.StdEncoding.EncodeToString([]byte("image data")), output.Image)
}

func TestGetLocationLabel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/labelmaker/location/789", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("image data"))
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	_, output, err := getLocationLabel(context.Background(), nil, GetLocationLabelInput{ID: "789"})
	assert.NoError(t, err)
	assert.Equal(t, base64.StdEncoding.EncodeToString([]byte("image data")), output.Image)
}

func TestGetUserSelf(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/users/self", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"u1","email":"test@example.com","name":"Test User","username":"testuser","isAdmin":true}`))
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	_, output, err := getUserSelf(context.Background(), nil, GetUserSelfInput{})
	assert.NoError(t, err)
	assert.Equal(t, "u1", output.ID)
	assert.Equal(t, "test@example.com", output.Email)
	assert.Equal(t, "Test User", output.Name)
	assert.Equal(t, "testuser", output.Username)
	assert.True(t, output.IsAdmin)
}

func TestUpdateUserSelf(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/users/self", r.URL.Path)
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"u1","email":"new@example.com","name":"New Name","username":"newuser"}`))
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	_, output, err := updateUserSelf(context.Background(), nil, UpdateUserSelfInput{Email: "new@example.com", Name: "New Name", Username: "newuser"})
	assert.NoError(t, err)
	assert.Equal(t, "new@example.com", output.Email)
	assert.Equal(t, "New Name", output.Name)
	assert.Equal(t, "newuser", output.Username)
}

func TestChangePassword(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/users/change-password", r.URL.Path)
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	_, _, err := changePassword(context.Background(), nil, ChangePasswordInput{CurrentPassword: "old", NewPassword: "new"})
	assert.NoError(t, err)
}

func TestDeleteUserSelf(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/users/self", r.URL.Path)
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	_, _, err := deleteUserSelf(context.Background(), nil, DeleteUserSelfInput{})
	assert.NoError(t, err)
}

func TestRegisterUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/users/register", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"u2","email":"reg@example.com","name":"Reg User","username":"reguser"}`))
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)

	_, output, err := registerUser(context.Background(), nil, RegisterUserInput{Email: "reg@example.com", Name: "Reg User", Username: "reguser", Password: "pass"})
	assert.NoError(t, err)
	assert.Equal(t, "u2", output.ID)
	assert.Equal(t, "reg@example.com", output.Email)
}
