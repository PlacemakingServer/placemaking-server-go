package templates

import "embed"

//go:embed assets/templates/*.html
var FS embed.FS
