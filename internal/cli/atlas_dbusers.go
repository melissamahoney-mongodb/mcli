// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"github.com/spf13/cobra"
)

func AtlasDBUsersBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "dbusers",
		Aliases: []string{"dbuser", "databaseUsers", "databaseUser"},
		Short:   "Manage database users for your project.",
		Long: `
The dbusers command lets you retrieve, create and modify the MongoDB users in your cluster.
Each user has a set of roles that provide access to the project’s databases. 
A user’s roles apply to all the clusters in the project.`,
	}
	cmd.AddCommand(AtlasDBUsersCreateBuilder())

	return cmd
}
