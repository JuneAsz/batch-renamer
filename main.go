package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var pathFlag = flag.String("p", "./", "Path flag: example/path/, defaults to current dir (./)")
var nameFlag = flag.String("n", "", "Name flag: pass something")
var printFlag = flag.Bool("print", false, "Print flag: pass it for details (default: not passed)")

func ParseFlags() { // parses the arguments, if nameFlag is missing - errors out
	flag.Parse()

	if *nameFlag == "" {
		fmt.Println("Error! -n (name) flag is required.")
		flag.Usage()
		os.Exit(1)
	}

}

func ReadDir(dir string, name string, printFlag bool) { // takes in path, target name, and printflag
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error! Files not present", err)
		os.Exit(1)
	}

	if printFlag {
		fmt.Println("Files before rename: ")
		PrintFiles(files)
	}

	err = RenameFiles(files, dir, name)
	if err != nil {
		fmt.Println("Error! ", err)
		os.Exit(1)
	}

	if printFlag {
		newFiles, _ := os.ReadDir(dir)
		fmt.Println("Files after rename: ")
		PrintFiles(newFiles)
	}

}

func RenameFiles(files []fs.DirEntry, dir string, name string) error {
	fmt.Println("Renaming files:")
	for i, f := range files {
		formattedName := fmt.Sprintf("%s_%d.txt", name, i+1)
		oldName := filepath.Join(dir, f.Name())
		newName := filepath.Join(dir, formattedName)
		err := os.Rename(oldName, newName)
		if err != nil {
			fmt.Printf("Error renaming '%s' to '%s': %v\n", f.Name(), formattedName, err)
			return err
		}
		fmt.Printf("Renamed: '%s' -> '%s'\n", f.Name(), formattedName)
	}

	return nil
}

func PrintFiles(files []fs.DirEntry) {
	for _, f := range files {
		fmt.Println("File: ", f.Name())
	}

}

func main() {
	ParseFlags()                              // init flags
	ReadDir(*pathFlag, *nameFlag, *printFlag) // do

}
