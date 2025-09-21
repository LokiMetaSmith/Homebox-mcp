package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// This struct represents a single item from your Homebox,
// based on the fields in the Homebox API.
type HomeboxItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	LocationID  string `json:"location_id,omitempty"`
}

type LabelSummary struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Color       string `json:"color,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}

type LocationSummary struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}

type ItemSummary struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	Description   string            `json:"description,omitempty"`
	Archived      bool              `json:"archived"`
	AssetID       string            `json:"assetId"`
	CreatedAt     string            `json:"createdAt"`
	ImageID       string            `json:"imageId,omitempty"`
	Insured       bool              `json:"insured"`
	Labels        []LabelSummary    `json:"labels"`
	Location      *LocationSummary  `json:"location,omitempty"`
	PurchasePrice float64           `json:"purchasePrice"`
	Quantity      int               `json:"quantity"`
	SoldTime      string            `json:"soldTime,omitempty"`
	ThumbnailID   string            `json:"thumbnailId,omitempty"`
	UpdatedAt     string            `json:"updatedAt"`
}

type ItemAttachment struct {
	ID        string          `json:"id"`
	CreatedAt string          `json:"createdAt"`
	MimeType  string          `json:"mimeType"`
	Path      string          `json:"path"`
	Primary   bool            `json:"primary"`
	Thumbnail *ItemAttachment `json:"thumbnail,omitempty"`
	Title     string          `json:"title"`
	Type      string          `json:"type"`
	UpdatedAt string          `json:"updatedAt"`
}

type ItemField struct {
	ID           string `json:"id"`
	BooleanValue bool   `json:"booleanValue"`
	Name         string `json:"name"`
	NumberValue  int    `json:"numberValue"`
	TextValue    string `json:"textValue"`
	Type         string `json:"type"`
}

type ItemOut struct {
	ID                      string            `json:"id"`
	Archived                bool              `json:"archived"`
	AssetID                 string            `json:"assetId"`
	Attachments             []ItemAttachment  `json:"attachments"`
	CreatedAt               string            `json:"createdAt"`
	Description             string            `json:"description"`
	Fields                  []ItemField       `json:"fields"`
	ImageID                 string            `json:"imageId,omitempty"`
	Insured                 bool              `json:"insured"`
	Labels                  []LabelSummary    `json:"labels"`
	LifetimeWarranty        bool              `json:"lifetimeWarranty"`
	Location                *LocationSummary  `json:"location,omitempty"`
	Manufacturer            string            `json:"manufacturer"`
	ModelNumber             string            `json:"modelNumber"`
	Name                    string            `json:"name"`
	Notes                   string            `json:"notes"`
	Parent                  *ItemSummary      `json:"parent,omitempty"`
	PurchaseFrom            string            `json:"purchaseFrom"`
	PurchasePrice           float64           `json:"purchasePrice"`
	PurchaseTime            string            `json:"purchaseTime"`
	Quantity                int               `json:"quantity"`
	SerialNumber            string            `json:"serialNumber"`
	SoldNotes               string            `json:"soldNotes"`
	SoldPrice               float64           `json:"soldPrice"`
	SoldTime                string            `json:"soldTime"`
	SoldTo                  string            `json:"soldTo"`
	SyncChildItemsLocations bool              `json:"syncChildItemsLocations"`
	ThumbnailID             string            `json:"thumbnailId,omitempty"`
	UpdatedAt               string            `json:"updatedAt"`
	WarrantyDetails         string            `json:"warrantyDetails"`
	WarrantyExpires         string            `json:"warrantyExpires"`
}

type LocationOutCount struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ItemCount   int    `json:"itemCount"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type LocationOut struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Parent      *LocationSummary  `json:"parent,omitempty"`
	Children    []LocationSummary `json:"children,omitempty"`
	TotalPrice  float64           `json:"totalPrice"`
	CreatedAt   string            `json:"createdAt"`
	UpdatedAt   string            `json:"updatedAt"`
}

type LabelOut struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type MaintenanceEntryWithDetails struct {
	ID            string `json:"id"`
	CompletedDate string `json:"completedDate,omitempty"`
	Cost          string `json:"cost,omitempty"`
	Description   string `json:"description,omitempty"`
	ItemID        string `json:"itemID"`
	ItemName      string `json:"itemName"`
	Name          string `json:"name"`
	ScheduledDate string `json:"scheduledDate,omitempty"`
}

type MaintenanceEntry struct {
	ID            string `json:"id"`
	CompletedDate string `json:"completedDate,omitempty"`
	Cost          string `json:"cost,omitempty"`
	Description   string `json:"description,omitempty"`
	Name          string `json:"name"`
	ScheduledDate string `json:"scheduledDate,omitempty"`
}

type ItemPath struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type PaginationResult_ItemSummary struct {
	Items    []ItemSummary `json:"items"`
	Page     int           `json:"page"`
	PageSize int           `json:"pageSize"`
	Total    int           `json:"total"`
}

type ActionAmountResult struct {
	Completed int `json:"completed"`
}

type Build struct {
	BuildTime string `json:"buildTime"`
	Commit    string `json:"commit"`
	Version   string `json:"version"`
}

type Latest struct {
	Date    string `json:"date"`
	Version string `json:"version"`
}

type APISummary struct {
	AllowRegistration bool     `json:"allowRegistration"`
	Build             Build    `json:"build"`
	Demo              bool     `json:"demo"`
	Health            bool     `json:"health"`
	LabelPrinting     bool     `json:"labelPrinting"`
	Latest            Latest   `json:"latest"`
	Message           string   `json:"message"`
	Title             string   `json:"title"`
	Versions          []string `json:"versions"`
}

type Currency struct {
	Code   string `json:"code"`
	Local  string `json:"local"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type ItemAttachmentToken struct {
	Token string `json:"token"`
}

type Group struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Currency  string `json:"currency"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type GroupStatistics struct {
	TotalItemPrice    float64 `json:"totalItemPrice"`
	TotalItems        int     `json:"totalItems"`
	TotalLabels       int     `json:"totalLabels"`
	TotalLocations    int     `json:"totalLocations"`
	TotalUsers        int     `json:"totalUsers"`
	TotalWithWarranty int     `json:"totalWithWarranty"`
}

type TotalsByOrganizer struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Total float64 `json:"total"`
}

