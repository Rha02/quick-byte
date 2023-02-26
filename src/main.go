package main

import (
	"fmt"
	"os"
	"syscall"

	"os/signal"

	"github.com/bwmarrin/discordgo"

	"github.com/joho/godotenv"

	"github.com/Rha02/src/handlers"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Create a new Discord session
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))

	if err != nil {
		println("error creating Discord session,", err)
		return
	}

	// Create a new handler repository
	repo := handlers.NewRepository()

	// Add the MessageHandler callback.
	dg.AddHandler(repo.MessageHandler)

	// Set the intents to receive all messages
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and listen to it.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	// Run the app until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Close the Discord session.
	dg.Close()
}
