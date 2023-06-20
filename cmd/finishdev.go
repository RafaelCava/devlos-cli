/*
Copyright © 2022 Rafael Cavalcante <ra.facavalcante@hotmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var finishdevCmd = &cobra.Command{
	Use:   "finishdev",
	Short: "Comando para finalizar o ambiente de desenvolvimento",
	Long: `Comando utilizado para finalizar o ambiente de desenvolvimento
	
	Selecione os serviços que vão ser finalizados`,
	Run: func(cmd *cobra.Command, args []string) {
		containersToRemove := []string{}
		output := SearchContainerUp()
		containers := FormatContainers(output)
		if len(containers) <= 0 {
			fmt.Println("Não há containers para remover")
			return
		}
		prompt := &survey.MultiSelect{
			Message: "Quais serviços você deseja finalizar ?",
			Options: containers,
		}
		survey.AskOne(prompt, &containersToRemove)
		for i := 0; i < len(containersToRemove); i++ {
			FinishContainer(containersToRemove[i])
			fmt.Printf("Container %s finalizado com sucesso\n", containersToRemove[i])
		}
	},
}

func SearchContainerUp() []byte {
	out, err := exec.Command("docker", "ps", "--format", "{{.Names}}").Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func FormatContainers(output []byte) []string {
	format := string(output)
	newFormat := strings.Split(format, "\n")
	newFormat = newFormat[:len(newFormat)-1]
	return newFormat
}

func FinishContainer(container string) {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()
	_, err := exec.Command("docker", "rm", "-f", container).Output()
	if err != nil {
		log.Fatal(err)
	}
	s.Stop()
}

func init() {
	rootCmd.AddCommand(finishdevCmd)
}
