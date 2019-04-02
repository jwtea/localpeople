package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const refreshInterval = 100 * time.Millisecond

type App struct {
	s         *Scan
	addrTable *tview.Table
	app       *tview.Application
}

func (a *App) update() {
	for {
		time.Sleep(refreshInterval)
		a.app.QueueUpdateDraw(func() {
			a.UpdateTable()
		})
	}
}

// UpdateTable contents
func (a *App) UpdateTable() {
	for i, addr := range a.s.Addresses {

		addr.Active = time.Now().Add(-1000 * time.Millisecond).Before(addr.Time)

		activeColor := tcell.ColorGreen

		if addr.Active == false {
			activeColor = tcell.ColorRed
		}

		textColor := tcell.ColorWhite
		a.addrTable.SetCell(i+1, 0,
			tview.NewTableCell(addr.IP).
				SetTextColor(textColor).
				SetAlign(tview.AlignCenter))
		a.addrTable.SetCell(i+1, 1,
			tview.NewTableCell(addr.MAC).
				SetTextColor(textColor).
				SetAlign(tview.AlignCenter))
		a.addrTable.SetCell(i+1, 2,
			tview.NewTableCell(fmt.Sprintf(addr.Time.Format(time.Stamp))).
				SetTextColor(textColor).
				SetAlign(tview.AlignCenter))
		a.addrTable.SetCell(i+1, 3,
			tview.NewTableCell(strconv.FormatBool(addr.Active)).
				SetTextColor(textColor).
				SetBackgroundColor(activeColor).
				SetAlign(tview.AlignCenter))
	}
}
func main() {

	scan := NewScan()

	app := &App{
		s:         scan,
		app:       tview.NewApplication(),
		addrTable: tview.NewTable().SetBorders(true),
	}

	app.StyleTable()

	go app.update()
	go scan.run()

	if err := app.app.SetRoot(app.addrTable, true).Run(); err != nil {
		panic(err)
	}

}

func (a *App) StyleTable() {
	a.addrTable.SetCell(0, 0, tview.NewTableCell("IP").
		SetTextColor(tcell.ColorYellow).
		SetAlign(tview.AlignCenter))
	a.addrTable.SetCell(0, 1, tview.NewTableCell("MAC").
		SetTextColor(tcell.ColorYellow).
		SetAlign(tview.AlignCenter))
	a.addrTable.SetCell(0, 2, tview.NewTableCell("TIME").
		SetTextColor(tcell.ColorYellow).
		SetAlign(tview.AlignCenter))
	a.addrTable.SetCell(0, 3, tview.NewTableCell("ACTIVE").
		SetTextColor(tcell.ColorYellow).
		SetAlign(tview.AlignCenter))
}
