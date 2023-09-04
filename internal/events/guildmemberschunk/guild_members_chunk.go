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

package guildmemberschunk

import (
	"github.com/adh-partnership/api/pkg/logger"
	"github.com/bwmarrin/discordgo"

	"github.com/vpaza/bot/internal/facility"
)

var log = logger.Logger.WithField("component", "events/guild_members_chunk")

func Handler(s *discordgo.Session, m *discordgo.GuildMembersChunk) {
	go func(s *discordgo.Session, m *discordgo.GuildMembersChunk) {
		fac, err := facility.FindFacility(&facility.Facility{DiscordID: m.GuildID})
		if err != nil {
			log.Errorf("Error finding facility for guild %s: %s", m.GuildID, err)
			return
		}

		log.Infof("(Chunk %s/%d/%d) Processing %d members", fac.Facility, m.ChunkIndex, m.ChunkCount, len(m.Members))
		defer log.Debugf("(Chunk %s/%d/%d) Completed processing chunk", fac.Facility, m.ChunkIndex, m.ChunkCount)
		for _, member := range m.Members {
			fac.ProcessMember(s, member)
		}
	}(s, m)
}
