package shape_test

import (
	"testing"

	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/shape"
	. "github.com/gregoryv/draw/shape"
)

func Test_write_allshapes(t *testing.T) {
	d := design.NewDiagram()
	vspace := 60

	actorLbl, actor := NewLabel("Actor"), NewActor()
	d.Place(actorLbl).At(20, 20)
	d.Place(actor).RightOf(actorLbl, vspace+70)

	var lastLabel, last shape.Shape = actorLbl, actor
	add := func(txt string, shape Shape) {
		label := NewLabel(txt)
		d.Place(label, shape).Below(lastLabel, vspace)
		d.VAlignCenter(last, shape)
		d.HAlignCenter(label, shape)
		lastLabel = label
		last = shape
	}

	add("Arrow", NewArrow(240, 0, 300, 0))
	add("Circle", NewCircle(20))
	add("Component", NewComponent("Component"))
	lcomp := NewComponent("Component")
	lcomp.SetHref("https://gregoryv.github.io/draw")
	add("Component(linked)", lcomp)
	add("Cylinder", NewCylinder(30, 40))
	add("Database", NewDatabase("database"))
	add("Diamond", NewDiamond())
	add("Dot", NewDot())
	add("ExitDot", NewExitDot())

	add("Hexagon", NewHexagon("Hexagon", 80, 40, 20))

	add("Internet", NewInternet())
	add("Label", NewLabel("label-text"))

	add("Line", NewLine(240, 0, 300, 0))
	add("Note", NewNote("This describes\nsomething..."))

	rec := NewRecord("record")
	rec.Fields = []string{"fields"}
	rec.Methods = []string{"methods"}
	add("Record", rec)

	add("Rect", NewRect("a rectangle"))
	add("State", NewState("active"))
	add("Triangle", NewTriangle())

	d.SaveAs("allshapes.svg")
}
