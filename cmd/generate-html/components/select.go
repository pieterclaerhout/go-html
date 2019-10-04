package components

import (
	"github.com/pieterclaerhout/go-html"
)

type Option struct {
	Value string
	Text  string
}

type MySelectBuilder struct {
	options  []Option
	selected string
}

func MySelect() *MySelectBuilder {
	return &MySelectBuilder{}
}

func (b *MySelectBuilder) Options(opts []Option) (r *MySelectBuilder) {
	b.options = opts
	return b
}

func (b *MySelectBuilder) Selected(selected string) (r *MySelectBuilder) {
	b.selected = selected
	return b
}

// <select>
//   <option value="1">
//     label 1
//   </option>
//   <option value="2" selected="true">
//     label 2
//   </option>
//   <option value="3">
//     label 3
//   </option>
// </select>

// HTML returns the HTML version of the component
func (b *MySelectBuilder) HTML(minified bool) string {

	options := []html.Block{}
	for _, op := range b.options {

		attrs := html.Value(op.Value)
		if op.Value == b.selected {
			attrs = attrs.Attr("selected", "true")
		}

		options = append(options, html.Option(attrs, html.Text(op.Text)))

	}

	root := html.Blocks{
		html.Select(nil, options...),
	}

	return render(root, minified)

}
