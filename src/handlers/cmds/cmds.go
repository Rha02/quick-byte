package cmds

import (
	"github.com/bwmarrin/discordgo"
)

// Command is a struct that holds information about a bot command
type Command struct {
	Name        string
	Description string
	Execute     func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
}

// Commands is a map of all the bot commands
var Commands = map[string]*Command{
	"ping": {
		Name:        "ping",
		Description: "Pong!",
		Execute:     Ping,
	},
}

// Ping is a command that replies with "Pong!"
func Ping(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	s.ChannelMessageSend(m.ChannelID, "Pong!")
}
