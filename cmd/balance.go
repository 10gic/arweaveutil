/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/10gic/arweaveutil/util"
	"github.com/everFinance/goar"

	"github.com/spf13/cobra"
)

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance <address>",
	Short: "Gets the specified address's balance",
	Long:  `Gets the specified address's balance`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var address = args[0]

		if !util.IsValidAddress(address) {
			fmt.Printf("invalid address %v\n", address)
			return
		}

		client := goar.NewClient(globalOptNodeUrl)
		balance, err := client.GetWalletBalance(address)
		if err != nil {
			fmt.Printf(`failed to get balance: %v\n`, err)
		}
		fmt.Printf("address %v, balance %v AR\n", address, balance.String())
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// balanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// balanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
