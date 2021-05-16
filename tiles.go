package engine

import (
	"log"
)

const (
	TileInvalid = iota
	TileGrass
	TileGrassTopLeftCorner
	TileGrassTopRightCorner
	TileGrassBottomLeftCorner
	TileGrassBottomRightCorner
	TileGrassTopSide
	TileGrassBottomSide
	TileGrassLeftSide
	TileGrassRightSide
	TileWater
	TileWaterTopLeftCorner
	TileWaterTopRightCorner
	TileWaterBottomLeftCorner
	TileWaterBottomRightCorner
	TileWaterTopSide
	TileWaterBottomSide
	TileWaterLeftSide
	TileWaterRightSide
	TileGrassBottomRightWaterCorner
	TileGrassBottomLeftWaterCorner
	TileGrassTopRightWaterCorner
	TileGrassTopLeftWaterCorner
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

var Tiles map[int]Tile

func InitTiles() {
	Tiles = map[int]Tile{
		TileGrass:                       {Id: TileGrass, Name: "grass", DisplayName: "grass", Sprite: "grass", Visible: true},
		TileGrassTopLeftCorner:          {Id: TileGrassTopLeftCorner, Name: "grassTopLeftCorner", DisplayName: "grass", Sprite: "grassTopLeftCorner", Visible: true},
		TileGrassTopRightCorner:         {Id: TileGrassTopRightCorner, Name: "grassTopRightCorner", DisplayName: "grass", Sprite: "grassTopRightCorner", Visible: true},
		TileGrassBottomLeftCorner:       {Id: TileGrassBottomLeftCorner, Name: "grassBottomLeftCorner", DisplayName: "grass", Sprite: "grassBottomLeftCorner", Visible: true},
		TileGrassBottomRightCorner:      {Id: TileGrassBottomRightCorner, Name: "grassBottomRightCorner", DisplayName: "grass", Sprite: "grassBottomRightCorner", Visible: true},
		TileGrassTopSide:                {Id: TileGrassTopSide, Name: "grassTopSide", DisplayName: "grass", Sprite: "grassTopSide", Visible: true},
		TileGrassBottomSide:             {Id: TileGrassBottomSide, Name: "grassBottomSide", DisplayName: "grass", Sprite: "grassBottomSide", Visible: true},
		TileGrassLeftSide:               {Id: TileGrassLeftSide, Name: "grassLeftSide", DisplayName: "grass", Sprite: "grassLeftSide", Visible: true},
		TileGrassRightSide:              {Id: TileGrassRightSide, Name: "grassRightSide", DisplayName: "grass", Sprite: "grassRightSide", Visible: true},
		TileWater:                       {Id: TileWater, Name: "water", DisplayName: "water", Sprite: "water", Visible: true},
		TileWaterTopLeftCorner:          {Id: TileWaterTopLeftCorner, Name: "waterTopLeftCorner", DisplayName: "water", Sprite: "waterTopLeftCorner", Visible: true},
		TileWaterTopRightCorner:         {Id: TileWaterTopRightCorner, Name: "waterTopRightCorner", DisplayName: "water", Sprite: "waterTopRightCorner", Visible: true},
		TileWaterBottomLeftCorner:       {Id: TileWaterBottomLeftCorner, Name: "waterBottomLeftCorner", DisplayName: "water", Sprite: "waterBottomLeftCorner", Visible: true},
		TileWaterBottomRightCorner:      {Id: TileWaterBottomRightCorner, Name: "waterBottomRightCorner", DisplayName: "water", Sprite: "waterBottomRightCorner", Visible: true},
		TileWaterTopSide:                {Id: TileWaterTopSide, Name: "waterTopSide", DisplayName: "water", Sprite: "waterTopSide", Visible: true},
		TileWaterBottomSide:             {Id: TileWaterBottomSide, Name: "waterBottomSide", DisplayName: "water", Sprite: "waterBottomSide", Visible: true},
		TileWaterLeftSide:               {Id: TileWaterLeftSide, Name: "waterLeftSide", DisplayName: "water", Sprite: "waterLeftSide", Visible: true},
		TileWaterRightSide:              {Id: TileWaterRightSide, Name: "waterRightSide", DisplayName: "water", Sprite: "waterRightSide", Visible: true},
		TileGrassTopLeftWaterCorner:     {Id: TileGrassTopLeftWaterCorner, Name: "grassTopLeftWaterCorner", DisplayName: "grass", Sprite: "grassTopLeftWaterCorner", Visible: true},
		TileGrassTopRightWaterCorner:    {Id: TileGrassTopRightWaterCorner, Name: "grassTopRightWaterCorner", DisplayName: "grass", Sprite: "grassTopRightWaterCorner", Visible: true},
		TileGrassBottomLeftWaterCorner:  {Id: TileGrassBottomLeftWaterCorner, Name: "grassBottomLeftWaterCorner", DisplayName: "grass", Sprite: "grassBottomLeftWaterCorner", Visible: true},
		TileGrassBottomRightWaterCorner: {Id: TileGrassBottomRightWaterCorner, Name: "grassBottomRightWaterCorner", DisplayName: "grass", Sprite: "grassBottomRightWaterCorner", Visible: true},
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
// 	if thisTile.Id == TileWater {
// 		// We compare using displayName rather than Id because variants of the same kind of Tile
// 		// ("grassLeftSide", "grassTopRightCorner", etc) will all have the same displayName ("grass")

// 		if leftTile != nil && leftTile.DisplayName == "grass" && !leftTile.IsBorderTile {
// 			if topTile != nil && topTile.DisplayName == "grass" {
// 				// We're at the top left corner of a body of water
// 				return tiles[TileWaterTopLeftCorner]
// 			}
// 			if bottomTile != nil && bottomTile.DisplayName == "grass" {
// 				// We're at the bottom left corner of a body of water
// 				return tiles[TileWaterBottomLeftCorner]
// 			}

// 			// We're on the left side of a body of water
// 			return tiles[TileWaterLeftSide]
// 		}

// 		if rightTile != nil && rightTile.DisplayName == "grass" && !rightTile.IsCornerTile {
// 			if topTile != nil && topTile.DisplayName == "grass" {
// 				// We're at the top right corner of a body of water
// 				return tiles[TileWaterTopRightCorner]
// 			}
// 			if bottomTile != nil && bottomTile.DisplayName == "grass" {
// 				// We're at the top right corner of a body of water
// 				return tiles[TileWaterBottomRightCorner]
// 			}

// 			// We're on the right side of a body of water.
// 			return tiles[TileWaterRightSide]
// 		}

// 		if topTile != nil && topTile.DisplayName == "grass" && !topTile.IsBorderTile {
// 			// We're on the top side of a body of water.
// 			return tiles[TileWaterTopSide]
// 		}

// 		if bottomTile != nil && bottomTile.DisplayName == "grass" && !bottomTile.IsBorderTile {
// 			// We're on the bottom side of a body of water.
// 			return tiles[TileWaterBottomSide]
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
// 		return tiles[TileWater]
// 	}

// 	// Right now we only add water boundary Tiles.
// 	return *thisTile
// }
