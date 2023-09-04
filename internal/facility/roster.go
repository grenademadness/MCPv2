package facility

import (
	"encoding/json"
	"fmt"

	"github.com/adh-partnership/api/pkg/database/models"

	"github.com/vpaza/bot/pkg/network"
)

func (f *Facility) GetRoster() []*models.User {
	status, content, err := network.Call(
		"GET",
		fmt.Sprintf("%s/v1/user/all", f.API),
		"application/json",
		nil,
		nil,
	)

	if err != nil {
		log.Errorf("Failed to get roster: %s", err)
		return nil
	}

	if status != 200 {
		log.Errorf("Failed to get roster: %s", content)
		return nil
	}

	var roster []*models.User
	err = json.Unmarshal(content, &roster)
	if err != nil {
		log.Errorf("Failed to unmarshal roster: %s", err)
		return nil
	}

	return roster
}
