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
	"fmt"

	"github.com/mongodb/mcli/internal/config"
	"github.com/mongodb/mcli/internal/convert"
	"github.com/mongodb/mcli/internal/flags"
	"github.com/mongodb/mcli/internal/messages"
	"github.com/mongodb/mcli/internal/search"
	"github.com/mongodb/mcli/internal/store"
	"github.com/mongodb/mcli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type cmClustersUpdateOpts struct {
	*globalOpts
	filename string
	fs       afero.Fs
	store    store.AutomationStore
}

func (opts *cmClustersUpdateOpts) init() error {
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

func (opts *cmClustersUpdateOpts) Run() error {
	newConfig, err := convert.NewClusterConfigFromFile(opts.fs, opts.filename)
	if err != nil {
		return err
	}
	current, err := opts.store.GetAutomationConfig(opts.ProjectID())

	if err != nil {
		return err
	}

	if !search.ClusterExists(current, newConfig.Name) {
		return fmt.Errorf("cluster '%s' doesn't exist", newConfig.Name)
	}

	err = newConfig.PatchAutomationConfig(current)

	if err != nil {
		return err
	}

	if err = opts.store.UpdateAutomationConfig(opts.ProjectID(), current); err != nil {
		return err
	}

	fmt.Print(messages.DeploymentStatus(config.OpsManagerURL(), opts.ProjectID()))

	return nil
}

// mcli cloud-manager cluster(s) update --projectId projectId --file myfile.yaml
func CloudManagerClustersUpdateBuilder() *cobra.Command {
	opts := &cmClustersUpdateOpts{
		globalOpts: newGlobalOpts(),
		fs:         afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a Cloud Manager cluster.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.filename, flags.File, flags.FileShort, "", "Filename to use to update the cluster")

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flags.File)

	return cmd
}
