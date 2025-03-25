package login

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func NewLogin(email, password, name string) *Login {
	return &Login{
		Email:    email,
		Password: password,
		Name:     name,
	}
}
