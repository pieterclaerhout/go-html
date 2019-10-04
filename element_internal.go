package html

func newElement(el string, attr Attributes, children []Block, opt ElementOption) Block {
	if len(children) == 0 {
		return Element{el, attr, nil, opt}
	}
	if len(children) == 1 {
		return Element{el, attr, children, opt}
	}
	return Element{el, attr, Blocks(children), opt}
}
