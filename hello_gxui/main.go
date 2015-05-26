package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)

	window := theme.CreateWindow(320, 240, "Hello")
	window.OnClose(driver.Terminate)

	label := theme.CreateLabel()
	label.SetText("Hello, world!")
	window.AddChild(label)
}

func main() {
	gl.StartDriver(appMain)
}
