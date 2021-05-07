package engine

import (
	"github.com/faiface/pixel"
	"log"
)

const (
	tileGrass = iota
	tileGrassTopLeftCorner
	tileGrassTopRightCorner
	tileGrassBottomLeftCorner
	tileGrassBottomRightCorner
	tileGrassTopSide
	tileGrassBottomSide
	tileGrassLeftSide
	tileGrassRightSide
	tileWater
	tileWaterTopLeftCorner
	tileWaterTopRightCorner
	tileWaterBottomLeftCorner
	tileWaterBottomRightCorner
	tileWaterTopSide
	tileWaterBottomSide
	tileWaterLeftSide
	tileWaterRightSide
	tileWaterBottomRightSpeck
	tileWaterBottomLeftSpeck
	tileWaterTopRightSpeck
	tileWaterTopLeftSpeck
)

type Tile struct {
	Id          int
	name        string
	displayName string
	sprite      *pixel.Sprite
	Visible     bool
}

var tiles map[int]Tile
var tilesBatch *pixel.Batch

func ClearTilesBatch() {
	tilesBatch.Clear()
}

func DrawTilesBatch(t pixel.Target) {
	tilesBatch.Draw(t)
}

func (t Tile) Draw(matrix pixel.Matrix) {
	t.sprite.Draw(tilesBatch, matrix)
}

func InitTiles() {
	tiles = map[int]Tile{
		tileGrass:                  {tileGrass, "grass", "grass", tileSprites["grass"], true},
		tileGrassTopLeftCorner:     {tileGrassTopLeftCorner, "grassTopLeftCorner", "grass", tileSprites["grassTopLeftCorner"], true},
		tileGrassTopRightCorner:    {tileGrassTopRightCorner, "grassTopRightCorner", "grass", tileSprites["grassTopRightCorner"], true},
		tileGrassBottomLeftCorner:  {tileGrassBottomLeftCorner, "grassBottomLeftCorner", "grass", tileSprites["grassBottomLeftCorner"], true},
		tileGrassBottomRightCorner: {tileGrassBottomRightCorner, "grassBottomRightCorner", "grass", tileSprites["grassBottomRightCorner"], true},
		tileGrassTopSide:           {tileGrassTopSide, "grassTopSide", "grass", tileSprites["grassTopSide"], true},
		tileGrassBottomSide:        {tileGrassBottomSide, "grassBottomSide", "grass", tileSprites["grassBottomSide"], true},
		tileGrassLeftSide:          {tileGrassLeftSide, "grassLeftSide", "grass", tileSprites["grassLeftSide"], true},
		tileGrassRightSide:         {tileGrassRightSide, "grassRightSide", "grass", tileSprites["grassRightSide"], true},
		tileWater:                  {tileWater, "water", "water", tileSprites["water"], true},
		tileWaterTopLeftCorner:     {tileWaterTopLeftCorner, "waterTopLeftCorner", "water", tileSprites["waterTopLeftCorner"], true},
		tileWaterTopRightCorner:    {tileWaterTopRightCorner, "waterTopRightCorner", "water", tileSprites["waterTopRightCorner"], true},
		tileWaterBottomLeftCorner:  {tileWaterBottomLeftCorner, "waterBottomLeftCorner", "water", tileSprites["waterBottomLeftCorner"], true},
		tileWaterBottomRightCorner: {tileWaterBottomRightCorner, "waterBottomRightCorner", "water", tileSprites["waterBottomRightCorner"], true},
		tileWaterTopSide:           {tileWaterTopSide, "waterTopSide", "water", tileSprites["waterTopSide"], true},
		tileWaterBottomSide:        {tileWaterBottomSide, "waterBottomSide", "water", tileSprites["waterBottomSide"], true},
		tileWaterLeftSide:          {tileWaterLeftSide, "waterLeftSide", "water", tileSprites["waterLeftSide"], true},
		tileWaterRightSide:         {tileWaterRightSide, "waterRightSide", "water", tileSprites["waterRightSide"], true},
		tileWaterBottomRightSpeck:  {tileWaterBottomRightSpeck, "waterBottomRightSpeck", "water", tileSprites["waterBottomRightSpeck"], true},
		tileWaterBottomLeftSpeck:   {tileWaterBottomLeftSpeck, "waterBottomLeftSpeck", "water", tileSprites["waterBottomLeftSpeck"], true},
		tileWaterTopRightSpeck:     {tileWaterTopRightSpeck, "waterTopRightSpeck", "water", tileSprites["waterTopRightSpeck"], true},
		tileWaterTopLeftSpeck:      {tileWaterTopLeftSpeck, "waterTopLeftSpeck", "water", tileSprites["waterTopLeftSpeck"], true},
	}

	tilesBatch = pixel.NewBatch(&pixel.TrianglesData{}, GTileSpritesheet)

	log.Println("Tiles initialized.")
}
