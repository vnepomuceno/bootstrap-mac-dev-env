package main

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

// getFormulasFromFile returns an array of strings by reading a file with
// formulas in each line of the file.
func getFormulasFromFile(filename string) []string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	formulasArray := strings.Split(string(file), "\n")
	log.Println(fmt.Sprintf("Loaded formulas to install: %s", formulasArray))

	return formulasArray
}

// installFormulas executes the command `brew install <formula>` for every formula
// in the `formulas.txt` file.
func installFormulas(dryRun bool) {
	formulas := getFormulasFromFile("./homebrew/formulas.txt")
	log.Println(fmt.Sprintf("Installing %d homebrew formulas...", len(formulas)))

	formulaSize := len(formulas)
	bar := progressbar.Default(int64(formulaSize))
	for i := 0; i < formulaSize; i++ {
		log.Println(fmt.Sprintf("Installing formula %s", formulas[i]))

		installCommand := exec.Command("brew", "install", formulas[i])
		var commandError error
		if !dryRun {
			commandError = installCommand.Run()
		} else {
			time.Sleep(1 * time.Second)
		}
		if commandError != nil {
			log.Fatal(commandError)
		}
		barError := bar.Add(1)
		if barError != nil {
			log.Fatal(barError)
		}
	}
}

func isDryRun() bool {
	log.Println(fmt.Sprintf("Arguments passed: %s", os.Args))
	if len(os.Args) == 1 {
		return false
	} else if os.Args[1] == "--dry-run" {
		log.Println("Running command as DRY RUN")
		return true
	} else {
		return false
	}
}

func main() {
	installFormulas(isDryRun())
}
