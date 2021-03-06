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
	"errors"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mcli/internal/flags"
	"github.com/mongodb/mcli/internal/json"
	"github.com/mongodb/mcli/internal/store"
	"github.com/mongodb/mcli/internal/usage"
	"github.com/spf13/cobra"
)

const (
	adminDB = "admin"
	roleSep = "@"
)

type atlasDBUsersCreateOpts struct {
	*globalOpts
	username string
	password string
	roles    []string
	store    store.DatabaseUserCreator
}

func (opts *atlasDBUsersCreateOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	s, err := store.New()

	if err != nil {
		return err
	}

	opts.store = s
	return nil
}

func (opts *atlasDBUsersCreateOpts) Run() error {
	user := opts.newDatabaseUser()
	result, err := opts.store.CreateDatabaseUser(user)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *atlasDBUsersCreateOpts) newDatabaseUser() *atlas.DatabaseUser {
	return &atlas.DatabaseUser{
		DatabaseName: adminDB,
		Roles:        opts.buildRoles(),
		GroupID:      opts.ProjectID(),
		Username:     opts.username,
		Password:     opts.password,
	}
}

func (opts *atlasDBUsersCreateOpts) buildRoles() []atlas.Role {
	rolesLen := len(opts.roles)
	roles := make([]atlas.Role, rolesLen)
	for i, roleP := range opts.roles {
		role := strings.Split(roleP, roleSep)
		roleName := role[0]
		databaseName := adminDB
		if len(role) > 1 {
			databaseName = role[1]
		}

		roles[i] = atlas.Role{
			RoleName:     roleName,
			DatabaseName: databaseName,
		}
	}
	return roles
}

func (opts *atlasDBUsersCreateOpts) Prompt() error {
	if opts.password != "" {
		return nil
	}
	prompt := &survey.Password{
		Message: "Password:",
	}
	return survey.AskOne(prompt, &opts.password)
}

// mcli atlas dbuser(s) create --username username --password password --role roleName@dbName [--projectId projectId]
func AtlasDBUsersCreateBuilder() *cobra.Command {
	opts := &atlasDBUsersCreateOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:       "create",
		Short:     "Create a database user for a project.",
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"atlasAdmin", "readWriteAnyDatabase", "readAnyDatabase", "clusterMonitor", "backup", "dbAdminAnyDatabase", "enableSharding"},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.init(); err != nil {
				return err
			}
			if len(args) == 0 && len(opts.roles) == 0 {
				return errors.New("no role specified for the user")
			}
			opts.roles = append(opts.roles, args...)
			return opts.Prompt()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.username, flags.Username, "", usage.Username)
	cmd.Flags().StringVar(&opts.password, flags.Password, "", usage.Password)
	cmd.Flags().StringSliceVar(&opts.roles, flags.Role, []string{}, usage.Roles)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flags.Username)

	return cmd
}
