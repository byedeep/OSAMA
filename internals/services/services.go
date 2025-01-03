package services

import (
	"flag"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var Token string

func init() {
	flag.StringVar(&Token, "t", "", "bot token")
	flag.Parse()
}

func Ready(s *discordgo.Session, m *discordgo.Ready) {
	s.UpdateGameStatus(0, "Ready to bomb")
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	fmt.Println("Sent message", m.Content)

	if m.Content == "Osama" {
		s.ChannelMessageSend(m.ChannelID, "Binladen")

	}

}
