package ui

import (
	// "github.com/Pyxis66/YallyScreen/interfaces"
	// "github.com/Pyxis66/YallyScreen/octoprintApis"
	"github.com/Pyxis66/YallyScreen/octoprintApis/dataModels"
	// "github.com/Pyxis66/YallyScreen/uiWidgets"
)


type customItemsPanel struct {
	CommonPanel
	items			[]dataModels.MenuItem
}

func CustomItemsPanel(
	ui				*UI,
	items			[]dataModels.MenuItem,
) *customItemsPanel {
	instance := &customItemsPanel {
		CommonPanel: NewCommonPanel("CustomItemsPanel", ui),
		items:       items,
	}
	instance.initialize()

	return instance
}

func (this *customItemsPanel) initialize() {
	defer this.Initialize()
	this.arrangeMenuItems(this.grid, this.items, 4)
}
