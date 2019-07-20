package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration struct {
	Restapi     Restapi `json:"restapi"`
	Streamapi   API     `json:"streamapi"`
	Exploitsapi API     `json:"exploitsapi"`
}

func LoadConfig() Configuration {
	jsonFile, err := os.Open("D:/code/go/shodan/config/config.json")

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

type API struct {
}

type Restapi struct {
	Search        Search        `json:"search"`
	Scan          Scan          `json:"scan"`
	NetworkAlerts NetworkAlerts `json:"networkAlerts"`
	Directory     Directory     `json:"directory"`
	Data          Data          `json:"data"`
	Organization  Organization  `json:"organization"`
	Account       Account       `json:"account"`
	DNS           DNS           `json:"dns"`
	Utility       Utility       `json:"utility"`
	Apistatus     Apistatus     `json:"apistatus"`
	Experimental  Experimental  `json:"experimental"`
}

type Account struct {
	AccountProfileURL string `json:"AccountProfileURL"`
}

type Apistatus struct {
	APIPlanInformationURL string `json:"APIPlanInformationURL"`
}

type DNS struct {
	DomainInfoURL string `json:"DomainInfoURL"`
	DNSLookupURL        string `json:"DNSLookupURL"`
	ReverseDNSLookupURL string `json:"ReverseDNSLookupURL"`
}

type Data struct {
	DataURL    string `json:"DataURL"`
	DatasetURL string `json:"DatasetURL"`
}

type Directory struct {
	QueryURL       string `json:"QueryURL"`
	QuerySearchURL string `json:"QuerySearchURL"`
	QueryTagsURL   string `json:"QueryTagsURL"`
}

type Experimental struct {
	HoneyscoreURL string `json:"HoneyscoreURL"`
}

type NetworkAlerts struct {
	AlertURL                 string `json:"AlertURL"`
	AlertIDInfoURL           string `json:"AlertIdInfoURL"`
	DeleteAlertURL           string `json:"DeleteAlertURL"`
	AlertInfoURL             string `json:"AlertInfoURL"`
	AlertTriggersURL         string `json:"AlertTriggersURL"`
	EnableTriggerURL         string `json:"EnableTriggerURL"`
	DisableTriggerURL        string `json:"DisableTriggerURL"`
	AddTriggerToWhitelistURL string `json:"AddTriggerToWhitelistURL"`
	RemoveFromWhitelistURL   string `json:"RemoveFromWhitelistURL"`
}

type Organization struct {
	InfoURL         string `json:"InfoURL"`
	AddNewMemberURL string `json:"AddNewMemberURL"`
	RemoveMemberURL string `json:"RemoveMemberURL"`
}

type Scan struct {
	ProtocolsURL  string `json:"ProtocolsURL"`
	ScanURL       string `json:"ScanURL"`
	ScanInternet  string `json:"ScanInternet"`
	ScanStatusURL string `json:"ScanStatusURL"`
}

type Search struct {
	HostInformationURL  string `json:"HostInformationURL"`
	HostCountURL        string `json:"HostCountURL"`
	HostSearchURL       string `json:"HostSearchURL"`
	HostSearchTokensURL string `json:"HostSearchTokensURL"`
	PortsURL            string `json:"PortsURL"`
}

type Utility struct {
	HTTPHeadersURL string `json:"HTTPHeadersURL"`
	MyIPAddressURL string `json:"MyIPAddressURL"`
}
