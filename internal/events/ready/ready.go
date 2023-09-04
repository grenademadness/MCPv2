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

package ready

import (
	"github.com/adh-partnership/api/pkg/logger"
	"github.com/bwmarrin/discordgo"
)

var log = logger.Logger.WithField("component", "events/ready")

func Handler(s *discordgo.Session, r *discordgo.Ready) {
	log.Infof("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)

	err := s.UpdateWatchStatus(0, "Falcon")
	if err != nil {
		log.Warnf("Error setting watch status: %s", err)
	}
}
