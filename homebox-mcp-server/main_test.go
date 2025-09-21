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
