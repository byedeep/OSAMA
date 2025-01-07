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

func Ready(s *discordgo.Session, m *discordgo.Ready) {
	s.UpdateWatchStatus(0, "9/11")
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

func CreateReply(s *discordgo.Session , keyword ,reply string )string{
	for _ , ExistingReply = range replies{
		if ExistingReply.Keyword == strings.ToLower(keyword){
			return fmt.Println("This keyword already exist")
		}
	NewReply := types.Reply{
		Keyword: strings.ToLower(keyword),
		Reply: reply,
	}
	replies = append(replies, NewReply)

	file , err := os.Open("../data/data.CSV", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil{
		return nil 
	}
	defer file.Close()
	
	writer := csv.NewWriter(file)
	defer writer.Flush()

	err
	}
}

func SlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate){
	if i.Type == discordgo.InteractionApplicationCommand{
		if i.ApplicationCommandData().Name == "Create"{

		var keyword,reply string

		options := i.ApplicationCommandData().Options
		for _, option := range options{
			switch option.Name{
			case "keyword":
				keyword = option.StringValue()
			case "reply":
				reply = option.StringValue()
			}
		}
		contentresponse = CreateReply(keyword, reply)
		createreply := CreateReply()
		err := s.InteractionResponse(i.Interaction , &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: contentresponse,
			},
		})
		if err != nil {
			fmt.Println("Errors in interaction!")
	}

}
}


func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	words := strings.Split(strings.ToLower(m.Content), " ")
	//Loop th all the words in the message
	for _, word := range words {
		//Loop through all the reply entries we have.
		for _, replyEntry := range replies {
			//Every reply entry is a message & a reply. So we deconstruct that into respective variables
			keyword := replyEntry.Keyword
			reply := replyEntry.Reply
			if word == keyword {
				s.ChannelMessageSend(m.ChannelID, reply)
			}

		}
	}
}
