package fyne

import fyne2 "fyne.io/fyne/v2"

type appRunner struct {
    w fyne2.Window
}

func NewAR(w fyne2.Window) *appRunner {
    return &appRunner{w: w}
}

func (ar *appRunner) Run() {
    ar.w.ShowAndRun()
}
