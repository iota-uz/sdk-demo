package assets

import (
	"embed"
	"github.com/benbjohnson/hashfs"
)

//go:embed *
var FS embed.FS

var HashFS = hashfs.NewFS(&FS)
