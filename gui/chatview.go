package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type SubmitMessHandler func (string)

type ChatView struct {
	submitMessHandler SubmitMessHandler
	window fyne.Window
	history *widget.Entry
}

func NewChatView(chatView fyne.App,handler SubmitMessHandler) *ChatView  {

	window := chatView.NewWindow("聊天窗口")

	history := widget.NewMultiLineEntry()
	history.ReadOnly=true
	history.Resize(fyne.NewSize(480,300))
	myChatView := ChatView{submitMessHandler:handler,history:history}
	input := widget.NewEntry()
	input.ReadOnly=false
	input.Resize(fyne.NewSize(460,20))
	send := widget.NewButton("send", func() {
		if len(input.Text)>0 {
			myChatView.submitMessHandler(input.Text)
			input.SetText("")
		}
	})
	group := widget.NewHBox(input,send)

	content := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),history,group)
	content.Resize(fyne.NewSize(480,320))
	window.SetContent(content)
	window.Resize(fyne.NewSize(480,320))
	myChatView.window = window

	return &myChatView
}

func (c *ChatView) Show()  {
	c.window.ShowAndRun()
}

func (c *ChatView) AddMessage(user string,msg string)  {
	c.history.SetText(c.history.Text+"\n"+user+":"+msg)

}