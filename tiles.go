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
	Id                  int
	Name                string
	DisplayName         string
	Sprite              string
	Visible             bool
	LeftNeighbor        *Tile
	RightNeighbor       *Tile
	TopNeighbor         *Tile
	BottomNeighbor      *Tile
	TopLeftNeighbor     *Tile
	TopRightNeighbor    *Tile
	BottomLeftNeighbor  *Tile
	BottomRightNeighbor *Tile
	Chunk               *Chunk
}

var tiles map[int]Tile

func InitTiles() {
	tiles = map[int]Tile{
		tileGrass:                       {Id: tileGrass, Name: "grass", DisplayName: "grass", Sprite: "grass", Visible: true},
		tileGrassTopLeftCorner:          {Id: tileGrassTopLeftCorner, Name: "grassTopLeftCorner", DisplayName: "grass", Sprite: "grassTopLeftCorner", Visible: true},
		tileGrassTopRightCorner:         {Id: tileGrassTopRightCorner, Name: "grassTopRightCorner", DisplayName: "grass", Sprite: "grassTopRightCorner", Visible: true},
		tileGrassBottomLeftCorner:       {Id: tileGrassBottomLeftCorner, Name: "grassBottomLeftCorner", DisplayName: "grass", Sprite: "grassBottomLeftCorner", Visible: true},
		tileGrassBottomRightCorner:      {Id: tileGrassBottomRightCorner, Name: "grassBottomRightCorner", DisplayName: "grass", Sprite: "grassBottomRightCorner", Visible: true},
		tileGrassTopSide:                {Id: tileGrassTopSide, Name: "grassTopSide", DisplayName: "grass", Sprite: "grassTopSide", Visible: true},
		tileGrassBottomSide:             {Id: tileGrassBottomSide, Name: "grassBottomSide", DisplayName: "grass", Sprite: "grassBottomSide", Visible: true},
		tileGrassLeftSide:               {Id: tileGrassLeftSide, Name: "grassLeftSide", DisplayName: "grass", Sprite: "grassLeftSide", Visible: true},
		tileGrassRightSide:              {Id: tileGrassRightSide, Name: "grassRightSide", DisplayName: "grass", Sprite: "grassRightSide", Visible: true},
		tileWater:                       {Id: tileWater, Name: "water", DisplayName: "water", Sprite: "water", Visible: true},
		tileWaterTopLeftCorner:          {Id: tileWaterTopLeftCorner, Name: "waterTopLeftCorner", DisplayName: "water", Sprite: "waterTopLeftCorner", Visible: true},
		tileWaterTopRightCorner:         {Id: tileWaterTopRightCorner, Name: "waterTopRightCorner", DisplayName: "water", Sprite: "waterTopRightCorner", Visible: true},
		tileWaterBottomLeftCorner:       {Id: tileWaterBottomLeftCorner, Name: "waterBottomLeftCorner", DisplayName: "water", Sprite: "waterBottomLeftCorner", Visible: true},
		tileWaterBottomRightCorner:      {Id: tileWaterBottomRightCorner, Name: "waterBottomRightCorner", DisplayName: "water", Sprite: "waterBottomRightCorner", Visible: true},
		tileWaterTopSide:                {Id: tileWaterTopSide, Name: "waterTopSide", DisplayName: "water", Sprite: "waterTopSide", Visible: true},
		tileWaterBottomSide:             {Id: tileWaterBottomSide, Name: "waterBottomSide", DisplayName: "water", Sprite: "waterBottomSide", Visible: true},
		tileWaterLeftSide:               {Id: tileWaterLeftSide, Name: "waterLeftSide", DisplayName: "water", Sprite: "waterLeftSide", Visible: true},
		tileWaterRightSide:              {Id: tileWaterRightSide, Name: "waterRightSide", DisplayName: "water", Sprite: "waterRightSide", Visible: true},
		tileGrassTopLeftWaterCorner:     {Id: tileGrassTopLeftWaterCorner, Name: "grassTopLeftWaterCorner", DisplayName: "grass", Sprite: "grassTopLeftWaterCorner", Visible: true},
		tileGrassTopRightWaterCorner:    {Id: tileGrassTopRightWaterCorner, Name: "grassTopRightWaterCorner", DisplayName: "grass", Sprite: "grassTopRightWaterCorner", Visible: true},
		tileGrassBottomLeftWaterCorner:  {Id: tileGrassBottomLeftWaterCorner, Name: "grassBottomLeftWaterCorner", DisplayName: "grass", Sprite: "grassBottomLeftWaterCorner", Visible: true},
		tileGrassBottomRightWaterCorner: {Id: tileGrassBottomRightWaterCorner, Name: "grassBottomRightWaterCorner", DisplayName: "grass", Sprite: "grassBottomRightWaterCorner", Visible: true},
	}

	log.Println("Tiles initialized.")
}

