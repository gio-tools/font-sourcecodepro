package sourcecodepro

import (
	"sync"

	"gio.tools/fonts/sourcecodepro/sourcecodeproblack"
	"gio.tools/fonts/sourcecodepro/sourcecodeproblackitalic"
	"gio.tools/fonts/sourcecodepro/sourcecodeprobold"
	"gio.tools/fonts/sourcecodepro/sourcecodeprobolditalic"
	"gio.tools/fonts/sourcecodepro/sourcecodeproextrabold"
	"gio.tools/fonts/sourcecodepro/sourcecodeproextrabolditalic"
	"gio.tools/fonts/sourcecodepro/sourcecodeproextralight"
	"gio.tools/fonts/sourcecodepro/sourcecodeproextralightitalic"
	"gio.tools/fonts/sourcecodepro/sourcecodeproitalic"
	"gio.tools/fonts/sourcecodepro/sourcecodeprolight"
	"gio.tools/fonts/sourcecodepro/sourcecodeprolightitalic"
	"gio.tools/fonts/sourcecodepro/sourcecodepromedium"
	"gio.tools/fonts/sourcecodepro/sourcecodepromediumitalic"
	"gio.tools/fonts/sourcecodepro/sourcecodeproregular"
	"gio.tools/fonts/sourcecodepro/sourcecodeprosemibold"
	"gio.tools/fonts/sourcecodepro/sourcecodeprosemibolditalic"

	"gioui.org/font"
	"gioui.org/font/opentype"
)

var (
	once       sync.Once
	collection []font.FontFace
)

func Collection() []font.FontFace {
	once.Do(func() {
		register(sourcecodeproblack.TTF)
		register(sourcecodeproblackitalic.TTF)
		register(sourcecodeprobold.TTF)
		register(sourcecodeprobolditalic.TTF)
		register(sourcecodeproextrabold.TTF)
		register(sourcecodeproextrabolditalic.TTF)
		register(sourcecodeproextralight.TTF)
		register(sourcecodeproextralightitalic.TTF)
		register(sourcecodeproitalic.TTF)
		register(sourcecodeprolight.TTF)
		register(sourcecodeprolightitalic.TTF)
		register(sourcecodepromedium.TTF)
		register(sourcecodepromediumitalic.TTF)
		register(sourcecodeproregular.TTF)
		register(sourcecodeprosemibold.TTF)
		register(sourcecodeprosemibolditalic.TTF)
		// Ensure that any outside appends will not reuse the backing store.
		n := len(collection)
		collection = collection[:n:n]
	})
	return collection
}

func register(src []byte) {
	faces, err := opentype.ParseCollection(src)
	if err != nil {
		panic("failed to parse font: " + err.Error())
	}
	collection = append(collection, faces[0])
}
