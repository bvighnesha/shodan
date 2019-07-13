package shodan

import (
	"shodan/config"
	"shodan/restapi"
)

const streamingAPIBaseURL = "https://stream.shodan.io"
const exploitAPIBaseURL = "https://exploits.shodan.io/api"
const restAPIURL = "https://api.shodan.io"

// mUZPi35ubzYPmbGkvkQkRDlDVnEIwS0c

type Shodan interface {
	RestAPI() restapi.Rest
	StreamAPI() *Stream
	ExploitsAPI() *Exploits
}

type shodan struct {
	apiKey string
	config config.Configuration
}

func (api *shodan) StreamAPI() *Stream {
	panic("implement me")
}

func (api *shodan) ExploitsAPI() *Exploits {
	panic("implement me")
}

func (api *shodan) RestAPI() restapi.Rest {
	return restapi.Configure(api.apiKey, api.config.Restapi)
}

/*func (api *shodan) StreamAPI() *Stream {
	return &Stream{configuration: Configuration{key: api.apiKey, url: streamingAPIBaseURL}, context:"/shodan"}
}

func (api *shodan) ExploitsAPI() *Exploits {
	return &Exploits{configuration: Configuration{key: api.apiKey, url: exploitAPIBaseURL}}
}*/

func Configure(apiKey string) Shodan {
	return &shodan{apiKey: apiKey, config: config.LoadConfig()}
}
