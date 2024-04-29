package notifications

import (
	"log"
	"os"
	"time"

	"github.com/gregdel/pushover"
)

func SendNotification(message string, title string) {

	for attempt := 1; attempt <= 3; attempt++ {
		log.Println("Trying to Send notification")

		app := pushover.New(os.Getenv("PT"))

		// Create a new recipient
		recipient := pushover.NewRecipient(os.Getenv("PU"))

		// Create the message to send
		toSend := pushover.NewMessageWithTitle(message, title)

		// Send the message to the recipient
		_, err := app.SendMessage(toSend, recipient)
		if err != nil {
			log.Printf("Unable to send message: %v\n", err)
		} else {
			log.Println("Message sent successfully")
			return
		}
		time.Sleep(10 * time.Second)
	}
	log.Println("Unable to send message: - giving up")

}
