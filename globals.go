package engine

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font"
)

const (
	TileSpritesheetFile      = "res/sprites/Tiles.png"
	CharacterSpritesheetFile = "res/sprites/characters.png"
	CharacterJsonFile        = "res/sprites/characters.json"
	DefaultCharacterSprite   = "manBlue_stand"
	DataDir                  = "data/"
	ChunksDir                = "data/chunks/"
	ChunkSize                = 30
	TileSize                 = 64
	WorldMaxAltitude         = 30
	ChunkLoadRadius          = 1
	MaxZoom                  = 3
	MinZoom                  = 0.5
)

var (
	GSpriteSheet         pixel.Picture
	GPlayer              *Player
	GTileSpritesheet     pixel.Picture
	GWorld               World
	GWindowCfg           pixelgl.WindowConfig
	GWindow              *pixelgl.Window
	GCharacterSpriteData []characterSpriteData
	GCharacterSprites    map[string]*pixel.Sprite
	GCamera              *Camera
	GGroundFriction      float64
	GFrameTime           float64
	GFramesThisSecond    int
	GFramesLastSecond    int
	GTextAtlas           *text.Atlas
	GChunkTxt            *text.Text
	GFontFace            font.Face
)
