package events

import "github.com/bwmarrin/discordgo"

func AddEvents(s *discordgo.Session) {
	s.AddHandler(handlerAddMember)
}
