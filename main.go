package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/cfi2017/dgcobra"
)

var (
	token   = flag.String("token", "", "token")
	handler *dgcobra.Handler
)

func main() {
	flag.Parse()
	if *token == "" {
		*token = os.Getenv("BOT_TOKEN")
		if *token == "" {
			log.Fatal("missing token")
		}
	}

	discord, err := discordgo.New("Bot " + *token)
	if err != nil {
		panic(err)
	}

	handler = dgcobra.NewHandler(discord)
	handler.AddPrefix("!")
	handler.RootFactory = rootFactory
	handler.Start()

	discord.AddHandlerOnce(onReady)

	err = discord.Open()
	if err != nil {
		panic(err)
	}

	log.Println("Bot is running. Press CTRL-C to exit.")
	waitForSig()
	err = discord.Close()
	if err != nil {
		panic(err)
	}

}

func onReady(_ *discordgo.Session, ready *discordgo.Ready) {
	handler.AddPrefix(fmt.Sprintf("<@!%s> ", ready.User.ID))
}

func waitForSig() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
