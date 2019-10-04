package components

import (
	"github.com/pieterclaerhout/go-html"
)

type PageHeader struct {
	Title    string
	Subtitle string
	IconURL  string
}

// <div class="page-header">
//   <table>
//     <tr>
//       <td>
//         <img src="https://en.gravatar.com/userimage/389631/32413201cb4a680e82ee7c4bdd30b795.jpeg" alt="Title" class="app-icon-left shadow">
//         </img>
//       </td>
//       <td>
//         <h1>
//           Title
//         </h1> d
//         <h3>
//           Subtitle
//         </h3>
//       </td>
//     </tr>
//   </table>
// </div>

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
