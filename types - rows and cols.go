package main

import (
	"math/rand"
)

// TODO - take editorial contraints into account
func (b *Block) ComputeRowsAndCols() {

	if b.Els == nil || len(b.Els) == 0 {
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

func (c *Corridor) RandomizeDirectionAndRowsCols(isLast bool) {

	p := c.Parent

	// The last corridor completes/complements the parental viewport
	// TODO: Differentiate between
	//   horizontally filling up and vertically filling up
	if isLast {
		c.Direction = Horizontal
		c.Fixed.Cols = p.Cols - p.ColsConsumed
		c.Cols = c.Fixed.Cols
		c.Rows = 2
	} else {
		direction := ExpandingTo(rand.Intn(2))
		if direction == Horizontal {
			c.Direction = Horizontal
			c.Fixed.Rows = 1 + rand.Intn(p.Rows)
			c.Rows = c.Fixed.Rows

			p.Rowsconsumed += c.Rows
		} else {
			c.Direction = Vertical
			c.Fixed.Cols = 1 + rand.Intn(p.Cols)
			c.Cols = c.Fixed.Cols

			p.ColsConsumed += c.Cols
		}
	}

}
