package streamapi

import "shodan/config"

type Stream struct {
	configuration config.Configuration
	context       string
}
