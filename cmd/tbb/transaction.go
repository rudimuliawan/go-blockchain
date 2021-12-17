package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-blockchain/internal/database"
	"os"
)

const flagFrom = "from"
const flagTo = "to"
const flagValue = "value"
const flagData = "data"

func transactionCmd() *cobra.Command {
	var trxCmd = &cobra.Command{
		Use:   "trx",
		Short: "Interact with transaction (add...)",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageError()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	trxCmd.AddCommand(trxAddCmd())

	return trxCmd
}

func trxAddCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "add",
		Short: "Add new Transaction to database.",
		Run: func(cmd *cobra.Command, args []string) {
			from, _ := cmd.Flags().GetString(flagFrom)
			to, _ := cmd.Flags().GetString(flagTo)
			value, _ := cmd.Flags().GetUint(flagValue)
			data, _ := cmd.Flags().GetString(flagData)

			trx := database.NewTransaction(database.NewAccount(from), database.NewAccount(to), value, data)

			state, err := database.NewStateFromDisk()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			defer state.Close()

			err = state.Add(trx)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			err = state.Persist()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			fmt.Println("Transaction successfully added to the ledger.")
		},
	}

	cmd.Flags().String(flagFrom, "", "From what account to send tokens")
	cmd.MarkFlagRequired(flagFrom)

	cmd.Flags().String(flagTo, "", "To what account to send tokens")
	cmd.MarkFlagRequired(flagTo)

	cmd.Flags().Uint(flagValue, 0, "How many tokens to send")
	cmd.MarkFlagRequired(flagValue)

	cmd.Flags().String(flagData, "", "Possible values: 'reward'")

	return cmd
}
