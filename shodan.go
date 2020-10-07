package shodan

import (
	"vighnesha.in/shodan/config"
	"vighnesha.in/shodan/exploitsapi"
	"vighnesha.in/shodan/restapi"
	"vighnesha.in/shodan/streamapi"
)

const streamingAPIBaseURL = "https://stream.shodan.io"
const exploitAPIBaseURL = "https://exploits.shodan.io/api"
const restAPIURL = "https://api.shodan.io"

// mUZPi35ubzYPmbGkvkQkRDlDVnEIwS0c

type Shodan interface {
	RestAPI() restapi.Rest
	StreamAPI() streamapi.Stream
	ExploitsAPI() exploitsapi.Exploits
}

type shodan struct {
	apiKey string
	config config.Configuration
}

func (api *shodan) StreamAPI() streamapi.Stream {
	return streamapi.Configure(api.apiKey, api.config.Streamapi)
}

func (api *shodan) ExploitsAPI() exploitsapi.Exploits {
	return exploitsapi.Configure(api.apiKey, api.config.Exploitsapi)
}

func (api *shodan) RestAPI() restapi.Rest {
	return restapi.Configure(api.apiKey, api.config.Restapi)
}

func Configure(apiKey string) Shodan {
	return &shodan{apiKey: apiKey, config: config.LoadConfig()}
}
