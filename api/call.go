package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Resp struct {
	Result int             `json:"Result"`
	ErrMsg string          `json:"ErrMsg"`
	Data   json.RawMessage `json:"Data"`
}

type Param struct {
	Type string `json:"TYPE"`
}

type Request struct {
	FuncName string `json:"func_name"`
	Action   string `json:"action"`
	Param    Param  `json:"param"`
}

func (i *Ikuai) Call(action, func_name, param string) (Resp, error) {
	data, err := json.Marshal(Request{func_name, action, Param{param}})
	if err != nil {
		return Resp{}, err
	}

	req, err := http.NewRequest("POST", i.url+"/Action/call", bytes.NewBuffer(data))
	if err != nil {
		return Resp{}, err
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

	if res.Result == 10014 {
		i.Login()
		return res, nil
	}

	return res, nil
}
