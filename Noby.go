package nobyapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://app.cotogoto.ai/webapi/noby.json"

// Call calls noby API.
func Call(req *NobyRequest) (*NobyResponse, error) {
	request, err := http.NewRequest("GET", URL, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	param := request.URL.Query()

	param.Add("appkey", req.Appkey)
	param.Add("text", req.Text)
	param.Add("persona", string(req.Persona))
	param.Add("lat", fmt.Sprint(req.Lat))
	param.Add("lng", fmt.Sprint(req.Lng))
	param.Add("mail", req.Mail)
	param.Add("pass", req.Pass)
	param.Add("ending", req.Ending)

	if req.Study {
		param.Add("study", "1")
	} else {
		param.Add("study", "0")
	}

	request.URL.RawQuery = param.Encode()

	client := http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var ret NobyResponse

	json.Unmarshal(body, &ret)

	return &ret, nil
}
