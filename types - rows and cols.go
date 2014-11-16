package main

import (
	"math/rand"
)

func (b *Block) ComputeRowsAndCols() {

	if b.Els == nil || len(b.Els) == 0 {
		panic("Set Elements first, then compute rows and cols number")
	}

	// editorial constraints
	if b.Fixed.Cols > 0 || b.Fixed.Rows > 0 {
		if b.Fixed.Rows > 0 {
			b.Rows = b.Fixed.Rows
			b.Cols = len(b.Els) / b.Rows
		}
		if b.Fixed.Cols > 0 {
			b.Cols = b.Fixed.Cols
			b.Rows = len(b.Els) / b.Cols
		}

	} else {
		// inherit from corridor
		par := b.Parent
		if par.Direction == Horizontal {
			b.Rows = par.Rows
			b.Cols = len(b.Els) / b.Rows
			if len(b.Els)%b.Rows > 0 {
				b.Cols++
			}

		} else {
			b.Cols = par.Cols
			b.Rows = len(b.Els) / b.Cols
			if len(b.Els)%b.Cols > 0 {
				b.Rows++
			}
		}
	}

}

func (c *Corridor) RandomizeDirectionAndRowsCols(islastCG bool) {

	p := c.Parent

	// The lastCG corridor completes/complements the parental viewport
	// TODO: Differentiate between
	//   horizontally filling up and vertically filling up
	if islastCG {
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
