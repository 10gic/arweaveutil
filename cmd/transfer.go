/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/10gic/arweaveutil/util"
	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// transferCmd represents the transfer command
var transferCmd = &cobra.Command{
	Use:   "transfer <target-address> <ar-amount>",
	Short: "Transfer AR to another address",
	Long:  `Transfer AR to another address`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(globalOptWallet) == 0 {
			fmt.Printf("Error: -w (--wallet) is requried for transfer command\n")
			_ = cmd.Help()
			os.Exit(1)
		}

		var targetAddress = args[0]
		var arAmount = args[1]

		if !util.IsValidAddress(targetAddress) {
			fmt.Printf("invalid target-address %v\n", targetAddress)
			return
		}

		winstonAmount, err := util.ARToWinston(arAmount)
		if err != nil {
			fmt.Printf("failed to convert amount from AR to Winston: %v\n", err)
			return
		}

		wallet, err := goar.NewWalletFromPath(globalOptWallet, globalOptNodeUrl)
		if err != nil {
			fmt.Printf("failed to generate wallet from path: %v\n", err)
			return
		}

		// Note that sending tokens to the owner address is not supported in Arweave
		// See https://docs.arweave.org/developers/server/http-api
		if wallet.Address == targetAddress {
			fmt.Printf("target address is same as owner address, can not send AR to yourself\n")
			return
		}

		reward, err := wallet.Client.GetTransactionPrice(nil, &targetAddress)
		if err != nil {
			fmt.Printf("failed to get transaction price: %v\n", err)
			return
		}

		tx := &types.Transaction{
			Format:   2,
			Target:   targetAddress,
			Quantity: winstonAmount,
			Tags:     []types.Tag{},
			Data:     "",
			DataSize: "0",
			Reward:   strconv.FormatInt(reward, 10),
		}

		id, err := wallet.SendTransaction(tx)
		if err != nil {
			fmt.Printf("failed to send transaction: %v\n", err)
			return
		}

		fmt.Printf("send %v AR from %v to %v, tx: %v\n", arAmount, wallet.Address, targetAddress, id)
	},
}

func init() {
	rootCmd.AddCommand(transferCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transferCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transferCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
