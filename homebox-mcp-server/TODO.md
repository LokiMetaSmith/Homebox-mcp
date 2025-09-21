# MCP Server TODO

This file lists the remaining endpoints from the Homebox API specification that have not yet been implemented in the MCP server.

## Item Attachments
- [x] `POST /v1/items/{id}/attachments` (Create Item Attachment - requires file upload)
- [x] `GET /v1/items/{id}/attachments/{attachment_id}` (Get Item Attachment)
- [x] `PUT /v1/items/{id}/attachments/{attachment_id}` (Update Item Attachment)
- [x] `DELETE /v1/items/{id}/attachments/{attachment_id}` (Delete Item Attachment)

## Item Import
- [x] `POST /v1/items/import` (Import Items - requires file upload)

## Groups
- [x] `GET /v1/groups`
- [x] `PUT /v1/groups`
- [x] `POST /v1/groups/invitations`
- [x] `GET /v1/groups/statistics`
- [x] `GET /v1/groups/statistics/labels`
- [x] `GET /v1/groups/statistics/locations`
- [x] `GET /v1/groups/statistics/purchase-price`

## Label Maker
- [x] `GET /v1/labelmaker/assets/{id}` (returns image)
- [x] `GET /v1/labelmaker/item/{id}` (returns image)
- [x] `GET /v1/labelmaker/location/{id}` (returns image)

## Maintenance
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

## Users
- [ ] `PUT /v1/users/change-password`
- [ ] `GET /v1/users/self`
- [ ] `PUT /v1/users/self`
- [ ] `DELETE /v1/users/self`
- [ ] `POST /v1/users/register`
