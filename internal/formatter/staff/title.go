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

package staff

import (
	"github.com/adh-partnership/api/pkg/database/dto"

	"github.com/vpaza/bot/internal/formatter/staff/all"
	"github.com/vpaza/bot/internal/formatter/staff/highest"
	"github.com/vpaza/bot/internal/formatter/staff/none"
)

func GetTitle(u *dto.UserResponse, staffformat, staffsep string) string {
	switch staffformat {
	case "all":
		return all.TitleFromUser(u, staffsep)
	case "none":
		return none.TitleFromUser(u, staffsep)
	case "highest":
		return highest.TitleFromUser(u, staffsep)
	default:
		return highest.TitleFromUser(u, staffsep)
	}
}
