package fyne

import (
    "fmt"

    fyne2 "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"

    guisettings "github.com/andrejkoleshko/VSRPP-LAB/lab11/internal/domain/gui_settings"
)

type window struct {
    w  fyne2.Window
    tw guisettings.TextWidget
}

func NewW(win fyne2.Window) *window {
    return &window{w: win}
}

func (w *window) Resize(ws guisettings.WindowSize) error {
    if ws.IsFull() {
        w.w.SetFullScreen(true)
    } else {
        w.w.Resize(fyne2.NewSize(float32(ws.Width()), float32(ws.Height())))
    }
    return nil
}

func (w *window) UpdateTemperature(t float32) error {
    w.tw.SetText(fmt.Sprintf("Температура: %.2f°C", t))
    return nil
}

func (w *window) SetTemperatureWidget(tw guisettings.TextWidget) error {
    w.tw = tw
    label := tw.Render().(*widget.Label)
    center := container.NewCenter(label)
    w.w.SetContent(center)
    return nil
}

func (w *window) Render() error {
    w.w.Show()
    return nil
}