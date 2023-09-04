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
