package network

import (
	"net/url"

	adhnetwork "github.com/adh-partnership/api/pkg/network"
)

func init() {
	adhnetwork.UserAgent = "ADH/bot"
}

func Call(method, req_url string, contenttype string, formdata map[string]string, headers map[string]string) (int, []byte, error) {
	u, err := url.Parse(req_url)
	if err != nil {
		return 0, nil, err
	}

	data := url.Values{}

	for k, v := range formdata {
		data.Set(k, v)
	}

	return adhnetwork.HandleWithHeaders(
		method,
		u.String(),
		contenttype,
		data.Encode(),
		headers,
	)
}
