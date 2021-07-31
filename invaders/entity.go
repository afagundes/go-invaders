package invaders

import (
	tl "github.com/JoelOtter/termloop"
	"io/ioutil"
)

func CreateCanvas(filename string) tl.Canvas {
	fileContent := GetEntityFromFile(filename)
	canvas := tl.CanvasFromString(string(fileContent))
	return canvas
}

func GetEntityFromFile(filename string) []byte {
	heroTxt, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return heroTxt
}
