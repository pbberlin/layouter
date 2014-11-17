package main

import (
	"github.com/drhodes/golorem"
	"github.com/pbberlin/tools/util"
	"math/rand"
)

func generateData02() {

	vp := Viewport{}
	mVp["vp2"] = &vp
	//fmt.Printf("%s \n", spew.Sdump(vp))

	vp.Cols = 5
	vp.Rows = 5

	vp.CSS = util.CSSColumnsWidth(vp.Cols) + "\n" + util.CSSRowsHeight(vp.Rows)

	nCorridors := 3
	vp.Corridors = make([]Corridor, nCorridors)

	c1 := &vp.Corridors[0]
	*c1 = Corridor{}
	c1.Cols = 5
	c1.Rows = 2
	c1.Direction = Horizontal

	c2 := &vp.Corridors[1]
	*c2 = Corridor{}
	c2.Cols = 2
	c2.Rows = 3
	c2.Direction = Vertical

	c3 := &vp.Corridors[2]
	*c3 = Corridor{}
	c3.Cols = 3
	c3.Rows = 3
	c3.Direction = Vertical

	var numEls = [][]int{
		[]int{2, 2, 1, 1, 4},
		[]int{2, 1, 2, 1},
		[]int{4, 1, 1, 1, 2},
	}

	var numCols = [][]int{
		[]int{1, 2, 1, 1, 2},
		[]int{2, 1, 1, 1},
		[]int{2, 1, 1, 1, 2},
	}

	for i1 := 0; i1 < len(vp.Corridors); i1++ {
		i1l := &vp.Corridors[i1] // the only way ...
		i1l.Parent = &vp
		i1l.Blocks = make([]Block, 5)
		if i1 == 1 {
			i1l.Blocks = make([]Block, 4)
		}
		for i2 := 0; i2 < len(i1l.Blocks); i2++ {
			i2l := &i1l.Blocks[i2]
			*i2l = Block{}
			i2l.Parent = i1l
			i2l.IdxEditorial = i2
			if rand.Intn(3) < 1 {
				i2l.Headline += lorem.Sentence(6, 16)
			}

			i2l.Fixed.Cols = numCols[i1][i2]

			i2l.Els = make([]El, numEls[i1][i2])
			for i3 := 0; i3 < len(i2l.Els); i3++ {
				i3l := &i2l.Els[i3]
				*i3l = El{}
				if rand.Intn(4) < 1 {
					i3l.Headline += lorem.Sentence(7, 12)
				}
				prefix := spf("%v.%v.%v: ", i1+1, i2+1, i3+1)

				for i := 0; i < 2; i++ {
					p := lorem.Paragraph(4, 5)
					if i == 0 {
						i3l.Body += spf("<p>%s %s</p>", prefix, p)

					} else {
						i3l.Body += spf("<p>%s</p>", p)

					}
				}

			}

			i2l.GenerateSizeSorting()
			i2l.ComputeRowsAndCols()
			i2l.RebalanceElements()

		}
		i1l.ColumnGroupingBlock()

	}
}

func init() {
	generateData02()
}
