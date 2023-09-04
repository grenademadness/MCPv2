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
	"errors"

	"github.com/adh-partnership/api/pkg/logger"
	"github.com/urfave/cli/v2"
)

func NewRootCommand() *cli.App {
	return &cli.App{
		Name:  "app",
		Usage: "PAZA Discord Bot",
		Commands: []*cli.Command{
			newStartCommand(),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "log-level",
				Value:   "info",
				Usage:   "Set the logging level",
				EnvVars: []string{"LOG_LEVEL"},
				Aliases: []string{"l"},
			},
			&cli.StringFlag{
				Name:    "log-format",
				Value:   "text",
				Usage:   "Set the logging format",
				EnvVars: []string{"LOG_FORMAT"},
				Aliases: []string{"f"},
			},
		},
		Before: func(c *cli.Context) error {
			format := c.String("log-format")
			if !logger.IsValidFormat(format) {
				return errors.New("invalid log format")
			}
			logger.NewLogger(format)

			if !logger.IsValidLogLevel(c.String("log-level")) {
				return errors.New("invalid log level")
			}

			l, _ := logger.ParseLogLevel(c.String("log-level"))
			logger.Logger.SetLevel(l)

			return nil
		},
	}
}
