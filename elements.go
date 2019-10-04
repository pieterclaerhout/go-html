package html

func Doctype(arg string) Block {
	return newElement("!DOCTYPE", Attributes{{arg, nil}}, nil, Void)
}

func Html(attr Attributes, children ...Block) Block {
	return newElement("html", attr, children, 0)
}

func Head(attr Attributes, children ...Block) Block {
	return newElement("head", attr, children, 0)
}

func Noscript(attr Attributes, children ...Block) Block {
	return newElement("noscript", attr, children, 0)
}

func Iframe(attr Attributes, children ...Block) Block {
	return newElement("iframe", attr, children, 0)
}

func Link(attr Attributes) Block {
	return newElement("link", attr, nil, SelfClose)
}

func Img(attr Attributes) Block {
	return newElement("img", attr, nil, SelfClose)
}

func Meta(attr Attributes, children ...Block) Block {
	return newElement("meta", attr, children, Void)
}

func Title(attr Attributes, children ...Block) Block {
	return newElement("title", attr, children, 0)
}

func Body(attr Attributes, children ...Block) Block {
	return newElement("body", attr, children, 0)
}

func Button(attr Attributes, children ...Block) Block {
	return newElement("button", attr, children, 0)
}

func Style(attr Attributes, children ...Block) Block {
	return newElement("style", attr, children, CSSElement)
}

func Script(attr Attributes, children ...Block) Block {
	return newElement("script", attr, children, JSElement)
}

func Textarea(attr Attributes, children ...Block) Block {
	return newElement("textarea", attr, children, NoWhitespace)
}

func Main(attr Attributes, children ...Block) Block {
	return newElement("main", attr, children, 0)
}

func Form(attr Attributes, children ...Block) Block {
	return newElement("form", attr, children, 0)
}

func Nav(attr Attributes, children ...Block) Block {
	return newElement("nav", attr, children, 0)
}

func Span(attr Attributes, children ...Block) Block {
	return newElement("span", attr, children, 0)
}

func I(attr Attributes, children ...Block) Block {
	return newElement("i", attr, children, 0)
}

func Div(attr Attributes, children ...Block) Block {
	return newElement("div", attr, children, 0)
}

func P(attr Attributes, children ...Block) Block {
	return newElement("p", attr, children, 0)
}

func Ul(attr Attributes, children ...Block) Block {
	return newElement("ul", attr, children, 0)
}

func Ol(attr Attributes, children ...Block) Block {
	return newElement("ol", attr, children, 0)
}

func Li(attr Attributes, children ...Block) Block {
	return newElement("li", attr, children, 0)
}

func A(attr Attributes, children ...Block) Block {
	return newElement("a", attr, children, 0)
}

func H1(attr Attributes, children ...Block) Block {
	return newElement("h1", attr, children, 0)
}

func H2(attr Attributes, children ...Block) Block {
	return newElement("h2", attr, children, 0)
}

func H3(attr Attributes, children ...Block) Block {
	return newElement("h3", attr, children, 0)
}

func H4(attr Attributes, children ...Block) Block {
	return newElement("h4", attr, children, 0)
}

func H5(attr Attributes, children ...Block) Block {
	return newElement("h5", attr, children, 0)
}

func H6(attr Attributes, children ...Block) Block {
	return newElement("h6", attr, children, 0)
}

func Pre(attr Attributes, children ...Block) Block {
	return newElement("pre", attr, children, NoWhitespace)
}

func Label(attr Attributes, children ...Block) Block {
	return newElement("label", attr, children, 0)
}

func Strong(attr Attributes, children ...Block) Block {
	return newElement("strong", attr, children, 0)
}

func Input(attr Attributes) Block {
	return newElement("input", attr, nil, SelfClose)
}

func Br() Block {
	return newElement("br", nil, nil, SelfClose)
}

func Hr(attr Attributes) Block {
	return newElement("hr", attr, nil, SelfClose)
}

func Select(attr Attributes, children ...Block) Block {
	return newElement("select", attr, children, 0)
}

