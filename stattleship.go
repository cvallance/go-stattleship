package stattleship

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var defaultparams map[string]string

func init() {
	defaultparams = map[string]string{
		"page":      "1",
		"page_size": "40",
	}
}

func combineParamsWithDefaults(params url.Values) url.Values {
	if params == nil {
		params = url.Values{}
	}

	for key, value := range defaultparams {
		if setvalue := params.Get(key); setvalue == "" {
			params.Set(key, value)
		}
	}

	return params
}

type StattleshipAPI struct {
	AccessToken string
}

func (api *StattleshipAPI) callEndPoint(sport string, league string, endpoint string, params url.Values, result interface{}) error {
	rawurl := fmt.Sprintf("https://www.stattleship.com/%v/%v/%v", sport, league, endpoint)
	baseurl, err := url.Parse(rawurl)
	if err != nil {
		return err
	}

	params = combineParamsWithDefaults(params)
	baseurl.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", baseurl.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Token token=%v", api.AccessToken))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/vnd.stattleship.com; version=1")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(result)
}

func (api *StattleshipAPI) Games(sport string, league string, params url.Values) (*GamesResult, error) {
	var result GamesResult
	err := api.callEndPoint(sport, league, "games", params, &result)
	return &result, err
}
