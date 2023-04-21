package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunScript(command string) {
	cmd := exec.Command("bash", "-c", command)
	println("RUN: " + command)

	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		os.Exit(0)
		return
	}

	// Run the command and capture its output
	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting for command to finish:", err)
		os.Exit(0)
		return
	}
}

func RunRemoteScript(targetMachine, command string) {
	cmd := fmt.Sprintf("ssh %s \"%s\"", targetMachine, command)
	RunScript(cmd)
}

func RunRemoteScriptSudo(targetMachine, command string) {
	cmd := fmt.Sprintf("ssh -t %s \"%s\"", targetMachine, command)
	println("\n")
	RunScript(cmd)
}

func RunScriptAndGetOutPut(command string) string {

	cmd := exec.Command("bash", "-c", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "\n"
	}

	return string(output)

}

func GetMasterInfo() (string, string) {
	hostname, _ := os.Hostname()
	command := fmt.Sprintf("cat /etc/hosts | grep %s", hostname)
	result := strings.Split(RunScriptAndGetOutPut(command), " ")
	return result[0], strings.ReplaceAll(result[len(result)-1], "\n", "")
}

func GetRemoteHostname(ip string) string {
	command := fmt.Sprintf("ssh %s \"hostname\"", ip)
	return strings.ReplaceAll(strings.ReplaceAll(RunScriptAndGetOutPut(command), " ", ""), "\n", "")
}
