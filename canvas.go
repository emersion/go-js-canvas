package canvas

import (
	"github.com/gopherjs/gopherjs/js"
)

type Canvas struct {
	Element *js.Object
	ctx *js.Object
	Width int
	Height int
}

func (c *Canvas) ClearRect(x, y, w, h int) {
	c.ctx.Call("clearRect", x, y, w, h)
}

func (c *Canvas) FillRect(x, y, w, h int) {
	c.ctx.Call("fillRect", x, y, w, h)
}

func NewCanvas(el *js.Object) *Canvas {
	return &Canvas{
		Element: el,
		ctx: el.Call("getContext", "2d"),
		Width: el.Get("width").Int(),
		Height: el.Get("height").Int(),
	}
}
