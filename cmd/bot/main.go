package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	discord "github.com/bwmarrin/discordgo"
	rando "github.com/jadefish/cata-rando"
)

var appID string
var token string
var guildID string

func init() {
	flag.StringVar(&appID, "a", "", "Application ID")
	flag.StringVar(&token, "t", "", "Bot access token")
	flag.StringVar(&guildID, "g", "", "Server (guild) ID")
	flag.Parse()
}

func main() {
	dg, err := discord.New("Bot " + token)

	if err != nil {
		log.Fatalln("error creating Discord session.", err)
		return
	}

	dg.AddHandler(ready)
	dg.AddHandler(interactionCreate)

	_, err = dg.ApplicationCommandCreate(appID, guildID, &discord.ApplicationCommand{
		Name:        "rando",
		Description: "Roll a random character from cata-rando",
	})

	if err != nil {
		log.Fatalln("error creating Discord command:", err)
		return
	}

	err = dg.Open()

	if err != nil {
		log.Fatalln("error opening Discord session:", err)
	}

	defer dg.Close()

	log.Println("Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func ready(s *discord.Session, r *discord.Ready) {
	log.Println("Bot is ready! 🤖")
}

func interactionCreate(s *discord.Session, i *discord.InteractionCreate) {
	// NOTE: responses must come within 3 seconds or Discord invalidates the interaction.
	content := fmt.Sprintf("%v", roll())

	err := s.InteractionRespond(i.Interaction, &discord.InteractionResponse{
		Type: discord.InteractionResponseChannelMessageWithSource,
		Data: &discord.InteractionResponseData{Content: content},
	})

	if err != nil {
		log.Println("error sending response:", err)

		s.FollowupMessageCreate(i.Interaction, true, &discord.WebhookParams{
			Content: "Something went wrong",
		})
		return
	}
}

func roll() string {
	r := rando.NewRoller()

	return r.Roll().String()
}
