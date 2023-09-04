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

package interactioncreate

import (
	"github.com/adh-partnership/api/pkg/logger"
	"github.com/bwmarrin/discordgo"

	"github.com/vpaza/bot/internal/commands"
)

var log = logger.Logger.WithField("component", "events/ready")

func Handler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Debugf("Interaction received: %s %s", i.Type, i.ApplicationCommandData().Name)

	h, ok := commands.FindHandler(i.ApplicationCommandData().Name)
	if !ok {
		log.Errorf("No handler found for command %s", i.ApplicationCommandData().Name)
		return
	}

	h(s, i)
}
