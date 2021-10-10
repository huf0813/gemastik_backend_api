package domain

type UserSignUp struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	IdentityNumber string `json:"identity_number"`
	Password       string `json:"password"`
}

type UserSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserClaimSupplier struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
}

type UserRepository interface {
	SignUp(value map[string]string) error
	ClaimSupplier(value map[string]string) error
}

type UserUseCase interface {
	SignUp(su *UserSignUp) error
	SignIn(si *UserSignIn) (string, error)
	ClaimSupplier(cs *UserClaimSupplier, userID int64) error
	GetProfile(userID int64, selectedColumns string) (map[string]interface{}, error)
}
