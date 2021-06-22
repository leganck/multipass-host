package multipass

import (
	"fmt"
	"os/exec"
	"strings"
)

const defaultInstanceName = "primary"
const commandName = "multipass"

func init()  {
	command := exec.Command("commandName")
	_, err := command.Output()
	if err != nil {
		panic(fmt.Sprintf("%v",err))
	}
}

func listAll() string {
	command := exec.Command(commandName, "list", "--format", "json")
	output, err := command.Output()
	if err != nil {
		panic(fmt.Sprintf("multipass list : %v", err))
	}
	return string(output)
}

func RunCommand(info *InstanceInfo, command string, args ...string) (string, error) {
	cmdStr := fmt.Sprintf("%s %s", command, strings.Join(args, " "))
	cmd := exec.Command(commandName, "exec", (*info).Name, "--", "bash", "-c", cmdStr)
	output, err := cmd.Output()
	if err !=nil {
		return "", fmt.Errorf("instance: %s, command: %s , error: %v",(*info).Name,command,err)
	}
	return string(output), nil
}

func DefaultRunCommand(command string, args ...string) (string, error) {
	instance := GetInstance(defaultInstanceName)
	return RunCommand(instance,command, args...)
}
