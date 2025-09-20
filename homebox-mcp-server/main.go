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
	"strconv"

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

// Group Inputs
type GetGroupInput struct{}
type UpdateGroupInput struct {
	Name     string `json:"name,omitempty"`
	Currency string `json:"currency,omitempty"`
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

//... (rest of the file is the same)
