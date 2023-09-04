package first_last

import (
	"fmt"

	"github.com/adh-partnership/api/pkg/database/dto"
	"github.com/vpaza/bot/internal/formatter/staff/all"
	"github.com/vpaza/bot/internal/formatter/staff/highest"
)

func GenerateNameFromUser(u *dto.UserResponse, staffformat, staffsep string) string {
	title := ""
	switch staffformat {
	case "all":
		title = all.TitleFromUser(u, staffsep)
	case "highest":
	default:
		title = highest.TitleFromUser(u, staffsep)
	}

	if len(fmt.Sprintf("%s %s | %s", u.FirstName, u.LastName, title)) > 32 {
		// Check length of FirstName + LastName Initial + Title
		if len(fmt.Sprintf("%s %s. | %s", u.FirstName, u.LastName[:1], title)) > 32 {
			diff := len(fmt.Sprintf("%s %s. | %s", u.FirstName, u.LastName[:1], title)) - 32 - 3
			return fmt.Sprintf("%s %s. | %s", u.FirstName[:len(u.FirstName)-diff], u.LastName[:1], title)
		}
		return fmt.Sprintf("%s %s. | %s", u.FirstName, u.LastName[:1], title)
	}

	return fmt.Sprintf("%s %s | %s", u.FirstName, u.LastName, title)
}
