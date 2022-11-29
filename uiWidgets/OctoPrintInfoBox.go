package uiWidgets

import (
	"fmt"

	"github.com/Pyxis66/YallyScreen/logger"
	"github.com/Pyxis66/YallyScreen/octoprintApis"
	// "github.com/Pyxis66/YallyScreen/octoprintApis/dataModels"
	"github.com/Pyxis66/YallyScreen/utils"
)


type OctoPrintInfoBox struct {
	*SystemInfoBox
}

func CreateOctoPrintInfoBox(
	client				*octoprintApis.Client,
	logoWidth			int,
) *OctoPrintInfoBox {
	logoHeight := int(float64(logoWidth) * 1.25)
	logoImage := utils.MustImageFromFileWithSize("logos/logo-octoprint.png", logoWidth, logoHeight)

	server := ""
	apiVersion := ""
	versionResponse, err := (&octoprintApis.VersionRequest{}).Do(client)
	if err != nil {
		logger.LogError("OctoPrintInfoBox.CreateOctoPrintInfoBox()", "VersionRequest.Do()", err)
	} else if versionResponse == nil {
		server = "Unknown?"
		apiVersion = "Unknown?"
	} else {
		server = versionResponse.Server
		apiVersion = versionResponse.API
	}

	base := CreateSystemInfoBox(
		client,
		logoImage,
		"OctoPrint",
		server,
		fmt.Sprintf("(API   %s)", apiVersion),   // Use 3 spaces... 1 space doesn't have enough kerning.
	)

	instance := &OctoPrintInfoBox {
		SystemInfoBox:			base,
	}

	return instance
}
