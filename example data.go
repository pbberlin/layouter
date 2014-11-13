package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/drhodes/golorem"
	"github.com/pbberlin/tools/util"
	"math/rand"
	"time"
)

var e1 = El{"Headline 01", "http://www.google.com", "Body 01 Body 01 Body 01 Body 01 Body 01"}
var e2 = El{"Headline 02", "http://www.microsoft.com", "Body 02 Body 02 Body 02"}

var vp = Viewport{}

func generateRandomData(nColsViewPort int) {
	rand.Seed(time.Now().UnixNano())

	//fmt.Printf("%s \n", spew.Sdump(vp))
	vp = Viewport{}
	if nColsViewPort < 1 {
		vp.Cols = rand.Intn(5) + 2

	} else {
		vp.Cols = nColsViewPort
	}
	vp.Rows = rand.Intn(5) + 2
	vp.CSS = util.CSSColumnsWidth(vp.Cols) + "\n" + util.CSSRowsHeight(vp.Rows)

	nCorridors := rand.Intn(3) + 2
	nCorridors = 3
	vp.Corridors = make([]Corridor, nCorridors)

	for i1 := 0; i1 < len(vp.Corridors); i1++ {
		i1l := &vp.Corridors[i1] // the only way ...
		*i1l = Corridor{}        // ... to change the value of the slice element
		i1l.Parent = &vp
		i1l.RandomizeDirectionAndRowsCols(i1 == len(vp.Corridors)-1)

		i1l.Blocks = make([]Block, 2+rand.Intn(4))
		for i2 := 0; i2 < len(i1l.Blocks); i2++ {
			i2l := &i1l.Blocks[i2]
			*i2l = Block{}
			i2l.Parent = i1l
			if rand.Intn(3) < 1 {
				i2l.Headline += lorem.Sentence(3, 6)
			}
			i2l.Els = make([]El, 3+rand.Intn(5))
			for i3 := 0; i3 < len(i2l.Els); i3++ {
				i3l := &i2l.Els[i3]
				*i3l = El{}
				if rand.Intn(3) < 1 {
					i3l.Headline += lorem.Sentence(4, 10)
				}
				prefix := spf("%v.%v.%v: ", i1+1, i2+1, i3+1)

				for i := 0; i < 1+rand.Intn(3); i++ {
					p := lorem.Paragraph(2, 5)
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

	}
}

func dumpAll(vp Viewport, level int) string {

	// remove some stuff to reduce redundant verbosity
	if level == 1 {
		for i1 := 0; i1 < len(vp.Corridors); i1++ {
			i1l := &vp.Corridors[i1]
			// remove the corridor to viewport pointers
			i1l.Parent = nil
			for i2 := 0; i2 < len(i1l.Blocks); i2++ {
				i2l := &i1l.Blocks[i2]
				// remove the block to corridor pointers
				i2l.Parent = nil
			}
		}
	}
	return spf("%s \n", spew.Sdump(vp))
}

func init() {
	generateRandomData(0)
}
