package engine

import (
	"log"
)

const (
	tileInvalid = iota
	tileGrass
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
	Name        string
	DisplayName string
	Sprite      string
	Visible     bool
}

var tiles map[int]Tile

func InitTiles() {
	tiles = map[int]Tile{
		tileGrass:                  {tileGrass, "grass", "grass", "grass", true},
		tileGrassTopLeftCorner:     {tileGrassTopLeftCorner, "grassTopLeftCorner", "grass", "grassTopLeftCorner", true},
		tileGrassTopRightCorner:    {tileGrassTopRightCorner, "grassTopRightCorner", "grass", "grassTopRightCorner", true},
		tileGrassBottomLeftCorner:  {tileGrassBottomLeftCorner, "grassBottomLeftCorner", "grass", "grassBottomLeftCorner", true},
		tileGrassBottomRightCorner: {tileGrassBottomRightCorner, "grassBottomRightCorner", "grass", "grassBottomRightCorner", true},
		tileGrassTopSide:           {tileGrassTopSide, "grassTopSide", "grass", "grassTopSide", true},
		tileGrassBottomSide:        {tileGrassBottomSide, "grassBottomSide", "grass", "grassBottomSide", true},
		tileGrassLeftSide:          {tileGrassLeftSide, "grassLeftSide", "grass", "grassLeftSide", true},
		tileGrassRightSide:         {tileGrassRightSide, "grassRightSide", "grass", "grassRightSide", true},
		tileWater:                  {tileWater, "water", "water", "water", true},
		tileWaterTopLeftCorner:     {tileWaterTopLeftCorner, "waterTopLeftCorner", "water", "waterTopLeftCorner", true},
		tileWaterTopRightCorner:    {tileWaterTopRightCorner, "waterTopRightCorner", "water", "waterTopRightCorner", true},
		tileWaterBottomLeftCorner:  {tileWaterBottomLeftCorner, "waterBottomLeftCorner", "water", "waterBottomLeftCorner", true},
		tileWaterBottomRightCorner: {tileWaterBottomRightCorner, "waterBottomRightCorner", "water", "waterBottomRightCorner", true},
		tileWaterTopSide:           {tileWaterTopSide, "waterTopSide", "water", "waterTopSide", true},
		tileWaterBottomSide:        {tileWaterBottomSide, "waterBottomSide", "water", "waterBottomSide", true},
		tileWaterLeftSide:          {tileWaterLeftSide, "waterLeftSide", "water", "waterLeftSide", true},
		tileWaterRightSide:         {tileWaterRightSide, "waterRightSide", "water", "waterRightSide", true},
		tileWaterBottomRightSpeck:  {tileWaterBottomRightSpeck, "waterBottomRightSpeck", "water", "waterBottomRightSpeck", true},
		tileWaterBottomLeftSpeck:   {tileWaterBottomLeftSpeck, "waterBottomLeftSpeck", "water", "waterBottomLeftSpeck", true},
		tileWaterTopRightSpeck:     {tileWaterTopRightSpeck, "waterTopRightSpeck", "water", "waterTopRightSpeck", true},
		tileWaterTopLeftSpeck:      {tileWaterTopLeftSpeck, "waterTopLeftSpeck", "water", "waterTopLeftSpeck", true},
	}

	log.Println("Tiles initialized.")
}
