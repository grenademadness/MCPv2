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
	"github.com/adh-partnership/api/pkg/logger"
	"github.com/bwmarrin/discordgo"

	"github.com/vpaza/bot/pkg/interactions"
	"github.com/vpaza/bot/pkg/utils"
)

var log = logger.Logger.WithField("component", "commands/events")

func Register() {
	interactions.AddCommand(&interactions.AppInteraction{
		Name:    "events",
		Handler: Handler,
		AppCommand: &discordgo.ApplicationCommand{
			Name:         "events",
			Description:  "Post information or positions of an event",
			DMPermission: utils.PointerOf(false),
		},
	})
	interactions.AddComponent(&interactions.AppInteraction{
		Name:    "event-select",
		Handler: EventSelectHandler,
	})
	interactions.AddComponent(&interactions.AppInteraction{
		Name:    "event-post-select",
		Handler: EventPostSelectHandler,
	})
}
