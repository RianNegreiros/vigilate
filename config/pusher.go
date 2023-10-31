package config

import (
	"os"

	"github.com/pusher/pusher-http-go"
)

func NewPusherClient() *pusher.Client {
	if os.Getenv("APP_ENV") == "docker" {
		return &pusher.Client{
			AppID:  "1",
			Key:    "pusherKey",
			Secret: "pusherSecret",
			Host:   "pusher:4001",
			Secure: false,
		}
	}

	return &pusher.Client{
		AppID:   os.Getenv("PUSHER_APP_ID"),
		Key:     os.Getenv("PUSHER_APP_KEY"),
		Secret:  os.Getenv("PUSHER_APP_SECRET"),
		Cluster: os.Getenv("PUSHER_APP_CLUSTER"),
		Secure:  true,
	}
}
