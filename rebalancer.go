package main

import (
	"fmt"
	_ "github.com/davecgh/go-spew/spew"
	"strings"
)

// RebalanceElements lumps all elements together into a slice of sentencesRaw.
// This slice is then been split by the number of columns.
// We aim for a perfectly even optical distribution of text across columns.
func (b *Block) RebalanceElements() []string {

	if b.Els == nil || len(b.Els) == 0 || b.Cols < 1 {
		panic("Set Elements and Cols first, then rebalance Elements")
	}

	type mixed struct {
		size       int
		isHeadline bool
		text       string
	}

	var all []mixed

	// all = append(all, make([]mixed, 10)...)

	cntr := 0
	sumSize := 0
	for _, v := range b.Els {

		// Even longer headlines should never by broken up
		if len(v.Headline) > 1 {
			all = append(all, make([]mixed, 1)...) // https://code.google.com/p/go-wiki/wiki/SliceTricks
			all[cntr].isHeadline = true
			all[cntr].text = v.Headline
			all[cntr].size = len(v.Headline)
			sumSize += len(v.Headline)
			cntr++
		}

		// Now split the body into smaller elements, that can more easily be balanced
		sentencesRaw := strings.SplitAfter(v.Body, ".")
		sentencsRefined := RecombineShortTokens(sentencesRaw, 15)
		sentencsRefined = SplitFurther(sentencsRefined, ",")
		sentencsRefined = RecombineShortTokens(sentencsRefined, 15)
		// We could now go on and split by
		// ? and ! and , and ;
		// We could exclude abbreviations such as inc. or and et. al.

		numRefinedSentences := len(sentencsRefined)
		all = append(all, make([]mixed, numRefinedSentences)...)
		for j := 0; j < numRefinedSentences; j++ {
			all[cntr].isHeadline = false
			all[cntr].text = sentencsRefined[j]
			all[cntr].size = len(sentencsRefined[j])
			sumSize += len(sentencsRefined[j])

			cntr++
		}
	}
	//fmt.Println(spf("%s \n", spew.Sdump(all)))

	// distribute
	colCounter := 0
	b.ElsPerCol = make([][]string, b.Cols)
	b.ElsPerColSizes = make([]int, b.Cols)
	idealColSize := sumSize / b.Cols
	fmt.Println(sumSize, b.Cols, idealColSize)

	colSizeAdder := 0
	colElCounter := 0
	for _, v := range all {
		colSizeAdder += v.size
		b.ElsPerCol[colCounter] = append(b.ElsPerCol[colCounter], make([]string, 1)...)
		b.ElsPerCol[colCounter][colElCounter] = v.text
		if v.isHeadline {
			b.ElsPerCol[colCounter][colElCounter] = spf("<h2>%s</h2>", b.ElsPerCol[colCounter][colElCounter])
		} else {
			//b.ElsPerCol[colCounter][colElCounter] = spf("%s<br>", b.ElsPerCol[colCounter][colElCounter])
		}
		colElCounter++
		b.ElsPerColSizes[colCounter] = colSizeAdder
		if colSizeAdder > idealColSize {
			colCounter++
			colSizeAdder = 0
			colElCounter = 0
		}

	}
	//fmt.Println(spf("%s \n", spew.Sdump(b.ElsPerCol)))

	return []string{}
}

// RecombineShortTokens removes empty tokens
// It also recombines all elements shorter than "atLeast"
func RecombineShortTokens(sentencesRaw []string, atLeast int) []string {

	sentencsRefined := make([]string, len(sentencesRaw))
	idx2 := 0
	hangover := ""
	for idxRaw, v := range sentencesRaw {

		if len(sentencesRaw[idxRaw]) < 1 {
			continue // remove empty tokens
		}
		if len(sentencesRaw[idxRaw]) < atLeast {
			hangover = hangover + v
			continue
		}

		sentencsRefined[idx2] = hangover + v
		hangover = ""
		idx2++
	}

	// a last hangover might be un-recombined:
	if len(hangover) > 0 {
		sentencsRefined[idx2-1] += hangover
		//pf("%s\n%s\n%s\n", hangover, sentencsRefined[idx2-1], sentencsRefined[idx2])
	}

	// constrain newly refined slice to non empty elements
	sentencsRefined = sentencsRefined[:idx2]

	return sentencsRefined

}

// SplitFurther takes an slice of strings
// and splits them by sep
// and returns a flattened array of string
func SplitFurther(sIn []string, sep string) []string {

	var sOut []string

	cntr := 0
	for _, v := range sIn {

		sLp := strings.SplitAfter(v, sep)
		sOut = append(sOut, make([]string, len(sLp))...)
		for j, k := range sLp {
			sOut[cntr+j] = k
		}
		cntr += len(sLp)
	}

	return sOut

}
