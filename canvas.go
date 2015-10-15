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

func (c *Canvas) SetFillStyle(fillStyle string) {
	c.ctx.Set("fillStyle", fillStyle)
}

func (c *Canvas) SetStrokeStyle(strokeStyle string) {
	c.ctx.Set("strokeStyle", strokeStyle)
}

func (c *Canvas) ClearRect(x, y, w, h int) {
	c.ctx.Call("clearRect", x, y, w, h)
}

func (c *Canvas) FillRect(x, y, w, h int) {
	c.ctx.Call("fillRect", x, y, w, h)
}

func (c *Canvas) StrokeRect(x, y, w, h int) {
	c.ctx.Call("strokeRect", x, y, w, h)
}

func New(el *js.Object) *Canvas {
	return &Canvas{
		Element: el,
		ctx: el.Call("getContext", "2d"),
		Width: el.Get("width").Int(),
		Height: el.Get("height").Int(),
	}
}
