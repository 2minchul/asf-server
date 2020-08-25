package asf_cmd

import (
	"log"
	"os/exec"
)

func StartAsf() {
	log.Println("start asf")
	_, _ = exec.LookPath("systemctl start asf")
}

func StopAsf() {
	log.Println("stop asf")
	_, _ = exec.LookPath("systemctl stop asf")
}
