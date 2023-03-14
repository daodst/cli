package cmd

import "github.com/spf13/cobra"

func DeployCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "deploy contract",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}
