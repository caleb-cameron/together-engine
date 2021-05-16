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
	tileGrassBottomRightWaterCorner
	tileGrassBottomLeftWaterCorner
	tileGrassTopRightWaterCorner
	tileGrassTopLeftWaterCorner
)

type Tile struct {
	Id           int
	Name         string
	DisplayName  string
	Sprite       string
	Visible      bool
	IsBorderTile bool
	IsCornerTile bool
}

var tiles map[int]Tile

func InitTiles() {
	tiles = map[int]Tile{
		tileGrass:                       {tileGrass, "grass", "grass", "grass", true, false, false},
		tileGrassTopLeftCorner:          {tileGrassTopLeftCorner, "grassTopLeftCorner", "grass", "grassTopLeftCorner", true, true, true},
		tileGrassTopRightCorner:         {tileGrassTopRightCorner, "grassTopRightCorner", "grass", "grassTopRightCorner", true, true, true},
		tileGrassBottomLeftCorner:       {tileGrassBottomLeftCorner, "grassBottomLeftCorner", "grass", "grassBottomLeftCorner", true, true, true},
		tileGrassBottomRightCorner:      {tileGrassBottomRightCorner, "grassBottomRightCorner", "grass", "grassBottomRightCorner", true, true, true},
		tileGrassTopSide:                {tileGrassTopSide, "grassTopSide", "grass", "grassTopSide", true, true, false},
		tileGrassBottomSide:             {tileGrassBottomSide, "grassBottomSide", "grass", "grassBottomSide", true, true, false},
		tileGrassLeftSide:               {tileGrassLeftSide, "grassLeftSide", "grass", "grassLeftSide", true, true, false},
		tileGrassRightSide:              {tileGrassRightSide, "grassRightSide", "grass", "grassRightSide", true, true, false},
		tileWater:                       {tileWater, "water", "water", "water", true, false, false},
		tileWaterTopLeftCorner:          {tileWaterTopLeftCorner, "waterTopLeftCorner", "water", "waterTopLeftCorner", true, true, true},
		tileWaterTopRightCorner:         {tileWaterTopRightCorner, "waterTopRightCorner", "water", "waterTopRightCorner", true, true, true},
		tileWaterBottomLeftCorner:       {tileWaterBottomLeftCorner, "waterBottomLeftCorner", "water", "waterBottomLeftCorner", true, true, true},
		tileWaterBottomRightCorner:      {tileWaterBottomRightCorner, "waterBottomRightCorner", "water", "waterBottomRightCorner", true, true, true},
		tileWaterTopSide:                {tileWaterTopSide, "waterTopSide", "water", "waterTopSide", true, true, false},
		tileWaterBottomSide:             {tileWaterBottomSide, "waterBottomSide", "water", "waterBottomSide", true, true, false},
		tileWaterLeftSide:               {tileWaterLeftSide, "waterLeftSide", "water", "waterLeftSide", true, true, false},
		tileWaterRightSide:              {tileWaterRightSide, "waterRightSide", "water", "waterRightSide", true, true, false},
		tileGrassTopLeftWaterCorner:     {tileGrassTopLeftWaterCorner, "grassTopLeftWaterCorner", "grass", "grassTopLeftWaterCorner", true, true, true},
		tileGrassTopRightWaterCorner:    {tileGrassTopRightWaterCorner, "grassTopRightWaterCorner", "grass", "grassTopRightWaterCorner", true, true, true},
		tileGrassBottomLeftWaterCorner:  {tileGrassBottomLeftWaterCorner, "grassBottomLeftWaterCorner", "grass", "grassBottomLeftWaterCorner", true, true, true},
		tileGrassBottomRightWaterCorner: {tileGrassBottomRightWaterCorner, "grassBottomRightWaterCorner", "grass", "grassBottomRightWaterCorner", true, true, true},
	}

	log.Println("Tiles initialized.")
}
