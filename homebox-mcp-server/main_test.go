package main

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetAssetLabel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/labelmaker/assets/123" {
			t.Errorf("Expected to request '/api/v1/labelmaker/assets/123', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test image data"))
	}))
	defer server.Close()

	os.Setenv("HOMEBOX_URL", server.URL)
	os.Setenv("HOMEBOX_TOKEN", "test-token")

	input := GetAssetLabelInput{ID: "123"}
	_, output, err := getAssetLabel(context.Background(), nil, input)
	if err != nil {
		t.Fatalf("getAssetLabel failed: %v", err)
	}

	expectedImage := base64.StdEncoding.EncodeToString([]byte("test image data"))
	if output.Image != expectedImage {
		t.Errorf("Expected image '%s', got '%s'", expectedImage, output.Image)
	}
}

func TestGetItemLabel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/labelmaker/item/456" {
			t.Errorf("Expected to request '/api/v1/labelmaker/item/456', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test item image data"))
	}))
	defer server.Close()

	os.Setenv("HOMEBOX_URL", server.URL)
	os.Setenv("HOMEBOX_TOKEN", "test-token")

	input := GetItemLabelInput{ID: "456"}
	_, output, err := getItemLabel(context.Background(), nil, input)
	if err != nil {
		t.Fatalf("getItemLabel failed: %v", err)
	}

	expectedImage := base64.StdEncoding.EncodeToString([]byte("test item image data"))
	if output.Image != expectedImage {
		t.Errorf("Expected image '%s', got '%s'", expectedImage, output.Image)
	}
}

func TestGetLocationLabel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/labelmaker/location/789" {
			t.Errorf("Expected to request '/api/v1/labelmaker/location/789', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test location image data"))
	}))
	defer server.Close()

	os.Setenv("HOMEBOX_URL", server.URL)
	os.Setenv("HOMEBOX_TOKEN", "test-token")

	input := GetLocationLabelInput{ID: "789"}
	_, output, err := getLocationLabel(context.Background(), nil, input)
	if err != nil {
		t.Fatalf("getLocationLabel failed: %v", err)
	}

	expectedImage := base64.StdEncoding.EncodeToString([]byte("test location image data"))
	if output.Image != expectedImage {
		t.Errorf("Expected image '%s', got '%s'", expectedImage, output.Image)
	}
}

func TestGetLabelImage_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	os.Setenv("HOMEBOX_URL", server.URL)
	os.Setenv("HOMEBOX_TOKEN", "test-token")

	_, err := getLabelImage("some/path")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}
