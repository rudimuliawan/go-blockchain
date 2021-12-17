package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var cmd = &cobra.Command{
		Use:   "tbb",
		Short: "The Blockchain Bar CLI",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cmd.AddCommand(versionCmd)
	cmd.AddCommand(balancesCmd())
	cmd.AddCommand(transactionCmd())

	err := cmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func incorrectUsageError() error {
	return fmt.Errorf("incorrect usage")
}
