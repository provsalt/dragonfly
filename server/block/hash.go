// Code generated by cmd/blockhash; DO NOT EDIT.

package block

const hashAir = 0
const hashAncientDebris = 1
const hashAndesite = 2
const hashBarrier = 3
const hashBasalt = 4
const hashBeacon = 5
const hashBedrock = 6
const hashBeetrootSeeds = 7
const hashBlueIce = 8
const hashBoneBlock = 9
const hashBricks = 10
const hashCake = 11
const hashCalcite = 12
const hashCarpet = 13
const hashCarrot = 14
const hashChest = 15
const hashChiseledQuartz = 16
const hashClay = 17
const hashCoalBlock = 18
const hashCoalOre = 19
const hashCobblestone = 20
const hashCocoaBean = 21
const hashConcrete = 22
const hashConcretePowder = 23
const hashCopperOre = 24
const hashCoral = 25
const hashCoralBlock = 26
const hashDiamondBlock = 27
const hashDiamondOre = 28
const hashDiorite = 29
const hashDirt = 30
const hashDirtPath = 31
const hashDoubleFlower = 32
const hashDoubleTallGrass = 33
const hashDragonEgg = 34
const hashDripstone = 35
const hashEmeraldBlock = 36
const hashEmeraldOre = 37
const hashEndBrickStairs = 38
const hashEndBricks = 39
const hashEndStone = 40
const hashFarmland = 41
const hashFire = 42
const hashFlower = 43
const hashGildedBlackstone = 44
const hashGlass = 45
const hashGlassPane = 46
const hashGlazedTerracotta = 47
const hashGlowstone = 48
const hashGoldBlock = 49
const hashGoldOre = 50
const hashGranite = 51
const hashGrass = 52
const hashGravel = 53
const hashInvisibleBedrock = 54
const hashIronBars = 55
const hashIronBlock = 56
const hashIronOre = 57
const hashKelp = 58
const hashLantern = 59
const hashLapisBlock = 60
const hashLapisOre = 61
const hashLava = 62
const hashLeaves = 63
const hashLight = 64
const hashLitPumpkin = 65
const hashLog = 66
const hashMelon = 67
const hashMelonSeeds = 68
const hashMossCarpet = 69
const hashNetherBrickFence = 70
const hashNetherGoldOre = 71
const hashNetherQuartzOre = 72
const hashNetherSprouts = 73
const hashNetherWart = 74
const hashNetheriteBlock = 75
const hashNetherrack = 76
const hashNoteBlock = 77
const hashObsidian = 78
const hashPlanks = 79
const hashPotato = 80
const hashPrismarine = 81
const hashPumpkin = 82
const hashPumpkinSeeds = 83
const hashQuartz = 84
const hashQuartzBricks = 85
const hashQuartzPillar = 86
const hashRawCopperBlock = 87
const hashRawGoldBlock = 88
const hashRawIronBlock = 89
const hashSand = 90
const hashSandstone = 91
const hashSeaLantern = 92
const hashShroomlight = 93
const hashSoulSand = 94
const hashSoulSoil = 95
const hashSponge = 96
const hashSporeBlossom = 97
const hashStainedGlass = 98
const hashStainedGlassPane = 99
const hashStainedTerracotta = 100
const hashStone = 101
const hashTallGrass = 102
const hashTerracotta = 103
const hashTorch = 104
const hashTuff = 105
const hashWater = 106
const hashWheatSeeds = 107
const hashWoodDoor = 108
const hashWoodFence = 109
const hashWoodFenceGate = 110
const hashWoodSlab = 111
const hashWoodStairs = 112
const hashWoodTrapdoor = 113
const hashWool = 114

func (Air) Hash() uint64 {
	return hashAir
}

func (AncientDebris) Hash() uint64 {
	return hashAncientDebris
}

func (a Andesite) Hash() uint64 {
	return hashAndesite | uint64(boolByte(a.Polished))<<7
}

func (Barrier) Hash() uint64 {
	return hashBarrier
}

func (b Basalt) Hash() uint64 {
	return hashBasalt | uint64(boolByte(b.Polished))<<7 | uint64(b.Axis)<<8
}

func (Beacon) Hash() uint64 {
	return hashBeacon
}

func (b Bedrock) Hash() uint64 {
	return hashBedrock | uint64(boolByte(b.InfiniteBurning))<<7
}

