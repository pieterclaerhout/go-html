package html

import (
	"bytes"
	"io"
)

func Render(blocks Block, minified bool) string {

	f := RenderString
	if minified {
		f = RenderMinifiedString
	}
	output, _ := f(blocks)
	return output

}

func RenderToWriter(w io.Writer, root Block) error {
	err := renderHTML(root, w, &renderCtx{})
	if err != nil {
		return err
	}
	return nil
}

func RenderMinified(root Block, w io.Writer) error {
	err := renderHTML(root, w, &renderCtx{minified: true})
	if err != nil {
		return err
	}
	return nil
}

func RenderString(root Block) (string, error) {
	buf := bytes.Buffer{}
	err := renderHTML(root, &buf, &renderCtx{})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func RenderMinifiedString(root Block) (string, error) {
	buf := bytes.Buffer{}
	err := renderHTML(root, &buf, &renderCtx{minified: true})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
