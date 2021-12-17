package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-blockchain/internal/database"
	"os"
)

func balancesCmd() *cobra.Command {
	balancesCmd := &cobra.Command{
		Use:   "balances",
		Short: "Interact with balances (list...).",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageError()
		},
	}

	balancesCmd.AddCommand(balanceListCmd)

	return balancesCmd
}

var balanceListCmd = &cobra.Command{
	Use:   "List",
	Short: "Lists all balances",
	Run: func(cmd *cobra.Command, args []string) {
		state, err := database.NewStateFromDisk()
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}
		defer state.Close()

		fmt.Println("Accounts Balances:")
		fmt.Println("___________________")
		fmt.Println("")

		for account, balance := range state.Balances {
			fmt.Println(fmt.Sprintf("%s: %d", account, balance))
		}
	},
}
