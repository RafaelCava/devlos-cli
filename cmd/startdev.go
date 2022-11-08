/*
Copyright © 2022 Rafael Cavalcante <ra.facavalcante@hotmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

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
		options := SearchProjects()
		prompt := &survey.MultiSelect{
			Message: "Quais serviços você deseja utilizar ?",
			Options: options,
		}
		survey.AskOne(prompt, &services)
		for i := 0; i < len(services); i++ {
			err := StartProject(services[i])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Projeto %s iniciado com sucesso\n", services[i])
		}
	},
}

var WORKDIR = "/home/tallos/Documentos/develop/tallos/"

func SearchProjects() []string {
	out, err := exec.Command("ls", WORKDIR).Output()
	if err != nil {
		log.Fatal(err)
	}
	format := string(out)
	newFormat := strings.Split(format, "\n")
	newFormat = newFormat[:len(newFormat)-1]
	return newFormat
}

func StartProject(project string) error {
	_, err := exec.Command("docker", "compose", "-f", WORKDIR+project+"/docker-compose.yml", "up", "-d").Output()
	if err != nil {
		return err
	}
	return nil
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
