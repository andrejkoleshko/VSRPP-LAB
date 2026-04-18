package fyne

import (
    fyne2 "fyne.io/fyne/v2"
    fyneApp "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"

    guisettings "github.com/andrejkoleshko/VSRPP-LAB/lab10/internal/domain/gui_settings"
)

type provider struct {
    a fyne2.App
    w fyne2.Window
}

func NewP() *provider {
    return &provider{
        a: fyneApp.New(),
    }
}

func (p *provider) CreateWindow(name string, size guisettings.WindowSize) (guisettings.Window, error) {
    w := p.a.NewWindow(name)
    p.w = w

    win := NewW(w)
    win.Resize(size)

    return win, nil
}

func (p *provider) GetAppRunner() guisettings.AppRunner {
    return NewAR(p.w)
}

func (p *provider) GetTextWidget(text string) guisettings.TextWidget {
    label := widget.NewLabel(text)
    return NewTW(label)
}
