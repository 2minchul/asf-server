package asf_cmd

import (
	"log"
	"os/exec"
)

func StartAsf() {
	log.Println("start asf")
	_, _ = exec.Command("systemctl", "start", "asf").CombinedOutput()
}

func StopAsf() {
	log.Println("stop asf")
	_, _ = exec.Command("systemctl", "stop", "asf").CombinedOutput()
}
