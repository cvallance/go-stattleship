package stattleship

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StattleshipAPI struct {
	AccessToken string
}

func (api *StattleshipAPI) callEndPoint(sport string, league string, endpoint string, params map[string]interface{}, result interface{}) error {
	url := fmt.Sprintf("https://www.stattleship.com/%v/%v/%v", sport, league, endpoint)

	req, err := http.NewRequest("GET", url, nil)
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

	//  out, err := os.Create("result.json")
	//  defer out.Close()
	//  io.Copy(out, resp.Body)

	return json.NewDecoder(resp.Body).Decode(result)
}

func (api *StattleshipAPI) Games(sport string, league string, params map[string]interface{}) (*GamesResult, error) {
	var result GamesResult
	err := api.callEndPoint(sport, league, "games", params, &result)
	return &result, err
}
