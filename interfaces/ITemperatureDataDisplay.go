package interfaces

import (
	// "github.com/Pyxis66/YallyScreen/octoprintApis"
	"github.com/Pyxis66/YallyScreen/octoprintApis/dataModels"
)

type ITemperatureDataDisplay interface {
	UpdateTemperatureData(temperatureData map[string]dataModels.TemperatureData)
}
