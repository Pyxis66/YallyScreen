package utils

import (
	// "fmt"
	// "sort"

	// "github.com/Pyxis66/YallyScreen/octoprintApis"
	"github.com/Pyxis66/YallyScreen/octoprintApis/dataModels"
	// "github.com/Pyxis66/YallyScreen/uiWidgets"
)


type FileResponsesSortedByDate []*dataModels.FileResponse

func (this FileResponsesSortedByDate) Len() int {
	 return len(this)
}

func (this FileResponsesSortedByDate) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this FileResponsesSortedByDate) Less(i, j int) bool {
	return this[j].Date.Time.Before(this[i].Date.Time)
}
