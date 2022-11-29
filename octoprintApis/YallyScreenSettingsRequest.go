package octoprintApis

import (
	"encoding/json"
	"fmt"

	"github.com/Pyxis66/YallyScreen/octoprintApis/dataModels"
)


// TODO: YallyScreenSettingsRequest seems like it's practically the same as PluginManagerInfoRequest
// Need to clean up and consolidate, or add comments as to why the two different classes.

type YallyScreenSettingsRequest struct {
	Command string `json:"command"`
}

func (this *YallyScreenSettingsRequest) Do(client *Client, uiState string) (*dataModels.YallyScreenSettingsResponse, error) {
	target := fmt.Sprintf("%s?command=get_settings", PluginPyxis66YallyScreenApiUri)
	bytes, err := client.doJsonRequest("GET", target, nil, ConnectionErrors, false)
	if err != nil {
		return nil, err
	}

	response := &dataModels.YallyScreenSettingsResponse{}
	if err := json.Unmarshal(bytes, response); err != nil {
		return nil, err
	}

	return response, err
}
