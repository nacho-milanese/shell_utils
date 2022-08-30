package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var autocreate bool //= true
var quiet bool      //= false
var dirName string

func main() {
	var exists bool

	exists = checkVenvDir(dirName)
	if exists || autocreate {
		createVirtualenv(dirName)
		runmeAfter(dirName, quiet)
	}

}

func init() {
	// This passes parameters to the script...
	flag.StringVar(&dirName, "d", "venv", "Specify virtualenv dir, default is venv")
	flag.BoolVar(&autocreate, "x", false, "Autocreate if doesn't exists?")
	flag.BoolVar(&quiet, "q", false, "do not print silly usage by default")

	flag.Usage = func() {
		fmt.Printf("Usage: \n")
		fmt.Printf("%s -x [-d VIRTUALENV_DIR]\n", os.Args[0])
	}
	flag.Parse()
}

func checkVenvDir(dirName string) bool {
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func createVirtualenv(dirName string) {
	python3, err_py := exec.LookPath("python3")
	if err_py != nil {
		log.Fatal("Python not found")
	}
	_, err := exec.Command(python3, "-m", "virtualenv", "-v", dirName).Output()
	if err != nil {
		log.Fatal(err)
	}
}

func runmeAfter(dirName string, quiet bool) {
	if !quiet {
		fmt.Printf("To activate, run: \nsource %s/bin/activate\n", dirName)
		fmt.Println("Then run 'deactivate' to stop the virtualenv")
	}
}
