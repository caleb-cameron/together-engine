package engine

var (
	ChunksDir              string
	ChunkSize              int
	ChunkLoadRadius        float64
	ChunkLoadPadding       int
	WorldMaxAltitude       int
	GWorld                 World
	TileSize               int
	DataDir                string
	GPlayer                *Player
	PlayerSpeed            = 200.0
	PlayerAcceleration     = 120.0
	DefaultCharacterSprite = "manBlue_stand"
)
