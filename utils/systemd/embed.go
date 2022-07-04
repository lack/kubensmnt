package systemd

import "embed"

//go:embed *.service *.conf
var _ embed.FS
