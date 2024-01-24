/*
Copyright © 2022 Rafael Cavalcante <ra.facavalcante@hotmail.com>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
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
	Aliases: []string{"up"},
}

var WORKDIR = os.Getenv("PATH_TALLOS")

func SearchProjects() []string {
	out, err := exec.Command("ls", WORKDIR).Output()
	if err != nil {
		fmt.Println("Erro ao executar o comando 'ls':", err)
		log.Fatal(err)
	}
	directoriesStringfy := string(out)
	newDirectoriesFormat := strings.Split(directoriesStringfy, "\n")
	newDirectoriesFormat = newDirectoriesFormat[:len(newDirectoriesFormat)-1]
	validDirectories := ValidateDir(newDirectoriesFormat)
	services := FindServices(validDirectories, "-parent")
	return services
}

func StartProject(project string) error {
	format := strings.Split(project, " -> ")
	if len(format) > 1 {
		s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
		s.Start()
		_, err := exec.Command("docker", "compose", "-f", WORKDIR+"/"+format[1]+"/docker-compose.yml", "up", "-d", format[0]).Output()
		if err != nil {
			return err
		}
		s.Stop()
		return nil
	} else {
		s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
		s.Start()
		_, err := exec.Command("docker", "compose", "-f", WORKDIR+"/"+project+"/docker-compose.yml", "up", "-d").Output()
		if err != nil {
			return err
		}
		s.Stop()
		return nil
	}
}

func FindServices(arr []string, regexPattern string) []string {
	var services []string
	regex := regexp.MustCompile(regexPattern)
	for _, item := range arr {
		if !regex.MatchString(item) {
			services = append(services, item)
		} else {
			out, err := exec.Command("docker", "compose", "-f", WORKDIR+"/"+item+"/docker-compose.yml", "config", "--services").Output()
			if err != nil {
				log.Fatal(err)
			}
			format := string(out)
			formatWithParent := format + item
			newFormat := strings.Split(formatWithParent, "\n")
			newFormat = newFormat[:len(newFormat)-1]
			for _, service := range newFormat {
				services = append(services, service+" -> "+item)
			}
		}
	}
	return services
}

func ValidateDir(directories []string) []string {
	var validDirectories []string
	for _, directory := range directories {
		files, err := ioutil.ReadDir(filepath.Join(WORKDIR, directory))
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			if !file.IsDir() && strings.HasPrefix(file.Name(), "docker-compose") {
				validDirectories = append(validDirectories, directory)
				break
			}
		}
	}
	return validDirectories
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
