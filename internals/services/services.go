package services

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/byedeep/osama/internals/types"
)

var Token string
var replies []types.Reply

func init() {
	flag.StringVar(&Token, "t", "", "bot token")
	flag.Parse()
}
func ReadFile(filename string) (err error) {
	file, err := os.Open(filename)
	var readReplies []types.Reply
	if err != nil {
		return err
	}
	defer file.Close()

	Reader := csv.NewReader(file)
	Records, err := Reader.ReadAll()
	fmt.Println(Records)
	if err != nil {
		return err
	}

	for _, Record := range Records {
		reply := types.Reply{
			Keyword: Record[0],
			Reply:   Record[1],
		}
		readReplies = append(readReplies, reply)
	}
	replies = readReplies
	return nil
}

func RelodeFile() {}

func Ready(s *discordgo.Session, m *discordgo.Ready) {
	s.UpdateWatchStatus(0, "9/11")
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	fmt.Println(m.Content)
	words := strings.Split(strings.ToLower(m.Content), " ")
	//Loop th all the words in the message
	for _, word := range words {
		//Loop through all the reply entries we have.
		for _, replyEntry := range replies {
			//Every reply entry is a message & a reply. So we deconstruct that into respective variables
			keyword := replyEntry.Keyword
			reply := replyEntry.Reply
			if word == keyword {
				fmt.Println(reply)
				s.ChannelMessageSend(m.ChannelID, reply)
			}

		}
	}
}
