package main

/*

	Viewport with six cols, unlimited rows
	Three corridors, first horizonally expanding,
	second and third vertically expanding.

	First corridor has height of one row  and contains two blocks.

	Seond corridor has width  of two cols and contains one block.
	Third corridor has width  of one col  and contains two blocks.

	The blocks contain several elements.

    _______________________________
    | XX  xxx | XXXXXXXX          |
    | xxx xxx | xxxxxxxxxxxxxxxxx |
    | xxx xxx | xxxxxxxxxxxxxxxxx |
    |-----------------------------|
    |--------------||-------------|
    | XXXX   xxxxx || XXXXXXXXX   |
    | xxxxx  xxxxx || xxxxxxxxxxx |
    |        xxxxx || xxxxxxxxxxx |
    | XXXX   xxxxx ||-------------|
    | xxxxx        || XXXXX       |
    | xxxxx  XXXXX || xxxxxxxxxxx |
    | xxxxx  xxxxx || xxxxxxxxxxx |
    -------------------------------


*/

// mVp is the major "cache" for viewport structures
var mVp = map[string]*Viewport{} // map of viewports

//
// An Element is a piece of unformatted textual or graphical content.
// Similar to RSS atom.
// It is small enough to be displayed into a part of a block
type El struct {
	Headline string
	URL      string
	Body     string
}

// This is an enum, but can not be restricted
// towards invalid values like  ExpandingTo(44)
type ExpandingTo int

const (
	Horizontal ExpandingTo = iota
	Vertical
)

// Dimensions  - a helper struct with fields repeating several times
type Dims struct {
	Direction ExpandingTo // constant Horizontal or Vertical
	// expanding direction - can inherit from surrounding corridor
	Cols         int
	Rows         int
	ColsConsumed int // dyn. counter for the lower levels
	Rowsconsumed int
}

// Block is a rectangle, bracketing a number of elements
type Block struct {
	Dims
	IdxEditorial int // index in the initial "editorial" Slice
	IdxGlobal    int // unique identification
	Parent       *Corridor
	Headline     string // a block may have its own Headline
	Subline      string // and Abstract
	Els          []El   // a slice of elements
	BySize       SSizeInfo
	Fixed        Dims // a shadow - with values set by the editor - not the computed ones

	ElsPerCol      [][]string // #Column - filled from Els by RebalanceElements()
	ElsPerColSizes []int
}

// Corridor is a stream of Blocks, that is Rectangles.
// It expands horizontal, with constant rows.
// Or grows vertical, constant cols.
type Corridor struct {
	Dims
	Blocks        []Block
	ColumnGroupsB [][]*Block // pointers to blocks in Blocks - dynamically allocated, according to corridor rows and cols
	Fixed         Dims       // a shadow - with values set by the editor - not the computed ones
	Parent        *Viewport
}

// Viewport contains the corridors
type Viewport struct {
	Dims
	CSS       string // dynamically generated CSS stuff
	Corridors []Corridor
}
