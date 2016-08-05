package stattleship

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"sync"
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

func (api *StattleshipAPI) Get(sport string, league string, endpoint string, params url.Values) (*interface{}, *HeaderDetails, error) {
	rawurl := fmt.Sprintf("https://www.stattleship.com/%v/%v/%v", sport, league, endpoint)
	baseurl, err := url.Parse(rawurl)
	if err != nil {
		return nil, nil, err
	}

	params = combineParamsWithDefaults(params)
	baseurl.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", baseurl.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Token token=%v", api.AccessToken))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/vnd.stattleship.com; version=1")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	headerdetails := createHeaderDetails(&resp.Header)

	var result interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)

	return &result, headerdetails, err
}

func createHeaderDetails(header *http.Header) *HeaderDetails {
	headerdetails := HeaderDetails{}

	perPage := header.Get("Per-Page")
	i, err := strconv.Atoi(perPage)
	if err == nil {
		headerdetails.PerPage = i
	}

	total := header.Get("Total")
	i, err = strconv.Atoi(total)
	if err == nil {
		headerdetails.Total = i
	}

	return &headerdetails
}

func (api *StattleshipAPI) GetAll(sport string, league string, endpoint string, params url.Values) (*interface{}, error) {
	if params == nil {
		params = url.Values{}
	}

	params.Set("page", "1")

	results, headerdetails, err := api.Get(sport, league, endpoint, params)
	if err != nil {
		return nil, err
	}

	total := headerdetails.Total
	perpage := headerdetails.PerPage
	if total < perpage {
		return results, nil
	}

	//If there are more pages to get, lets do it concurrently
	resultschan := make(chan *interface{})

	numpages := int(math.Ceil(float64(total) / float64(perpage)))

	var waitgroup sync.WaitGroup
	waitgroup.Add(numpages - 1)

	//Note: starting at 2 because we've already got the first page
	for i := 2; i <= numpages; i++ {
		loopparams, _ := url.ParseQuery(params.Encode()) //copy the original params
		loopparams.Set("page", strconv.Itoa(i))
		go func(loopparams url.Values) {
			loopresult, _, err := api.Get(sport, league, endpoint, loopparams)
			if err == nil {
				resultschan <- loopresult
			}
			waitgroup.Done()
		}(loopparams)
	}

	go func() {
		waitgroup.Wait()
		close(resultschan)
	}()

	for loopresult := range resultschan {
		merged := mergeResults(results, loopresult)
		results = &merged
	}

	return results, nil
}

func mergeResults(results *interface{}, tomerge *interface{}) interface{} {
	resultsobj := (*results).(map[string]interface{})
	tomergeobj := (*tomerge).(map[string]interface{})

	for key, value := range tomergeobj {
		if _, ok := resultsobj[key]; !ok {
			resultsobj[key] = value
			continue
		}

		tomergecol := value.([]interface{})
		resultscol := resultsobj[key].([]interface{})
		for _, item := range tomergecol {
			itemobj := item.(map[string]interface{})
			if _, ok := itemobj["id"]; !ok {
				resultscol = append(resultscol, item)
				continue
			}

			itemid := itemobj["id"].(string)
			//does this item already exist in the results?
			alreadyexists := false
			for _, resitem := range resultscol {
				resitemobj := resitem.(map[string]interface{})
				if resitemobj["id"] == itemid {
					alreadyexists = true
					break
				}
			}

			if !alreadyexists {
				resultscol = append(resultscol, item)
				fmt.Println(len(resultscol))
			}
		}
		resultsobj[key] = resultscol
	}

	return resultsobj
}
