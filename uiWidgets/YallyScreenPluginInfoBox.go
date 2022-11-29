package uiWidgets

import (
	// "fmt"

	"github.com/Pyxis66/YallyScreen/logger"
	"github.com/Pyxis66/YallyScreen/octoprintApis"
	// "github.com/Pyxis66/YallyScreen/octoprintApis/dataModels"
	"github.com/Pyxis66/YallyScreen/utils"
)

type YallyScreenPluginInfoBox struct {
	*SystemInfoBox
}

func CreateYallyScreenPluginInfoBox(
	client							*octoprintApis.Client,
	uiState							string,
	octoPrintPluginIsInstalled		bool,
) *YallyScreenPluginInfoBox {
	logoImage := utils.MustImageFromFile("logos/puzzle-piece.png")
	str1 := "YallyScreen plugin"

	str2 := ""
	if octoPrintPluginIsInstalled {
		pluginManagerInfoResponse, err := (&octoprintApis.PluginManagerInfoRequest{}).Do(client, uiState)
		if err != nil {
			logger.LogError("CreateYallyScreenPluginInfoBox()", "PluginManagerInfoRequest.Do()", err)
			str2 = "Error"
		} else {
			found := false
			for i := 0; i < len(pluginManagerInfoResponse.Plugins) && !found; i++ {
				plugin := pluginManagerInfoResponse.Plugins[i]
				if plugin.Key == "pyxis66_yallyscreen" {
					found = true
					str2 = plugin.Version
				}
			}

			if !found {
				// OK, the plugin is there, we just can't get the info from a GET request.
				// Default to displaying, "Present"
				str2 = "Present"
			}
		}
	} else {
		str2 = "Not installed"
	}

	base := CreateSystemInfoBox(
		client,
		logoImage,
		str1,
		str2,
		"",
	)

	instance := &YallyScreenPluginInfoBox {
		SystemInfoBox:			base,
	}

	return instance
}
