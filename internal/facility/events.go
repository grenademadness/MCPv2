package facility

import (
	"encoding/json"
	"fmt"

	"github.com/adh-partnership/api/pkg/database/dto"

	"github.com/vpaza/bot/pkg/network"
)

func (f *Facility) GetEvents() []*dto.EventsResponse {
	status, content, err := network.Call(
		"GET",
		fmt.Sprintf("%s/v1/events?limit=5", f.API),
		"application/json",
		nil,
		nil,
	)
	if err != nil {
		log.Errorf("Failed to get events for %s: %s", f.Facility, err)
		return nil
	}

	if status != 200 {
		log.Errorf("Failed to get events for %s: %s", f.Facility, content)
		return nil
	}

	var events []*dto.EventsResponse
	err = json.Unmarshal(content, &events)
	if err != nil {
		log.Errorf("Failed to unmarshal events for %s: %s", f.Facility, err)
		return nil
	}

	return events
}

func (f *Facility) GetEvent(id int64) *dto.EventsResponse {
	status, content, err := network.Call(
		"GET",
		fmt.Sprintf("%s/v1/events/%d", f.API, id),
		"application/json",
		nil,
		nil,
	)
	if err != nil {
		log.Errorf("Failed to get event for %s: %s", f.Facility, err)
		return nil
	}

	if status != 200 {
		log.Errorf("Failed to get event for %s: %s", f.Facility, content)
		return nil
	}

	var event *dto.EventsResponse
	err = json.Unmarshal(content, &event)
	if err != nil {
		log.Errorf("Failed to unmarshal event for %s: %s", f.Facility, err)
		return nil
	}

	return event
}
