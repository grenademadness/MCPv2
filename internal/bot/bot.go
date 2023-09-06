/*
 * Copyright Daniel Hawton
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package bot

import (
	"github.com/adh-partnership/api/pkg/logger"
	"github.com/bwmarrin/discordgo"

	"github.com/vpaza/bot/internal/commands"
	"github.com/vpaza/bot/internal/events"
	"github.com/vpaza/bot/pkg/config"
)

var (
	log     = logger.Logger.WithField("component", "bot")
	session *discordgo.Session
)

func Start() (*discordgo.Session, error) {
	var err error
	log.Infof("Starting bot")
	session, err = discordgo.New("Bot " + config.Cfg.Discord.Token)
	if err != nil {
		log.Errorf("Failed to create discord client: %s", err)
		return nil, err
	}

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentGuildMembers

	log.Infof("Registering events")
	events.AddEvents(session)

	log.Infof("Opening websocket to discord...")
	err = session.Open()
	if err != nil {
		return nil, err
	}

	log.Infof("Setting up commands")
	commands.SetupCommands()

	return session, nil
}

func GetSession() *discordgo.Session {
	return session
}

func FindBotMessageIn(channelid string) (*discordgo.Message, error) {
	messages, err := session.ChannelMessages(channelid, 100, "", "", "")
	if err != nil {
		return nil, err
	}

	for _, m := range messages {
		if m.Author.ID == session.State.User.ID {
			return m, nil
		}
	}

	return nil, nil
}
