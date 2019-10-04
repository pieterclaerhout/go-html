package components

import (
	"github.com/pieterclaerhout/go-html"
)

type UserProfile struct {
	Username  string
	AvatarURL string
}

// <div class="userProfile">
//   <h1 class="profileName">
//     pclaerhout
//   </h1>
//   <img src="https://en.gravatar.com/userimage/389631/32413201cb4a680e82ee7c4bdd30b795.jpeg" class="profileImage" alt="pclaerhout">
//   </img>
// </div>

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
