package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
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
	for _, v := range b.Els {

		// Even longer headlines should never by broken up
		if len(v.Headline) > 1 {
			all = append(all, make([]mixed, 1)...) // https://code.google.com/p/go-wiki/wiki/SliceTricks
			all[cntr].isHeadline = true
			all[cntr].text = v.Headline
			all[cntr].size = len(v.Headline)
			cntr++
		}

		// Now split the body into smaller elements, that can more easily be balanced
		sentencesRaw := strings.SplitAfter(v.Body, ".") // this could be fine tuned with ? and ! and comma/semiccolon - and exclusion of inc. and et. al.

		sentencsRefined := RefineTokenizedSlice(sentencesRaw, 15)
		sentencsRefined := SplitFurther(sentencesRaw, ",")

		numRefinedSentences := len(sentencsRefined)
		all = append(all, make([]mixed, numRefinedSentences)...)
		for j := 0; j < numRefinedSentences; j++ {
			all[cntr].isHeadline = false
			all[cntr].text = sentencsRefined[j]
			all[cntr].size = len(sentencsRefined[j])
			cntr++
		}
	}
	fmt.Println(spf("%s \n", spew.Sdump(all)))

	return []string{}
}

// RefineTokenizedSlice removes empty tokens
// It also recombines all elements shorter than "atLeast"
func RefineTokenizedSlice(sentencesRaw []string, atLeast int) []string {

	sentencsRefined := make([]string, len(sentencesRaw))
	idx2 := 0
	hangover := ""
	for idxRaw, v := range sentencesRaw {
		if len(sentencesRaw[idxRaw]) < 1 {
			continue // remove empty tokens
		}
		if len(sentencesRaw[idxRaw]) < atLeast {
			hangover += v
			continue
		}

		sentencsRefined[idx2] = strings.TrimSpace(hangover + v)
		hangover = ""
		idx2++
	}

	// a last hangover might be un-recombined:
	if len(hangover) > 0 {
		idx2++
		sentencsRefined[idx2] = strings.TrimSpace(hangover)
	}

	// constrain newly refined slice to non empty elements
	sentencsRefined = sentencsRefined[:idx2]

	return sentencsRefined

}

func SplitFurther(sentencesRaw []string, sep string) []string {

	sentencsRefined := make([]string, len(sentencesRaw))

	return sentencsRefined

}
