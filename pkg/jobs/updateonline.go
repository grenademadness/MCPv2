package jobs

import (
	"fmt"
	"strings"
	"time"

	"github.com/adh-partnership/api/pkg/logger"
	"github.com/adh-partnership/api/pkg/network/vatsim"
	"github.com/bwmarrin/discordgo"

	"github.com/vpaza/bot/internal/bot"
	"github.com/vpaza/bot/internal/facility"
	"github.com/vpaza/bot/pkg/utils"
)

var uolog = logger.Logger.WithField("component", "jobs/updateonline")

func UpdateOnline() {
	if discord == nil || !discord.DataReady {
		log.Infof("Discord data not ready, skipping UpdateGuilds job")
		return
	}

	data, err := vatsim.GetData()
	if err != nil {
		uolog.Errorf("Error getting data from VATSIM: %s", err)
		return
	}
	for _, f := range facility.FacCfg {
		if f.OnlineChannel == "" {
			continue
		}
		go updateOnline(f, data)
	}
}

func updateOnline(f *facility.Facility, data *vatsim.VATSIMData) {
	uolog.Debugf("Starting UpdateOnline job for %s", f.Facility)
	start := time.Now()

	var online = make(map[string][]string)

	for _, p := range f.Positions {
		online[p.Name] = []string{}
	}

	for _, c := range data.Controllers {
		group := findGroup(f, c.Callsign)
		if group != "" {
			online[group] = append(
				online[group],
				fmt.Sprintf(
					"%s - %s - %s - %s",
					c.Callsign,
					getOI(f, fmt.Sprint(c.CID)),
					c.Frequency,
					time.Since(*c.LogonTime).Round(time.Second).String(),
				),
			)
		}
	}

	message := &discordgo.MessageEmbed{
		Author:    &discordgo.MessageEmbedAuthor{},
		Color:     0x00ff00,
		Fields:    []*discordgo.MessageEmbedField{},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     "Online Controllers",
		Footer: &discordgo.MessageEmbedFooter{
			Text: "END OF LINE.",
		},
	}

	for _, p := range f.Positions {
		if len(online[p.Name]) == 0 {
			online[p.Name] = []string{f.GetNoControllersOnlineMessage()}
		}
		message.Fields = append(
			message.Fields,
			&discordgo.MessageEmbedField{
				Name:   p.Name,
				Value:  strings.Join(online[p.Name], "\n"),
				Inline: false,
			},
		)
	}

	post, err := bot.FindBotMessageIn(f.OnlineChannel)
	if err != nil {
		uolog.Errorf("Error finding bot message in %s: %s", f.OnlineChannel, err)
		return
	}

	if post == nil {
		bot.GetSession().ChannelMessageSendEmbed(f.OnlineChannel, message)
	} else {
		bot.GetSession().ChannelMessageEditEmbed(f.OnlineChannel, post.ID, message)
	}

	uolog.Debugf("Finished UpdateOnline job for %s (took: %s)", f.Facility, time.Since(start))
}

func findGroup(f *facility.Facility, callsign string) string {
	parts := strings.Split(callsign, "_")
	for _, p := range f.Positions {
		if utils.Contains(p.Callsigns.Prefix, parts[0]) && utils.Contains(p.Callsigns.Suffix, parts[len(parts)-1]) {
			return p.Name
		}
	}
	return ""
}

func getOI(f *facility.Facility, cid string) string {
	u, err := f.FindUserByCID(cid)
	if err != nil {
		return cid
	}

	return u.OperatingInitials
}
