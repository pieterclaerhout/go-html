package html

import (
	"html/template"
)

type UnsafeString string

func (UnsafeString) RenderHTML() Block {
	return nil
}

type Text string

func (Text) RenderHTML() Block {
	return nil
}

type Comment string

func (Comment) RenderHTML() Block {
	return nil
}

type CSS template.CSS

func (CSS) RenderHTML() Block {
	return nil
}

type HTML template.HTML

func (HTML) RenderHTML() Block {
	return nil
}

type HTMLAttr template.HTMLAttr

func (HTMLAttr) RenderHTML() Block {
	return nil
}

type JS template.JS

func (JS) RenderHTML() Block {
	return nil
}

func (b JS) renderString() string {
	return string(b)
}

type JSStr template.JSStr

func (JSStr) RenderHTML() Block {
	return nil
}

func (b JSStr) renderString() string {
	return string(b)
}

type URL template.URL

func (URL) RenderHTML() Block {
	return nil
}
