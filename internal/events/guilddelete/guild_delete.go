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

package guilddelete

import (
	"github.com/adh-partnership/api/pkg/logger"
	"github.com/bwmarrin/discordgo"

	"github.com/vpaza/bot/internal/commands"
	"github.com/vpaza/bot/internal/facility"
)

var log = logger.Logger.WithField("component", "events/guild_delete")

func Handler(s *discordgo.Session, m *discordgo.GuildDelete) {
	log.Infof("Left guild %s", m.Guild.Name)

	_, err := facility.FindFacility(&facility.Facility{DiscordID: m.Guild.ID})
	if err == nil {
		log.Warnf("Facility %s still exists, but we left the guild. This is unexpected.", m.Guild.ID)
	}

	log.Infof("Unregistering commands")
	err = commands.Unregister(s, m.Guild.ID)
	if err != nil {
		log.Errorf("Failed to unregister commands: %s", err)
	}
}
