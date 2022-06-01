package squares

import (
	"image/color"
	"io"

	svg "github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
)

// RandomGradientColorSVG builds a square image with with x colors selected at random for each quadrant.
// the background color stays the same the other colors get mixed in a gradient color from the first one to the last one.
func RandomGradientColorSVG(w io.Writer, colors, gColors []color.RGBA, gv colors.GradientVector, width, height, xsquares int, prob float64) {

	var gradientColors []svg.Offcolor
	gradientColors = make([]svg.Offcolor, len(gColors))
	percentage := uint8(0)
	step := uint8(100 / len(gColors))
	for i, c := range gColors {
		gradientColors[i] = svg.Offcolor{
			Offset:  percentage,
			Color:   draw.RGBToHex(c.R, c.G, c.B),
			Opacity: 1,
		}
		percentage += step
	}

	canvas := svg.New(w)
	canvas.Start(width, height, draw.DefaultSvgAttributes()...)
	canvas.Def()
	canvas.LinearGradient("gradientColors", gv.X1, gv.Y1, gv.X2, gv.Y2, gradientColors)
	canvas.DefEnd()
	canvas.Rect(0, 0, width, height, "fill:url(#gradientColors)")

	squares := xsquares
	quadrantSize := width / squares
	colorMap := make(map[int]color.RGBA)
	colorIndex := make(map[int]int)
	for yQ := 0; yQ < squares; yQ++ {
		y := yQ * quadrantSize
		colorMap = make(map[int]color.RGBA)
		colorIndex = make(map[int]int)
		for xQ := 0; xQ <= squares+1; xQ++ {
			x := xQ * quadrantSize
			if _, ok := colorMap[xQ]; !ok {
				colorIndex[xQ] = draw.RandomIndexFromArrayWithFreq(colors, prob)
				colorMap[xQ] = colors[colorIndex[xQ]]
			}
			if colorIndex[xQ] == 0 {
				fill := draw.FillFromRGBA(colorMap[xQ])
				canvas.Rect(x, y, quadrantSize, quadrantSize, fill)
			}

		}
	}
	canvas.End()
}
