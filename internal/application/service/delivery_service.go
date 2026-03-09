package service

import (
	"context"

	"github.com/qwikshelf/api/internal/domain/entity"
	"github.com/qwikshelf/api/internal/domain/repository"
)

type DeliveryService struct {
	pincodeRepo repository.PincodeRepository
}

// NewDeliveryService creates a new instance of DeliveryService.
func NewDeliveryService(pincodeRepo repository.PincodeRepository) *DeliveryService {
	return &DeliveryService{pincodeRepo: pincodeRepo}
}

// CheckServiceability returns the area details if the pincode is serviceable.
func (s *DeliveryService) CheckServiceability(ctx context.Context, pincode string) (*entity.ServiceableArea, error) {
	return s.pincodeRepo.GetByPincode(ctx, pincode)
}

// ListServiceableAreas returns all serviceable areas.
func (s *DeliveryService) ListServiceableAreas(ctx context.Context) ([]entity.ServiceableArea, error) {
	return s.pincodeRepo.List(ctx)
}

// UpdateServiceableArea updates an existing serviceable area's configuration.
func (s *DeliveryService) UpdateServiceableArea(ctx context.Context, area *entity.ServiceableArea) error {
	return s.pincodeRepo.Update(ctx, area)
}
