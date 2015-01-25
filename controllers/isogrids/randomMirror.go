package isogrids

import (
	"image/color"
	"net/http"

	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/extract"
	"github.com/taironas/tinygraphs/write"
)

func RandomMirror(w http.ResponseWriter, r *http.Request) {

	theme := extract.Theme(r)
	colorMap := colors.MapOfColorThemes()
	var err error
	var bg, fg color.RGBA
	if bg, err = extract.Background(r); err != nil {
		bg = colorMap["base"][0]
	}
	if fg, err = extract.Foreground(r); err != nil {
		fg = colorMap["base"][1]
	}

	if val, ok := colorMap[theme]; ok {
		bg = val[0]
		fg = val[1]
	}

	size := extract.Size(r)
	write.ImageSVG(w)
	draw.IsogridsRandomMirror(w, "", bg, fg, size)
}
