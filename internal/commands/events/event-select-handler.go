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
	"encoding/json"

	"github.com/bwmarrin/discordgo"
)

type EventSelectData struct {
	EventID int64  `json:"event_id"`
	Type    string `json:"type"`
}

func EventSelectHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Infof("Received event-select component from %s in %s #%s", i.Member.User.ID, i.GuildID, i.ChannelID)

	// Convert i.MessageComponentData().Values[0] to int32
	var val map[string]interface{}
	err := json.Unmarshal([]byte(i.MessageComponentData().Values[0]), &val)
	if err != nil {
		log.Errorf("Failed to parse event data: %s", err)
		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "Error finding event",
			},
		})
		if err != nil {
			log.Warnf("Error responding to channel: %s", err)
		}
		return
	}

	announcementstruct := &EventSelectData{
		EventID: int64(val["event_id"].(float64)),
		Type:    "announce",
	}
	announcementdata, err := json.Marshal(announcementstruct)
	if err != nil {
		log.Errorf("Failed to marshal announcement data: %s", err)
		return
	}

	positionstruct := &EventSelectData{
		EventID: int64(val["event_id"].(float64)),
		Type:    "position",
	}
	positiondata, err := json.Marshal(positionstruct)
	if err != nil {
		log.Errorf("Failed to marshal position data: %s", err)
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: "Select post type for event " + val["event_title"].(string) + ":",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							CustomID:    "event-post-select",
							Placeholder: "Select a post type",
							Options: []discordgo.SelectMenuOption{
								{
									Label:       "Announcement",
									Value:       string(announcementdata),
									Description: "Post an announcement for this event",
								},
								{
									Label:       "Positions",
									Value:       string(positiondata),
									Description: "Post the positions for this event",
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Warnf("Error responding to channel: %s", err)
	}
}
