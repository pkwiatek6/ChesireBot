package main

import (
	"flag"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

var (
	discord *discordgo.Session
	// Token for the bot
	Token          = flag.String("t", "", "Bot acess token")
	GuildID        = flag.String("GID", "", "Test Guild ID. IF not passed - bot registers commands globally")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

func init() {
	var err error
	discord, err := discordgo.New("Bot " + "authentication token")
	if err != nil {
		log.Fatal().Msgf("Error creating Discord session: %v", err)
	}
	log.Debug().Msg("Cibbectuib to Discord established")
}

func main() {

}
