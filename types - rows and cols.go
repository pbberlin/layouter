package main

import (
	"math/rand"
)

// TODO - take editorial contraints into account
func (b *Block) ComputeRowsAndCols() {

	if b.Els == nil {
		panic("Set Elements first, then compute rows and cols number")
	}

	p := b.Parent

	if p.Direction == Horizontal {
		b.Rows = p.Rows
		b.Cols = len(b.Els) / b.Rows
		if len(b.Els)%b.Rows > 0 {
			b.Cols++
		}

	} else {
		b.Cols = p.Cols
		b.Rows = len(b.Els) / b.Cols
		if len(b.Els)%b.Cols > 0 {
			b.Rows++
		}
	}

}

func (c *Corridor) RandomizeDirectionAndRowsCols() {
	// Corridor direction
	direction := ExpandingTo(rand.Intn(2))
	if direction == Horizontal {
		c.Direction = Horizontal
		c.Fixed.Rows = 2 + rand.Intn(3)
		c.Rows = c.Fixed.Rows
	} else {
		c.Direction = Vertical
		c.Fixed.Cols = 2 + rand.Intn(4)
		c.Cols = c.Fixed.Cols
	}
}
