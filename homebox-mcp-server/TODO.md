# MCP Server TODO

This file lists the remaining endpoints from the Homebox API specification that have not yet been implemented in the MCP server.

## Items
- [x] `GET /v1/items`
- [x] `POST /v1/items`
- [x] `GET /v1/items/{id}`
- [x] `PUT /v1/items/{id}`
- [x] `DELETE /v1/items/{id}`
- [x] `POST /v1/items/{id}/duplicate`
- [x] `GET /v1/items/{id}/path`
- [x] `GET /v1/items/export`
- [x] `GET /v1/items/by-asset-id/{id}`
- [x] `GET /v1/items/fields`
- [x] `GET /v1/items/fields/{name}`

## Item Attachments
- [x] `POST /v1/items/{id}/attachments` (Create Item Attachment - requires file upload)
- [x] `GET /v1/items/{id}/attachments/{attachment_id}` (Get Item Attachment)
- [x] `PUT /v1/items/{id}/attachments/{attachment_id}` (Update Item Attachment)
- [x] `DELETE /v1/items/{id}/attachments/{attachment_id}` (Delete Item Attachment)

## Item Import
- [ ] `POST /v1/items/import` (Import Items - requires file upload)

## Locations
- [x] `GET /v1/locations`
- [x] `POST /v1/locations`
- [x] `GET /v1/locations/{id}`
- [x] `PUT /v1/locations/{id}`
- [x] `DELETE /v1/locations/{id}`

## Labels
- [x] `GET /v1/labels`
- [x] `POST /v1/labels`
- [x] `GET /v1/labels/{id}`
- [x] `PUT /v1/labels/{id}`
- [x] `DELETE /v1/labels/{id}`

## Groups
- [x] `GET /v1/groups`
- [x] `PUT /v1/groups`
- [ ] `POST /v1/groups/invitations`
- [x] `GET /v1/groups/statistics`
- [x] `GET /v1/groups/statistics/labels`
- [x] `GET /v1/groups/statistics/locations`
- [x] `GET /v1/groups/statistics/purchase-price`

## Label Maker
- [ ] `GET /v1/labelmaker/assets/{id}` (returns image)
- [ ] `GET /v1/labelmaker/item/{id}` (returns image)
- [ ] `GET /v1/labelmaker/location/{id}` (returns image)

## Maintenance
- [x] `GET /v1/items/{item_id}/maintenance`
- [x] `POST /v1/items/{item_id}/maintenance`
- [x] `PUT /v1/maintenance/{id}`
- [x] `DELETE /v1/maintenance/{id}`

## Notifiers
- [x] `GET /v1/notifiers`
- [x] `POST /v1/notifiers`
- [x] `POST /v1/notifiers/test`
- [x] `PUT /v1/notifiers/{id}`
- [x] `DELETE /v1/notifiers/{id}`

## Products
- [x] `GET /v1/products/search-from-barcode`

## QR Code
- [x] `GET /v1/qrcode` (returns image)

## Reporting
- [x] `GET /v1/reporting/bill-of-materials`

## Actions
- [x] `POST /v1/actions/create-missing-thumbnails`
- [x] `POST /v1/actions/ensure-asset-ids`
- [x] `POST /v1/actions/ensure-import-refs`
- [x] `POST /v1/actions/set-primary-photos`
- [x] `POST /v1/actions/zero-item-time-fields`

## Users
- [ ] `PUT /v1/users/change-password`
- [ ] `GET /v1/users/self`
- [ ] `PUT /v1/users/self`
- [ ] `DELETE /v1/users/self`
- [ ] `POST /v1/users/register`

## System
- [x] `GET /v1/status`
- [x] `GET /v1/currency`
