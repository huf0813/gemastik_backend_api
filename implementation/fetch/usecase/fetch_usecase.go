package usecase

import "github.com/huf0813/gemastik_api_backend_supabase/domain"

type FetchUseCase struct {
	FetchRepositorySupabase domain.FetchRepository
}

func NewFetchUseCase(f domain.FetchRepository) domain.FetchUseCase {
	return &FetchUseCase{FetchRepositorySupabase: f}
}

func (f *FetchUseCase) FetchValue(table string, queries map[string]string) ([]map[string]interface{}, error) {
	result, err := f.FetchRepositorySupabase.FetchValue(table, queries)
	if err != nil {
		return nil, err
	}
	return result, nil
}
