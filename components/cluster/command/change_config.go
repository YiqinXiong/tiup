// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package command

import (
	"github.com/spf13/cobra"
)

func newChangeConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "change-config <cluster-name> <config.yaml>",
		Short: "Change TiDB cluster config.\nWill use configuration defined by config.yaml.\n",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return cmd.Help()
			}

			clusterName := args[0]
			clusterReport.ID = scrubClusterName(clusterName)
			teleCommand = append(teleCommand, scrubClusterName(clusterName))
			configFile := args[1]

			return cm.ChangeConfig(clusterName, configFile, skipConfirm)
		},
	}

	return cmd
}
