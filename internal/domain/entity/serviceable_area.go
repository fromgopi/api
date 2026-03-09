package entity

import "time"

// ServiceableArea represents a geographical area (pincode) where delivery is supported.
type ServiceableArea struct {
	ID                    int64     `json:"id"`
	Pincode               string    `json:"pincode"`
	WarehouseID           *int64    `json:"warehouse_id"` // Normalized to Nullable int64 in case warehouse is deleted
	IsActive              bool      `json:"is_active"`
	MinOrderAmount        float64   `json:"min_order_amount"`
	DeliveryCharge        float64   `json:"delivery_charge"`
	EstimatedDeliveryText string    `json:"estimated_delivery_text"`
	CreatedAt             time.Time `json:"created_at"`
}
