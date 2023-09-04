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

package guildcreate

import (
	"github.com/adh-partnership/api/pkg/logger"
	"github.com/bwmarrin/discordgo"

	"github.com/vpaza/bot/internal/commands"
	"github.com/vpaza/bot/internal/facility"
)

var log = logger.Logger.WithField("component", "events/guild_create")

func Handler(s *discordgo.Session, m *discordgo.GuildCreate) {
	log.Infof("Joined guild %s", m.Guild.Name)

	fac, err := facility.FindFacility(&facility.Facility{DiscordID: m.Guild.ID})
	if err != nil {
		log.Errorf("Error finding facility for guild %s: %s", m.Guild.ID, err)
		return
	}

	log.Infof("Changing nickname to %s", fac.BotName)
	err = s.GuildMemberNickname(m.Guild.ID, "@me", fac.BotName)
	if err != nil {
		log.Errorf("Error changing nickname in guild %s: %s", fac.Facility, err)
	}

	log.Infof("Registering commands")
	err = commands.RegisterCommands(s, m.Guild.ID)
	if err != nil {
		log.Errorf("Failed to register commands: %s", err)
	}
}
