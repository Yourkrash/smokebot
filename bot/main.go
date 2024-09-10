package main

import (
	"smokebot/bot/handlers"
	"flag"

	// "fmt"
	"log"
	"os"
	"time"

	"github.com/vitaliy-ukiru/fsm-telebot/v2"
	"github.com/vitaliy-ukiru/fsm-telebot/v2/pkg/storage/memory"
	"github.com/vitaliy-ukiru/telebot-filter/dispatcher"

	// tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tele "gopkg.in/telebot.v3"
)

var debug = flag.Bool("debug", false, "log debug info")

func main() {
	flag.Parse()

	bot, err := tele.NewBot(tele.Settings{
		Token:   os.Getenv("TOKEN_TG_BOT"),
		Poller:  &tele.LongPoller{Timeout: 10 * time.Second},
		Verbose: *debug,
	})
	if err != nil {
		log.Fatalln(err)
	}

	g := bot.Group()
	m := fsm.New(memory.NewStorage())

	g.Use(m.WrapContext)

	dp := dispatcher.NewDispatcher(g)

	handlers.InitHandlers(m, dp)

	bot.Start()
}
