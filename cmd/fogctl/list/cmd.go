package list

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cheyang/fog/cluster"
	"github.com/cheyang/fog/persist"
	"github.com/cheyang/fog/util"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:     "list <name>",
		Short:   "list a cluster",
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("cluster name missing")
			}

			if len(args) > 2 {
				return errors.New("remove command takes exactly 1 argument")
			}

			for _, name := range args {
				storePath, err := util.GetStorePath(name)
				if err != nil {
					return err
				}

				if _, err := os.Stat(storePath); os.IsNotExist(err) {
					return fmt.Errorf("Failed to find the storage of cluster %s in %s",
						name,
						filepath.Join(pwd, ".fog"))
				}
				storage := persist.NewFilestore(storePath)

				err = cluster.List(storage)

				if err != nil {
					return err
				}
			}

			return nil
		},
	}
)
