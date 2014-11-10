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

func generateRandomData() {
	rand.Seed(time.Now().UnixNano())

	//fmt.Printf("%s \n", spew.Sdump(vp))
	vp = Viewport{}
	vp.Cols = rand.Intn(5) + 2
	vp.CSS = util.CSSColumnsWidth(vp.Cols)

	nCorridors := rand.Intn(3) + 2
	vp.Corridors = make([]Corridor, nCorridors)

	for i1 := 0; i1 < len(vp.Corridors); i1++ {
		i1l := &vp.Corridors[i1] // the only way ...
		*i1l = Corridor{}        // ... to change the value of the slice element
		i1l.RandomizeDirectionAndRowsCols()

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
				i3l.Body += spf("%v.%v.%v: ", i1+1, i2+1, i3+1)
				i3l.Body += lorem.Paragraph(2, 5)
			}

			i2l.GenerateSizeSorting()
			i2l.ComputeRowsAndCols()

		}

	}
}

func dumpAll(vp Viewport, level int) string {

	if level == 1 {
		// remove the block to corridor pointers
		for i1 := 0; i1 < len(vp.Corridors); i1++ {
			i1l := &vp.Corridors[i1]
			for i2 := 0; i2 < len(i1l.Blocks); i2++ {
				i2l := &i1l.Blocks[i2]
				i2l.Parent = nil
			}
		}
	}
	return spf("%s \n", spew.Sdump(vp))
}

func init() {
	generateRandomData()
}
