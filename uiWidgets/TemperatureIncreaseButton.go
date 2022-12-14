package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/Pyxis66/YallyScreen/logger"
	"github.com/Pyxis66/YallyScreen/octoprintApis"
	// "github.com/Pyxis66/YallyScreen/octoprintApis/dataModels"
	"github.com/Pyxis66/YallyScreen/utils"
)

type TemperatureIncreaseButton struct {
	*gtk.Button

	client							*octoprintApis.Client
	temperatureAmountStepButton		*TemperatureAmountStepButton
	selectHotendStepButton			*SelectToolStepButton
	isIncrease						bool
}

func CreateTemperatureIncreaseButton(
	client							*octoprintApis.Client,
	temperatureAmountStepButton		*TemperatureAmountStepButton,
	selectHotendStepButton			*SelectToolStepButton,
	isIncrease						bool,
) *TemperatureIncreaseButton {
	var base *gtk.Button
	if isIncrease {
		base = utils.MustButtonImageStyle("Increase", "increase.svg", "", nil)
	} else {
		base = utils.MustButtonImageStyle("Decrease", "decrease.svg", "", nil)
	}

	instance := &TemperatureIncreaseButton{
		Button:							base,
		client:							client,
		temperatureAmountStepButton:	temperatureAmountStepButton,
		selectHotendStepButton:			selectHotendStepButton,
		isIncrease:						isIncrease,
	}
	_, err := instance.Button.Connect("clicked", instance.handleClicked)
	if err != nil {
		logger.LogError("PANIC!!! - CreateTemperatureIncreaseButton()", "instance.Button.Connect()", err)
		panic(err)
	}

	return instance
}

func (this *TemperatureIncreaseButton) handleClicked() {
	value := this.temperatureAmountStepButton.Value()
	tool := this.selectHotendStepButton.Value()
	target, err := utils.GetToolTarget(this.client, tool)
	if err != nil {
		logger.LogError("TemperatureIncreaseButton.handleClicked()", "GetToolTarget()", err)
		return
	}

	if this.isIncrease {
		target += value
	} else {
		target -= value
	}

	if target < 0 {
		target = 0
	}

	// TODO: should the target be checked for a max temp?
	// If so, how to calculate what the max should be?

	logger.Infof("TemperatureIncreaseButton.handleClicked() - setting target temperature for %s to %1.f°C.", tool, target)

	err = utils.SetToolTarget(this.client, tool, target)
	if err != nil {
		logger.LogError("TemperatureIncreaseButton.handleClicked()", "GetToolTarget()", err)
	}
}
