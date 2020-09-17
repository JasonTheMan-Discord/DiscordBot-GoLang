package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	Token := ""

	fmt.Println("Main event successfuly recieved")
	client, err := discordgo.New("Bot " + Token)
	if err != nil {
		panic(err)
	}

	client.AddHandler(ready)

	client.AddHandler(messageCreate)

	err = client.Open()
	if err != nil {
		panic(err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fmt.Println("Shutting down the system...\n... Successful")
	// Cleanly close down the Discord session.
	client.Close()
}

// The ready event, just like in discordjs
func ready(s *discordgo.Session, event *discordgo.Ready) {
	err := s.UpdateStatus(0, "Accepting Orders | Illusion")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Ready event successfuly recieved")
}

func messageCreate(session *discordgo.Session, msg *discordgo.MessageCreate) {
	// Let's ignore all messages sent by the bot itself
	if msg.Author.ID == session.State.User.ID {
		return
	}

	if msg.Content == "ping" {
		session.ChannelMessageSend(msg.ChannelID, "Pong!")
	}
}
