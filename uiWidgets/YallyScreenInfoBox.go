package uiWidgets

import (
	// "fmt"
	"strings"

	"github.com/Pyxis66/YallyScreen/octoprintApis"
	// "github.com/Pyxis66/YallyScreen/octoprintApis/dataModels"
	"github.com/Pyxis66/YallyScreen/utils"
)

type YallyScreenInfoBox struct {
	*SystemInfoBox
}

func CreateYallyScreenInfoBox(
	client				*octoprintApis.Client,
	yallyScreenVersion	string,
) *YallyScreenInfoBox {
	logoImage := utils.MustImageFromFile("logos/yallyscreen-isometric-90%.png")

	str2 := ""
	str3 := ""
	stringArray := strings.Split(yallyScreenVersion, " ")
	if len(stringArray) == 2 {
		str2 = stringArray[0]
		str3 = stringArray[1]
	} else {
		str2 = yallyScreenVersion
		str3 = ""
	}

	base := CreateSystemInfoBox(
		client,
		logoImage,
		"YallyScreen",
		str2,
		str3,
	)

	instance := &YallyScreenInfoBox {
		SystemInfoBox:			base,
	}

	return instance
}
