package octoprintApis

import (
	"bytes"
	"encoding/json"
	// "fmt"

	"github.com/Pyxis66/YallyScreen/logger"
)


type RunZOffsetCalibrationRequest struct {
	Command string `json:"command"`
}

func (this *RunZOffsetCalibrationRequest) Do(client *Client) error {
	this.Command = "run_zoffset_calibration"

	buffer := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buffer).Encode(this); err != nil {
		logger.LogError("RunZOffsetCalibrationRequest.Do()", "json.NewEncoder(params).Encode(this)", err)
		return err
	}

	_, err := client.doJsonRequest("POST", PluginPyxis66YallyScreenApiUri, buffer, ConnectionErrors, true)
	return err
}
