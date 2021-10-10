package domain

type Fetch struct{}

type FetchRepository interface {
	FetchValue(table string, queries map[string]string) ([]map[string]interface{}, error)
}

type FetchUseCase interface {
	FetchValue(table string, queries map[string]string) ([]map[string]interface{}, error)
}

