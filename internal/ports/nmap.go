package ports

import (
	"fmt"
	"os/exec"
)

func RunNmap(host string) {
	fmt.Println("Running nmap...")
	cmd := exec.Command("nmap" /*"-p-", "-sV", "-A",*/, host)

	out, err := cmd.Output()

	if err != nil {
		fmt.Println("Error on run nmap: ", err)
		return
	}

	fmt.Println("nmap output:\n", string(out))
	fmt.Println()
}
