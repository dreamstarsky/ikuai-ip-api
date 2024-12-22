package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Resp struct {
	Result int         `json:"Result"`
	ErrMsg string      `json:"ErrMsg"`
	Data   interface{} `json:"Data"`
}

func (i *Ikuai) Call(url, action, func_name string, param any) (Resp, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return Resp{}, err
	}

	req, err := http.NewRequest("POST", i.url+"/Action/call", bytes.NewBuffer(data))
	if err != nil {
		return Resp{}, err
	}

	if i.client == nil {
		return Resp{}, errors.New("not login yet")
	}

	resp, err := i.client.Do(req)
	if err != nil {
		return Resp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Resp{}, errors.New("call failed with status " + resp.Status)
	}

	var res Resp
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return Resp{}, err
	}

	return res, nil
}
