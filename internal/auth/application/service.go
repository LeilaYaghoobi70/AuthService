package application

type Service interface {
	IsValidateToken(email string) (bool, error)
	GenerateToken(email string) (string, error)
}

type service struct {
}

func AuthService() Service {
	return service{}
}

func (s service) GenerateToken(email string) (string, error) {
	return GenerateToken(email)
}
func (s service) IsValidateToken(token string) (bool, error) {
	return TokenIsValid(token)
}