func (b BeetrootSeeds) Hash() uint64 {
	return hashBeetrootSeeds | uint64(b.Growth)<<7
}

func (BlueIce) Hash() uint64 {
	return hashBlueIce
}

func (b BoneBlock) Hash() uint64 {
	return hashBoneBlock | uint64(b.Axis)<<7
}

func (Bricks) Hash() uint64 {
	return hashBricks
}

func (c Cake) Hash() uint64 {
	return hashCake | uint64(c.Bites)<<7
}

func (c Calcite) Hash() uint64 {
	return hashCalcite
}

func (c Carpet) Hash() uint64 {
	return hashCarpet | uint64(c.Colour.Uint8())<<7
}

func (c Carrot) Hash() uint64 {
	return hashCarrot | uint64(c.Growth)<<7
}

func (c Chest) Hash() uint64 {
	return hashChest | uint64(c.Facing)<<7
}

func (ChiseledQuartz) Hash() uint64 {
	return hashChiseledQuartz
}

func (c Clay) Hash() uint64 {
	return hashClay
}

func (CoalBlock) Hash() uint64 {
	return hashCoalBlock
}

func (c CoalOre) Hash() uint64 {
	return hashCoalOre | uint64(c.Type.Uint8())<<7
}

func (c Cobblestone) Hash() uint64 {
	return hashCobblestone | uint64(boolByte(c.Mossy))<<7
}

func (c CocoaBean) Hash() uint64 {
	return hashCocoaBean | uint64(c.Facing)<<7 | uint64(c.Age)<<9
}

func (c Concrete) Hash() uint64 {
	return hashConcrete | uint64(c.Colour.Uint8())<<7
}

func (c ConcretePowder) Hash() uint64 {
	return hashConcretePowder | uint64(c.Colour.Uint8())<<7
}

func (c CopperOre) Hash() uint64 {
	return hashCopperOre | uint64(c.Type.Uint8())<<7
}

func (c Coral) Hash() uint64 {
	return hashCoral | uint64(c.Type.Uint8())<<7 | uint64(boolByte(c.Dead))<<10
}

func (c CoralBlock) Hash() uint64 {
	return hashCoralBlock | uint64(c.Type.Uint8())<<7 | uint64(boolByte(c.Dead))<<10
}

func (DiamondBlock) Hash() uint64 {
	return hashDiamondBlock
}

func (d DiamondOre) Hash() uint64 {
	return hashDiamondOre | uint64(d.Type.Uint8())<<7
}

func (d Diorite) Hash() uint64 {
	return hashDiorite | uint64(boolByte(d.Polished))<<7
}

func (d Dirt) Hash() uint64 {
	return hashDirt | uint64(boolByte(d.Coarse))<<7
}

func (DirtPath) Hash() uint64 {
	return hashDirtPath
}

func (d DoubleFlower) Hash() uint64 {
	return hashDoubleFlower | uint64(boolByte(d.UpperPart))<<7 | uint64(d.Type.Uint8())<<8
}

func (d DoubleTallGrass) Hash() uint64 {
	return hashDoubleTallGrass | uint64(boolByte(d.UpperPart))<<7 | uint64(d.Type.Uint8())<<8
}

func (DragonEgg) Hash() uint64 {
	return hashDragonEgg
}

func (d Dripstone) Hash() uint64 {
	return hashDripstone
}

func (EmeraldBlock) Hash() uint64 {
	return hashEmeraldBlock
}

func (e EmeraldOre) Hash() uint64 {
	return hashEmeraldOre | uint64(e.Type.Uint8())<<7
}

func (s EndBrickStairs) Hash() uint64 {
	return hashEndBrickStairs | uint64(boolByte(s.UpsideDown))<<7 | uint64(s.Facing)<<8
}

func (EndBricks) Hash() uint64 {
	return hashEndBricks
}

func (EndStone) Hash() uint64 {
	return hashEndStone
}

func (f Farmland) Hash() uint64 {
	return hashFarmland | uint64(f.Hydration)<<7
}

func (f Fire) Hash() uint64 {
	return hashFire | uint64(f.Type.Uint8())<<7 | uint64(f.Age)<<8
}

func (f Flower) Hash() uint64 {
	return hashFlower | uint64(f.Type.Uint8())<<7
}

