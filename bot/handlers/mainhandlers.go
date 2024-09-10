package handlers

import (
	pb "smokebot/dbservice/proto"
	"smokebot/bot/keyboards"
	"context"
	"log"

	"github.com/vitaliy-ukiru/fsm-telebot/v2"
	"github.com/vitaliy-ukiru/telebot-filter/dispatcher"
	tele "gopkg.in/telebot.v3"
)

const (
	MainState    fsm.State = "main"
	MainSettings fsm.State = "mainsett"
	MainEvents   fsm.State = "mainevent"
)

func InitHandlers(m *fsm.Manager, dp *dispatcher.Dispatcher) {

	m.Handle(
		dp,
		tele.OnText,
		fsm.DefaultState,
		func(c tele.Context, state fsm.Context) error {
			if err := state.SetState(context.TODO(), MainState); err != nil {
				log.Fatal("error bot /start")
			}
			return c.Send("Добро пожаловать в прекрасного подпивасного бота", keyboards.DefaultMenu)
		},
	)

	m.Handle(
		dp,
		&keyboards.BtnSettings,
		MainState,
		func(c tele.Context, state fsm.Context) error {
			if err := state.SetState(context.TODO(), MainSettings); err != nil {
				log.Fatal("bad switch from Main to my_state")
			}
			st, _ := state.State(context.TODO())
			return c.Send(st.GoString())
		},
	)

	m.Handle(
		dp,
		&keyboards.BtnEvents,
		MainState,
		func(c tele.Context, state fsm.Context) error {
			if err := state.SetState(context.TODO(), MainEvents); err != nil {
				log.Fatal("bad switch from my_state to Main")
			}
			c.Delete()
			return c.Send("1")
		},
	)
}
