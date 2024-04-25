package interface_to_air

import (
	"fmt"
	"os/exec"
)

func SendIRCommand(status string) error {
	device := "/dev/lirc0"

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
