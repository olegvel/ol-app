package web

import (
	"embed"
)

//go:embed *
var web embed.FS

func GetWebFS() embed.FS {
	return web
}
