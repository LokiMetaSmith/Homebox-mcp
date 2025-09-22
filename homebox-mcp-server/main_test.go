package main

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAssetLabel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/labelmaker/assets/123", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("image data"))
	}))
	defer server.Close()

	os.Setenv("HOMEBOX_URL", server.URL)
	os.Setenv("HOMEBOX_TOKEN", "test-token")

	_, output, err := getAssetLabel(context.Background(), nil, GetAssetLabelInput{ID: "123"})
	assert.NoError(t, err)
	assert.Equal(t, base64.StdEncoding.EncodeToString([]byte("image data")), output.Image)
}

func TestChangePassword(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/users/change-password", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"user-123","name":"Test User","email":"test@example.com"}`))
	}))
	defer server.Close()

	os.Setenv("HOMEBOX_URL", server.URL)
	os.Setenv("HOMEBOX_TOKEN", "test-token")

	_, output, err := changePassword(context.Background(), nil, ChangePasswordInput{CurrentPassword: "old", NewPassword: "new"})
	assert.NoError(t, err)
	assert.Equal(t, "user-123", output.ID)
}

func TestGetCurrentUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/users/self", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"user-123","name":"Test User","email":"test@example.com"}`))
	}))
	defer server.Close()

	os.Setenv("HOMEBOX_URL", server.URL)
	os.Setenv("HOMEBOX_TOKEN", "test-token")

	_, output, err := getCurrentUser(context.Background(), nil, GetCurrentUserInput{})
	assert.NoError(t, err)
	assert.Equal(t, "user-123", output.ID)
}

func TestUpdateCurrentUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/users/self", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"user-123","name":"Updated User","email":"updated@example.com"}`))
	}))
	defer server.Close()

	os.Setenv("HOMEBOX_URL", server.URL)
	os.Setenv("HOMEBOX_TOKEN", "test-token")

	_, output, err := updateCurrentUser(context.Background(), nil, UpdateUserInput{Name: "Updated User", Email: "updated@example.com"})
	assert.NoError(t, err)
	assert.Equal(t, "Updated User", output.Name)
}

func TestDeleteCurrentUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/users/self", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	os.Setenv("HOMEBOX_URL", server.URL)
	os.Setenv("HOMEBOX_TOKEN", "test-token")

	_, _, err := deleteCurrentUser(context.Background(), nil, DeleteCurrentUserInput{})
	assert.NoError(t, err)
}

func TestRegisterUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/users/register", r.URL.Path)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"id":"user-123","name":"New User","email":"new@example.com"}`))
	}))
	defer server.Close()

	os.Setenv("HOMEBOX_URL", server.URL)
	os.Setenv("HOMEBOX_TOKEN", "test-token")

	_, output, err := registerUser(context.Background(), nil, RegisterUserInput{Name: "New User", Email: "new@example.com", Password: "password"})
	assert.NoError(t, err)
	assert.Equal(t, "New User", output.Name)
}

func TestGetItemLabel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/labelmaker/item/456", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("image data"))
	}))
	defer server.Close()

	os.Setenv("HOMEBOX_URL", server.URL)
	os.Setenv("HOMEBOX_TOKEN", "test-token")

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

	os.Setenv("HOMEBOX_URL", server.URL)
	os.Setenv("HOMEBOX_TOKEN", "test-token")

	_, output, err := getLocationLabel(context.Background(), nil, GetLocationLabelInput{ID: "789"})
	assert.NoError(t, err)
	assert.Equal(t, base64.StdEncoding.EncodeToString([]byte("image data")), output.Image)
}
