package cdsqueue

import (
	"github.com/gdamore/tcell"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeCommonControls(widget.Refresh)

	widget.SetKeyboardChar("j", widget.Next, "Select next workflow")
	widget.SetKeyboardChar("k", widget.Prev, "Select previous workflow")
	widget.SetKeyboardChar("l", widget.NextSource, "Select next filter")
	widget.SetKeyboardChar("h", widget.PrevSource, "Select previous filter")
	widget.SetKeyboardChar("o", widget.openWorkflow, "Open workflow in browser")

	widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select next workflow")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select previous workflow")
	widget.SetKeyboardKey(tcell.KeyRight, widget.NextSource, "Select next filter")
	widget.SetKeyboardKey(tcell.KeyLeft, widget.PrevSource, "Select previous filter")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openWorkflow, "Open workflow in browser")
	widget.SetKeyboardKey(tcell.KeyEsc, widget.Unselect, "Clear selection")
}
