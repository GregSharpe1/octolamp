package wled

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// struct DeviceConfig {

type DeviceConfig struct {
	Schema    string
	IPAddress string
	State     State
}

type State struct {
	State      bool `json:"on"`
	Brightness int  `json:"bri"`
	Transition int  `json:"transition"`
}

func apiEndpoint(config DeviceConfig, state string) string {
	baseEndpoint := (config.Schema + "://" + config.IPAddress + "/json")

	if state == "state" {
		return baseEndpoint + "/state"
	}
	return baseEndpoint
}

// func GetState(config DeviceConfig) (map[string]string, error) {
func GetState(config DeviceConfig) {
	resp, err := http.Get(apiEndpoint(config, "state"))
	if err != nil {
		logrus.Fatalf("error when getting state: %v", err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatalf("error reading json response: %v", readErr)
	}

	deviceState := State{}
	jsonErr := json.Unmarshal(body, &deviceState)
	if jsonErr != nil {
		logrus.Fatalf("error read json: %v", jsonErr)
	}

	fmt.Printf("getState value: %t", deviceState.State)
}

func (d *DeviceConfig) ChangeColour(config DeviceConfig, colour string) {
	GetState(config)
	fmt.Printf("Device State: %t", d.State.State)
}
