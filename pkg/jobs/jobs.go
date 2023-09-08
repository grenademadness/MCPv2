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

package jobs

import (
	"time"

	"github.com/adh-partnership/api/pkg/logger"
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
)

var (
	s       *gocron.Scheduler
	discord *discordgo.Session
	log     = logger.Logger.WithField("component", "jobs")
)

func BuildJobs() {
	s = gocron.NewScheduler(time.UTC)

	log.Infof("Registering UpdateGuilds Job")
	_, err := s.Every(1).Minutes().SingletonMode().Do(UpdateGuilds)
	if err != nil {
		log.Errorf("Failed to schedule UpdateGuilds: %s", err)
	}

	log.Infof("Registering UpdateOnline Job")
	_, err = s.Every(30).Seconds().SingletonMode().Do(UpdateOnline)
	if err != nil {
		log.Errorf("Failed to schedule UpdateOnline: %s", err)
	}
}

func Start(sess *discordgo.Session) {
	log.Infof("Starting jobs")
	s.StartAsync()
	discord = sess
}