func Option(attr Attributes, children ...Block) Block {
	return newElement("option", attr, children, NoWhitespace)
}

func Table(attr Attributes, children ...Block) Block {
	return newElement("table", attr, children, 0)
}

func Thead(attr Attributes, children ...Block) Block {
	return newElement("thead", attr, children, 0)
}

func Th(attr Attributes, children ...Block) Block {
	return newElement("th", attr, children, 0)
}

func Tbody(attr Attributes, children ...Block) Block {
	return newElement("tbody", attr, children, 0)
}

func Tr(attr Attributes, children ...Block) Block {
	return newElement("tr", attr, children, 0)
}

func Td(attr Attributes, children ...Block) Block {
	return newElement("td", attr, children, 0)
}

func Tfoot(attr Attributes, children ...Block) Block {
	return newElement("tfoot", attr, children, 0)
}

func Article(attr Attributes, children ...Block) Block {
	return newElement("article", attr, children, 0)
}

func Section(attr Attributes, children ...Block) Block {
	return newElement("section", attr, children, 0)
}

func Wbr(attr Attributes, children ...Block) Block {
	return newElement("wbr", attr, children, 0)
}

func Var(attr Attributes, children ...Block) Block {
	return newElement("var", attr, children, 0)
}

func Video(attr Attributes, children ...Block) Block {
	return newElement("video", attr, children, 0)
}

func Source(attr Attributes) Block {
	return newElement("source", attr, nil, SelfClose)
}

func Track(attr Attributes) Block {
	return newElement("track", attr, nil, SelfClose)
}

func U(attr Attributes, children ...Block) Block {
	return newElement("u", attr, children, 0)
}

func Time(attr Attributes, children ...Block) Block {
	return newElement("time", attr, children, 0)
}

func Template(attr Attributes, children ...Block) Block {
	return newElement("template", attr, children, 0)
}

func Svg(attr Attributes, children ...Block) Block {
	return newElement("svg", attr, children, 0)
}

func Sup(attr Attributes, children ...Block) Block {
	return newElement("sup", attr, children, 0)
}

func Sub(attr Attributes, children ...Block) Block {
	return newElement("sub", attr, children, 0)
}

func Summary(attr Attributes, children ...Block) Block {
	return newElement("summary", attr, children, 0)
}

func Strike(attr Attributes, children ...Block) Block {
	return newElement("strike", attr, children, 0)
}

func Small(attr Attributes, children ...Block) Block {
	return newElement("small", attr, children, 0)
}

func Samp(attr Attributes, children ...Block) Block {
	return newElement("samp", attr, children, 0)
}

func S(attr Attributes, children ...Block) Block {
	return newElement("s", attr, children, 0)
}

func Ruby(attr Attributes, children ...Block) Block {
	return newElement("ruby", attr, children, 0)
}

func Rt(attr Attributes, children ...Block) Block {
	return newElement("rt", attr, children, 0)
}

func Rp(attr Attributes, children ...Block) Block {
	return newElement("rp", attr, children, 0)
}

func Q(attr Attributes, children ...Block) Block {
	return newElement("q", attr, children, 0)
}

func Progress(attr Attributes) Block {
	return newElement("progress", attr, nil, 0)
}

func Picture(attr Attributes, children ...Block) Block {
	return newElement("picture", attr, children, 0)
}

func Param(attr Attributes) Block {
	return newElement("param", attr, nil, SelfClose)
}

func Output(attr Attributes) Block {
	return newElement("output", attr, nil, 0)
}

func Optgroup(attr Attributes, children ...Block) Block {
	return newElement("optgroup", attr, children, 0)
}

func Object(attr Attributes) Block {
	return newElement("object", attr, nil, 0)
}

func Noframes(attr Attributes, children ...Block) Block {
	return newElement("noframes", attr, children, 0)
}

func Meter(attr Attributes, children ...Block) Block {
	return newElement("meter", attr, children, 0)
}

func Mark(attr Attributes, children ...Block) Block {
	return newElement("mark", attr, children, 0)
}

