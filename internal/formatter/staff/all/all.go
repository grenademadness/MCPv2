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

package all

import (
	"strings"

	"github.com/adh-partnership/api/pkg/database/dto"

	"github.com/vpaza/bot/pkg/utils"
)

func TitleFromUser(u *dto.UserResponse, sep string) string {
	title := []string{}

	if utils.Contains[string](u.Roles, "atm") {
		title = append(title, "ATM")
	}
	if utils.Contains[string](u.Roles, "datm") {
		title = append(title, "DATM")
	}
	if utils.Contains[string](u.Roles, "ta") {
		title = append(title, "TA")
	}
	if utils.Contains[string](u.Roles, "ec") {
		title = append(title, "EC")
	}
	if utils.Contains[string](u.Roles, "fe") {
		title = append(title, "FE")
	}
	if utils.Contains[string](u.Roles, "wm") {
		title = append(title, "WM")
	}
	if utils.Contains[string](u.Roles, "ins") {
		title = append(title, "INS")
	}
	if utils.Contains[string](u.Roles, "mtr") {
		title = append(title, "MTR")
	}
	if utils.Contains[string](u.Roles, "events") {
		title = append(title, "AEC")
	}
	if utils.Contains[string](u.Roles, "facilities") {
		title = append(title, "AFE")
	}
	if utils.Contains[string](u.Roles, "web") {
		title = append(title, "AWM")
	}

	return strings.Join(title, sep)
}
