package services

import (
	"flag"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/byedeep/osama/constants"
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

	words := strings.Split(strings.ToLower(m.Content), " ")

	for _, word := range words {
		for message, reply := range constants.MessageReplies {
			if word == message {
				s.ChannelMessageSend(m.ChannelID, reply)

			}
		}
	}

}
