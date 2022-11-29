package uiWidgets

import (
	"github.com/Pyxis66/YallyScreen/logger"
	"github.com/Pyxis66/YallyScreen/octoprintApis"
	// "github.com/Pyxis66/YallyScreen/octoprintApis/dataModels"
	"github.com/Pyxis66/YallyScreen/utils"
)

type FlowRateStepButton struct {
	*StepButton
	client			*octoprintApis.Client
}

func CreateFlowRateStepButton(
	client			*octoprintApis.Client,
) *FlowRateStepButton {
	base, err := CreateStepButton(
		1,
		Step{"Normal (100%)", "speed-normal.svg", nil, 100},
		Step{"Fast (125%)",   "speed-fast.svg",   nil, 125},
		Step{"Slow (75%)",    "speed-slow.svg",   nil,  75},
	)
	if err != nil {
		logger.LogError("PANIC!!! - CreateFlowRateStepButton()", "CreateStepButton()", err)
		panic(err)
	}

	instance := &FlowRateStepButton{
		StepButton:		base,
		client:			client,
	}

	return instance
}

func (this *FlowRateStepButton) Value() int {
	return this.StepButton.Value().(int)
}

func (this *FlowRateStepButton) SendChangeFlowRate() error {
	return utils.SetFlowRate(this.client, this.Value())
}
