package html

// Blocks defines a list of HTML blocks
type Blocks []Block

// Add adds a block to the list
func (b *Blocks) Add(block Block) {
	*b = append(*b, block)
}

// AddBlocks adds multiple blocks to the list
func (b *Blocks) AddBlocks(blocks Blocks) {
	*b = append(*b, blocks...)
}

// RenderHTML renders the block as HTML
func (Blocks) RenderHTML() Block {
	return nil
}
