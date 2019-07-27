package streamapi

import (
	"shodan/config"
)

type Stream interface {
	Data() Data
	NetworkAlerts() NetworkAlert
}

type stream struct {
	key    string
	config config.Streamapi
}

func (stream *stream) Data() Data {
	return &data{stream.key, stream.config.StreamData}
}

func (stream *stream) NetworkAlerts() NetworkAlert {
	return &networkAlert{stream.key, stream.config.StreamNetworkAlerts}
}

func Configure(apiKey string, configuration config.Streamapi) Stream {
	return &stream{key: apiKey, config: configuration}
}





