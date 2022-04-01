/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/10gic/arweaveutil/util"
	"github.com/everFinance/goar"
	"math/big"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// transferPstCmd represents the transferPst command
var transferPstCmd = &cobra.Command{
	Use:   "transfer-pst <pst-symbol-or-address> <target-address> <pst-amount>",
	Short: "Transfer PST token to another address",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		if len(globalOptWallet) == 0 {
			fmt.Printf("Error: -w (--wallet) is requried for transfer command\n")
			_ = cmd.Help()
			os.Exit(1)
		}

		var pstSymbolOrAddress = args[0]
		var targetAddress = args[1]
		var pstAmount = args[2]

		if value, ok := util.PopularPSTAddress[strings.ToUpper(pstSymbolOrAddress)]; ok {
			pstSymbolOrAddress = value
		} else {
			// Not found in PopularPSTAddress, so pstSymbolOrAddress must be a pst address
			if !util.IsValidAddress(pstSymbolOrAddress) {
				fmt.Printf("invalid pst-symbol-or-address %v\n", pstSymbolOrAddress)
				return
			}
		}

		if !util.IsValidAddress(targetAddress) {
			fmt.Printf("invalid target-address %v\n", targetAddress)
			return
		}

		// Convert pstAmount to big int
		pstAmountBigInt := new(big.Int)
		pstAmountBigInt, ok := pstAmountBigInt.SetString(pstAmount, 10)
		if !ok {
			fmt.Printf("invalid pst-amount %v\n", pstAmount)
			return
		}

		wallet, err := goar.NewWalletFromPath(globalOptWallet, globalOptNodeUrl)
		if err != nil {
			fmt.Printf("failed to generate wallet from path: %v\n", err)
			return
		}

		tx, err := wallet.SendPst(pstSymbolOrAddress, targetAddress, pstAmountBigInt, nil, 0)
		if err != nil {
			fmt.Printf("failed to send pst: %v\n", err)
			return
		}

		fmt.Printf("send %v PST (%v) from %v to %v, tx: %v\n", pstAmount, pstSymbolOrAddress, wallet.Address, targetAddress, tx.ID)
	},
}

func init() {
	rootCmd.AddCommand(transferPstCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transferPstCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transferPstCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
