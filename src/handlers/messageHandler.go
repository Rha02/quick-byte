package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/Rha02/src/handlers/cmds"
)

// MessageHandler handles the messageCreate event
func (r *Repository) MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message doesn't start with "qb! ", then it is not a command intended for the bot
	if !strings.HasPrefix(m.Content, "qb! ") {
		return
	}

	// Split the message into command and arguments
	args := strings.Split(m.Content, " ")

	// Check if the command is valid
	cmd, ok := cmds.Commands[args[1]]
	if !ok {
		s.ChannelMessageSend(m.ChannelID, "Invalid command!")
		return
	}

	// Execute the command
	cmd.Execute(s, m, args)
}
