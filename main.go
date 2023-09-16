package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("goshell % ")
		text, err := reader.ReadString('\n')
		if err != nil {
			printError("error occured when reading input", err)
		}

		text = strings.Replace(text, "\n", "", 1)

		cmdList := strings.Fields(text)
		if len(cmdList) == 1 {
			cmd := cmdList[0]
			if strings.Compare("exit", cmd) == 0 {
				// exit command
				break
			} else {
				path, err := exec.LookPath(cmd)
				if err != nil {
					printError("error when looking for path", err)
					continue
				}
				args := []string{path}
				var procAttr os.ProcAttr
				procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
				if process, err := os.StartProcess(path, args, &procAttr); err != nil {
					printError("error when running your process", err)
				} else {
					process.Wait()
				}
			}

		} else {
			cmd := cmdList[0]
			// check for background process
			if strings.Compare("&", cmdList[len(cmdList)-1]) == 0 {
				fmt.Println("working on background things")
			} else {
				path, err := exec.LookPath(cmd)
				if err != nil {
					printError("error when looking for path", err)
					continue
				}
				args := cmdList
				args[0] = path
				var procAttr os.ProcAttr
				procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
				if process, err := os.StartProcess(path, args, &procAttr); err != nil {
					printError("error when running your process", err)
				} else {
					process.Wait()
				}
			}
		}
	}
}

func printError(message string, err error) {
	fmt.Println(message)
	fmt.Println(err.Error())
}
