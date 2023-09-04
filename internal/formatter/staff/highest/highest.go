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

package highest

import (
	"github.com/adh-partnership/api/pkg/database/dto"

	"github.com/vpaza/bot/pkg/utils"
)

func TitleFromUser(u *dto.UserResponse, _ string) string {
	if utils.Contains[string](u.Roles, "atm") {
		return "ATM"
	}
	if utils.Contains[string](u.Roles, "datm") {
		return "DATM"
	}
	if utils.Contains[string](u.Roles, "ta") {
		return "TA"
	}
	if utils.Contains[string](u.Roles, "ec") {
		return "EC"
	}
	if utils.Contains[string](u.Roles, "fe") {
		return "FE"
	}
	if utils.Contains[string](u.Roles, "wm") {
		return "WM"
	}
	if utils.Contains[string](u.Roles, "ins") {
		return "INS"
	}
	if utils.Contains[string](u.Roles, "mtr") {
		return "MTR"
	}
	if utils.Contains[string](u.Roles, "events") {
		return "AEC"
	}
	if utils.Contains[string](u.Roles, "facilities") {
		return "AFE"
	}
	if utils.Contains[string](u.Roles, "web") {
		return "AWM"
	}

	return ""
}
