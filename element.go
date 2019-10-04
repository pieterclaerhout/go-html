package html

// Element defines a HTML element
type Element struct {
	Type       string        // The type of the element
	Attributes               // The attributes for the element
	Children   Blocks        // The children for the element
	Options    ElementOption // The options for the element
}

// Elem returns a new ad-hoc element with the given name, attributes and children
func Elem(el string, attr Attributes, children ...Block) Block {
	return newElement(el, attr, children, 0)
}

// RenderHTML renders the element as HTML
func (Element) RenderHTML() Block {
	return nil
}