func (GildedBlackstone) Hash() uint64 {
	return hashGildedBlackstone
}

func (Glass) Hash() uint64 {
	return hashGlass
}

func (GlassPane) Hash() uint64 {
	return hashGlassPane
}

func (t GlazedTerracotta) Hash() uint64 {
	return hashGlazedTerracotta | uint64(t.Colour.Uint8())<<7 | uint64(t.Facing)<<11
}

func (Glowstone) Hash() uint64 {
	return hashGlowstone
}

func (GoldBlock) Hash() uint64 {
	return hashGoldBlock
}

func (g GoldOre) Hash() uint64 {
	return hashGoldOre | uint64(g.Type.Uint8())<<7
}

func (g Granite) Hash() uint64 {
	return hashGranite | uint64(boolByte(g.Polished))<<7
}

func (Grass) Hash() uint64 {
	return hashGrass
}

func (Gravel) Hash() uint64 {
	return hashGravel
}

func (InvisibleBedrock) Hash() uint64 {
	return hashInvisibleBedrock
}

func (IronBars) Hash() uint64 {
	return hashIronBars
}

func (IronBlock) Hash() uint64 {
	return hashIronBlock
}

func (i IronOre) Hash() uint64 {
	return hashIronOre | uint64(i.Type.Uint8())<<7
}

func (k Kelp) Hash() uint64 {
	return hashKelp | uint64(k.Age)<<7
}

func (l Lantern) Hash() uint64 {
	return hashLantern | uint64(boolByte(l.Hanging))<<7 | uint64(l.Type.Uint8())<<8
}

func (LapisBlock) Hash() uint64 {
	return hashLapisBlock
}

func (l LapisOre) Hash() uint64 {
	return hashLapisOre | uint64(l.Type.Uint8())<<7
}

func (l Lava) Hash() uint64 {
	return hashLava | uint64(boolByte(l.Still))<<7 | uint64(l.Depth)<<8 | uint64(boolByte(l.Falling))<<16
}

func (l Leaves) Hash() uint64 {
	return hashLeaves | uint64(l.Wood.Uint8())<<7 | uint64(boolByte(l.Persistent))<<10 | uint64(boolByte(l.ShouldUpdate))<<11
}

func (l Light) Hash() uint64 {
	return hashLight | uint64(l.Level)<<7
}

func (l LitPumpkin) Hash() uint64 {
	return hashLitPumpkin | uint64(l.Facing)<<7
}

func (l Log) Hash() uint64 {
	return hashLog | uint64(l.Wood.Uint8())<<7 | uint64(boolByte(l.Stripped))<<10 | uint64(l.Axis)<<11
}

func (Melon) Hash() uint64 {
	return hashMelon
}

func (m MelonSeeds) Hash() uint64 {
	return hashMelonSeeds | uint64(m.Growth)<<7 | uint64(m.Direction)<<15
}

func (m MossCarpet) Hash() uint64 {
	return hashMossCarpet
}

func (NetherBrickFence) Hash() uint64 {
	return hashNetherBrickFence
}

func (NetherGoldOre) Hash() uint64 {
	return hashNetherGoldOre
}

func (NetherQuartzOre) Hash() uint64 {
	return hashNetherQuartzOre
}

func (n NetherSprouts) Hash() uint64 {
	return hashNetherSprouts
}

func (n NetherWart) Hash() uint64 {
	return hashNetherWart | uint64(n.Age)<<7
}

func (NetheriteBlock) Hash() uint64 {
	return hashNetheriteBlock
}

func (Netherrack) Hash() uint64 {
	return hashNetherrack
}

func (n NoteBlock) Hash() uint64 {
	return hashNoteBlock
}

func (o Obsidian) Hash() uint64 {
	return hashObsidian | uint64(boolByte(o.Crying))<<7
}

func (p Planks) Hash() uint64 {
	return hashPlanks | uint64(p.Wood.Uint8())<<7
}

func (p Potato) Hash() uint64 {
	return hashPotato | uint64(p.Growth)<<7
}

func (p Prismarine) Hash() uint64 {
	return hashPrismarine | uint64(p.Type.Uint8())<<7
}

func (p Pumpkin) Hash() uint64 {
	return hashPumpkin | uint64(boolByte(p.Carved))<<7 | uint64(p.Facing)<<8
}

