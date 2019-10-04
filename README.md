## go-html

[![Go Report Card](https://goreportcard.com/badge/github.com/pieterclaerhout/go-html)](https://goreportcard.com/report/github.com/pieterclaerhout/go-html)
[![Documentation](https://godoc.org/github.com/pieterclaerhout/go-html?status.svg)](http://godoc.org/github.com/pieterclaerhout/go-html)
[![license](https://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://github.com/pieterclaerhout/go-html/raw/master/LICENSE)
[![GitHub version](https://badge.fury.io/gh/pieterclaerhout%2Fgo-html.svg)](https://badge.fury.io/gh/pieterclaerhout%2Fgo-html)
[![GitHub issues](https://img.shields.io/github/issues/pieterclaerhout/go-html.svg)](https://github.com/pieterclaerhout/go-html/issues)

This is a [Golang](https://golang.org) library which allows you to build HTML using Go code.

It's mainly designed as a way to create re-usable HTML components, similar to how [React](https://reactjs.org) components work.

# Usage

Let's start with an easy example. This is the generated HTML:

```html
<div class="userProfile">
  <h1 class="profileName">
    pclaerhout
  </h1>
  <img src="https://en.gravatar.com/userimage/389631/32413201cb4a680e82ee7c4bdd30b795.jpeg" class="profileImage" alt="pclaerhout">
  </img>
</div>
```

This is the Go code which generates this HTML:

```go
package components

import (
	"github.com/pieterclaerhout/go-html"
)

type UserProfile struct {
	Username  string
	AvatarURL string
}

// HTML returns the HTML version of the component
func (userProfile *UserProfile) HTML(minified bool) string {

	root := html.Blocks{
		html.Div(
			html.Class("userProfile"),
			html.H1(
				html.Class("profileName"),
				html.Text(userProfile.Username),
			),
			html.Img(
				html.Class("profileImage").Src(userProfile.AvatarURL).Alt(userProfile.Username),
			),
		),
	}

	return html.Render(root, minified)

}
```

Let's assume you want to create this blob of HTML:

```html
<div class="page-header">
  <table>
    <tr>
      <td>
        <img src="https://en.gravatar.com/userimage/389631/32413201cb4a680e82ee7c4bdd30b795.jpeg" alt="Title" class="app-icon-left shadow">
        </img>
      </td>
      <td>
        <h1>
          Title
        </h1> d
        <h3>
          Subtitle
        </h3>
      </td>
    </tr>
  </table>
</div>
```

To express the following HTML code in Go code using this library, you write:

```go
package components

import (
	"github.com/pieterclaerhout/go-html"
)

type PageHeader struct {
	Title    string
	Subtitle string
	IconURL  string
}

// HTML returns the HTML version of the component
func (pageHeader *PageHeader) HTML(minified bool) string {

	root := html.Blocks{
		html.Div(
			html.Class("page-header"),
			html.Table(nil,
				html.Tr(nil,
					html.Td(nil,
						html.Img(html.Src(pageHeader.IconURL).Class("app-icon-left shadow").Alt(pageHeader.Title)),
					),
					html.Td(nil,
						html.H1(nil, html.Text(pageHeader.Title)),
						html.H3(nil, html.Text(pageHeader.Subtitle)),
					),
				),
			),
		),
	}

	return html.Render(root, minified)

}
```

You can also make it slightly more complex, e.g. to build a HTML `select` box:

```html
<select>
  <option value="1">label 1</option>
  <option value="2" selected="true">label 2</option>
  <option value="3">label 3</option>
</select>
```

The equivalent Go code:

```go
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

	return html.Render(root, minified)

}
```

# Credits

This library is a fork of [https://github.com/mbertschler/blocks](https://github.com/mbertschler/blocks) with additional features and incompatible changes.