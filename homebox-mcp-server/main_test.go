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

func TestGetNotifiers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/notifiers", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id":"1","name":"Test Notifier","url":"http://test.com","isActive":true}]`))
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	_, output, err := getNotifiers(context.Background(), nil, GetNotifiersInput{})
	assert.NoError(t, err)
	assert.Len(t, output.Notifiers, 1)
	assert.Equal(t, "1", output.Notifiers[0].ID)
	assert.Equal(t, "Test Notifier", output.Notifiers[0].Name)
}

func TestCreateNotifier(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/notifiers", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"2","name":"New Notifier","url":"http://new.com","isActive":false}`))
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	_, output, err := createNotifier(context.Background(), nil, CreateNotifierInput{Name: "New Notifier", URL: "http://new.com", IsActive: false})
	assert.NoError(t, err)
	assert.Equal(t, "2", output.ID)
	assert.Equal(t, "New Notifier", output.Name)
}

func TestUpdateNotifier(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/notifiers/2", r.URL.Path)
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"2","name":"Updated Notifier","url":"http://updated.com","isActive":true}`))
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	_, output, err := updateNotifier(context.Background(), nil, UpdateNotifierInput{ID: "2", Name: "Updated Notifier", URL: "http://updated.com", IsActive: true})
	assert.NoError(t, err)
	assert.Equal(t, "2", output.ID)
	assert.Equal(t, "Updated Notifier", output.Name)
}

func TestDeleteNotifier(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/notifiers/2", r.URL.Path)
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	_, _, err := deleteNotifier(context.Background(), nil, DeleteNotifierInput{ID: "2"})
	assert.NoError(t, err)
}

func TestTestNotifier(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/notifiers/test", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	_, _, err := testNotifier(context.Background(), nil, TestNotifierInput{URL: "http://test.com"})
	assert.NoError(t, err)
}

func TestUpdateMaintenanceEntry(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/maintenance/3", r.URL.Path)
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"3","name":"Updated Maintenance"}`))
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	_, output, err := updateMaintenanceEntry(context.Background(), nil, UpdateMaintenanceEntryInput{ID: "3", Name: "Updated Maintenance"})
	assert.NoError(t, err)
	assert.Equal(t, "3", output.ID)
	assert.Equal(t, "Updated Maintenance", output.Name)
}

func TestDeleteMaintenanceEntry(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/maintenance/3", r.URL.Path)
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	_, _, err := deleteMaintenanceEntry(context.Background(), nil, DeleteMaintenanceEntryInput{ID: "3"})
	assert.NoError(t, err)
}

func TestSearchFromBarcode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/products/search-from-barcode" && r.URL.Query().Get("data") == "12345" && r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[{"id":"1","name":"Product 1"}]`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	input := SearchFromBarcodeInput{
		Data: "12345",
	}

	_, output, err := searchFromBarcode(context.Background(), nil, input)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(output.Items) != 1 || output.Items[0].ID != "1" {
		t.Errorf("Expected 1 product with ID 1, got %+v", output.Items)
	}
}

func TestCreateQRCode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/qrcode" && r.URL.Query().Get("data") == "test-data" && r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("image-data"))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	input := CreateQRCodeInput{
		Data: "test-data",
	}

	_, output, err := createQRCode(context.Background(), nil, input)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedImage := base64.StdEncoding.EncodeToString([]byte("image-data"))
	if output.Image != expectedImage {
		t.Errorf("Expected %s, got %s", expectedImage, output.Image)
	}
}

func TestExportBillOfMaterials(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/reporting/bill-of-materials" && r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("csv-data"))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	_, output, err := exportBillOfMaterials(context.Background(), nil, ExportBillOfMaterialsInput{})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if output.Data != "csv-data" {
		t.Errorf("Expected csv-data, got %s", output.Data)
	}
}

func TestCreateItemAttachment(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/items/item123/attachments", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"id":"att123","title":"test.png"}`))
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	input := CreateItemAttachmentInput{
		ItemID:      "item123",
		FileContent: base64.StdEncoding.EncodeToString([]byte("dummy content")),
		FileName:    "test.png",
	}

	_, output, err := createItemAttachment(context.Background(), nil, input)
	assert.NoError(t, err)
	assert.Equal(t, "att123", output.ID)
	assert.Equal(t, "test.png", output.Title)
}

func TestGetItemAttachment(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/items/item123/attachments/att123", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("dummy content"))
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	input := GetItemAttachmentInput{
		ItemID:       "item123",
		AttachmentID: "att123",
	}

	_, output, err := getItemAttachment(context.Background(), nil, input)
	assert.NoError(t, err)
	assert.Equal(t, base64.StdEncoding.EncodeToString([]byte("dummy content")), output.FileContent)
}

func TestUpdateItemAttachment(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/items/item123/attachments/att123", r.URL.Path)
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"att123","title":"updated.png"}`))
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	input := UpdateItemAttachmentInput{
		ItemID:       "item123",
		AttachmentID: "att123",
		Title:        "updated.png",
	}

	_, output, err := updateItemAttachment(context.Background(), nil, input)
	assert.NoError(t, err)
	assert.Equal(t, "att123", output.ID)
	assert.Equal(t, "updated.png", output.Title)
}

func TestDeleteItemAttachment(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/items/item123/attachments/att123", r.URL.Path)
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	t.Setenv("HOMEBOX_URL", server.URL)
	t.Setenv("HOMEBOX_TOKEN", "test-token")

	input := DeleteItemAttachmentInput{
		ItemID:       "item123",
		AttachmentID: "att123",
	}

	_, _, err := deleteItemAttachment(context.Background(), nil, input)
	assert.NoError(t, err)
}
