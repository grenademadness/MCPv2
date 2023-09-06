package events

import (
	"fmt"
	"strings"
	"time"

	"github.com/adh-partnership/api/pkg/database/dto"
	"github.com/bwmarrin/discordgo"
)

func postAnnouncement(s *discordgo.Session, i *discordgo.InteractionCreate, event *dto.EventsResponse) {
	message := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Starting At",
				Value:  fmt.Sprintf("<t:%d:f>", event.StartDate.Unix()),
				Inline: true,
			},
			{
				Name:   "Finishing At",
				Value:  fmt.Sprintf("<t:%d:f>", event.EndDate.Unix()),
				Inline: true,
			},
			{
				Name:   "Description",
				Value:  event.Description,
				Inline: false,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: event.Banner,
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     event.Title,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "END OF LINE.",
		},
	}
	s.ChannelMessageSendEmbed(i.ChannelID, message)
}

func postPositions(s *discordgo.Session, i *discordgo.InteractionCreate, event *dto.EventsResponse) {
	if len(event.Positions) > 25 {
		postPositionsSplit(s, i, event)
		return
	}

	enroutePositions := map[string]string{}
	terminalRadarPositions := map[string]string{}
	cabPositions := map[string]string{}

	for _, p := range event.Positions {
		parts := strings.Split(p.Position, "_")
		switch parts[len(parts)-1] {
		case "CTR", "FSS":
			enroutePositions[p.Position] = getControllerFromUser(p.User)
		case "APP", "DEP":
			terminalRadarPositions[p.Position] = getControllerFromUser(p.User)
		default:
			cabPositions[p.Position] = getControllerFromUser(p.User)
		}
	}

	// Part 1 -- Enroute Positions
	message := &discordgo.MessageEmbed{
		Color:       0x0000ff,
		Title:       event.Title,
		Description: "Event Position Assignments",
		Footer: &discordgo.MessageEmbedFooter{
			Text: "END OF LINE.",
		},
		Image: &discordgo.MessageEmbedImage{
			URL: event.Banner,
		},
		Fields: []*discordgo.MessageEmbedField{},
	}
	for _, v := range event.Positions {
		message.Fields = append(message.Fields, &discordgo.MessageEmbedField{
			Name:   v.Position,
			Value:  getControllerFromUser(v.User),
			Inline: true,
		})
	}
	s.ChannelMessageSendEmbed(i.ChannelID, message)
}

func postPositionsSplit(s *discordgo.Session, i *discordgo.InteractionCreate, event *dto.EventsResponse) {
	enroutePositions := map[string]string{}
	terminalRadarPositions := map[string]string{}
	cabPositions := map[string]string{}

	for _, p := range event.Positions {
		parts := strings.Split(p.Position, "_")
		switch parts[len(parts)-1] {
		case "CTR", "FSS":
			enroutePositions[p.Position] = getControllerFromUser(p.User)
		case "APP", "DEP":
			terminalRadarPositions[p.Position] = getControllerFromUser(p.User)
		default:
			cabPositions[p.Position] = getControllerFromUser(p.User)
		}
	}

	// Part 1 -- Enroute Positions
	message := &discordgo.MessageEmbed{
		Color:       0x0000ff,
		Title:       event.Title,
		Description: "Event Position Assignments",
		Fields:      []*discordgo.MessageEmbedField{},
	}
	for k, v := range enroutePositions {
		message.Fields = append(message.Fields, &discordgo.MessageEmbedField{
			Name:   k,
			Value:  v,
			Inline: true,
		})
	}
	s.ChannelMessageSendEmbed(i.ChannelID, message)

	// Part 2 -- Terminal Radar Positions
	message = &discordgo.MessageEmbed{
		Color:  0x0000ff,
		Fields: []*discordgo.MessageEmbedField{},
	}
	for k, v := range terminalRadarPositions {
		message.Fields = append(message.Fields, &discordgo.MessageEmbedField{
			Name:   k,
			Value:  v,
			Inline: true,
		})
	}
	s.ChannelMessageSendEmbed(i.ChannelID, message)

	// Part 3 -- Cab Positions
	message = &discordgo.MessageEmbed{
		Footer: &discordgo.MessageEmbedFooter{
			Text: "END OF LINE.",
		},
		Image: &discordgo.MessageEmbedImage{
			URL: event.Banner,
		},
		Color:  0x0000ff,
		Fields: []*discordgo.MessageEmbedField{},
	}
	for k, v := range cabPositions {
		message.Fields = append(message.Fields, &discordgo.MessageEmbedField{
			Name:   k,
			Value:  v,
			Inline: true,
		})
	}
	s.ChannelMessageSendEmbed(i.ChannelID, message)
}

func getControllerFromUser(user *dto.UserResponse) string {
	if user == nil {
		return "Unassigned"
	}

	if user.DiscordID == "" {
		return fmt.Sprintf("%s %s - %s", user.FirstName, user.LastName, user.OperatingInitials)
	}

	return fmt.Sprintf("<@%s>", user.DiscordID)
}
