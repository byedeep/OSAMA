package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/byedeep/osama/internals/services"
)

func main() {
	dg, err := discordgo.New("Bot " + services.Token)
	if err != nil {
		fmt.Println("Error creating session")
		return
	}
	err = services.ReadFile("../data/data.CSV")
	if err != nil {
		fmt.Println(err)
		return
	}
	dg.AddHandler(services.Ready)

	dg.AddHandler(services.MessageCreate)

	dg.Identify.Intents = discordgo.IntentGuildMessages | discordgo.IntentGuildMessageTyping

	/*	SlashCommand := &discordgo.ApplicationCommand{
			Name: "Create",
			Description: "/Create <Keyword> <Reply>",
		}

		_, err := dg.ApplicationCommandCreate()*/

	err = dg.Open()
	if err != nil {
		fmt.Println("Error connecting")
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()

}
