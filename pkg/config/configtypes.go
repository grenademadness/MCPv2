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

package config

type Config struct {
	Database ConfigDatabase `json:"database"`
	Discord  ConfigDiscord  `json:"discord"`
}

type ConfigDatabase struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	Username    string `json:"user"`
	Password    string `json:"password"`
	Database    string `json:"database"`
	AutoMigrate bool   `json:"auto_migrate"`
	CACert      string `json:"ca_cert"`
}

type ConfigDiscord struct {
	AppID string `json:"app_id"`
	Token string `json:"token"`
}

type FacilityConfig struct {
	Facility    string `json:"facility"`
	DiscordID   string `json:"discord_id"`
	Description string `json:"description"`
	NameFormat  string `json:"name_format"`
	RosterAPI   string `json:"roster_api"`
	API         string `json:"api"`
	Roles       []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		If   []struct {
			Condition string `json:"condition"`
			Value     string `json:"value"`
		} `json:"if"`
	} `json:"roles"`
}
