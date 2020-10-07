package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration struct {
	Restapi     Restapi     `json:"restapi,omitempty"`
	Streamapi   Streamapi   `json:"streamapi,omitempty"`
	Exploitsapi Exploitsapi `json:"exploitsapi,omitempty"`
}

func LoadConfig() Configuration {
	jsonFile, err := os.Open("config/config.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	var config Configuration
	data, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("error:", err)
	}
	return config
}

type Exploitsapi struct {
	SearchURL string `json:"SearchURL,omitempty"`
	CountURL  string `json:"CountURL,omitempty"`
}

type Restapi struct {
	Search        Search        `json:"search,omitempty"`
	Scan          Scan          `json:"scan,omitempty"`
	NetworkAlerts NetworkAlerts `json:"networkAlerts,omitempty"`
	Directory     Directory     `json:"directory,omitempty"`
	Data          Data          `json:"data,omitempty"`
	Organization  Organization  `json:"organization,omitempty"`
	Account       Account       `json:"account,omitempty"`
	DNS           DNS           `json:"dns,omitempty"`
	Utility       Utility       `json:"utility,omitempty"`
	Apistatus     Apistatus     `json:"apistatus,omitempty"`
	Experimental  Experimental  `json:"experimental,omitempty"`
}

type Account struct {
	AccountProfileURL string `json:"AccountProfileURL,omitempty"`
}

type Apistatus struct {
	APIPlanInformationURL string `json:"APIPlanInformationURL,omitempty"`
}

type DNS struct {
	DomainInfoURL       string `json:"DomainInfoURL,omitempty"`
	DNSLookupURL        string `json:"DNSLookupURL,omitempty"`
	ReverseDNSLookupURL string `json:"ReverseDNSLookupURL,omitempty"`
}

type Data struct {
	DataURL    string `json:"DataURL,omitempty"`
	DatasetURL string `json:"DatasetURL,omitempty"`
}

type Directory struct {
	QueryURL       string `json:"QueryURL,omitempty"`
	QuerySearchURL string `json:"QuerySearchURL,omitempty"`
	QueryTagsURL   string `json:"QueryTagsURL,omitempty"`
}

type Experimental struct {
	HoneyscoreURL string `json:"HoneyscoreURL,omitempty"`
}

type NetworkAlerts struct {
	AlertURL                 string `json:"AlertURL,omitempty"`
	AlertIDInfoURL           string `json:"AlertIdInfoURL,omitempty"`
	DeleteAlertURL           string `json:"DeleteAlertURL,omitempty"`
	AlertInfoURL             string `json:"AlertInfoURL,omitempty"`
	AlertTriggersURL         string `json:"AlertTriggersURL,omitempty"`
	EnableTriggerURL         string `json:"EnableTriggerURL,omitempty"`
	DisableTriggerURL        string `json:"DisableTriggerURL,omitempty"`
	AddTriggerToWhitelistURL string `json:"AddTriggerToWhitelistURL,omitempty"`
	RemoveFromWhitelistURL   string `json:"RemoveFromWhitelistURL,omitempty"`
}

type Organization struct {
	InfoURL         string `json:"InfoURL,omitempty"`
	AddNewMemberURL string `json:"AddNewMemberURL,omitempty"`
	RemoveMemberURL string `json:"RemoveMemberURL,omitempty"`
}

type Scan struct {
	ProtocolsURL  string `json:"ProtocolsURL,omitempty"`
	ScanURL       string `json:"ScanURL,omitempty"`
	ScanInternet  string `json:"ScanInternet,omitempty"`
	ScanStatusURL string `json:"ScanStatusURL,omitempty"`
}

type Search struct {
	HostInformationURL  string `json:"HostInformationURL,omitempty"`
	HostCountURL        string `json:"HostCountURL,omitempty"`
	HostSearchURL       string `json:"HostSearchURL,omitempty"`
	HostSearchTokensURL string `json:"HostSearchTokensURL,omitempty"`
	PortsURL            string `json:"PortsURL,omitempty"`
}

type Utility struct {
	HTTPHeadersURL string `json:"HTTPHeadersURL,omitempty"`
	MyIPAddressURL string `json:"MyIPAddressURL,omitempty"`
}

type Streamapi struct {
	StreamData          StreamData          `json:"streamData,omitempty"`
	StreamNetworkAlerts StreamNetworkAlerts `json:"streamNetworkAlerts,omitempty"`
}

type StreamData struct {
	BannersURL          string `json:"BannersURL,omitempty"`
	FiltereByASNURL     string `json:"FiltereByASNURL,omitempty"`
	FiltereByCountryURL string `json:"FiltereByCountry,omitempty"`
	FiltereByPortsURL   string `json:"FiltereByPorts,omitempty"`
}

type StreamNetworkAlerts struct {
	AllNetworkAlertsURL string `json:"AllNetworkAlertsURL,omitempty"`
	FiltereByAlertIDURL string `json:"FiltereByAlertIDURL,omitempty"`
}
