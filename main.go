package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func main() {
	discord, err := discordgo.New("Bot " + "authentication token")
	if err != nil {
		log.Fatal().Msgf("Unable to connect to bot: %v", err)
	}
}
