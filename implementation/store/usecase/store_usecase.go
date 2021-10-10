package usecase

import (
	"fmt"
	"github.com/huf0813/gemastik_api_backend_supabase/domain"
)

type StoreUseCase struct {
	StoreRepositorySupabase domain.StoreRepository
}

func NewStoreUseCase(s domain.StoreRepository) domain.StoreUseCase {
	return &StoreUseCase{StoreRepositorySupabase: s}
}

func (s *StoreUseCase) ClaimStore(c *domain.CreateStoreRequest, userID int64) error {
	value := map[string]string{
		"name":        c.Name,
		"description": c.Description,
		"address":     c.Address,
		"phone":       c.Phone,
		"thumbnail":   c.Thumbnail,
		"time_closed": c.TimeClosed,
		"time_opened": c.TimeOpened,
		"user_id":     fmt.Sprintf("%d", userID),
	}
	if err := s.StoreRepositorySupabase.ClaimStore(value); err != nil {
		return err
	}

	return nil
}
