package usecase

import (
	"errors"
	"fmt"
	"github.com/huf0813/gemastik_api_backend_supabase/domain"
	"github.com/huf0813/gemastik_api_backend_supabase/infrastructure"
	"github.com/huf0813/gemastik_api_backend_supabase/middleware"
	"github.com/huf0813/gemastik_api_backend_supabase/utility"
)

type UserUseCase struct {
	UserRepositorySupabase  domain.UserRepository
	FetchRepositorySupabase domain.FetchRepository
	AppDriver               *infrastructure.DriverAppService
}

func NewUserUseCase(u domain.UserRepository, f domain.FetchRepository, a *infrastructure.DriverAppService) domain.UserUseCase {
	return &UserUseCase{UserRepositorySupabase: u, FetchRepositorySupabase: f, AppDriver: a}
}

func (u *UserUseCase) SignUp(su *domain.UserSignUp) error {
	hashedPassword, err := utility.NewHashValue(su.Password)
	if err != nil {
		return err
	}
	value := map[string]string{
		"name":            su.Name,
		"email":           su.Email,
		"identity_number": su.IdentityNumber,
		"password":        hashedPassword,
	}
	if err := u.UserRepositorySupabase.SignUp(value); err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) SignIn(su *domain.UserSignIn) (string, error) {
	queries := map[string]string{
		"select": "email,password,name,id,suppliers(id)",
		"email":  fmt.Sprintf("eq.%s", su.Email),
	}
	result, err := u.FetchRepositorySupabase.FetchValue("users", queries)
	if err != nil {
		return "", err
	}
	if len(result) <= 0 {
		return "", errors.New("row is empty")
	}

	temp := result[0]
	password := temp["password"].(string)
	if err := utility.NewCompareValue(password, su.Password); err != nil {
		return "", err
	}

	name := temp["name"].(string)
	userID := int64(temp["id"].(float64))
	token, err := middleware.NewToken(u.AppDriver.Secret, name, userID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Bearer %s", token), nil
}

func (u *UserUseCase) ClaimSupplier(cs *domain.UserClaimSupplier, userID int64) error {
	value := map[string]string{
		"name":        cs.Name,
		"address":     cs.Address,
		"phone":       cs.Phone,
		"description": cs.Description,
		"thumbnail":   cs.Thumbnail,
		"user_id":     fmt.Sprintf("%d", userID),
	}
	if err := u.UserRepositorySupabase.ClaimSupplier(value); err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) GetProfile(userID int64, selectedColumns string) (map[string]interface{}, error) {
	queries := map[string]string{
		"select": selectedColumns,
		"id":     fmt.Sprintf("eq.%d", userID),
	}
	result, err := u.FetchRepositorySupabase.FetchValue("users", queries)
	if err != nil {
		return nil, err
	}
	if len(result) <= 0 {
		return nil, errors.New("row is empty")
	}

	return result[0], nil
}
