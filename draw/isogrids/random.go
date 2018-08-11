package isogrids

import (
	"image/color"
	"io"

	svg "github.com/ajstarks/svgo"
	"github.com/taironas/tinygraphs/draw"
)

// Random builds an image with 10x10 grids of half diagonals with random background.
// The image is symetric in the middle vertical axis.
func Random(w io.Writer, colors []color.RGBA, width, height, lines int, prob float64) {
	canvas := svg.New(w)
	canvas.Start(width, height, draw.DefaultSvgAttributes()...)

	fringeSize := width / lines
	distance := distanceTo3rdPoint(fringeSize)
	fringeSize = distance
	lines = width / fringeSize

	for xL := 0; xL < lines; xL++ {
		for yL := -1; yL <= lines; yL++ {
			var x1, x2, y1, y2, y3 int
			if (xL % 2) == 0 {
				x1, y1, x2, y2, _, y3 = right1stTriangle(xL, yL, fringeSize, distance)
			} else {
				x1, y1, x2, y2, _, y3 = left1stTriangle(xL, yL, fringeSize, distance)
			}
			xs := []int{x2, x1, x2}
			ys := []int{y1, y2, y3}
			canvas.Polygon(xs, ys, draw.FillFromRGBA(draw.RandomColorFromArrayWithFreq(colors, prob)))

			var x11, x12, y11, y12, y13 int
			if (xL % 2) == 0 {
				x11, y11, x12, y12, _, y13 = left2ndTriangle(xL, yL, fringeSize, distance)

				// we make sure that the previous triangle and this one touch each other in this point.
				y12 = y3
			} else {
				x11, y11, x12, y12, _, y13 = right2ndTriangle(xL, yL, fringeSize, distance)

				// we make sure that the previous triangle and this one touch each other in this point.
				y12 = y1 + fringeSize
			}
			xs1 := []int{x12, x11, x12}
			ys1 := []int{y11, y12, y13}
			canvas.Polygon(xs1, ys1, draw.FillFromRGBA(draw.RandomColorFromArrayWithFreq(colors, prob)))
		}
	}
	canvas.End()
}

// RandomGradient creates an isogrids svg image with half diagonals.
// colors are filled at random in the image with a frequency that decreases
// from left to right.
func RandomGradient(w io.Writer, colors []color.RGBA, width, height, lines int) {
	canvas := svg.New(w)
	canvas.Start(width, height, draw.DefaultSvgAttributes()...)

	fringeSize := width / lines
	distance := distanceTo3rdPoint(fringeSize)
	fringeSize = distance
	lines = width / fringeSize

	for xL := 0; xL < lines; xL++ {
		percentage := 100 - int(float64(xL)/float64(lines)*100)
		for yL := -1; yL < lines; yL++ {
			var x1, x2, y1, y2, y3 int
			if (xL % 2) == 0 {
				x1, y1, x2, y2, _, y3 = right1stTriangle(xL, yL, fringeSize, distance)
			} else {
				x1, y1, x2, y2, _, y3 = left1stTriangle(xL, yL, fringeSize, distance)
			}
			xs := []int{x2, x1, x2}
			ys := []int{y1, y2, y3}
			canvas.Polygon(xs, ys, draw.FillFromRGBA(draw.ColorByPercentage(colors, percentage)))

			var x11, x12, y11, y12, y13 int
			if (xL % 2) == 0 {
				x11, y11, x12, y12, _, y13 = left2ndTriangle(xL, yL, fringeSize, distance)

				// we make sure that the previous triangle and this one touch each other in this point.
				y12 = y3
			} else {
				x11, y11, x12, y12, _, y13 = right2ndTriangle(xL, yL, fringeSize, distance)

				// in order to have a perfect hexagon,
				// we make sure that the previous triangle and this one touch each other in this point.
				y12 = y1 + fringeSize
			}
			xs1 := []int{x12, x11, x12}
			ys1 := []int{y11, y12, y13}
			canvas.Polygon(xs1, ys1, draw.FillFromRGBA(draw.ColorByPercentage(colors, percentage)))
		}
	}
	canvas.End()
}

// RandomMirror builds an image with 10x10 grids of half diagonals
func RandomMirror(w io.Writer, colors []color.RGBA, size int, prob float64) {
	canvas := svg.New(w)
	canvas.Start(size, size, draw.DefaultSvgAttributes()...)

	lines := 10
	fringeSize := size / lines
	distance := distanceTo3rdPoint(fringeSize)
	fringeSize = distance
	lines = size / fringeSize

	for xL := 0; xL <= lines/2; xL++ {
		for yL := -1; yL < lines; yL++ {
			var x1, x2, y1, y2, y3 int
			if (xL % 2) == 0 {
				x1, y1, x2, y2, _, y3 = right1stTriangle(xL, yL, fringeSize, distance)
			} else {
				x1, y1, x2, y2, _, y3 = left1stTriangle(xL, yL, fringeSize, distance)
			}
			xs := []int{x2, x1, x2}
			ys := []int{y1, y2, y3}

			fill1 := draw.FillFromRGBA(draw.RandomColorFromArrayWithFreq(colors, prob))
			canvas.Polygon(xs, ys, fill1)

			xs = mirrorCoordinates(xs, lines, fringeSize, 0)
			canvas.Polygon(xs, ys, fill1)

			var x11, x12, y11, y12, y13 int
			if (xL % 2) == 0 {
				x11, y11, x12, y12, _, y13 = left2ndTriangle(xL, yL, fringeSize, distance)

				// we make sure that the previous triangle and this one touch each other in this point.
				y12 = y3
			} else {
				x11, y11, x12, y12, _, y13 = right2ndTriangle(xL, yL, fringeSize, distance)

				// in order to have a perfect hexagon,
				// we make sure that the previous triangle and this one touch each other in this point.
				y12 = y1 + fringeSize
			}
			xs1 := []int{x12, x11, x12}
			ys1 := []int{y11, y12, y13}

			fill2 := draw.FillFromRGBA(draw.RandomColorFromArrayWithFreq(colors, prob))
			canvas.Polygon(xs1, ys1, fill2)

			xs1 = mirrorCoordinates(xs1, lines, fringeSize, 0)
			canvas.Polygon(xs1, ys1, fill2)

		}
	}
	canvas.End()
}
