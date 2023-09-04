package facility

import (
	"errors"
	"os"

	"github.com/adh-partnership/api/pkg/logger"
	"gopkg.in/yaml.v2"
)

type Facility struct {
	Facility            string `json:"facility"`
	DiscordID           string `json:"discord_id"`
	Description         string `json:"description"`
	NameFormat          string `json:"name_format"`
	StaffFormat         string `json:"staff_format"`
	StaffTitleSeperator string `json:"staff_title_seperator"`
	RosterAPI           string `json:"roster_api"`
	API                 string `json:"api"`
	Roles               []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		If   []struct {
			Condition string `json:"condition"`
			Value     string `json:"value"`
		} `json:"if"`
	} `json:"roles"`
}

var FacCfg map[string]*Facility

var (
	ErrorFacilityExists   = errors.New("facility already exists")
	ErrorFacilityNotFound = errors.New("facility not found")
	log                   = logger.Logger.WithField("component", "facility")
)

func init() {
	FacCfg = make(map[string]*Facility)
}

func FindFacility(f *Facility) (*Facility, error) {
	for _, cfg := range FacCfg {
		if cfg.DiscordID == f.DiscordID {
			return cfg, nil
		}
	}

	return nil, ErrorFacilityNotFound
}

func ParseFacilityConfig(file string) (*Facility, error) {
	config, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	cfg := &Facility{}
	err = yaml.Unmarshal(config, cfg)
	if err != nil {
		return nil, err
	}

	if cfg.StaffFormat == "all" && cfg.StaffTitleSeperator == "" {
		cfg.StaffTitleSeperator = "/"
		logger.Logger.WithField("component", "config").Warnf("Staff format is set to 'all' but no staff title seperator is set. Defaulting to '/'")
	}

	if _, ok := FacCfg[cfg.Facility]; ok {
		return nil, ErrorFacilityExists
	}

	FacCfg[cfg.Facility] = cfg

	return cfg, nil
}
