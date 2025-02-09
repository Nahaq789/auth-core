package user

type UserType string

const (
	Guest    UserType = "guest"
	Standard UserType = "standard"
)

func (ut UserType) String() string {
	return string(ut)
}

func NewUserType(userType string) UserType {
	switch userType {
	case "guest":
		return Guest
	case "standard":
		return Standard
	default:
		return Guest
	}
}
