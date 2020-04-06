package world

// Structure represents a structure which may be placed in the world. It has fixed dimensions.
type Structure interface {
	// Dimensions returns the dimensions of the structure. It returns an int array with the width, height and
	// length respectively.
	Dimensions() [3]int
	// At returns the block at a specific location in the structure. When the structure is placed in the
	// world, this method is called for every location within the dimensions of the structure.
	// At may return nil to not place any block at the position. Returning Air will set any block at that
	// position to air, but returning nil will not do anything.
	At(x, y, z int) Block
}