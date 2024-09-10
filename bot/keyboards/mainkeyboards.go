package keyboards

import (
	tele "gopkg.in/telebot.v3"
)

var (
	DefaultMenu  = &tele.ReplyMarkup{ResizeKeyboard: true, OneTimeKeyboard: true}
	BtnEvents    = DefaultMenu.Text("События")
	BtnSubscribe = DefaultMenu.Text("Подписки")
	BtnSettings  = DefaultMenu.Text("Настройки")
	BtnBack      = DefaultMenu.Text("Назад")
)

func CreateDefaultMenu() {
	DefaultMenu.Reply(
		DefaultMenu.Row(BtnEvents, BtnSettings),
		DefaultMenu.Row(BtnSubscribe, BtnBack),
	)
}

func init() {
	CreateDefaultMenu()
}