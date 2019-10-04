package components

import (
	"github.com/pieterclaerhout/go-html"
)

func render(blocks html.Block, minified bool) string {

	f := html.RenderString
	if minified {
		f = html.RenderMinifiedString
	}
	output, _ := f(blocks)
	return output

}
