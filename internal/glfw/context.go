/*
Copyright 2014 Hajime Hoshi

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package glfw

import (
	"github.com/hajimehoshi/ebiten"
)

type graphicsContext struct {
	canvas *canvas
}

var _ ebiten.GraphicsContext = new(graphicsContext)

func (c *graphicsContext) Clear() {
	c.canvas.use(func() {
		c.canvas.graphicsContext.Clear()
	})
}

func (c *graphicsContext) Fill(r, g, b uint8) {
	c.canvas.use(func() {
		c.canvas.graphicsContext.Fill(r, g, b)
	})
}

func (c *graphicsContext) Texture(id ebiten.TextureID) (d ebiten.Drawer) {
	c.canvas.use(func() {
		d = &drawer{
			canvas:      c.canvas,
			innerDrawer: c.canvas.graphicsContext.Texture(id),
		}
	})
	return
}

func (c *graphicsContext) RenderTarget(id ebiten.RenderTargetID) (d ebiten.Drawer) {
	c.canvas.use(func() {
		d = &drawer{
			canvas:      c.canvas,
			innerDrawer: c.canvas.graphicsContext.RenderTarget(id),
		}
	})
	return
}

func (c *graphicsContext) PopRenderTarget() {
	c.canvas.use(func() {
		c.canvas.graphicsContext.PopRenderTarget()
	})
}

func (c *graphicsContext) PushRenderTarget(id ebiten.RenderTargetID) {
	c.canvas.use(func() {
		c.canvas.graphicsContext.PushRenderTarget(id)
	})
}

type drawer struct {
	canvas      *canvas
	innerDrawer ebiten.Drawer
}

var _ ebiten.Drawer = new(drawer)

func (d *drawer) Draw(parts []ebiten.TexturePart, geo ebiten.GeometryMatrix, color ebiten.ColorMatrix) {
	d.canvas.use(func() {
		d.innerDrawer.Draw(parts, geo, color)
	})
}