package interface_to_air

import (
	"fmt"
	"log"
	"os/exec"
)

func SendIRCommand(status string) error {
	device := "/dev/lirc0"
	device2 := "/dev/lirc1"
	err := actualSend(status, device)
	if err != nil {
		return actualSend(status, device2)
	}
	return err
}

func actualSend(status string, device string) error {
	log.Println("executing status: ", status)
	cmd := exec.Command("ir-ctl", "-d", device, "--send="+status)

	if err := cmd.Run(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf(
				"ir-ctl failed with status %v: %s", exiterr.ExitCode(), exiterr.Stderr,
			)
		}
		return fmt.Errorf("ir-ctl failed: %v", err)
	}

	fmt.Println("IR command sent successfully")
	return nil
}
