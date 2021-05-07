package engine

import (
	"encoding/json"
	"github.com/faiface/pixel"
	"io/ioutil"
	"log"
	"os"
)

var tileMapping map[string]pixel.Vec
var tileSprites map[string]*pixel.Sprite

type characterSpriteData struct {
	Name        string
	X           int
	Y           int
	Width       int
	Height      int
	FrameX      int
	FrameY      int
	FrameWidth  int
	FrameHeight int
}

func initTileSpriteMap() {
	tileMapping = map[string]pixel.Vec{
		"grassTopLeftCorner":     {0, 12},
		"grassTopSide":           {1, 12},
		"grassTopRightCorner":    {2, 12},
		"grassLeftSide":          {0, 11},
		"grass":                  {1, 11},
		"grassRightSide":         {2, 11},
		"grassBottomLeftCorner":  {0, 10},
		"grassBottomSide":        {1, 10},
		"grassBottomRightCorner": {2, 10},
		"waterTopLeftCorner":     {10, 12},
		"waterTopSide":           {11, 12},
		"waterTopRightCorner":    {12, 12},
		"waterLeftSide":          {10, 11},
		"water":                  {11, 11},
		"waterRightSide":         {12, 11},
		"waterBottomLeftCorner":  {10, 10},
		"waterBottomSide":        {11, 10},
		"waterBottomRightCorner": {12, 10},
		"waterBottomRightSpeck":  {13, 12},
		"waterBottomLeftSpeck":   {14, 12},
		"waterTopRightSpeck":     {13, 11},
		"waterTopLeftSpeck":      {14, 11},
	}

	tileSprites = map[string]*pixel.Sprite{}

	for name, vec := range tileMapping {
		tileSprites[name] = pixel.NewSprite(GTileSpritesheet, pixel.R(vec.X*TileSize, vec.Y*TileSize, (vec.X+1)*TileSize, (vec.Y+1)*TileSize))
	}

	log.Println("Sprites initialized.")
}

func initCharacterSpriteMap() {
	jsonFile, err := os.Open(CharacterJsonFile)

	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	bytesData, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(bytesData, &GCharacterSpriteData)

	if err != nil {
		panic(err)
	}

	GCharacterSprites = make(map[string]*pixel.Sprite, len(GCharacterSpriteData))

	for _, spriteData := range GCharacterSpriteData {
		GCharacterSprites[spriteData.Name] = pixel.NewSprite(GSpriteSheet, pixel.R(float64(spriteData.X), float64(spriteData.Y), float64(spriteData.X+spriteData.Width), float64(spriteData.Y+spriteData.Height)))
	}
}

func LoadSpriteSheet() {
	var err error
	GTileSpritesheet, err = loadPicture(TileSpritesheetFile)

	if err != nil {
		panic(err)
	}

	GSpriteSheet, err = loadPicture(CharacterSpritesheetFile)

	if err != nil {
		panic(err)
	}

	log.Println("Spritesheets loaded.")

	initTileSpriteMap()
	initCharacterSpriteMap()
}
