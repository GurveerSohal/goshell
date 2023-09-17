package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
)

type LogFile struct {
	mu sync.Mutex
	f *os.File
}

func (logFile *LogFile) writeToFile(s string) {
	logFile.mu.Lock()
	defer logFile.mu.Unlock()
	logFile.f.WriteString(s)
}

type Container struct{
	mu sync.Mutex
	table map[int]string 
}

func (c *Container) add(name string, pid int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.table[pid] = name
}

func (c *Container) remove(pid int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.table, pid)
}
func main() {
	pid := os.Getpid()

	f, err := os.Create(fmt.Sprintf(".log%d", pid))
	if err != nil {
		printError("error when opening log file", err)
		return
	}

	f.WriteString(fmt.Sprintf("Writing logs for goshell, pid: %d...\n", pid))

	logFile := LogFile{
		f: f,
	}

	c := Container{
		table: make(map[int]string),
	}

	// this will close the file too once every process is done writing to it
	defer killRemainingProcesses(&c, &logFile)

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
				logFile.writeToFile("Exiting the shell...\n")
				break
			} else if strings.Compare("list", cmd) == 0 {
				logFile.writeToFile("Printed the processes table...\n")
				if (len(c.table) == 0) {
					fmt.Println("No processes currently running")
				}
				fmt.Printf("PID\tName\t\n")
				for k, v := range c.table {
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
					logFile.writeToFile(fmt.Sprintf("Running a process...\n%s\n", text))
					process.Wait()
					logFile.writeToFile(fmt.Sprintf("Done with a process...\n%s\npid: %d\n", text, process.Pid))
				}
			}

		} else {
			cmd := cmdList[0]
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
				go runBackgroundProcess(path, args, &procAttr, text, &c, &logFile)
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
					logFile.writeToFile(fmt.Sprintf("Running a process...\n%s\n", text))
					process.Wait()
					logFile.writeToFile(fmt.Sprintf("Done with a process...\n%s\npid: %d\n", text, process.Pid))
				}
			}
		}
	}
}

func runBackgroundProcess(path string, args []string, procAttr *os.ProcAttr, name string, c *Container, logFile *LogFile) {
	if process, err := os.StartProcess(path, args, procAttr); err != nil {
		printError("error when running your process", err)
	} else {
		c.add(name, process.Pid)
		logFile.writeToFile(fmt.Sprintf("Running a process...\n%s\n", name))
		process.Wait()
		logFile.writeToFile(fmt.Sprintf("Done with a process...\n%s\npid: %d\n", name, process.Pid))
		c.remove(process.Pid)
	}
}

func killRemainingProcesses(c *Container, logFile *LogFile) {
	defer logFile.f.Close()
	if (len(c.table) == 0) {
		return
	}

	// for some reason these processes don't write to the log file
	// need to figure out why but it's a nitpick
	for k := range c.table {
		if err := syscall.Kill(k ,syscall.SIGTERM); err != nil {
			printError("error when killing a process", err)
		}
	}
}

func printError(message string, err error) {
	fmt.Println(message)
	fmt.Println(err.Error())
}