func (t *Tile) DecideTileType() {
}

func (t *Tile) IsLeftEdge() bool {
	return t.LeftNeighbor != nil && t.LeftNeighbor.DisplayName != t.DisplayName
}

func (t *Tile) IsRightEdge() bool {
	return t.RightNeighbor != nil && t.RightNeighbor.DisplayName != t.DisplayName
}

func (t *Tile) IsTopEdge() bool {
	return t.TopNeighbor != nil && t.TopNeighbor.DisplayName != t.DisplayName
}

func (t *Tile) IsBottomEdge() bool {
	return t.BottomNeighbor != nil && t.BottomNeighbor.DisplayName != t.DisplayName
}

// func (c *Chunk) decideTileType(thisTile, leftTile, rightTile, topTile, bottomTile,
// 	topLeftCornerTile, topRightCornerTile, bottomLeftCornerTile, bottomRightCornerTile *Tile) Tile {
// 	// return *thisTile
// 	if thisTile.Id == tileWater {
// 		// We compare using displayName rather than Id because variants of the same kind of Tile
// 		// ("grassLeftSide", "grassTopRightCorner", etc) will all have the same displayName ("grass")

// 		if leftTile != nil && leftTile.DisplayName == "grass" && !leftTile.IsBorderTile {
// 			if topTile != nil && topTile.DisplayName == "grass" {
// 				// We're at the top left corner of a body of water
// 				return tiles[tileWaterTopLeftCorner]
// 			}
// 			if bottomTile != nil && bottomTile.DisplayName == "grass" {
// 				// We're at the bottom left corner of a body of water
// 				return tiles[tileWaterBottomLeftCorner]
// 			}

// 			// We're on the left side of a body of water
// 			return tiles[tileWaterLeftSide]
// 		}

// 		if rightTile != nil && rightTile.DisplayName == "grass" && !rightTile.IsCornerTile {
// 			if topTile != nil && topTile.DisplayName == "grass" {
// 				// We're at the top right corner of a body of water
// 				return tiles[tileWaterTopRightCorner]
// 			}
// 			if bottomTile != nil && bottomTile.DisplayName == "grass" {
// 				// We're at the top right corner of a body of water
// 				return tiles[tileWaterBottomRightCorner]
// 			}

// 			// We're on the right side of a body of water.
// 			return tiles[tileWaterRightSide]
// 		}

// 		if topTile != nil && topTile.DisplayName == "grass" && !topTile.IsBorderTile {
// 			// We're on the top side of a body of water.
// 			return tiles[tileWaterTopSide]
// 		}

// 		if bottomTile != nil && bottomTile.DisplayName == "grass" && !bottomTile.IsBorderTile {
// 			// We're on the bottom side of a body of water.
// 			return tiles[tileWaterBottomSide]
// 		}

// 		if topRightCornerTile != nil && topRightCornerTile.DisplayName == "grass" && !topRightCornerTile.IsBorderTile {
// 			// There's grass at the top right, use the speck Tile so we don't break the coastline
// 			return tiles[tileGrassTopRightWaterCorner]
// 		}
// 		if topLeftCornerTile != nil && topLeftCornerTile.DisplayName == "grass" && !topLeftCornerTile.IsBorderTile {
// 			// There's grass at the top left, use the speck Tile so we don't break the coastline
// 			return tiles[tileGrassTopLeftWaterCorner]
// 		}
// 		if bottomRightCornerTile != nil && bottomRightCornerTile.DisplayName == "grass" && !bottomRightCornerTile.IsBorderTile {
// 			// There's grass at the bottom right, use the speck Tile so we don't break the coastline
// 			return tiles[tileGrassBottomRightWaterCorner]
// 		}
// 		if bottomLeftCornerTile != nil && bottomLeftCornerTile.DisplayName == "grass" && !bottomLeftCornerTile.IsBorderTile {
// 			// There's grass at the bottom left, use the speck Tile so we don't break the coastline
// 			return tiles[tileGrassBottomLeftWaterCorner]
// 		}

// 		// We're not on any edge, just floating at sea <3
// 		return tiles[tileWater]
// 	}

// 	// Right now we only add water boundary Tiles.
// 	return *thisTile
// }
