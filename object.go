package engine

import "github.com/faiface/pixel"

var ObjectTypes map[string]Object

type Object struct {
	Id          string
	DisplayName string
	SpriteName  string
	Position    pixel.Vec
}

func InitializeObjects() {
	ObjectTypes = map[string]Object{
		"tree": {
			Id:          "objTree",
			DisplayName: "tree",
			SpriteName:  "tree",
		},
	}

}
