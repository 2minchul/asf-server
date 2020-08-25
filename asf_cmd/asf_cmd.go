package asf_cmd

import (
	"log"
	"os/exec"
)

func StartAsf() {
	log.Println("start asf")
	out, _ := exec.Command("systemctl", "start", "asf").CombinedOutput()
	log.Println(string(out))
}

func StopAsf() {
	log.Println("stop asf")
	out, _ := exec.Command("systemctl", "stop", "asf").CombinedOutput()
	log.Println(string(out))
}
