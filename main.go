package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	argsWithoutProg := os.Args[1:]
	wg.Add(len(argsWithoutProg))

	for _, arg := range argsWithoutProg {
		go func(cmdArg string, wg *sync.WaitGroup) {
			cmdInfo := execCommand(cmdArg)

			fmt.Print(cmdInfo.stdout)
			fmt.Println("Time taken: ", cmdInfo.timeTaken)

			wg.Done()
		}(arg, &wg)
	}

	wg.Wait()
}

type CommandInfo struct {
	stdout    string
	timeTaken time.Duration
}

func execCommand(commandWithArgs string) CommandInfo {
	cmdWithArgs := strings.Split(commandWithArgs, " ")
	command := cmdWithArgs[0]
	args := cmdWithArgs[1:]

	start := time.Now()
	cmd := exec.Command(command, args...)
	stdout, err := cmd.Output()
	end := time.Now()

	if err != nil {
		fmt.Println(err.Error())
		return CommandInfo{}
	}

	return CommandInfo{stdout: string(stdout), timeTaken: end.Sub(start)}
}
