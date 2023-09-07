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
