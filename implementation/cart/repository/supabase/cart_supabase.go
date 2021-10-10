package supabase

import (
	"encoding/json"
	"fmt"
	"github.com/huf0813/gemastik_api_backend_supabase/domain"
	"github.com/huf0813/gemastik_api_backend_supabase/infrastructure"
	"net/http"
)

type CartRepositorySupabase struct {
	DB        *infrastructure.DriverSupabase
	TableCart string
}

func NewCartRepositorySupabase(db *infrastructure.DriverSupabase) domain.CartRepository {
	return &CartRepositorySupabase{DB: db, TableCart: "carts"}
}

func (c *CartRepositorySupabase) DeleteProductFromCart(userID, productID, cartID int64) error {
	request, err := c.DB.RequestFormula(c.TableCart, http.MethodDelete, nil)
	if err != nil {
		return err
	}
	queries := request.URL.Query()
	queries.Add("id", fmt.Sprintf("eq.%d", cartID))
	queries.Add("user_id", fmt.Sprintf("eq.%d", userID))
	queries.Add("product_id", fmt.Sprintf("eq.%d", productID))
	request.URL.RawQuery = queries.Encode()

	if _, err := c.DB.ExecuteRequestFormula(request); err != nil {
		return err
	}

	return nil
}

func (c *CartRepositorySupabase) AddProductToCart(value map[string]string) error {
	valueByte, err := json.Marshal(value)
	if err != nil {
		return err
	}

	request, err := c.DB.RequestFormula(c.TableCart, http.MethodPost, valueByte)
	if err != nil {
		return err
	}

	if _, err := c.DB.ExecuteRequestFormula(request); err != nil {
		return err
	}

	return nil
}
