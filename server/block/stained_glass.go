package block

import (
	"github.com/df-mc/dragonfly/server/world"
)

// StainedGlass is a decorative, fully transparent solid block that is dyed into a different colour.
type StainedGlass struct {
	transparent
	solid
	clicksAndSticks

	// Colour specifies the colour of the block.
	Colour Colour
}

// BreakInfo ...
func (g StainedGlass) BreakInfo() BreakInfo {
	// TODO: Silk touch.
	return newBreakInfo(0.3, alwaysHarvestable, nothingEffective, simpleDrops())
}

// EncodeItem ...
func (g StainedGlass) EncodeItem() (name string, meta int16) {
	return "minecraft:stained_glass", int16(g.Colour.Uint8())
}

// EncodeBlock ...
func (g StainedGlass) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:stained_glass", map[string]interface{}{"color": g.Colour.String()}
}

// allStainedGlass returns stained glass blocks with all possible colours.
func allStainedGlass() []world.Block {
	b := make([]world.Block, 0, 16)
	for _, c := range Colours() {
		b = append(b, StainedGlass{Colour: c})
	}
	return b
}
