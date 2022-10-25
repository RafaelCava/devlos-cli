/*
Copyright © 2022 Rafael Cavalcante <ra.facavalcante@hotmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// startdevCmd represents the startdev command
var startdevCmd = &cobra.Command{
	Use:   "startdev",
	Short: "Comando para subir o ambiente de desenvolvimento",
	Long: `Comando utilizado para selecionar os serviços para desenvolvimento
	
	Utilizar previamente o comando de configuração da CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		services := []string{}
		options := []string{"Megasac-api", "Broadcast-api", "Tallos-chat", "Tallos-SSO"}
		prompt := &survey.MultiSelect{
			Message: "Quais serviços você deseja utilizar ?",
			Options: options,
		}
		survey.AskOne(prompt, &services)
		fmt.Println(options)
		fmt.Println(services)
	},
}

func init() {
	rootCmd.AddCommand(startdevCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startdevCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startdevCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
