package events

import (
	"encoding/json"

	"github.com/bwmarrin/discordgo"
	"github.com/vpaza/bot/internal/facility"
)

func EventPostSelectHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Infof("Received event-select component from %s in %s #%s", i.Member.User.ID, i.GuildID, i.ChannelID)
	f, err := facility.FindFacility(
		&facility.Facility{
			DiscordID: i.GuildID,
		},
	)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Content: "Error finding facility",
			},
		})
		return
	}

	val := &EventSelectData{}
	err = json.Unmarshal([]byte(i.MessageComponentData().Values[0]), &val)
	if err != nil {
		log.Errorf("Failed to parse event data: %s", err)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Content: "Error finding event",
			},
		})
		return
	}

	event := f.GetEvent(val.EventID)
	if event == nil {
		log.Errorf("Failed to find event with id %d", val.EventID)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Content: "Error finding event",
			},
		})
		return
	}

	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: "Creating post for event " + event.Title,
		},
	})

	switch val.Type {
	case "announce":
		postAnnouncement(s, i, event)
	case "position":
		postPositions(s, i, event)
	}
}
