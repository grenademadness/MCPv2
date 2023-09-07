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

package events

import (
	"github.com/bwmarrin/discordgo"

	"github.com/vpaza/bot/internal/facility"
	"github.com/vpaza/bot/pkg/utils"
)

func Handler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Infof("Received events command from %s in %s #%s", i.Member.User.ID, i.GuildID, i.ChannelID)
	f, err := facility.FindFacility(
		&facility.Facility{
			DiscordID: i.GuildID,
		},
	)
	if err != nil {
		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "Error finding facility",
			},
		})
		if err != nil {
			log.Warnf("Error respond to events channel: %s", err)
		}
		return
	}

	events := f.GetEvents()
	if len(events) == 0 {
		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "No events found",
			},
		})
		if err != nil {
			log.Warnf("Error respond to events channel: %s", err)
		}
		return
	}

	options := make([]discordgo.SelectMenuOption, len(events))
	for idx, event := range events {
		options[idx] = discordgo.SelectMenuOption{
			Label: event.Title,
			Value: utils.MapJSON(map[string]interface{}{
				"event_id":    event.ID,
				"event_title": event.Title,
			}),
			Description: utils.Trim(event.Description, 32),
		}
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: "Select an event:",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							CustomID:    "event-select",
							Placeholder: "Select an event",
							Options:     options,
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Errorf("Error responding to events command: %s", err)
	}
}
