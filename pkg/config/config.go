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

import (
	"errors"
	"os"

	"sigs.k8s.io/yaml"
)

var Cfg *Config
var FacCfg map[string]*FacilityConfig

var (
	ErrorFacilityExists   = errors.New("facility already exists")
	ErrorFacilityNotFound = errors.New("facility not found")
)

func init() {
	FacCfg = make(map[string]*FacilityConfig)
}

func ParseConfig(file string) (*Config, error) {
	config, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(config, cfg)
	if err != nil {
		return nil, err
	}

	Cfg = cfg

	return cfg, nil
}

func ParseFacilityConfig(file string) (*FacilityConfig, error) {
	config, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	cfg := &FacilityConfig{}
	err = yaml.Unmarshal(config, cfg)
	if err != nil {
		return nil, err
	}

	if _, ok := FacCfg[cfg.Facility]; ok {
		return nil, ErrorFacilityExists
	}

	FacCfg[cfg.Facility] = cfg

	return cfg, nil
}

func FindFacilityForGuildId(id string) (*FacilityConfig, error) {
	for _, cfg := range FacCfg {
		if cfg.DiscordID == id {
			return cfg, nil
		}
	}

	return nil, ErrorFacilityNotFound
}
