package repository

import (
	"context"

	"github.com/qwikshelf/api/internal/domain/entity"
)

// PincodeRepository defines the interface for managing serviceable areas.
type PincodeRepository interface {
	GetByPincode(ctx context.Context, pincode string) (*entity.ServiceableArea, error)
	List(ctx context.Context) ([]entity.ServiceableArea, error)
	Update(ctx context.Context, area *entity.ServiceableArea) error
}
