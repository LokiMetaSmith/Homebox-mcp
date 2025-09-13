# MCP Server TODO

This file lists the remaining endpoints from the Homebox API specification that have not yet been implemented in the MCP server.

## Item Attachments
- [ ] `POST /v1/items/{id}/attachments` (Create Item Attachment - requires file upload)
- [ ] `GET /v1/items/{id}/attachments/{attachment_id}` (Get Item Attachment)
- [ ] `PUT /v1/items/{id}/attachments/{attachment_id}` (Update Item Attachment)
- [ ] `DELETE /v1/items/{id}/attachments/{attachment_id}` (Delete Item Attachment)

## Item Import
- [ ] `POST /v1/items/import` (Import Items - requires file upload)

## Groups
- [ ] `GET /v1/groups`
- [ ] `PUT /v1/groups`
- [ ] `POST /v1/groups/invitations`
- [ ] `GET /v1/groups/statistics`
- [ ] `GET /v1/groups/statistics/labels`
- [ ] `GET /v1/groups/statistics/locations`
- [ ] `GET /v1/groups/statistics/purchase-price`

## Label Maker
- [ ] `GET /v1/labelmaker/assets/{id}` (returns image)
- [ ] `GET /v1/labelmaker/item/{id}` (returns image)
- [ ] `GET /v1/labelmaker/location/{id}` (returns image)

## Maintenance
- [ ] `PUT /v1/maintenance/{id}`
- [ ] `DELETE /v1/maintenance/{id}`

## Notifiers
- [ ] `GET /v1/notifiers`
- [ ] `POST /v1/notifiers`
- [ ] `POST /v1/notifiers/test`
- [ ] `PUT /v1/notifiers/{id}`
- [ ] `DELETE /v1/notifiers/{id}`

## Products
- [ ] `GET /v1/products/search-from-barcode`

## QR Code
- [ ] `GET /v1/qrcode` (returns image)

## Reporting
- [ ] `GET /v1/reporting/bill-of-materials`

## Users
- [ ] `PUT /v1/users/change-password`
- [ ] `GET /v1/users/self`
- [ ] `PUT /v1/users/self`
- [ ] `DELETE /v1/users/self`
- [ ] `POST /v1/users/register`
