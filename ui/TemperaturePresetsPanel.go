package ui

import (
	// "github.com/Pyxis66/YallyScreen/interfaces"
	"github.com/Pyxis66/YallyScreen/logger"
	"github.com/Pyxis66/YallyScreen/octoprintApis"
	// "github.com/Pyxis66/YallyScreen/octoprintApis/dataModels"
	"github.com/Pyxis66/YallyScreen/uiWidgets"
	// "github.com/Pyxis66/YallyScreen/utils"
)


var temperaturePresetsPanelInstance *temperaturePresetsPanel

type temperaturePresetsPanel struct {
	CommonPanel

	selectHotendStepButton	*uiWidgets.SelectToolStepButton

}

func TemperaturePresetsPanel(
	ui						*UI,
	selectHotendStepButton	*uiWidgets.SelectToolStepButton,
) *temperaturePresetsPanel {
	if temperaturePresetsPanelInstance == nil {
		instance := &temperaturePresetsPanel {
			CommonPanel:			NewCommonPanel("temperaturePresetsPanel", ui),
			selectHotendStepButton:	selectHotendStepButton,
		}
		instance.initialize()
		temperaturePresetsPanelInstance = instance
	}

	return temperaturePresetsPanelInstance
}

func (this *temperaturePresetsPanel) initialize() {
	defer this.Initialize()
	this.createAllOffButton()
	this.createTemperaturePresetButtons()
}

func (this *temperaturePresetsPanel) createAllOffButton() {
	allOffButton := uiWidgets.CreateCoolDownButton(this.UI.Client, this.UI.GoToPreviousPanel)
	this.AddButton(allOffButton)
}

func (this *temperaturePresetsPanel) createTemperaturePresetButtons() {
	settings, err := (&octoprintApis.SettingsRequest{}).Do(this.UI.Client)
	if err != nil {
		logger.LogError("TemperaturePresetsPanel.getTemperaturePresets()", "Do(SettingsRequest)", err)
		return
	}

	// 12 (max) - Back button - All Off button = 10 available slots to display.
	const maxSlots = 10

	count := 0
	for _, temperaturePreset := range settings.Temperature.TemperaturePresets {
		if count < maxSlots {
			temperaturePresetButton := uiWidgets.CreateTemperaturePresetButton(
				this.UI.Client,
				this.selectHotendStepButton,
				"heat-up.svg",
				temperaturePreset,
				this.UI.GoToPreviousPanel,
			)
			this.AddButton(temperaturePresetButton)
			count++
		}
	}
}
