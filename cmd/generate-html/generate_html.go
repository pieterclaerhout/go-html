package main

import (
	// "fmt"

	"github.com/pieterclaerhout/go-html"
	"github.com/pieterclaerhout/go-html/cmd/generate-html/components"
	"github.com/pieterclaerhout/go-log"
)

func main() {

	minimize := false

	userProfile := &components.UserProfile{
		Username:  "pclaerhout",
		AvatarURL: "https://en.gravatar.com/userimage/389631/32413201cb4a680e82ee7c4bdd30b795.jpeg",
	}
	log.Info(userProfile.HTML(minimize))

	pageHeader := &components.PageHeader{
		Title:    "Title",
		Subtitle: "Subtitle",
		IconURL:  "https://en.gravatar.com/userimage/389631/32413201cb4a680e82ee7c4bdd30b795.jpeg",
	}
	log.Info(pageHeader.HTML(minimize))

	selectBox := components.MySelect().Options([]components.Option{
		{"1", "label 1"},
		{"2", "label 2"},
		{"3", "label 3"},
	}).Selected("2")
	log.Info(selectBox.HTML(minimize))

	// root := html.Blocks{
	// 	// Option 1: directly add an element
	// 	html.Doctype("html"),
	// 	html.Html(nil,
	// 		// Option 2: struct that implements Block interface (RenderHTML() Block)
	// 		HeadBlock{html.Name("key").Content("super")},
	// 		// Option 3: function that returns a Block
	// 		BodyBlock("Hello, world! :) <br>"),
	// 	),
	// }
	// out, err := html.RenderString(root)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println(out)

	// out, err = html.RenderMinifiedString(root)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println(out)
}

type HeadBlock struct {
	html.Attributes
}

func (h HeadBlock) RenderHTML() html.Block {
	return html.Head(nil,
		html.Meta(h.Attributes),
	)
}

func BodyBlock(in string) html.Block {
	return html.Body(nil,
		html.Main(html.Class("main-class\" href=\"/evil/link"),
			html.H1(nil,
				html.Text(in),
				html.Br(),
				html.UnsafeString(in),
			),
		),
	)
}
