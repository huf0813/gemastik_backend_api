package utility

import "golang.org/x/crypto/bcrypt"

func NewHashValue(source string) (string, error) {
	p := []byte(source)
	bytes, err := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashedPassword := string(bytes)
	return hashedPassword, nil
}

func NewCompareValue(hashedValue, target string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(target))
}
