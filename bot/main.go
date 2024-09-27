package main

import (
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"smokebot/bot/handlers"
	pb "smokebot/dbservice/proto"

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

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewMovieServiceClient(conn)

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
    
	handle := handlers.New(m, dp, &client)
	handle.InitHandlers()

	bot.Start()
}
