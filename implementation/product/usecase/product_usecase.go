package usecase

import (
	"fmt"
	"github.com/huf0813/gemastik_api_backend_supabase/domain"
)

type ProductUseCase struct {
	ProductRepositorySupabase domain.ProductRepository
	FetchRepositorySupabase   domain.FetchRepository
}

func NewProductUseCase(p domain.ProductRepository, f domain.FetchRepository) domain.ProductUseCase {
	return &ProductUseCase{ProductRepositorySupabase: p, FetchRepositorySupabase: f}
}

func (p *ProductUseCase) CreateProduct(create *domain.CreateProductRequest, userID int64) error {
	queries := map[string]string{
		"select":  "id",
		"user_id": fmt.Sprintf("eq.%d", userID),
	}
	supplier, err := p.FetchRepositorySupabase.FetchValue("suppliers", queries)
	if err != nil {
		return err
	}
	supplierID := int64(0)
	if len(supplier) > 0 {
		tempSupplier := supplier[0]
		supplierID = int64(tempSupplier["id"].(float64))
	}

	value := map[string]string{
		"name":            create.Name,
		"quality":         create.Quality,
		"description":     create.Description,
		"thumbnail":       create.Thumbnail,
		"price":           fmt.Sprintf("%d", create.Price),
		"product_type_id": fmt.Sprintf("%d", create.ProductTypeID),
		"supplier_id":     fmt.Sprintf("%d", supplierID),
	}
	if err := p.ProductRepositorySupabase.CreateProduct(value); err != nil {
		return err
	}

	return nil
}

func (p *ProductUseCase) UpdateProduct(update *domain.CreateProductRequest, productID, userID int64) error {
	queries := map[string]string{
		"select":  "id",
		"user_id": fmt.Sprintf("eq.%d", userID),
	}
	supplier, err := p.FetchRepositorySupabase.FetchValue("suppliers", queries)
	if err != nil {
		return err
	}
	supplierID := int64(0)
	if len(supplier) > 0 {
		tempSupplier := supplier[0]
		supplierID = int64(tempSupplier["id"].(float64))
	}

	value := map[string]string{
		"name":            update.Name,
		"quality":         update.Quality,
		"description":     update.Description,
		"thumbnail":       update.Thumbnail,
		"price":           fmt.Sprintf("%d", update.Price),
		"product_type_id": fmt.Sprintf("%d", update.ProductTypeID),
	}
	if err := p.ProductRepositorySupabase.UpdateProduct(value, productID, supplierID); err != nil {
		return err
	}

	return nil
}

func (p *ProductUseCase) DeleteProduct(productID, userID int64) error {
	queries := map[string]string{
		"select":  "id",
		"user_id": fmt.Sprintf("eq.%d", userID),
	}
	supplier, err := p.FetchRepositorySupabase.FetchValue("suppliers", queries)
	if err != nil {
		return err
	}
	supplierID := int64(0)
	if len(supplier) > 0 {
		tempSupplier := supplier[0]
		supplierID = int64(tempSupplier["id"].(float64))
	}

	fmt.Println(supplierID)

	if err := p.ProductRepositorySupabase.DeleteProduct(productID, supplierID); err != nil {
		return err
	}
	return nil
}
