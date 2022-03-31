/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/10gic/arweaveutil/util"
	"github.com/spf13/cobra"
	"os"
)

// genWalletCmd represents the genWallet command
var genWalletCmd = &cobra.Command{
	Use:   "gen-wallet",
	Short: "Generate a wallet (jwk file)",
	Long:  `Generate a wallet, save the generated jwk file in current directory`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := util.GenerateRsaJwk()
		if err != nil {
			fmt.Printf(`failed to generate RSA jwk: %v\n`, err)
			return
		}
		bytes, _ := json.Marshal(key)

		// Compute the Arweave address
		address, err := util.JwkToAddress(bytes)
		if err != nil {
			fmt.Printf(`failed to get address from jwk: %v\n`, err)
			return
		}

		// Save keyfile
		fileName := "arweave-keyfile-" + address + ".json"
		err = os.WriteFile(fileName, bytes, 0644)
		if err != nil {
			fmt.Printf(`failed to write file %s: %v\n`, fileName, err)
			return
		}
		fmt.Printf("keyfile %v saved\n", fileName)
	},
}

func init() {
	rootCmd.AddCommand(genWalletCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genWalletCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genWalletCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
