package facade

type IUser interface {
	Login(phone, code int) (*User, error)
	Register(phone, code int) (*User, error)
}

type IUserFacade interface {
	LoginOrRegister(phone, code int) (*User, error)
}

type User struct {
	Name string
}

func newUser(name string) IUser {
	return User{Name: name}
}

func (u User) Login(phone, code int) (*User, error) {
	return &User{Name: "test login"}, nil
}

func (u User) Register(phone, code int) (*User, error) {
	return &User{Name: "test register"}, nil
}

type UserService struct {
	User IUser
}

func NewUserService(name string) UserService {
	return UserService{
		User: newUser(name),
	}
}

func (u UserService) LoginOrRegister(phone, code int) (*User, error) {
	user, err := u.User.Login(phone, code)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	return u.User.Register(phone, code)
}
