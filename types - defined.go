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
	Cols int
	Rows int
	// Horizontal bool
	// Vertical   bool
}

// Block is a rectangle, bracketing a number of elements
type Block struct {
	Dims
	Parent   *Corridor
	Headline string // a block may have its own Headline
	Subline  string // and Abstract
	Els      []El   // a slice of elements, the pointer enabled access while ranging - http://stackoverflow.com/questions/15945030/change-values-while-iterating-in-golang
	BySize   SSizeInfo
	Fixed    Dims // a shadow - with values set by the editor - not the computed ones
}

// Corridor is a stream of Rectangles
// It expands horizontal, with constant rows.
// Or grows vertical, constant cols.
type Corridor struct {
	Dims
	Blocks []Block
	Fixed  Dims // a shadow - with values set by the editor - not the computed ones
}

// Viewport contains the corridors
type Viewport struct {
	Dims
	Corridors []Corridor
}
