package login

// LoginService defines the interface for the login service
type LoginService interface {
	Login(username, password string) error
}
