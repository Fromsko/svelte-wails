package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "myproject",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Frameless:        true, // 无边框模式(True)
		DisableResize:    true,
		// AlwaysOnTop:      true,
		Windows: &windows.Options{
			WebviewIsTransparent:              true, // 网页透明
			WindowIsTranslucent:               true, // 窗口完全透明
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: true, // 禁用窗口装饰
			WebviewGpuIsDisabled:              true, // 禁用GPU
			WebviewUserDataPath:               "",
			Theme:                             windows.SystemDefault,
		},
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
