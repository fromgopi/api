package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/qwikshelf/api/internal/domain/entity"
	"github.com/qwikshelf/api/internal/domain/repository"
)

type pincodeRepository struct {
	db *DB
}

// NewPincodeRepository creates a new instance of PincodeRepository.
func NewPincodeRepository(db *DB) repository.PincodeRepository {
	return &pincodeRepository{db: db}
}

func (r *pincodeRepository) GetByPincode(ctx context.Context, pincode string) (*entity.ServiceableArea, error) {
	query := `
		SELECT id, pincode, warehouse_id, is_active, min_order_amount, delivery_charge, estimated_delivery_text, created_at
		FROM serviceable_pincodes
		WHERE pincode = $1 AND is_active = true
	`
	var area entity.ServiceableArea
	err := r.db.Pool.QueryRow(ctx, query, pincode).Scan(
		&area.ID,
		&area.Pincode,
		&area.WarehouseID,
		&area.IsActive,
		&area.MinOrderAmount,
		&area.DeliveryCharge,
		&area.EstimatedDeliveryText,
		&area.CreatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get serviceable pincode: %w", err)
	}
	return &area, nil
}

func (r *pincodeRepository) List(ctx context.Context) ([]entity.ServiceableArea, error) {
	query := `
		SELECT id, pincode, warehouse_id, is_active, min_order_amount, delivery_charge, estimated_delivery_text, created_at
		FROM serviceable_pincodes
		ORDER BY pincode ASC
	`
	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list serviceable pincodes: %w", err)
	}
	defer rows.Close()

	var areas []entity.ServiceableArea
	for rows.Next() {
		var area entity.ServiceableArea
		if err := rows.Scan(
			&area.ID,
			&area.Pincode,
			&area.WarehouseID,
			&area.IsActive,
			&area.MinOrderAmount,
			&area.DeliveryCharge,
			&area.EstimatedDeliveryText,
			&area.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan serviceable pincode: %w", err)
		}
		areas = append(areas, area)
	}
	return areas, nil
}

func (r *pincodeRepository) Update(ctx context.Context, area *entity.ServiceableArea) error {
	query := `
		UPDATE serviceable_pincodes
		SET warehouse_id = $1, is_active = $2, min_order_amount = $3, delivery_charge = $4, estimated_delivery_text = $5
		WHERE pincode = $6
	`
	_, err := r.db.Pool.Exec(ctx, query,
		area.WarehouseID,
		area.IsActive,
		area.MinOrderAmount,
		area.DeliveryCharge,
		area.EstimatedDeliveryText,
		area.Pincode,
	)
	if err != nil {
		return fmt.Errorf("failed to update serviceable pincode: %w", err)
	}
	return nil
}