func Map(attr Attributes, children ...Block) Block {
	return newElement("map", attr, children, 0)
}

func Legend(attr Attributes, children ...Block) Block {
	return newElement("legend", attr, children, 0)
}

func Kbd(attr Attributes, children ...Block) Block {
	return newElement("kbd", attr, children, 0)
}

func Ins(attr Attributes, children ...Block) Block {
	return newElement("ins", attr, children, 0)
}

func Header(attr Attributes, children ...Block) Block {
	return newElement("header", attr, children, 0)
}

func Frameset(attr Attributes, children ...Block) Block {
	return newElement("frameset", attr, children, 0)
}

func Footer(attr Attributes, children ...Block) Block {
	return newElement("footer", attr, children, 0)
}

func Figure(attr Attributes, children ...Block) Block {
	return newElement("figure", attr, children, 0)
}

func Figcaption(attr Attributes, children ...Block) Block {
	return newElement("figcaption", attr, children, 0)
}

func Fieldset(attr Attributes, children ...Block) Block {
	return newElement("fieldset", attr, children, 0)
}

func Embed(attr Attributes) Block {
	return newElement("embed", attr, nil, SelfClose)
}

func Em(attr Attributes, children ...Block) Block {
	return newElement("em", attr, children, 0)
}

func Dt(attr Attributes, children ...Block) Block {
	return newElement("dt", attr, children, 0)
}

func Dl(attr Attributes, children ...Block) Block {
	return newElement("dl", attr, children, 0)
}

func Dir(attr Attributes, children ...Block) Block {
	return newElement("dir", attr, children, 0)
}

func Dialog(attr Attributes, children ...Block) Block {
	return newElement("dialog", attr, children, 0)
}

func Dfn(attr Attributes, children ...Block) Block {
	return newElement("dfn", attr, children, 0)
}

func Details(attr Attributes, children ...Block) Block {
	return newElement("details", attr, children, 0)
}

func Del(attr Attributes, children ...Block) Block {
	return newElement("del", attr, children, 0)
}

func Dd(attr Attributes, children ...Block) Block {
	return newElement("dd", attr, children, 0)
}

func Datalist(attr Attributes, children ...Block) Block {
	return newElement("datalist", attr, children, 0)
}

func Colgroup(attr Attributes, children ...Block) Block {
	return newElement("colgroup", attr, children, 0)
}

func Col(attr Attributes) Block {
	return newElement("col", attr, nil, SelfClose)
}

func Cite(attr Attributes, children ...Block) Block {
	return newElement("cite", attr, children, 0)
}

func Caption(attr Attributes, children ...Block) Block {
	return newElement("caption", attr, children, 0)
}

func Canvas(attr Attributes, children ...Block) Block {
	return newElement("canvas", attr, children, 0)
}

func Blockquote(attr Attributes, children ...Block) Block {
	return newElement("blockquote", attr, children, 0)
}

func Big(attr Attributes, children ...Block) Block {
	return newElement("big", attr, children, 0)
}

func Bdo(attr Attributes, children ...Block) Block {
	return newElement("bdo", attr, children, 0)
}

func Bdi(attr Attributes, children ...Block) Block {
	return newElement("bdi", attr, children, 0)
}

func Base(attr Attributes) Block {
	return newElement("base", attr, nil, SelfClose)
}

func B(attr Attributes, children ...Block) Block {
	return newElement("b", attr, children, 0)
}

func Audio(attr Attributes, children ...Block) Block {
	return newElement("audio", attr, children, 0)
}

func Aside(attr Attributes, children ...Block) Block {
	return newElement("aside", attr, children, 0)
}

func Area(attr Attributes) Block {
	return newElement("area", attr, nil, SelfClose)
}

func Address(attr Attributes, children ...Block) Block {
	return newElement("address", attr, children, 0)
}

func Acronym(attr Attributes, children ...Block) Block {
	return newElement("acronym", attr, children, 0)
}

func Abbr(attr Attributes, children ...Block) Block {
	return newElement("abbr", attr, children, 0)
}
