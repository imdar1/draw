package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/go-design/xy"
)

func NewTriangle(x, y int, class string) *Triangle {
	return &Triangle{
		Pos:   xy.Position{x, y},
		class: class,
	}
}

type Triangle struct {
	Pos   xy.Position
	class string
}

func (tri *Triangle) String() string {
	return fmt.Sprintf("triangle at %v", tri.Pos)
}

func (tri *Triangle) WriteSvg(out io.Writer) error {
	w, err := newTagPrinter(out)
	// the path is drawn as if it points straight to the right
	w.printf(`<path class="%s" d="M%v,%v l-8,-4 l 0,8 Z" />`,
		tri.class, tri.Pos.X, tri.Pos.Y)
	return *err
}
