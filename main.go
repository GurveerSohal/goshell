package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)


type Container struct{
	mu sync.Mutex
	Table map[int]string 
}

func (c *Container) add(name string, pid int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Table[pid] = name
}

func (c *Container) remove(pid int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.Table, pid)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	c := Container{
		Table: make(map[int]string),
	}

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
			} else if strings.Compare("list", cmd) == 0 {
				if (len(c.Table) == 0) {
					fmt.Println("No processes currently running")
				}
				fmt.Printf("PID\tName\t\n")
				for k, v := range c.Table {
					fmt.Printf("%d\t%s\t\n", k, v)
				}
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
			if strings.Compare("kill", cmd) == 0 {
				
			}
			// check for background process
			if strings.Compare("&", cmdList[len(cmdList)-1]) == 0 {
				path, err := exec.LookPath(cmd)
				if err != nil {
					printError("error when looking for path", err)
					continue
				}
				args := cmdList
				args[0] = path
				var procAttr os.ProcAttr
				procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
				go runBackgroundProcess(path, args, &procAttr, text, &c)
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

func runBackgroundProcess(path string, args []string, procAttr *os.ProcAttr, name string, c *Container) {
	if process, err := os.StartProcess(path, args, procAttr); err != nil {
		printError("error when running your process", err)
	} else {
		c.add(name, process.Pid)
		fmt.Println(c.Table[process.Pid])
		process.Wait()
		fmt.Printf("done with process %d\n", process.Pid)
		c.remove(process.Pid)
	}
}

func printError(message string, err error) {
	fmt.Println(message)
	fmt.Println(err.Error())
}