func (p PumpkinSeeds) Hash() uint64 {
	return hashPumpkinSeeds | uint64(p.Growth)<<7 | uint64(p.Direction)<<15
}

func (q Quartz) Hash() uint64 {
	return hashQuartz | uint64(boolByte(q.Smooth))<<7
}

func (QuartzBricks) Hash() uint64 {
	return hashQuartzBricks
}

func (q QuartzPillar) Hash() uint64 {
	return hashQuartzPillar | uint64(q.Axis)<<7
}

func (RawCopperBlock) Hash() uint64 {
	return hashRawCopperBlock
}

func (RawGoldBlock) Hash() uint64 {
	return hashRawGoldBlock
}

func (RawIronBlock) Hash() uint64 {
	return hashRawIronBlock
}

func (s Sand) Hash() uint64 {
	return hashSand | uint64(boolByte(s.Red))<<7
}

func (s Sandstone) Hash() uint64 {
	return hashSandstone | uint64(s.Type.Uint8())<<7 | uint64(boolByte(s.Red))<<9
}

func (SeaLantern) Hash() uint64 {
	return hashSeaLantern
}

func (Shroomlight) Hash() uint64 {
	return hashShroomlight
}

func (SoulSand) Hash() uint64 {
	return hashSoulSand
}

func (SoulSoil) Hash() uint64 {
	return hashSoulSoil
}

func (s Sponge) Hash() uint64 {
	return hashSponge | uint64(boolByte(s.Wet))<<7
}

func (s SporeBlossom) Hash() uint64 {
	return hashSporeBlossom
}

func (g StainedGlass) Hash() uint64 {
	return hashStainedGlass | uint64(g.Colour.Uint8())<<7
}

func (p StainedGlassPane) Hash() uint64 {
	return hashStainedGlassPane | uint64(p.Colour.Uint8())<<7
}

func (t StainedTerracotta) Hash() uint64 {
	return hashStainedTerracotta | uint64(t.Colour.Uint8())<<7
}

func (s Stone) Hash() uint64 {
	return hashStone | uint64(boolByte(s.Smooth))<<7
}

func (g TallGrass) Hash() uint64 {
	return hashTallGrass | uint64(g.Type.Uint8())<<7
}

func (Terracotta) Hash() uint64 {
	return hashTerracotta
}

func (t Torch) Hash() uint64 {
	return hashTorch | uint64(t.Facing)<<7 | uint64(t.Type.Uint8())<<10
}

func (t Tuff) Hash() uint64 {
	return hashTuff
}

func (w Water) Hash() uint64 {
	return hashWater | uint64(boolByte(w.Still))<<7 | uint64(w.Depth)<<8 | uint64(boolByte(w.Falling))<<16
}

func (s WheatSeeds) Hash() uint64 {
	return hashWheatSeeds | uint64(s.Growth)<<7
}

func (d WoodDoor) Hash() uint64 {
	return hashWoodDoor | uint64(d.Wood.Uint8())<<7 | uint64(d.Facing)<<10 | uint64(boolByte(d.Open))<<12 | uint64(boolByte(d.Top))<<13 | uint64(boolByte(d.Right))<<14
}

func (w WoodFence) Hash() uint64 {
	return hashWoodFence | uint64(w.Wood.Uint8())<<7
}

func (f WoodFenceGate) Hash() uint64 {
	return hashWoodFenceGate | uint64(f.Wood.Uint8())<<7 | uint64(f.Facing)<<10 | uint64(boolByte(f.Open))<<12 | uint64(boolByte(f.Lowered))<<13
}

func (s WoodSlab) Hash() uint64 {
	return hashWoodSlab | uint64(s.Wood.Uint8())<<7 | uint64(boolByte(s.Top))<<10 | uint64(boolByte(s.Double))<<11
}

func (s WoodStairs) Hash() uint64 {
	return hashWoodStairs | uint64(s.Wood.Uint8())<<7 | uint64(boolByte(s.UpsideDown))<<10 | uint64(s.Facing)<<11
}

func (t WoodTrapdoor) Hash() uint64 {
	return hashWoodTrapdoor | uint64(t.Wood.Uint8())<<7 | uint64(t.Facing)<<10 | uint64(boolByte(t.Open))<<12 | uint64(boolByte(t.Top))<<13
}

func (w Wool) Hash() uint64 {
	return hashWool | uint64(w.Colour.Uint8())<<7
}
