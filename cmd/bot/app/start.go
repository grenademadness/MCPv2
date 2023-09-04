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

package app

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/adh-partnership/api/pkg/database"
	"github.com/adh-partnership/api/pkg/logger"
	"github.com/urfave/cli/v2"

	"github.com/vpaza/bot/pkg/config"
	"github.com/vpaza/bot/pkg/database/models"
	"github.com/vpaza/bot/pkg/jobs"
)

var log = logger.Logger.WithField("component", "start")

func newStartCommand() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "Start discord bot",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Value:   "config.yaml",
				Usage:   "Load configuration from `FILE`",
				Aliases: []string{"c"},
				EnvVars: []string{"CONFIG"},
			},
			&cli.StringFlag{
				Name:    "facility-configs-path",
				Usage:   "Path of facility configs",
				EnvVars: []string{"FACILITY_CONFIGS_PATH"},
				Value:   "facilities",
			},
		},
		Action: func(c *cli.Context) error {
			log.Infof("Starting bot...")
			log.Infof("config=%s", c.String("config"))

			log.Infof("Loading configuration...")
			_, err := config.ParseConfig(c.String("config"))
			if err != nil {
				return err
			}
			// Walk the facility configs path and load them
			filepath.WalkDir(c.String("facility-configs-path"), func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if d.IsDir() && path != c.String("facility-configs-path") {
					log.Infof("Skipping %s as it's a directory", path)
					return filepath.SkipDir
				}

				if strings.HasSuffix(path, "yaml") || strings.HasSuffix(path, "yml") {
					log.Infof("Loading facility config %s", path)
					_, err = config.ParseFacilityConfig(path)
					if err != nil {
						return err
					}
				}
				return nil
			})

			log.Debugf("FacCfg=%+v", config.FacCfg)

			log.Infof("Building database connection...")
			err = database.Connect(database.DBOptions{
				Host:     config.Cfg.Database.Host,
				Port:     config.Cfg.Database.Port,
				User:     config.Cfg.Database.Username,
				Password: config.Cfg.Database.Password,
				Database: config.Cfg.Database.Database,
				CACert:   config.Cfg.Database.CACert,
				Driver:   "mysql",
				Logger:   logger.Logger,
			})
			if err != nil {
				return err
			}

			log.Infof("Running migrations...")
			err = database.DB.AutoMigrate(
				&models.Bot{},
			)
			if err != nil {
				return err
			}

			log.Infof("Building jobs...")
			jobs.BuildJobs()

			log.Infof("Starting jobs async...")
			jobs.Start()

			return nil
		},
	}
}
