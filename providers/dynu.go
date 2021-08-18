package providers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/zx06/ddnss/common"
)

type DynuProvider struct {
	Domain string
	ApiKey string
	IpV4   string
}

type domainInfo struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	UnicodeName string `json:"unicodeName"`
	IpV4Address string `json:"ipv4Address"`
}

type domainInfoResp struct {
	Domains []domainInfo `json:"domains"`
}

type updateDnsRequest struct {
	Name              string `json:"name"`
	IpV4              bool   `json:"ipv4"`
	IpV4Address       string `json:"ipv4Address"`
	IpV4WildcardAlias bool   `json:"ipv4WildcardAlias"`
}

func NewDynuProvider(domain, apiKey, ipV4 string) (*DynuProvider, error) {
	if domain == "" || apiKey == "" {
		return nil, errors.New("DynuProvider: domain and apiKey must be set")
	}
	return &DynuProvider{
		Domain: domain,
		ApiKey: apiKey,
		IpV4:   ipV4,
	}, nil
}

func (p *DynuProvider) Update() bool {
	var err error
	if p.IpV4 == "" {
		p.IpV4, err = common.GetIPV4()
		if err != nil {
			return false
		}
	}
	domains, err := dynuGetDomains(context.Background(), p.ApiKey)
	if err != nil {
		log.Printf("dynuGetDomains error: %v\n", err)
		return false

	}
	for _, d := range domains {
		if d.UnicodeName == p.Domain {
			err = dynuUpdateDDNS(context.Background(), d.Id, p.ApiKey, p.Domain, p.IpV4)
			if err != nil {
				log.Printf("dynuUpdateDDNS error: %v\n", err)
				return false
			}
			return true
		}
	}
	return false
}

func dynuGetDomains(ctx context.Context, apiKey string) ([]domainInfo, error) {
	const apiURL = "https://api.dynu.com/v2/dns"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Key", apiKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var d domainInfoResp
		err = json.NewDecoder(resp.Body).Decode(&d)
		if err != nil {
			return nil, err
		}
		return d.Domains, nil
	}
	return nil, errors.New("dynuGetDomains: " + resp.Status)
}

func dynuUpdateDDNS(ctx context.Context, id uint64, apiKey, domain, ipV4 string) error {
	var apiURL = "https://api.dynu.com/v2/dns/" + strconv.FormatUint(id, 10)
	var data = updateDnsRequest{
		Name:              domain,
		IpV4:              true,
		IpV4Address:       ipV4,
		IpV4WildcardAlias: true,
	}
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL, strings.NewReader(string(body)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Key", apiKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return nil
	} else {
		return errors.New("dynuUpdateDDNS: " + resp.Status)
	}

}
