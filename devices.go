package gowink

import (
	"io/ioutil"
)

func (w *Wink) GetDevices() ([]byte, error) {
	uri, err := w.GetUri("/users/me/wink_devices", nil)
	if err != nil {
		return nil, err
	}

	resp, err := w.DoRequest("GET", uri, nil)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
