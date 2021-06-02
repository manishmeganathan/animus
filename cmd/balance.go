package cmd

import (
	"fmt"

	"github.com/manishmeganathan/animus/blockchain"
	"github.com/spf13/cobra"
)

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		address, _ := cmd.Flags().GetString("address")

		chain, err := blockchain.AnimateBlockChain()
		if err != nil {
			fmt.Println("Animus Blockchain does not exist! Use 'animus chain create' to create one.")
			chain.Database.Close()
			return
		}

		defer chain.Database.Close()

		balance := 0
		unspenttxos := chain.AccumulateUTXO(address)

		for _, output := range unspenttxos {
			balance += output.Value
		}

		fmt.Printf("Balance of %s: %d\n", address, balance)
	},
}

func init() {
	// Create the 'balance' command
	rootCmd.AddCommand(balanceCmd)

	// Add the 'address' flag to the 'balance' command and mark it as required
	balanceCmd.Flags().StringP("address", "a", "", "address for which to retrieve balance")
	balanceCmd.MarkFlagRequired("address")

}