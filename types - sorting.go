package main

import (
	"sort"
)

//
//
type SizeInfo struct {
	Size  int
	Index int // index in main slice
}
type SSizeInfo []SizeInfo

func (e SSizeInfo) Len() int {
	return len(e)
}
func (e SSizeInfo) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
func (e SSizeInfo) Less(i, j int) bool {
	return e[i].Size < e[j].Size
}

//
//

// GenerateSizeSorting fills the slice BySize with data from slice Els
// Then sorts this data by the combined lengths of headline and body
func (b *Block) GenerateSizeSorting() {

	if b.Els == nil {
		panic("Set Elements first, then call sorting method")
	}

	b.BySize = make(SSizeInfo, len(b.Els)) // we must pre-allocate the destinaton slice
	for i, v := range b.Els {
		b.BySize[i].Index = i
		b.BySize[i].Size = len(v.Headline) + len(v.Body)
	}
	sort.Sort(b.BySize)
}