type ValueOverTimeEntry struct {
	Date  string  `json:"date"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type ValueOverTime struct {
	End          string               `json:"end"`
	Entries      []ValueOverTimeEntry `json:"entries"`
	Start        string               `json:"start"`
	ValueAtEnd   float64              `json:"valueAtEnd"`
	ValueAtStart float64              `json:"valueAtStart"`
}

type NotifierOut struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	IsActive  bool   `json:"isActive"`
	UserID    string `json:"userId"`
	GroupID   string `json:"groupId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type BarcodeProduct struct {
	Barcode         string        `json:"barcode"`
	ImageBase64     string        `json:"imageBase64"`
	ImageURL        string        `json:"imageURL"`
	Item            CreateItemInput `json:"item"`
	Manufacturer    string        `json:"manufacturer"`
	ModelNumber     string        `json:"modelNumber"`
	Notes           string        `json:"notes"`
	SearchEngineName string        `json:"search_engine_name"`
}

// Input for the create_item tool.
type CreateItemInput struct {
	Name        string   `json:"name" jsonschema:"minLength:1,maxLength:255"`
	Description string   `json:"description,omitempty" jsonschema:"maxLength:1000"`
	LabelIDs    []string `json:"labelIds,omitempty"`
	LocationID  string   `json:"locationId,omitempty"`
	ParentID    string   `json:"parentId,omitempty"`
	Quantity    int      `json:"quantity,omitempty"`
}

// Input for the get_item tool.
type GetItemInput struct {
	ID string `json:"id" jsonschema:"required"`
}

// Input for the update_item tool.
type UpdateItemInput struct {
	ID                      string      `json:"id" jsonschema:"required"`
	Archived                bool        `json:"archived,omitempty"`
	AssetID                 string      `json:"assetId,omitempty"`
	Description             string      `json:"description,omitempty"`
	Fields                  []ItemField `json:"fields,omitempty"`
	Insured                 bool        `json:"insured,omitempty"`
	LabelIDs                []string    `json:"labelIds,omitempty"`
	LifetimeWarranty        bool        `json:"lifetimeWarranty,omitempty"`
	LocationID              string      `json:"locationId,omitempty"`
	Manufacturer            string      `json:"manufacturer,omitempty"`
	ModelNumber             string      `json:"modelNumber,omitempty"`
	Name                    string      `json:"name"`
	Notes                   string      `json:"notes,omitempty"`
	ParentID                string      `json:"parentId,omitempty"`
	PurchaseFrom            string      `json:"purchaseFrom,omitempty"`
	PurchasePrice           float64     `json:"purchasePrice,omitempty"`
	PurchaseTime            string      `json:"purchaseTime,omitempty"`
	Quantity                int         `json:"quantity,omitempty"`
	SerialNumber            string      `json:"serialNumber,omitempty"`
	SoldNotes               string      `json:"soldNotes,omitempty"`
	SoldPrice               float64     `json:"soldPrice,omitempty"`
	SoldTime                string      `json:"soldTime,omitempty"`
	SoldTo                  string      `json:"soldTo,omitempty"`
	SyncChildItemsLocations bool        `json:"syncChildItemsLocations,omitempty"`
	WarrantyDetails         string      `json:"warrantyDetails,omitempty"`
	WarrantyExpires         string      `json:"warrantyExpires,omitempty"`
}

// Input for the delete_item tool.
type DeleteItemInput struct {
	ID string `json:"id" jsonschema:"required"`
}

// Output for the delete_item tool.
type DeleteItemOutput struct{}

// Input for the get_items tool. It takes no parameters.
type GetItemsInput struct{}

// Output for the get_items tool. It returns a list of items.
type GetItemsOutput struct {
	Items []HomeboxItem `json:"items"`
}

// Input for get_locations tool.
type GetLocationsInput struct{}

// Output for get_locations tool.
type GetLocationsOutput struct {
	Locations []LocationOutCount `json:"locations"`
}

// Input for create_location tool.
type CreateLocationInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	ParentID    string `json:"parentId,omitempty"`
}

// Input for get_location tool.
type GetLocationInput struct {
	ID string `json:"id" jsonschema:"required"`
}

