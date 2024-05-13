package external_cam

var requestedCam bool

func init() {
	requestedCam = false
}
func GetExternalCam() string {
	defer releaseExternalCam()
	if requestedCam == true {
		return "yes"
	}
	return "no"
}
func RequestExternalCam() {
	requestedCam = true
}

func releaseExternalCam() {
	requestedCam = false
}
