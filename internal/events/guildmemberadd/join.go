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

package guildmemberadd

import (
	"github.com/adh-partnership/api/pkg/logger"
	"github.com/bwmarrin/discordgo"

	"github.com/vpaza/bot/internal/facility"
)

var log = logger.Logger.WithField("component", "events/guild_member_add")

func Handler(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	guild := m.GuildID
	f, err := facility.FindFacility(&facility.Facility{
		DiscordID: guild,
	})
	if err == facility.ErrorFacilityNotFound {
		log.Errorf("No facility found for guild %s, what guild is this?", guild)
		return
	}

	log.Infof("User %s joined %s guild, processing", m.Member.User.Username, f.Facility)
	f.ProcessMember(s, m.Member)
}
