package facility

import (
	"github.com/adh-partnership/api/pkg/database/dto"
	"github.com/vpaza/bot/internal/formatter/names/first_last"
	"github.com/vpaza/bot/internal/formatter/names/first_last_initial"
)

func (f *Facility) GenerateNameFromUser(u *dto.UserResponse) string {
	switch f.NameFormat {
	case "first_last_initial":
		return first_last_initial.GenerateNameFromUser(u, f.StaffFormat, f.StaffTitleSeperator)
	case "first_last":
	default:
		return first_last.GenerateNameFromUser(u, f.StaffFormat, f.StaffTitleSeperator)
	}

	return ""
}
