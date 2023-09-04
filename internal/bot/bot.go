package bot

import (
	"github.com/adh-partnership/api/pkg/logger"
	"github.com/bwmarrin/discordgo"

	"github.com/vpaza/bot/internal/events"
	"github.com/vpaza/bot/pkg/config"
)

var log = logger.Logger.WithField("component", "bot")

func Start() error {
	log.Infof("Starting bot")
	session, err := discordgo.New("Bot " + config.Cfg.Discord.Token)
	if err != nil {
		log.Errorf("Failed to create discord client: %s", err)
		return err
	}

	session.Identify.Intents = discordgo.IntentsAll

	events.AddEvents(session)

}
