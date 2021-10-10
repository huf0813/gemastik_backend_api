package domain

type CreateStoreRequest struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	TimeOpened  string `json:"time_opened"`
	TimeClosed  string `json:"time_closed"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
}

type StoreRepository interface {
	ClaimStore(value map[string]string) error
}

type StoreUseCase interface {
	ClaimStore(c *CreateStoreRequest, userID int64) error
}
