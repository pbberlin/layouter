package main

//
// ColumnGrouping() puts blocks into groups
// so that each group hat maximum row depth
func (c *Corridor) ColumnGroupingBlock() {

	corrMaxCols := 1000
	corrMaxRows := 1000
	if c.Direction == Vertical {
		corrMaxCols = c.Cols
	} else {
		corrMaxRows = c.Rows
	}
	pf(" -----------  %v %v \n", corrMaxRows, corrMaxCols)

	cgMaxCols := 0 // dynamic, depending on *first* block of each new col group
	cgMaxRows := 0 // dynamic, depending on *maximum height* in a row of col groups

	cgRowsCntr := 0 // row counter *inside* column group
	cgColsCntr := 0 // col counter *inside* column group

	_ = cgMaxRows
	_ = corrMaxRows

	//
	for idxBlock := 0; idxBlock < len(c.Blocks); idxBlock++ {
		lpBlock := &c.Blocks[idxBlock] // again the pointer thing...

		//
		// begin and end of column group
		if lpBlock.Rows >= corrMaxRows {
			c.ColumnGroupsB = append(c.ColumnGroupsB, make([][]*Block, 1)...) // https://code.google.com/p/go-wiki/wiki/SliceTricks
			appendBlock(c.ColumnGroupsB, lpBlock)
			pf(" ----cg-be -  %v %v\n", len(c.ColumnGroupsB), len(c.ColumnGroupsB[len(c.ColumnGroupsB)-1]))
			cgRowsCntr = 0
			cgColsCntr = 0
			cgMaxCols = 0
			continue
		}

		// begin of new column group - open to amends
		if cgRowsCntr == 0 &&
			cgRowsCntr+lpBlock.Rows < corrMaxRows {
			c.ColumnGroupsB = append(c.ColumnGroupsB, make([][]*Block, 1)...) // https://code.google.com/p/go-wiki/wiki/SliceTricks
			appendBlock(c.ColumnGroupsB, lpBlock)
			pf(" ----cg-b  -  %v %v\n", len(c.ColumnGroupsB), len(c.ColumnGroupsB[len(c.ColumnGroupsB)-1]))
			cgRowsCntr += lpBlock.Rows
			cgColsCntr += lpBlock.Cols
			cgMaxCols = lpBlock.Cols
			continue
		}

		// continue column group
		if cgRowsCntr != 0 &&
			(cgRowsCntr+lpBlock.Rows <= corrMaxRows ||
				cgColsCntr+lpBlock.Cols <= cgMaxCols) &&
			cgColsCntr+lpBlock.Cols <= corrMaxCols { // the last condition holds
			appendBlock(c.ColumnGroupsB, lpBlock)
			pf(" ----cg-c  -  %v %v\n", len(c.ColumnGroupsB), len(c.ColumnGroupsB[len(c.ColumnGroupsB)-1]))
			cgRowsCntr += lpBlock.Rows
			cgColsCntr += lpBlock.Cols
			if cgColsCntr >= cgMaxCols {
				cgColsCntr = 0
			}
			continue
		} else {
			// END of column group
			// overflow of current column group
			appendBlock(c.ColumnGroupsB, lpBlock)
			pf(" ----cg- e -  %v %v\n", len(c.ColumnGroupsB), len(c.ColumnGroupsB[len(c.ColumnGroupsB)-1]))
			cgRowsCntr = 0
			cgColsCntr = 0
			cgMaxCols = 0
			continue
		}

		panic("if cases should soak up all possibilities")

	}

}

func appendBlock(cg [][]*Block, b *Block) {
	lastCG := cg[len(cg)-1]
	lastCG = append(lastCG, make([]*Block, 1)...)
	lastCG[len(lastCG)-1] = b
	cg[len(cg)-1] = lastCG

}
