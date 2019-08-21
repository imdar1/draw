package design

import "github.com/gregoryv/go-design/shape"

func ExampleClassDiagram() {
	diagram := NewDiagram()
	diagramRec := shape.NewRecordOf(diagram)
	record := shape.NewRecordOf(shape.Record{})
	adjuster := shape.NewRecordOf(shape.Adjuster{})

	diagram.Place(diagramRec).At(10, 30)
	diagram.Place(record).RightOf(diagramRec)
	diagram.Place(adjuster).RightOf(record)

	diagram.HAlignCenter(diagramRec, record, adjuster)
	diagram.SaveAs("img/class_example.svg")
}
