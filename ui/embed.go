package ui

import (
	"embed"
	"io/fs"
)

//go:embed all:dist
var DistDir embed.FS

// DistDirFS contains the embedded dist directory files (without the "dist" prefix)
var DistDirFS, _ = fs.Sub(DistDir, "dist")
