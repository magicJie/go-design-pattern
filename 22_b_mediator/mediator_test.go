package mediator

import "testing"

func TestDemo(t *testing.T) {
	usernameInput := Input("username input")
	passwordInput := Input("password input")
	repeaPwdInput := Input("repeat password input")

	selection := Selection("登录")
	d := &Dialog{
		Selection:           &selection,
		UsernameInput:       &usernameInput,
		PasswordInput:       &passwordInput,
		RepeatPasswordInput: &repeaPwdInput,
	}
	d.HandleEvent(&selection)

	regSelection := Selection("注册")
	d.Selection = &regSelection
	d.HandleEvent(&regSelection)
}