// Input for update_location tool.
type UpdateLocationInput struct {
	ID          string `json:"id" jsonschema:"required"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	ParentID    string `json:"parentId,omitempty"`
}

// Input for delete_location tool.
type DeleteLocationInput struct {
	ID string `json:"id" jsonschema:"required"`
}

// Output for delete_location tool.
type DeleteLocationOutput struct{}

// Input for get_labels tool.
type GetLabelsInput struct{}

// Output for get_labels tool.
type GetLabelsOutput struct {
	Labels []LabelOut `json:"labels"`
}

// Input for create_label tool.
type CreateLabelInput struct {
	Name        string `json:"name" jsonschema:"minLength:1,maxLength:255"`
	Description string `json:"description,omitempty" jsonschema:"maxLength:1000"`
	Color       string `json:"color,omitempty"`
}

// Input for get_label tool.
type GetLabelInput struct {
	ID string `json:"id" jsonschema:"required"`
}

// Input for update_label tool.
type UpdateLabelInput struct {
	ID          string `json:"id" jsonschema:"required"`
	Name        string `json:"name" jsonschema:"minLength:1,maxLength:255"`
	Description string `json:"description,omitempty" jsonschema:"maxLength:1000"`
	Color       string `json:"color,omitempty"`
}

// Input for delete_label tool.
type DeleteLabelInput struct {
	ID string `json:"id" jsonschema:"required"`
}

// Output for delete_label tool.
type DeleteLabelOutput struct{}

// Input for get_maintenance_log tool.
type GetMaintenanceLogInput struct {
	ItemID string `json:"item_id" jsonschema:"required"`
}

// Output for get_maintenance_log tool.
type GetMaintenanceLogOutput struct {
	Entries []MaintenanceEntryWithDetails `json:"entries"`
}

// Input for create_maintenance_entry tool.
type CreateMaintenanceEntryInput struct {
	ItemID        string `json:"item_id" jsonschema:"required"`
	Name          string `json:"name" jsonschema:"required"`
	CompletedDate string `json:"completedDate,omitempty"`
	Cost          string `json:"cost,omitempty"`
	Description   string `json:"description,omitempty"`
	ScheduledDate string `json:"scheduledDate,omitempty"`
}

// Input for duplicate_item tool.
type DuplicateItemInput struct {
	ID               string `json:"id" jsonschema:"required"`
	CopyAttachments  bool   `json:"copyAttachments,omitempty"`
	CopyCustomFields bool   `json:"copyCustomFields,omitempty"`
	CopyMaintenance  bool   `json:"copyMaintenance,omitempty"`
	CopyPrefix       string `json:"copyPrefix,omitempty"`
}

// Input for get_item_path tool.
type GetItemPathInput struct {
	ID string `json:"id" jsonschema:"required"`
}

// Output for get_item_path tool.
type GetItemPathOutput struct {
	Path []ItemPath `json:"path"`
}

// Input for export_items tool.
type ExportItemsInput struct{}

// Output for export_items tool.
type ExportItemsOutput struct {
	CSVData string `json:"csv_data"`
}

// Input for get_item_fields tool.
type GetItemFieldsInput struct{}

// Output for get_item_fields tool.
type GetItemFieldsOutput struct {
	Fields []string `json:"fields"`
}

// Input for get_item_field_values tool.
type GetItemFieldValuesInput struct{}

// Output for get_item_field_values tool.
type GetItemFieldValuesOutput struct {
	Values []string `json:"values"`
}

// Input for get_item_by_asset_id tool.
type GetItemByAssetIDInput struct {
	ID string `json:"id" jsonschema:"required"`
}

// Action Inputs
type CreateMissingThumbnailsInput struct{}
type EnsureAssetIDsInput struct{}
type EnsureImportRefsInput struct{}
type SetPrimaryPhotosInput struct{}
type ZeroItemTimeFieldsInput struct{}
type GetStatusInput struct{}
type GetCurrencyInput struct{}

// Maintenance Inputs
type UpdateMaintenanceEntryInput struct {
	ID            string `json:"id" jsonschema:"required"`
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
	CompletedDate string `json:"completedDate,omitempty"`
	Cost          string `json:"cost,omitempty"`
	ScheduledDate string `json:"scheduledDate,omitempty"`
}

type DeleteMaintenanceEntryInput struct {
	ID string `json:"id" jsonschema:"required"`
}

type DeleteMaintenanceEntryOutput struct{}

// Item Attachment Inputs
type DeleteItemAttachmentInput struct {
	ItemID       string `json:"item_id" jsonschema:"required"`
	AttachmentID string `json:"attachment_id" jsonschema:"required"`
}

type DeleteItemAttachmentOutput struct{}

type UpdateItemAttachmentInput struct {
	ItemID       string `json:"item_id" jsonschema:"required"`
	AttachmentID string `json:"attachment_id" jsonschema:"required"`
	Primary      bool   `json:"primary,omitempty"`
	Title        string `json:"title,omitempty"`
	Type         string `json:"type,omitempty"`
}

type GetItemAttachmentInput struct {
	ItemID       string `json:"item_id" jsonschema:"required"`
	AttachmentID string `json:"attachment_id" jsonschema:"required"`
}

type GetItemAttachmentOutput struct {
	FileContent string `json:"file_content"`
}

type CreateItemAttachmentInput struct {
	ItemID      string `json:"item_id" jsonschema:"required"`
	FileContent string `json:"file_content" jsonschema:"required,description:Base64 encoded file content"`
	FileName    string `json:"file_name" jsonschema:"required"`
	Type        string `json:"type,omitempty"`
	Primary     bool   `json:"primary,omitempty"`
}

// Input for the import_items tool.
type ImportItemsInput struct {
	FileContent string `json:"file_content" jsonschema:"required,description:Base64 encoded file content"`
	FileName    string `json:"file_name" jsonschema:"required"`
}

// Output for the import_items tool.
type ImportItemsOutput struct {
	Completed int `json:"completed"`
}

// Group Inputs
type GetGroupInput struct{}
type UpdateGroupInput struct {
	Name     string `json:"name,omitempty"`
	Currency string `json:"currency,omitempty"`
}
type CreateGroupInvitationInput struct {
    Email string `json:"email" jsonschema:"required,format:email"`
}
type GroupInvitation struct {
    ID      string `json:"id"`
    Email   string `json:"email"`
    Expires string `json:"expires,omitempty"`
}
type GetGroupStatisticsInput struct{}
type GetLabelStatisticsInput struct{}
type GetLocationStatisticsInput struct{}
type GetPurchasePriceStatisticsInput struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

// Notifier Inputs
type GetNotifiersInput struct{}
type CreateNotifierInput struct {
	Name     string `json:"name" jsonschema:"required"`
	URL      string `json:"url" jsonschema:"required"`
	IsActive bool   `json:"isActive,omitempty"`
}
type UpdateNotifierInput struct {
	ID       string `json:"id" jsonschema:"required"`
	Name     string `json:"name,omitempty"`
	URL      string `json:"url,omitempty"`
	IsActive bool   `json:"isActive,omitempty"`
}
type DeleteNotifierInput struct {
	ID string `json:"id" jsonschema:"required"`
}
type TestNotifierInput struct {
	URL string `json:"url" jsonschema:"required"`
}

// Product Inputs
type SearchFromBarcodeInput struct {
	Data string `json:"data" jsonschema:"required"`
}

// QR Code Inputs
type CreateQRCodeInput struct {
	Data string `json:"data" jsonschema:"required"`
}

// Reporting Inputs
type ExportBillOfMaterialsInput struct{}


// getItems is the implementation of the "get_items" tool.
func getItems(ctx context.Context, req *mcp.CallToolRequest, input GetItemsInput) (*mcp.CallToolResult, GetItemsOutput, error) {
	// Get the Homebox API URL and token from environment variables.
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, GetItemsOutput{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	// Create a new HTTP request to the Homebox API.
	httpReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/items", homeboxURL), nil)
	if err != nil {
		return nil, GetItemsOutput{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)

	// Execute the request.
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, GetItemsOutput{}, err
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, GetItemsOutput{}, err
	}

	// Unmarshal the JSON response into a slice of HomeboxItem structs.
	var items []HomeboxItem
	if err := json.Unmarshal(body, &items); err != nil {
		return nil, GetItemsOutput{}, err
	}

	// Return the items.
	return nil, GetItemsOutput{Items: items}, nil
}

// createItem is the implementation of the "create_item" tool.
func createItem(ctx context.Context, req *mcp.CallToolRequest, input CreateItemInput) (*mcp.CallToolResult, ItemSummary, error) {
	// Get the Homebox API URL and token from environment variables.
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, ItemSummary{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	// Marshal the input to JSON
	reqBody, err := json.Marshal(input)
	if err != nil {
		return nil, ItemSummary{}, err
	}

	// Create a new HTTP request to the Homebox API.
	httpReq, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/items", homeboxURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, ItemSummary{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)
	httpReq.Header.Set("Content-Type", "application/json")

	// Execute the request.
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, ItemSummary{}, err
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ItemSummary{}, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, ItemSummary{}, fmt.Errorf("failed to create item, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Unmarshal the JSON response into an ItemSummary struct.
	var itemSummary ItemSummary
	if err := json.Unmarshal(body, &itemSummary); err != nil {
		return nil, ItemSummary{}, err
	}

	// Return the created item summary.
	return nil, itemSummary, nil
}

// getItem is the implementation of the "get_item" tool.
func getItem(ctx context.Context, req *mcp.CallToolRequest, input GetItemInput) (*mcp.CallToolResult, ItemOut, error) {
	// Get the Homebox API URL and token from environment variables.
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, ItemOut{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	// Create a new HTTP request to the Homebox API.
	httpReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/items/%s", homeboxURL, input.ID), nil)
	if err != nil {
		return nil, ItemOut{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)

	// Execute the request.
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, ItemOut{}, err
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ItemOut{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ItemOut{}, fmt.Errorf("failed to get item, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Unmarshal the JSON response into an ItemOut struct.
	var itemOut ItemOut
	if err := json.Unmarshal(body, &itemOut); err != nil {
		return nil, ItemOut{}, err
	}

	// Return the item.
	return nil, itemOut, nil
}

// updateItem is the implementation of the "update_item" tool.
func updateItem(ctx context.Context, req *mcp.CallToolRequest, input UpdateItemInput) (*mcp.CallToolResult, ItemOut, error) {
	// Get the Homebox API URL and token from environment variables.
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, ItemOut{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	// Marshal the input to JSON
	reqBody, err := json.Marshal(input)
	if err != nil {
		return nil, ItemOut{}, err
	}

	// Create a new HTTP request to the Homebox API.
	httpReq, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v1/items/%s", homeboxURL, input.ID), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, ItemOut{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)
	httpReq.Header.Set("Content-Type", "application/json")

	// Execute the request.
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, ItemOut{}, err
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ItemOut{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ItemOut{}, fmt.Errorf("failed to update item, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Unmarshal the JSON response into an ItemOut struct.
	var itemOut ItemOut
	if err := json.Unmarshal(body, &itemOut); err != nil {
		return nil, ItemOut{}, err
	}

	// Return the updated item.
	return nil, itemOut, nil
}

// deleteItem is the implementation of the "delete_item" tool.
func deleteItem(ctx context.Context, req *mcp.CallToolRequest, input DeleteItemInput) (*mcp.CallToolResult, DeleteItemOutput, error) {
	// Get the Homebox API URL and token from environment variables.
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, DeleteItemOutput{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	// Create a new HTTP request to the Homebox API.
	httpReq, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/items/%s", homeboxURL, input.ID), nil)
	if err != nil {
		return nil, DeleteItemOutput{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)

	// Execute the request.
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, DeleteItemOutput{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return nil, DeleteItemOutput{}, fmt.Errorf("failed to delete item, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Return an empty output.
	return nil, DeleteItemOutput{}, nil
}

// getLocations is the implementation of the "get_locations" tool.
func getLocations(ctx context.Context, req *mcp.CallToolRequest, input GetLocationsInput) (*mcp.CallToolResult, GetLocationsOutput, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, GetLocationsOutput{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	httpReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/locations", homeboxURL), nil)
	if err != nil {
		return nil, GetLocationsOutput{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, GetLocationsOutput{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, GetLocationsOutput{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, GetLocationsOutput{}, fmt.Errorf("failed to get locations, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var locations []LocationOutCount
	if err := json.Unmarshal(body, &locations); err != nil {
		return nil, GetLocationsOutput{}, err
	}

	return nil, GetLocationsOutput{Locations: locations}, nil
}

// createLocation is the implementation of the "create_location" tool.
func createLocation(ctx context.Context, req *mcp.CallToolRequest, input CreateLocationInput) (*mcp.CallToolResult, LocationSummary, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, LocationSummary{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	reqBody, err := json.Marshal(input)
	if err != nil {
		return nil, LocationSummary{}, err
	}

	httpReq, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/locations", homeboxURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, LocationSummary{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, LocationSummary{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, LocationSummary{}, err
	}

	if resp.StatusCode != http.StatusOK { // Note: API spec says 200, not 201
		return nil, LocationSummary{}, fmt.Errorf("failed to create location, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var locationSummary LocationSummary
	if err := json.Unmarshal(body, &locationSummary); err != nil {
		return nil, LocationSummary{}, err
	}

	return nil, locationSummary, nil
}

// getLocation is the implementation of the "get_location" tool.
func getLocation(ctx context.Context, req *mcp.CallToolRequest, input GetLocationInput) (*mcp.CallToolResult, LocationOut, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, LocationOut{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	httpReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/locations/%s", homeboxURL, input.ID), nil)
	if err != nil {
		return nil, LocationOut{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, LocationOut{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, LocationOut{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, LocationOut{}, fmt.Errorf("failed to get location, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var locationOut LocationOut
	if err := json.Unmarshal(body, &locationOut); err != nil {
		return nil, LocationOut{}, err
	}

	return nil, locationOut, nil
}

// updateLocation is the implementation of the "update_location" tool.
func updateLocation(ctx context.Context, req *mcp.CallToolRequest, input UpdateLocationInput) (*mcp.CallToolResult, LocationOut, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, LocationOut{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	reqBody, err := json.Marshal(input)
	if err != nil {
		return nil, LocationOut{}, err
	}

	httpReq, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v1/locations/%s", homeboxURL, input.ID), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, LocationOut{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, LocationOut{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, LocationOut{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, LocationOut{}, fmt.Errorf("failed to update location, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var locationOut LocationOut
	if err := json.Unmarshal(body, &locationOut); err != nil {
		return nil, LocationOut{}, err
	}

	return nil, locationOut, nil
}

// deleteLocation is the implementation of the "delete_location" tool.
func deleteLocation(ctx context.Context, req *mcp.CallToolRequest, input DeleteLocationInput) (*mcp.CallToolResult, DeleteLocationOutput, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, DeleteLocationOutput{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	httpReq, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/locations/%s", homeboxURL, input.ID), nil)
	if err != nil {
		return nil, DeleteLocationOutput{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, DeleteLocationOutput{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return nil, DeleteLocationOutput{}, fmt.Errorf("failed to delete location, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	return nil, DeleteLocationOutput{}, nil
}

// getLabels is the implementation of the "get_labels" tool.
func getLabels(ctx context.Context, req *mcp.CallToolRequest, input GetLabelsInput) (*mcp.CallToolResult, GetLabelsOutput, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, GetLabelsOutput{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	httpReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/labels", homeboxURL), nil)
	if err != nil {
		return nil, GetLabelsOutput{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, GetLabelsOutput{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, GetLabelsOutput{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, GetLabelsOutput{}, fmt.Errorf("failed to get labels, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var labels []LabelOut
	if err := json.Unmarshal(body, &labels); err != nil {
		return nil, GetLabelsOutput{}, err
	}

	return nil, GetLabelsOutput{Labels: labels}, nil
}

// createLabel is the implementation of the "create_label" tool.
func createLabel(ctx context.Context, req *mcp.CallToolRequest, input CreateLabelInput) (*mcp.CallToolResult, LabelSummary, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, LabelSummary{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	reqBody, err := json.Marshal(input)
	if err != nil {
		return nil, LabelSummary{}, err
	}

	httpReq, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/labels", homeboxURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, LabelSummary{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, LabelSummary{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, LabelSummary{}, err
	}

	if resp.StatusCode != http.StatusOK { // Note: API spec says 200, not 201
		return nil, LabelSummary{}, fmt.Errorf("failed to create label, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var labelSummary LabelSummary
	if err := json.Unmarshal(body, &labelSummary); err != nil {
		return nil, LabelSummary{}, err
	}

	return nil, labelSummary, nil
}

// getLabel is the implementation of the "get_label" tool.
func getLabel(ctx context.Context, req *mcp.CallToolRequest, input GetLabelInput) (*mcp.CallToolResult, LabelOut, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, LabelOut{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	httpReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/labels/%s", homeboxURL, input.ID), nil)
	if err != nil {
		return nil, LabelOut{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, LabelOut{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, LabelOut{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, LabelOut{}, fmt.Errorf("failed to get label, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var labelOut LabelOut
	if err := json.Unmarshal(body, &labelOut); err != nil {
		return nil, LabelOut{}, err
	}

	return nil, labelOut, nil
}

// updateLabel is the implementation of the "update_label" tool.
func updateLabel(ctx context.Context, req *mcp.CallToolRequest, input UpdateLabelInput) (*mcp.CallToolResult, LabelOut, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, LabelOut{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	reqBody, err := json.Marshal(input)
	if err != nil {
		return nil, LabelOut{}, err
	}

	httpReq, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v1/labels/%s", homeboxURL, input.ID), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, LabelOut{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, LabelOut{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, LabelOut{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, LabelOut{}, fmt.Errorf("failed to update label, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var labelOut LabelOut
	if err := json.Unmarshal(body, &labelOut); err != nil {
		return nil, LabelOut{}, err
	}

	return nil, labelOut, nil
}

// deleteLabel is the implementation of the "delete_label" tool.
func deleteLabel(ctx context.Context, req *mcp.CallToolRequest, input DeleteLabelInput) (*mcp.CallToolResult, DeleteLabelOutput, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, DeleteLabelOutput{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	httpReq, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/labels/%s", homeboxURL, input.ID), nil)
	if err != nil {
		return nil, DeleteLabelOutput{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, DeleteLabelOutput{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return nil, DeleteLabelOutput{}, fmt.Errorf("failed to delete label, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	return nil, DeleteLabelOutput{}, nil
}

// getMaintenanceLog is the implementation of the "get_maintenance_log" tool.
func getMaintenanceLog(ctx context.Context, req *mcp.CallToolRequest, input GetMaintenanceLogInput) (*mcp.CallToolResult, GetMaintenanceLogOutput, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, GetMaintenanceLogOutput{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	httpReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/items/%s/maintenance", homeboxURL, input.ItemID), nil)
	if err != nil {
		return nil, GetMaintenanceLogOutput{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, GetMaintenanceLogOutput{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, GetMaintenanceLogOutput{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, GetMaintenanceLogOutput{}, fmt.Errorf("failed to get maintenance log, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var entries []MaintenanceEntryWithDetails
	if err := json.Unmarshal(body, &entries); err != nil {
		return nil, GetMaintenanceLogOutput{}, err
	}

	return nil, GetMaintenanceLogOutput{Entries: entries}, nil
}

// createMaintenanceEntry is the implementation of the "create_maintenance_entry" tool.
func createMaintenanceEntry(ctx context.Context, req *mcp.CallToolRequest, input CreateMaintenanceEntryInput) (*mcp.CallToolResult, MaintenanceEntry, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, MaintenanceEntry{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	type MaintenanceEntryCreate struct {
		Name          string `json:"name" jsonschema:"required"`
		CompletedDate string `json:"completedDate,omitempty"`
		Cost          string `json:"cost,omitempty"`
		Description   string `json:"description,omitempty"`
		ScheduledDate string `json:"scheduledDate,omitempty"`
	}

	reqPayload := MaintenanceEntryCreate{
		Name:          input.Name,
		CompletedDate: input.CompletedDate,
		Cost:          input.Cost,
		Description:   input.Description,
		ScheduledDate: input.ScheduledDate,
	}

	reqBody, err := json.Marshal(reqPayload)
	if err != nil {
		return nil, MaintenanceEntry{}, err
	}

	httpReq, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/items/%s/maintenance", homeboxURL, input.ItemID), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, MaintenanceEntry{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, MaintenanceEntry{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, MaintenanceEntry{}, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, MaintenanceEntry{}, fmt.Errorf("failed to create maintenance entry, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var entry MaintenanceEntry
	if err := json.Unmarshal(body, &entry); err != nil {
		return nil, MaintenanceEntry{}, err
	}

	return nil, entry, nil
}

// duplicateItem is the implementation of the "duplicate_item" tool.
func duplicateItem(ctx context.Context, req *mcp.CallToolRequest, input DuplicateItemInput) (*mcp.CallToolResult, ItemOut, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, ItemOut{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	reqBody, err := json.Marshal(input)
	if err != nil {
		return nil, ItemOut{}, err
	}

	httpReq, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/items/%s/duplicate", homeboxURL, input.ID), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, ItemOut{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, ItemOut{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ItemOut{}, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, ItemOut{}, fmt.Errorf("failed to duplicate item, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var itemOut ItemOut
	if err := json.Unmarshal(body, &itemOut); err != nil {
		return nil, ItemOut{}, err
	}

	return nil, itemOut, nil
}

// getItemPath is the implementation of the "get_item_path" tool.
func getItemPath(ctx context.Context, req *mcp.CallToolRequest, input GetItemPathInput) (*mcp.CallToolResult, GetItemPathOutput, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, GetItemPathOutput{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	httpReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/items/%s/path", homeboxURL, input.ID), nil)
	if err != nil {
		return nil, GetItemPathOutput{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, GetItemPathOutput{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, GetItemPathOutput{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, GetItemPathOutput{}, fmt.Errorf("failed to get item path, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var itemPath []ItemPath
	if err := json.Unmarshal(body, &itemPath); err != nil {
		return nil, GetItemPathOutput{}, err
	}

	return nil, GetItemPathOutput{Path: itemPath}, nil
}

// exportItems is the implementation of the "export_items" tool.
func exportItems(ctx context.Context, req *mcp.CallToolRequest, input ExportItemsInput) (*mcp.CallToolResult, ExportItemsOutput, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, ExportItemsOutput{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	httpReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/items/export", homeboxURL), nil)
	if err != nil {
		return nil, ExportItemsOutput{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, ExportItemsOutput{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ExportItemsOutput{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ExportItemsOutput{}, fmt.Errorf("failed to export items, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	return nil, ExportItemsOutput{CSVData: string(body)}, nil
}

// importItems is the implementation of the "import_items" tool.
func importItems(ctx context.Context, req *mcp.CallToolRequest, input ImportItemsInput) (*mcp.CallToolResult, ImportItemsOutput, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")
	if homeboxURL == "" || homeboxToken == "" {
		return nil, ImportItemsOutput{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN must be set")
	}

	// Decode the base64 file content
	fileContent, err := base64.StdEncoding.DecodeString(input.FileContent)
	if err != nil {
		return nil, ImportItemsOutput{}, fmt.Errorf("failed to decode file content: %w", err)
	}

	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a new form file
	part, err := writer.CreateFormFile("file", input.FileName)
	if err != nil {
		return nil, ImportItemsOutput{}, fmt.Errorf("failed to create form file: %w", err)
	}

	// Write the file content to the form file
	_, err = io.Copy(part, bytes.NewReader(fileContent))
	if err != nil {
		return nil, ImportItemsOutput{}, fmt.Errorf("failed to write file content to form file: %w", err)
	}

	// Close the writer
	err = writer.Close()
	if err != nil {
		return nil, ImportItemsOutput{}, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	// Create the HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, "POST", homeboxURL+"/api/v1/items/import", body)
	if err != nil {
		return nil, ImportItemsOutput{}, fmt.Errorf("failed to create request: %w", err)
	}

	// Set the content type and auth headers
	httpReq.Header.Set("Content-Type", writer.FormDataContentType())
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, ImportItemsOutput{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, ImportItemsOutput{}, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(respBody))
	}

	// Decode the response body
	var result ActionAmountResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, ImportItemsOutput{}, fmt.Errorf("failed to decode response body: %w", err)
	}

	// Return the output
	return nil, ImportItemsOutput{
		Completed: result.Completed,
	}, nil
}

// getItemFields is the implementation of the "get_item_fields" tool.
func getItemFields(ctx context.Context, req *mcp.CallToolRequest, input GetItemFieldsInput) (*mcp.CallToolResult, GetItemFieldsOutput, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, GetItemFieldsOutput{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	httpReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/items/fields", homeboxURL), nil)
	if err != nil {
		return nil, GetItemFieldsOutput{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, GetItemFieldsOutput{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, GetItemFieldsOutput{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, GetItemFieldsOutput{}, fmt.Errorf("failed to get item fields, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var fields []string
	if err := json.Unmarshal(body, &fields); err != nil {
		return nil, GetItemFieldsOutput{}, err
	}

	return nil, GetItemFieldsOutput{Fields: fields}, nil
}

// getItemFieldValues is the implementation of the "get_item_field_values" tool.
func getItemFieldValues(ctx context.Context, req *mcp.CallToolRequest, input GetItemFieldValuesInput) (*mcp.CallToolResult, GetItemFieldValuesOutput, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, GetItemFieldValuesOutput{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	httpReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/items/fields/values", homeboxURL), nil)
	if err != nil {
		return nil, GetItemFieldValuesOutput{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, GetItemFieldValuesOutput{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, GetItemFieldValuesOutput{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, GetItemFieldValuesOutput{}, fmt.Errorf("failed to get item field values, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var values []string
	if err := json.Unmarshal(body, &values); err != nil {
		return nil, GetItemFieldValuesOutput{}, err
	}

	return nil, GetItemFieldValuesOutput{Values: values}, nil
}

// getItemByAssetID is the implementation of the "get_item_by_asset_id" tool.
func getItemByAssetID(ctx context.Context, req *mcp.CallToolRequest, input GetItemByAssetIDInput) (*mcp.CallToolResult, PaginationResult_ItemSummary, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return nil, PaginationResult_ItemSummary{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	httpReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/assets/%s", homeboxURL, input.ID), nil)
	if err != nil {
		return nil, PaginationResult_ItemSummary{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, PaginationResult_ItemSummary{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, PaginationResult_ItemSummary{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, PaginationResult_ItemSummary{}, fmt.Errorf("failed to get item by asset id, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var result PaginationResult_ItemSummary
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, PaginationResult_ItemSummary{}, err
	}

	return nil, result, nil
}

// createMissingThumbnails is the implementation of the "create_missing_thumbnails" tool.
func createMissingThumbnails(ctx context.Context, req *mcp.CallToolRequest, input CreateMissingThumbnailsInput) (*mcp.CallToolResult, ActionAmountResult, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")
	if homeboxURL == "" || homeboxToken == "" {
		return nil, ActionAmountResult{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN must be set")
	}
	httpReq, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/actions/create-missing-thumbnails", homeboxURL), nil)
	if err != nil {
		return nil, ActionAmountResult{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, ActionAmountResult{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ActionAmountResult{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ActionAmountResult{}, fmt.Errorf("failed to create missing thumbnails, status code: %d, body: %s", resp.StatusCode, string(body))
	}
	var result ActionAmountResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, ActionAmountResult{}, err
	}
	return nil, result, nil
}

// ensureAssetIDs is the implementation of the "ensure_asset_ids" tool.
func ensureAssetIDs(ctx context.Context, req *mcp.CallToolRequest, input EnsureAssetIDsInput) (*mcp.CallToolResult, ActionAmountResult, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")
	if homeboxURL == "" || homeboxToken == "" {
		return nil, ActionAmountResult{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN must be set")
	}
	httpReq, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/actions/ensure-asset-ids", homeboxURL), nil)
	if err != nil {
		return nil, ActionAmountResult{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, ActionAmountResult{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ActionAmountResult{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ActionAmountResult{}, fmt.Errorf("failed to ensure asset ids, status code: %d, body: %s", resp.StatusCode, string(body))
	}
	var result ActionAmountResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, ActionAmountResult{}, err
	}
	return nil, result, nil
}

// ensureImportRefs is the implementation of the "ensure_import_refs" tool.
func ensureImportRefs(ctx context.Context, req *mcp.CallToolRequest, input EnsureImportRefsInput) (*mcp.CallToolResult, ActionAmountResult, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")
	if homeboxURL == "" || homeboxToken == "" {
		return nil, ActionAmountResult{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN must be set")
	}
	httpReq, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/actions/ensure-import-refs", homeboxURL), nil)
	if err != nil {
		return nil, ActionAmountResult{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, ActionAmountResult{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ActionAmountResult{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ActionAmountResult{}, fmt.Errorf("failed to ensure import refs, status code: %d, body: %s", resp.StatusCode, string(body))
	}
	var result ActionAmountResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, ActionAmountResult{}, err
	}
	return nil, result, nil
}

// setPrimaryPhotos is the implementation of the "set_primary_photos" tool.
func setPrimaryPhotos(ctx context.Context, req *mcp.CallToolRequest, input SetPrimaryPhotosInput) (*mcp.CallToolResult, ActionAmountResult, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")
	if homeboxURL == "" || homeboxToken == "" {
		return nil, ActionAmountResult{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN must be set")
	}
	httpReq, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/actions/set-primary-photos", homeboxURL), nil)
	if err != nil {
		return nil, ActionAmountResult{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, ActionAmountResult{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ActionAmountResult{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ActionAmountResult{}, fmt.Errorf("failed to set primary photos, status code: %d, body: %s", resp.StatusCode, string(body))
	}
	var result ActionAmountResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, ActionAmountResult{}, err
	}
	return nil, result, nil
}

// zeroItemTimeFields is the implementation of the "zero_item_time_fields" tool.
func zeroItemTimeFields(ctx context.Context, req *mcp.CallToolRequest, input ZeroItemTimeFieldsInput) (*mcp.CallToolResult, ActionAmountResult, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")
	if homeboxURL == "" || homeboxToken == "" {
		return nil, ActionAmountResult{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN must be set")
	}
	httpReq, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/actions/zero-item-time-fields", homeboxURL), nil)
	if err != nil {
		return nil, ActionAmountResult{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, ActionAmountResult{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ActionAmountResult{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ActionAmountResult{}, fmt.Errorf("failed to zero item time fields, status code: %d, body: %s", resp.StatusCode, string(body))
	}
	var result ActionAmountResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, ActionAmountResult{}, err
	}
	return nil, result, nil
}

// getStatus is the implementation of the "get_status" tool.
func getStatus(ctx context.Context, req *mcp.CallToolRequest, input GetStatusInput) (*mcp.CallToolResult, APISummary, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")
	if homeboxURL == "" || homeboxToken == "" {
		return nil, APISummary{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN must be set")
	}
	httpReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/status", homeboxURL), nil)
	if err != nil {
		return nil, APISummary{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, APISummary{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, APISummary{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, APISummary{}, fmt.Errorf("failed to get status, status code: %d, body: %s", resp.StatusCode, string(body))
	}
	var result APISummary
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, APISummary{}, err
	}
	return nil, result, nil
}

// getCurrency is the implementation of the "get_currency" tool.
func getCurrency(ctx context.Context, req *mcp.CallToolRequest, input GetCurrencyInput) (*mcp.CallToolResult, Currency, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")
	if homeboxURL == "" || homeboxToken == "" {
		return nil, Currency{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN must be set")
	}
	httpReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/currency", homeboxURL), nil)
	if err != nil {
		return nil, Currency{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, Currency{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, Currency{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, Currency{}, fmt.Errorf("failed to get currency, status code: %d, body: %s", resp.StatusCode, string(body))
	}
	var result Currency
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, Currency{}, err
	}
	return nil, result, nil
}

// createGroupInvitation is the implementation of the "create_group_invitation" tool.
func createGroupInvitation(ctx context.Context, req *mcp.CallToolRequest, input CreateGroupInvitationInput) (*mcp.CallToolResult, GroupInvitation, error) {
    homeboxURL := os.Getenv("HOMEBOX_URL")
    homeboxToken := os.Getenv("HOMEBOX_TOKEN")

    if homeboxURL == "" || homeboxToken == "" {
        return nil, GroupInvitation{}, fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
    }

    // Marshal the input to JSON
    reqBody, err := json.Marshal(input)
    if err != nil {
        return nil, GroupInvitation{}, err
    }

    // Create a new HTTP request to the Homebox API.
    httpReq, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/groups/invitations", homeboxURL), bytes.NewBuffer(reqBody))
    if err != nil {
        return nil, GroupInvitation{}, err
    }
    httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)
    httpReq.Header.Set("Content-Type", "application/json")

    // Execute the request.
    client := &http.Client{}
    resp, err := client.Do(httpReq)
    if err != nil {
        return nil, GroupInvitation{}, err
    }
    defer resp.Body.Close()

    // Read the response body.
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, GroupInvitation{}, err
    }

    if resp.StatusCode != http.StatusCreated {
        return nil, GroupInvitation{}, fmt.Errorf("failed to create group invitation, status code: %d, body: %s", resp.StatusCode, string(body))
    }

    // Unmarshal the JSON response into a GroupInvitation struct.
    var invitation GroupInvitation
    if err := json.Unmarshal(body, &invitation); err != nil {
        return nil, GroupInvitation{}, err
    }

    // Return the created invitation.
    return nil, invitation, nil
}

// Label Maker Inputs
type GetAssetLabelInput struct {
	ID string `json:"id" jsonschema:"required"`
}

type GetItemLabelInput struct {
	ID string `json:"id" jsonschema:"required"`
}

type GetLocationLabelInput struct {
	ID string `json:"id" jsonschema:"required"`
}

// Label Maker Output
type GetLabelOutput struct {
	Image string `json:"image" jsonschema:"required,description:Base64 encoded image data"`
}

// getLabelImage is a helper function to get a label from a given path.
func getLabelImage(path string) (string, error) {
	homeboxURL := os.Getenv("HOMEBOX_URL")
	homeboxToken := os.Getenv("HOMEBOX_TOKEN")

	if homeboxURL == "" || homeboxToken == "" {
		return "", fmt.Errorf("HOMEBOX_URL and HOMEBOX_TOKEN environment variables must be set")
	}

	httpReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/%s", homeboxURL, path), nil)
	if err != nil {
		return "", err
	}
	httpReq.Header.Set("Authorization", "Bearer "+homeboxToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to get label, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(body), nil
}

// getAssetLabel is the implementation of the "get_asset_label" tool.
func getAssetLabel(ctx context.Context, req *mcp.CallToolRequest, input GetAssetLabelInput) (*mcp.CallToolResult, GetLabelOutput, error) {
	image, err := getLabelImage(fmt.Sprintf("labelmaker/assets/%s", input.ID))
	if err != nil {
		return nil, GetLabelOutput{}, err
	}
	return nil, GetLabelOutput{Image: image}, nil
}

// getItemLabel is the implementation of the "get_item_label" tool.
func getItemLabel(ctx context.Context, req *mcp.CallToolRequest, input GetItemLabelInput) (*mcp.CallToolResult, GetLabelOutput, error) {
	image, err := getLabelImage(fmt.Sprintf("labelmaker/item/%s", input.ID))
	if err != nil {
		return nil, GetLabelOutput{}, err
	}
	return nil, GetLabelOutput{Image: image}, nil
}

// getLocationLabel is the implementation of the "get_location_label" tool.
func getLocationLabel(ctx context.Context, req *mcp.CallToolRequest, input GetLocationLabelInput) (*mcp.CallToolResult, GetLabelOutput, error) {
	image, err := getLabelImage(fmt.Sprintf("labelmaker/location/%s", input.ID))
	if err != nil {
		return nil, GetLabelOutput{}, err
	}
	return nil, GetLabelOutput{Image: image}, nil
}

func main() {
	// Create a new MCP server.
	server := mcp.NewServer(&mcp.Implementation{Name: "homebox-mcp-server", Version: "v0.0.1"}, nil)

	// Item tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_items",
		Description: "Retrieves all items from the Homebox inventory.",
	}, getItems)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_item",
		Description: "Creates a new item in the Homebox inventory.",
	}, createItem)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_item",
		Description: "Retrieves a single item from the Homebox inventory by its ID.",
	}, getItem)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_item",
		Description: "Updates an existing item in the Homebox inventory.",
	}, updateItem)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_item",
		Description: "Deletes an item from the Homebox inventory.",
	}, deleteItem)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "duplicate_item",
		Description: "Duplicates an existing item.",
	}, duplicateItem)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_item_path",
		Description: "Retrieves the path of an item.",
	}, getItemPath)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "export_items",
		Description: "Exports all items as a CSV string.",
	}, exportItems)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "import_items",
		Description: "Imports items from a CSV file.",
	}, importItems)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_item_fields",
		Description: "Gets all custom field names.",
	}, getItemFields)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_item_field_values",
		Description: "Gets all custom field values.",
	}, getItemFieldValues)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_item_by_asset_id",
		Description: "Retrieves an item by its asset ID.",
	}, getItemByAssetID)

	// Location tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_locations",
		Description: "Retrieves all locations from the Homebox inventory.",
	}, getLocations)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_location",
		Description: "Creates a new location in the Homebox inventory.",
	}, createLocation)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_location",
		Description: "Retrieves a single location from the Homebox inventory by its ID.",
	}, getLocation)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_location",
		Description: "Updates an existing location in the Homebox inventory.",
	}, updateLocation)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_location",
		Description: "Deletes a location from the Homebox inventory.",
	}, deleteLocation)

	// Label tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_labels",
		Description: "Retrieves all labels from the Homebox inventory.",
	}, getLabels)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_label",
		Description: "Creates a new label in the Homebox inventory.",
	}, createLabel)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_label",
		Description: "Retrieve a single label from the Homebox inventory by its ID.",
	}, getLabel)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_label",
		Description: "Updates an existing label in the Homebox inventory.",
	}, updateLabel)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_label",
		Description: "Deletes a label from the Homebox inventory.",
	}, deleteLabel)

	// Label Maker tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_asset_label",
		Description: "Gets a label for an asset.",
	}, getAssetLabel)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_item_label",
		Description: "Gets a label for an item.",
	}, getItemLabel)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_location_label",
		Description: "Gets a label for a location.",
	}, getLocationLabel)

	// Maintenance tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_maintenance_log",
		Description: "Retrieves the maintenance log for a specific item.",
	}, getMaintenanceLog)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_maintenance_entry",
		Description: "Creates a new maintenance entry for an item.",
	}, createMaintenanceEntry)

	// Action tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_missing_thumbnails",
		Description: "Creates thumbnails for items that are missing them.",
	}, createMissingThumbnails)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "ensure_asset_ids",
		Description: "Ensures all items in the database have an asset ID.",
	}, ensureAssetIDs)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "ensure_import_refs",
		Description: "Ensures all items in the database have an import ref.",
	}, ensureImportRefs)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "set_primary_photos",
		Description: "Sets the first photo of each item as the primary photo.",
	}, setPrimaryPhotos)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "zero_item_time_fields",
		Description: "Resets all item date fields to the beginning of the day.",
	}, zeroItemTimeFields)

	// Status and Currency tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_status",
		Description: "Gets application status information.",
	}, getStatus)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_currency",
		Description: "Gets currency information.",
	}, getCurrency)
	
    // Group tools
	mcp.AddTool(server, &mcp.Tool{
	    Name:        "create_group_invitation",
	    Description: "Creates a new group invitation.",
	}, createGroupInvitation)

	// Start the server, which will listen for connections on stdin/stdout.
	log.Println("Starting Homebox MCP server...")
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("MCP server error: %v", err)
	}
}
