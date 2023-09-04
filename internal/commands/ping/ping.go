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

package ping

import (
	"github.com/adh-partnership/api/pkg/logger"
	"github.com/bwmarrin/discordgo"

	"github.com/vpaza/bot/pkg/utils"
)

var log = logger.Logger.WithField("component", "commands/ping")

func Register() (name string, handler func(*discordgo.Session, *discordgo.InteractionCreate), appCommand *discordgo.ApplicationCommand) {
	return "ping", Handler, &discordgo.ApplicationCommand{
		Name:         "ping",
		Description:  "Ping the bot",
		DMPermission: utils.PointerOf(false),
	}
}

func Handler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Infof("Received ping command from %s in %s #%s", i.Member.User.ID, i.GuildID, i.ChannelID)
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: "Pong!",
		},
	})
	if err != nil {
		log.Errorf("Error responding to ping command: %s", err)
	}
}
