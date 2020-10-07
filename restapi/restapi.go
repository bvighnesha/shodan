package restapi

import (
	"vighnesha.in/shodan/config"
)

type Rest interface {
	Search() Search
	Scan() Scan
	Network() Network
	Directory() Directory
	Data() Data
	Organization() Organization
	Account() Account
	DNS() DNS
	Utility() Utility
	APIStatus() APIStatus
	Experimental() Experimental
}

type rest struct {
	key    string
	config config.Restapi
}

func (rest *rest) Search() Search {
	return &search{rest.key, rest.config.Search}
}

func (rest *rest) Scan() Scan {
	return &scan{rest.key, rest.config.Scan}
}

func (rest *rest) Network() Network {
	return &network{rest.key, rest.config.NetworkAlerts}
}

func (rest *rest) Directory() Directory {
	return &directory{rest.key, rest.config.Directory}
}

func (rest *rest) Data() Data {
	return &data{rest.key, rest.config.Data}
}

func (rest *rest) Organization() Organization {
	return &organization{rest.key, rest.config.Organization}
}

func (rest *rest) Account() Account {
	return &account{rest.key, rest.config.Account}
}

func (rest *rest) DNS() DNS {
	return &dns{rest.key, rest.config.DNS}
}

func (rest *rest) Utility() Utility {
	return &utility{rest.key, rest.config.Utility}
}

func (rest *rest) APIStatus() APIStatus {
	return &apiStatus{rest.key, rest.config.Apistatus}
}

func (rest *rest) Experimental() Experimental {
	return &experimental{rest.key, rest.config.Experimental}
}

func Configure(apiKey string, configuration config.Restapi) Rest {
	return &rest{key: apiKey, config: configuration}
}
