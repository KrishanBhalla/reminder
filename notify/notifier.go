package notify

import (
	"github.com/cschomburg/go-pushbullet"
	"github.com/gen2brain/beeep"
)

// Notifier is a generic interface for implementing a notification
// TODO: Consider updating to use github.com/nikoksr/notify?
type Notifier interface {
	Notify(title, message string) error
}

// Desktop wraps  github.com/gen2brain/beeep
type Desktop struct {
}

// Notify creates a desktop notification
func (d *Desktop) Notify(title, message string) error {
	return beeep.Notify(title, message, "assets/reminder.png")
}

// Pushbullet struct holds necessary data to communicate with the Pushbullet API.
type Pushbullet struct {
	client          *pushbullet.Client
	deviceNicknames []string
}

// NewPushbullet returns a new instance of a Pushbullet notification
// service. For more information about Pushbullet api token:
//    -> https://docs.pushbullet.com/#api-overview
func NewPushbullet(apiToken string, deviceNicknames ...string) *Pushbullet {
	client := pushbullet.New(apiToken)

	pb := &Pushbullet{
		client:          client,
		deviceNicknames: deviceNicknames,
	}

	return pb
}

// Notify will send the push notification to all receivers.
// you will need Pushbullet installed on the relevant devices
// (android, chrome, firefox, windows)
// see https://www.pushbullet.com/apps
func (pb Pushbullet) Notify(title, message string) error {
	for _, deviceNickname := range pb.deviceNicknames {
		dev, err := pb.client.Device(deviceNickname)
		if err != nil {
			return err
		}
		err = dev.PushNote(title, message)
		if err != nil {
			return err
		}
	}
	return nil
}
