package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func (i *Ikuai) Login() error {
	data, err := json.Marshal(struct {
		Username string `json:"username"`
		Passwd   string `json:"passwd"`
		Pass     string `json:"pass"`
	}{
		Username: i.username,
		Passwd:   i.passwd,
		Pass:     i.pass,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", i.url+"/Action/login", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	if i.client == nil {
		i.client = &http.Client{}
	}

	resp, err := i.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("login failed with status " + resp.Status)
	}

	var res Resp
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}

	if res.ErrMsg != "Success" {
		return errors.New("login failed with ErrMsg " + res.ErrMsg)
	}

	return nil
}
