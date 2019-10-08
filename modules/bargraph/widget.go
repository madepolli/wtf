package bargraph

/**************
This is a demo bargraph that just populates some random date/val data
*/

import (
	"github.com/rivo/tview"
	"math/rand"
	"time"

	"github.com/wtfutil/wtf/view"
)

var (
	ok      = true
	started = false
)

// Widget define wtf widget to register widget later
type Widget struct {
	view.BarGraph

	app *tview.Application
}

// NewWidget Make new instance of widget
func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		BarGraph: view.NewBarGraph(app, "Sample Bar Graph", settings.common),

		app: app,
	}

	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// MakeGraph - Load the dead drop stats
func MakeGraph(widget *Widget) {

	//this could come from config
	const lineCount = 8
	var stats [lineCount]view.Bar

	barTime := time.Now()
	for i := 0; i < lineCount; i++ {
		barTime = barTime.Add(time.Duration(time.Minute))

		bar := view.Bar{
			Label:      barTime.Format("15:04"),
			Percent:    rand.Intn(100-5) + 5,
			LabelColor: "red",
		}

		stats[i] = bar
	}

	widget.BarGraph.BuildBars(stats[:])

}

// Refresh & update after interval time
func (widget *Widget) Refresh() {

	if widget.Disabled() {
		return
	}

	widget.View.Clear()

	widget.app.QueueUpdateDraw(func() {
		display(widget)
	})

}

/* -------------------- Unexported Functions -------------------- */

func display(widget *Widget) {
	MakeGraph(widget)
}
