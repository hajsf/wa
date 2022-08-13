package public

import (
	"embed"
)

// Views is our static web server layouts, views with dynamic content and partials content that is a static view.
//go:embed libs/scripts views layouts scripts styles
var Views embed.FS
