package main

//
// ColumnGrouping() puts blocks into groups
// so that each group hat maximum row depth
func (c *Corridor) ColumnGrouping() {

	if c.Direction == Vertical {
		return
	}

	nRowsCorridor := c.Rows
	nColsGroup := 0 // dynamic, depending on last opening column group

	rowsPrev := 0
	colsPrev := 0

	for idxBlock, lpBlock := range c.Blocks {

		//
		// begin and end of column group
		if lpBlock.Rows >= nRowsCorridor {
			c.ColumnGroups = append(c.ColumnGroups, make([][]int, 1)...) // https://code.google.com/p/go-wiki/wiki/SliceTricks
			appendCurrentColumnGroup(c.ColumnGroups, idxBlock)
			pf(" --cg-be -  %v \n", len(c.ColumnGroups))
			rowsPrev = 0
			colsPrev = 0
			nColsGroup = 0
			continue
		}

		// begin of column group
		if rowsPrev == 0 &&
			rowsPrev+lpBlock.Rows < nRowsCorridor {
			c.ColumnGroups = append(c.ColumnGroups, make([][]int, 1)...) // https://code.google.com/p/go-wiki/wiki/SliceTricks
			appendCurrentColumnGroup(c.ColumnGroups, idxBlock)
			pf(" ----cg-b  -  %v \n", len(c.ColumnGroups))
			rowsPrev += lpBlock.Rows
			colsPrev += lpBlock.Cols
			nColsGroup = lpBlock.Cols
			continue
		}

		// continue column group
		if rowsPrev != 0 &&
			(rowsPrev+lpBlock.Rows <= nRowsCorridor ||
				colsPrev+lpBlock.Cols <= nColsGroup) {
			appendCurrentColumnGroup(c.ColumnGroups, idxBlock)
			pf(" ----cg-c  -  %v \n", len(c.ColumnGroups))
			rowsPrev += lpBlock.Rows
			colsPrev += lpBlock.Cols
			if colsPrev >= nColsGroup {
				colsPrev = 0
			}
			continue
		} else {
			// overflow of current column group
			// => start a new one
			appendCurrentColumnGroup(c.ColumnGroups, idxBlock)
			pf(" ----cg- e -  %v \n", len(c.ColumnGroups))
			rowsPrev = 0
			colsPrev = 0
			nColsGroup = 0
			continue
		}

		panic("if cases should soak up all possibilities")

	}

}

func appendCurrentColumnGroup(cg [][]int, idxBlock int) {
	lastCG := cg[len(cg)-1]
	lastCG = append(lastCG, make([]int, 1)...)
	lastCG[len(lastCG)-1] = idxBlock
	cg[len(cg)-1] = lastCG

}
