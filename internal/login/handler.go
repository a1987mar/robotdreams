package login

type Login struct {
	Email    string `json:"emeil"`
	Passwors string `json:"password"`
	Name     string `json:"name"`
}

func NewLogin(email, password, name string) *Login {
	return &Login{
		Email:    email,
		Passwors: password,
		Name:     name,
	}
}
