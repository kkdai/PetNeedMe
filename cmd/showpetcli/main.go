package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	logging "github.com/op/go-logging"
	"github.com/spf13/cobra"

	pnm "github.com/kkdai/pnm"
	pet "github.com/kkdai/pnm/opendata"
)

func printPet(p *pet.Pet) {
	fmt.Println(" -----------------------------------------")
	fmt.Println(" 名字:", p.Name)
	fmt.Println(" 敘述:", p.Note)
	fmt.Println(" 品種:", p.Type)
	fmt.Println(" 圖片:", p.ImageName)
	fmt.Println(" -----------------------------------------")
}

func main() {
	var format = logging.MustStringFormatter("%{level} %{message}")
	logging.SetFormatter(format)
	logging.SetLevel(logging.INFO, "PetNeedMeCLI")

	rootCmd := &cobra.Command{
		Use:   "PetNeedMeCLI",
		Short: "Display All",
		Run: func(cmd *cobra.Command, args []string) {
			pets := pnm.NewPets()

			scanner := bufio.NewScanner(os.Stdin)
			quit := false

			for !quit {
				fmt.Print("Show Me Pets> ")

				if !scanner.Scan() {
					break
				}

				line := scanner.Text()
				parts := strings.Split(line, " ")
				cmd := parts[0]
				args := parts[1:]

				switch cmd {
				case "quit":
					quit = true
				case "n":
					printPet(pets.GetNextPet())
				case "p":
					printPet(pets.GetPreviousPet())
				case "d": //show next dog
					printPet(pets.GetNextDog())
				case "c": //show next cat
					printPet(pets.GetNextCat())
				default:
					fmt.Println("Unrecognized command:", cmd, args)
				}
			}
		},
	}

	rootCmd.Execute()
}
