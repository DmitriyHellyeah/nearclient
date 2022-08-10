package nearclient

import "time"

type NearConfig struct {
	Host           string
	WaitingTimeout time.Duration
}
