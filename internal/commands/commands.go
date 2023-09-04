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

package commands

import (
	"github.com/adh-partnership/api/pkg/logger"
	"github.com/bwmarrin/discordgo"

	"github.com/vpaza/bot/internal/commands/ping"
)

type AppCommand struct {
	Name       string
	ID         string
	Handler    func(*discordgo.Session, *discordgo.InteractionCreate)
	AppCommand *discordgo.ApplicationCommand
}

var (
	commands map[string]*AppCommand
	log      = logger.Logger.WithField("component", "commands")
)

func init() {
	commands = make(map[string]*AppCommand)
}

func SetupCommands() {
	AddCommand(ping.Register())
}

func AddCommand(name string, handler func(*discordgo.Session, *discordgo.InteractionCreate), appCommand *discordgo.ApplicationCommand) {
	commands[name] = &AppCommand{
		Name:       name,
		Handler:    handler,
		AppCommand: appCommand,
	}
}

func FindHandler(cmd string) (func(*discordgo.Session, *discordgo.InteractionCreate), bool) {
	if command, ok := commands[cmd]; ok {
		return command.Handler, true
	}
	return nil, false
}

func RegisterCommands(s *discordgo.Session, guildid string) error {
	// Clean all commands associated with us
	registeredCommands, err := s.ApplicationCommands(s.State.User.ID, guildid)
	if err != nil {
		return err
	}
	for _, registeredCommand := range registeredCommands {
		log.Warnf("Unregistering command %s, we didn't register", registeredCommand.Name)
		err := s.ApplicationCommandDelete(s.State.User.ID, guildid, registeredCommand.ID)
		if err != nil {
			return err
		}
	}

	for _, command := range commands {
		log.Infof("Registering command %s", command.Name)
		appCommand, err := s.ApplicationCommandCreate(s.State.User.ID, guildid, command.AppCommand)
		if err != nil {
			return err
		}
		command.ID = appCommand.ID
	}

	return nil
}

func Unregister(s *discordgo.Session, guildid string) error {
	for _, command := range commands {
		if command.ID != "" {
			log.Infof("Unregistering command %s", command.Name)
			err := s.ApplicationCommandDelete(s.State.User.ID, guildid, command.ID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
