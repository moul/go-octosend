package octosend

import (
	"encoding/json"
	"fmt"
)

func (a *OctosendAPI) GetSpoolerMessage(spoolerToken string) (*SpoolerMessage, error) {
	resp, err := a.GetRequest(fmt.Sprintf("3.0/spooler/%s/message", spoolerToken))
	if err != nil {
		return nil, err
	}

	var result SpoolerMessage
	err = json.Unmarshal(resp, &result)
	return &result, err
}

func (a *OctosendAPI) GetDomain(domain string) (*Domain, error) {
	resp, err := a.GetRequest(fmt.Sprintf("3.0/domain/%s", domain))
	if err != nil {
		return nil, err
	}

	var result Domain
	err = json.Unmarshal(resp, &result)
	return &result, err
}
