package transmission

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeCommonControls(nil)

	widget.SetKeyboardChar("j", widget.Prev, "Select previous item")
	widget.SetKeyboardChar("k", widget.Next, "Select next item")
	widget.SetKeyboardChar("u", widget.Unselect, "Clear selection")

	widget.SetKeyboardKey(tcell.KeyCtrlD, widget.deleteSelectedTorrent, "Delete selected torrent")
	widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select next item")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.pauseUnpauseTorrent, "Pause/unpause torrent")
	widget.SetKeyboardKey(tcell.KeyEsc, widget.Unselect, "Clear selection")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select previous item")
}
