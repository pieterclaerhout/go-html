package html

import (
	"bytes"
	"fmt"
	"html"
	"html/template"
	"io"
)

var indentation = 2

type stringRenderer interface {
	renderString() string
}

type renderCtx struct {
	level    int
	item     int
	minified bool
}

func (c *renderCtx) enter() (item int) {
	item = c.item
	c.level++
	c.item = 0
	return item
}

func (c *renderCtx) next() {
	c.item++
}

func (c *renderCtx) exit(item int) {
	c.level--
	c.item = item
}

func renderHTML(c Block, w io.Writer, ctx *renderCtx) error {
	switch el := c.(type) {
	case nil:
		return nil
	case UnsafeString:
		if !ctx.minified {
			w.Write(bytes.Repeat([]byte{' '}, ctx.level*indentation))
		}
		w.Write([]byte(el))
		if !ctx.minified {
			w.Write([]byte{'\n'})
		}
		ctx.next()
	case Text:
		if !ctx.minified {
			w.Write(bytes.Repeat([]byte{' '}, ctx.level*indentation))
		}
		w.Write([]byte(template.HTMLEscapeString(string(el))))
		if !ctx.minified {
			w.Write([]byte{'\n'})
		}
		ctx.next()
	case stringRenderer:
		if !ctx.minified {
			w.Write(bytes.Repeat([]byte{' '}, ctx.level*indentation))
		}
		w.Write([]byte(el.renderString()))
		if !ctx.minified {
			w.Write([]byte{'\n'})
		}
		ctx.next()
	case Comment:
		if !ctx.minified {
			w.Write(bytes.Repeat([]byte{' '}, ctx.level*indentation))
		}
		w.Write([]byte("<!--" + el + "-->"))
		if !ctx.minified {
			w.Write([]byte{'\n'})
		}
		ctx.next()
	case Element:
		if !ctx.minified {
			w.Write(bytes.Repeat([]byte{' '}, ctx.level*indentation))
		}
		var attr string
		for _, v := range el.Attributes {
			if v.Value == nil {
				attr += " " + v.Key
				continue
			}
			attr += " " + v.Key + "=\"" + html.EscapeString(fmt.Sprint(v.Value)) + "\""
		}
		w.Write([]byte("<" + el.Type + attr))
		if el.Options&SelfClose != 0 {
			w.Write([]byte("/>"))
		} else {
			w.Write([]byte(">"))
		}
		if len(el.Children) > 0 {
			if !ctx.minified && el.Options&NoWhitespace == 0 {
				w.Write([]byte{'\n'})
			}
			item := ctx.enter()
			var min bool
			if el.Options&NoWhitespace != 0 {
				min = ctx.minified
				ctx.minified = true
			}
			for _, e := range el.Children {
				renderHTML(e, w, ctx)
			}
			if el.Options&NoWhitespace != 0 {
				ctx.minified = min
			}
			ctx.exit(item)
		}
		if el.Options&Void+el.Options&SelfClose == 0 {
			if !ctx.minified && el.Options&NoWhitespace == 0 && len(el.Children) > 0 {
				w.Write(bytes.Repeat([]byte{' '}, ctx.level*indentation))
			}
			w.Write([]byte("</" + el.Type + ">"))
		}
		if !ctx.minified {
			w.Write([]byte{'\n'})
		}
		ctx.next()
	case Blocks:
		for _, e := range el {
			renderHTML(e, w, ctx)
		}
	case Block:
		c := el.RenderHTML()
		renderHTML(c, w, ctx)
	default:
		if !ctx.minified {
			w.Write(bytes.Repeat([]byte{' '}, ctx.level*indentation))
		}
		fmt.Fprintf(w, "{{ ERROR value=%#v\n }}", c)
		if !ctx.minified {
			w.Write([]byte{'\n'})
		}
		ctx.next()
	}
	return nil
}
