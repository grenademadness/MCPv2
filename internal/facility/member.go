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

package facility

import (
	"slices"

	"github.com/adh-partnership/api/pkg/database/dto"
	"github.com/bwmarrin/discordgo"

	"github.com/vpaza/bot/internal/formatter/names/firstcid"
	"github.com/vpaza/bot/internal/formatter/names/firstlast"
	"github.com/vpaza/bot/internal/formatter/names/firstlastinitial"
	"github.com/vpaza/bot/pkg/utils"
)

func (f *Facility) GenerateNameFromUser(u *dto.UserResponse) string {
	switch f.NameFormat {
	case "first_cid":
		return firstcid.GenerateNameFromUser(u, f.StaffFormat, f.StaffTitleSeparator)
	case "first_last_initial":
		return firstlastinitial.GenerateNameFromUser(u, f.StaffFormat, f.StaffTitleSeparator)
	case "first_last":
	default:
		return firstlast.GenerateNameFromUser(u, f.StaffFormat, f.StaffTitleSeparator)
	}

	return ""
}

func (f *Facility) ProcessMember(s *discordgo.Session, m *discordgo.Member) {
	user, err := f.FindUserByDiscordID(m.User.ID)
	if err != nil && err != ErrUserNotFound {
		log.Errorf("Failed to find user %s: %s", m.User.ID, err)
		return
	}

	if user != nil {
		name := f.GenerateNameFromUser(user)
		err := s.GuildMemberNickname(m.GuildID, m.User.ID, name)
		if err != nil {
			log.Errorf("Failed to set nickname for %s: %s", m.User.Username, err)
		}
	}
	f.ProcessMemberRoles(s, m)
}

func (f *Facility) ProcessMemberRoles(s *discordgo.Session, m *discordgo.Member) {
	user, err := f.FindUserByDiscordID(m.User.ID)
	if err != nil && err != ErrUserNotFound {
		log.Errorf("Failed to find user %s: %s", m.User.ID, err)
		return
	}

	for _, role := range f.Roles {
		roleDisplay := role.ID
		if role.Name != "" {
			roleDisplay = role.Name
		}

		shouldHave := false
		for _, condition := range role.If {
			if f.checkCondition(user, &condition.Condition, &condition.Value) {
				shouldHave = true
				break
			}
		}

		if shouldHave {
			if !slices.Contains(m.Roles, role.ID) {
				log.Infof("Adding role %s to %s", roleDisplay, m.User.Username)
				err := s.GuildMemberRoleAdd(m.GuildID, m.User.ID, role.ID)
				if err != nil {
					log.Errorf("Failed to add role %s to %s: %s", roleDisplay, m.User.Username, err)
				}
			}
		} else {
			if slices.Contains(m.Roles, role.ID) {
				log.Infof("Removing role %s from %s", roleDisplay, m.User.Username)
				err := s.GuildMemberRoleRemove(m.GuildID, m.User.ID, role.ID)
				if err != nil {
					log.Errorf("Failed to remove role %s from %s: %s", roleDisplay, m.User.Username, err)
				}
			}
		}
	}
}

func (f *Facility) checkCondition(user *dto.UserResponse, condition, value *string) bool {
	switch *condition {
	case "controller_type":
		if user == nil {
			return false
		}
		return user.ControllerType == *value
	case "has_role":
		if user == nil {
			return false
		}
		return utils.Contains[string](user.Roles, *value)
	case "rating":
		if user == nil {
			return false
		}
		return user.Rating == *value
	case "unknown":
		switch *value {
		case "true":
			return user == nil
		case "false":
			return user != nil
		default:
			log.Warnf("Invalid role condition value for guild %s: %s value %s. Expecting true or false.", f.Facility, *condition, *value)
			return false
		}
	default:
		log.Warnf("Invalid role condition for guild %s: %s", f.Facility, *condition)
		return false
	}
}
