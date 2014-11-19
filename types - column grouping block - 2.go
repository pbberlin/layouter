package main

//import "github.com/pbberlin/tools/util"

//
// ColumnGrouping() puts blocks into groups
// so that each group hat maximum row depth
func (c *Corridor) ColumnGroupingBlock() {

	//
	// max values and counters for the corridor
	//    note the two granularities - first atomic rows and cols
	//    second - counters for *column group* rows and cols
	corrMaxCols := 1000
	corrMaxRows := 1000
	if c.Direction == Vertical {
		corrMaxCols = c.Cols
	} else {
		corrMaxRows = c.Rows
	}
	corrRowsCntr := 0     // corridor wide counter
	corrColsCntr := 0     // corridor wide counter
	corrClGrRowsCntr := 0 // column groups rows - i.e. one magnitude bigger than clgrMaxRows
	corrclgrColsCntr := 0 // column groups rows
	_ = corrclgrColsCntr  //   but we never use it
	pf("\n----COR----%v--%v--\n", corrMaxRows, corrMaxCols)

	//
	// the max values and counters *inside* the column group
	clgrMaxCols := 0 // dynamic, depending on *first* block of each new col group
	clgrMaxRows := 0 // dynamic, depending on *start height* in a row of col groups

	clgrRowsCntr := 0 // row counter *inside* column group
	clgrColsCntr := 0 // col counter *inside* column group

	var clgrIsOpened bool

	rowsCorrective := 0 // special variable for special case

	//
	for idxBlock := 0; idxBlock < len(c.Blocks); idxBlock++ {
		lpBlock := &c.Blocks[idxBlock] // again - for iterating over structs - changing origin - not copies

		// common precomputations
		clgrRowsCntr += lpBlock.Rows
		clgrColsCntr += lpBlock.Cols

		// corrRowsCntr only on newline
		rowsCorrective = -lpBlock.Rows
		corrColsCntr += lpBlock.Cols

		//pf("      lp start idx %v iOp:%v \n", idxBlock, clgrIsOpened)
		bl1 := "<="
		bl2 := "<="
		bl3 := "<="
		bl4 := "<="
		if clgrRowsCntr > clgrMaxRows {
			bl1 = "> "
		}
		if clgrColsCntr > clgrMaxCols {
			bl2 = "> "
		}
		if corrRowsCntr > corrMaxRows {
			bl3 = "> "
		}
		if corrColsCntr > corrMaxCols {
			bl4 = "> "
		}
		pf("\t %v%v%v %v%v%v   corr %v%v%v %v%v%v  ",
			clgrRowsCntr, bl1, clgrMaxRows,
			clgrColsCntr, bl2, clgrMaxCols,
			corrRowsCntr, bl3, corrMaxRows,
			corrColsCntr, bl4, corrMaxCols,
		)

		// we start out from the most special cases:
		// append horizontal - one row
		//              xxxxxxxxxx
		//              xxxx | new
		if true &&
			clgrIsOpened &&
			// even if we append to the right, need checking vertical constraints
			// both are tautological from previous loop:
			corrRowsCntr+rowsCorrective <= corrMaxRows && // fitting into corridor height
			clgrRowsCntr+rowsCorrective <= clgrMaxRows && // fitting into CG row height
			//
			corrColsCntr <= corrMaxCols && // smaller than corridor
			clgrColsCntr <= clgrMaxCols && // fitting into CG width
			true {

			// re-adjust row count
			// clgrRowsCntr = clgrRowsCntr + rowsCorrective // Wrong thought
			// corrRowsCntr = corrRowsCntr + rowsCorrective // only on newline

			clgrIsOpened = true
			appendBlock(c.ColumnGroupsB, lpBlock)
			pf("\t\tcg-c hor B%v - %v %v\n", idxBlock, len(c.ColumnGroupsB), len(c.ColumnGroupsB[len(c.ColumnGroupsB)-1]))
			continue
		}

		// append vertical - one row
		if true &&
			clgrIsOpened &&
			corrRowsCntr <= corrMaxRows && // fitting into corridor height
			clgrRowsCntr <= clgrMaxRows && // fitting into CG row height
			true {
			// re-adjust cg col count
			clgrColsCntr = 0
			corrColsCntr = corrColsCntr - lpBlock.Cols
			clgrIsOpened = true
			appendBlock(c.ColumnGroupsB, lpBlock)
			pf("\t\tcg-c ver B%v - %v %v\n", idxBlock, len(c.ColumnGroupsB), len(c.ColumnGroupsB[len(c.ColumnGroupsB)-1]))
			continue
		}

		// We only reach this position, if above appending
		// is neither horizontally nor vertically possible
		// therefore
		clgrIsOpened = false

		//
		// CG line break
		var newRow string = "       "
		if true &&
			idxBlock == 0 || // first CG line ?
			corrColsCntr > corrMaxCols && // new CG line ? - *wrap around condition*
				true {

			clgrMaxRows = lpBlock.Rows // set CG max rows - depending on newline
			corrRowsCntr += lpBlock.Rows

			// reset counters
			corrColsCntr = lpBlock.Cols

			clgrColsCntr = lpBlock.Cols
			clgrRowsCntr = lpBlock.Rows

			corrClGrRowsCntr++ // but unused

			newRow = "new row"

		}

		var expandableHor bool = c.Direction == Horizontal && corrColsCntr <= corrMaxCols
		var expandableVer bool = c.Direction == Vertical && corrRowsCntr <= corrMaxRows
		dirMsg := "Hor"
		if expandableVer {
			dirMsg = "Ver"
		}

		// Now open new column group
		if true && (expandableHor || expandableVer) &&
			true {
			clgrColsCntr = lpBlock.Cols
			clgrRowsCntr = lpBlock.Rows

			clgrMaxCols = lpBlock.Cols // set CG max cols - depending on first block
			c.ColumnGroupsB = appendColumnGroup(c.ColumnGroupsB)
			appendBlock(c.ColumnGroupsB, lpBlock)
			pf("\t\tcg-b     B%v - %v %v - %v - %v - %v %v\n", idxBlock, len(c.ColumnGroupsB),
				len(c.ColumnGroupsB[len(c.ColumnGroupsB)-1]), dirMsg, newRow, clgrMaxRows, clgrMaxCols)
			clgrIsOpened = true // remember an open CG

			continue
		}

		panic("if cases should soak up all possibilities \n  a block may be too large or tall ")

	}

}

func appendColumnGroup(cg [][]*Block) [][]*Block {
	// https://code.google.com/p/go-wiki/wiki/SliceTricks
	retCg := append(cg, make([][]*Block, 1)...)
	return retCg
}

func appendBlock(cg [][]*Block, b *Block) {
	lastCG := cg[len(cg)-1]
	lastCG = append(lastCG, make([]*Block, 1)...)
	lastCG[len(lastCG)-1] = b
	cg[len(cg)-1] = lastCG

}
