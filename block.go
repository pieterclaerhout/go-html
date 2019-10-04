package html

type Block interface {
	RenderHTML() Block
}
