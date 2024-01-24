/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// normalizeCmd represents the normalize command
var normalizeCmd = &cobra.Command{
	Use:   "normalize",
	Short: "Comando para normalizar uma branch selecionada",
	Long: `Comando criado para normalizar uma branch selecionada
	Após executar esse comando será aberto a seleção dos repos.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("normalize called")
	},
	Aliases: []string{"norm"},
}

func init() {
	rootCmd.AddCommand(normalizeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// normalizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// normalizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
