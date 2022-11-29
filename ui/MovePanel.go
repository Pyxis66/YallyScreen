package ui

import (
	// "github.com/gotk3/gotk3/gtk"

	// "github.com/Pyxis66/YallyScreen/interfaces"
	// "github.com/Pyxis66/YallyScreen/octoprintApis"
	"github.com/Pyxis66/YallyScreen/octoprintApis/dataModels"
	"github.com/Pyxis66/YallyScreen/uiWidgets"
	// "github.com/Pyxis66/YallyScreen/utils"
)

var movePanelInstance *movePanel

type movePanel struct {
	CommonPanel
	amountToMoveStepButton    *uiWidgets.AmountToMoveStepButton
}

func MovePanel(
	ui				*UI,
) *movePanel {
	if movePanelInstance == nil {
		instance := &movePanel {
			CommonPanel: NewCommonPanel("MovePanel", ui),
		}
		instance.initialize()
		movePanelInstance = instance
	}

	return movePanelInstance
}

func (this *movePanel) initialize() {
	defer this.Initialize()

	// Create the step button first, since it is needed by some of the other controls.
	this.amountToMoveStepButton = uiWidgets.CreateAmountToMoveStepButton()

	xAxisInverted, yAxisInverted, zAxisInverted := false, false, false
	if this.UI.Settings != nil {
		xAxisInverted = this.UI.Settings.XAxisInverted
		yAxisInverted = this.UI.Settings.YAxisInverted
		zAxisInverted = this.UI.Settings.ZAxisInverted
	}

	if xAxisInverted {
		this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "X-", "move-x-.svg", dataModels.XAxis,  1), 0, 1, 1, 1)
		this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "X+", "move-x+.svg", dataModels.XAxis, -1), 2, 1, 1, 1)
	} else {
		this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "X-", "move-x-.svg", dataModels.XAxis, -1), 0, 1, 1, 1)
		this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "X+", "move-x+.svg", dataModels.XAxis,  1), 2, 1, 1, 1)
	}

	if yAxisInverted {
		this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "Y+", "move-y+.svg", dataModels.YAxis, -1), 1, 0, 1, 1)
		this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "Y-", "move-y-.svg", dataModels.YAxis,  1), 1, 2, 1, 1)
	} else {
		this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "Y+", "move-y+.svg", dataModels.YAxis,  1), 1, 0, 1, 1)
		this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "Y-", "move-y-.svg", dataModels.YAxis, -1), 1, 2, 1, 1)
	}

	if zAxisInverted {
		this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "Z+", "move-z+.svg", dataModels.ZAxis, -1), 3, 0, 1, 1)
		this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "Z-", "move-z-.svg", dataModels.ZAxis,  1), 3, 1, 1, 1)
	} else {
		this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "Z+", "move-z+.svg", dataModels.ZAxis,  1), 3, 0, 1, 1)
		this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "Z-", "move-z-.svg", dataModels.ZAxis, -1), 3, 1, 1, 1)
	}

	this.Grid().Attach(this.amountToMoveStepButton, 1, 1, 1, 1)
}
