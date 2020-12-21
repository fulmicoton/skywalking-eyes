// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package commands

import (
	headercommand "github.com/apache/skywalking-eyes/license-eye/commands/header"
	"github.com/apache/skywalking-eyes/license-eye/internal/logger"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	verbosity string
)

// Root represents the base command when called without any subcommands
var Root = &cobra.Command{
	Use:           "license-eye command [flags]",
	Long:          "A full-featured license guard to check and fix license headers and dependencies' licenses",
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		level, err := logrus.ParseLevel(verbosity)
		if err != nil {
			return err
		}
		logger.Log.SetLevel(level)
		return nil
	},
	Version: version,
}

// Execute sets flags to the root command appropriately.
// This is called by main.main(). It only needs to happen once to the Root.
func Execute() error {
	Root.PersistentFlags().StringVarP(&verbosity, "verbosity", "v", logrus.InfoLevel.String(), "log level (debug, info, warn, error, fatal, panic")

	Root.AddCommand(headercommand.Header)

	return Root.Execute()
}