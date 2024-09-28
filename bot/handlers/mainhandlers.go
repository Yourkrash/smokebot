package handlers

import (
	"context"
	"log"
	"smokebot/bot/keyboards"
	pb "smokebot/dbservice/proto"

	"github.com/vitaliy-ukiru/fsm-telebot/v2"
	"github.com/vitaliy-ukiru/telebot-filter/dispatcher"
	tele "gopkg.in/telebot.v3"
)

const (
	MainState    fsm.State = "main"
	MainSettings fsm.State = "mainsett"
	MainEvents   fsm.State = "mainevent"
)

type HandleRunner struct {
	m      *fsm.Manager
	dp     *dispatcher.Dispatcher
	client pb.RegServiceClient
}

func New(m *fsm.Manager,
	dp *dispatcher.Dispatcher, client pb.RegServiceClient) *HandleRunner {
	return &HandleRunner{m, dp, client}
}

func (run *HandleRunner) InitHandlers() {
	run.m.Handle(
		run.dp,
		tele.OnText,
		fsm.DefaultState,
		func(c tele.Context, state fsm.Context) error {
			_, error := run.client.RegUser(context.TODO(), &pb.RegUserRequest{
				User: &pb.User{
					IdUser:    c.Sender().ID,
					FirstName: c.Sender().FirstName,
					LastName:  c.Sender().LastName,
				}})
			if error != nil {
				log.Fatal("error db run")
				return c.Send("Ошибка Подключения")
			}
			if err := state.SetState(context.TODO(), MainState); err != nil {
				log.Fatal("error bot /start")
			}
			return c.Send("Добро пожаловать в прекрасного бота", keyboards.DefaultMenu)
		},
	)

	run.m.Handle(
		run.dp,
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

	run.m.Handle(
		run.dp,
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
