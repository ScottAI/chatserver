package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type LoginHandler func(string)


type LoginView struct {
	loginHandler LoginHandler
	window fyne.Window
}

func NewLoginView(login fyne.App,handler LoginHandler) *LoginView {
	myApp := LoginView{loginHandler:handler}
	window := login.NewWindow("登录")
	input := widget.NewEntry()
	input.ReadOnly = false
	input.Resize(fyne.NewSize(24,5))
	window.SetContent(widget.NewVBox(
			widget.NewLabel("Please input your name："),
			input,
			widget.NewButton("login", func() {
				if len(input.Text) >0 {
					myApp.loginHandler(input.Text)
					myApp.window.Hide()
					login.Quit()
				}
			}),
		))
	window.Resize(fyne.NewSize(24,24))
	myApp.window = window
	return &myApp
}

func (v *LoginView) Show()  {
	v.window.ShowAndRun()
}




