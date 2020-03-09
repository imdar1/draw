package design

import (
	"io"
	"testing"

	"github.com/gregoryv/asserter"
)

type myOwn int
type myStr struct{ f string }

func Test_NewVRecord_types(t *testing.T) {
	ok := func(v interface{}, exp string) {
		vr := NewVRecord(v)
		got := vr.Title
		if got != exp {
			t.Error("got: ", got, "exp: ", exp)
		}
	}
	ok(myOwn(1), "design.myOwn int")
	ok(myStr{}, "design.myStr struct")
	ok((*io.Reader)(nil), "io.Reader interface")
}

func TestVRecord(t *testing.T) {
	r := NewStruct(VRecord{})
	before := len(r.Fields)
	r.TitleOnly()
	got := len(r.Fields)
	assert := asserter.New(t)
	assert(got != before).Error("Did not hide fields")
}

func TestNewStruct(t *testing.T) {
	x := struct {
		Field string
	}{}
	s := NewStruct(x)
	if len(s.Fields) != 1 {
		t.Error("Expected one field")
	}
}

func TestNewInterface(t *testing.T) {
	i := NewInterface((*io.Writer)(nil))
	if len(i.Methods) != 1 {
		t.Error("Expected one method")
	}
}

func mustCatchPanic(t asserter.T) {
	t.Helper()
	e := recover()
	if e == nil {
		t.Error("should panic")
	}
}

type C struct{}

func TestVRecord_ComposedOf(t *testing.T) {
	ok := func(a, b interface{}) {
		t.Helper()
		A := NewStruct(a)
		B := NewStruct(b)
		if !A.ComposedOf(&B) {
			t.Fail()
		}
	}
	ok(struct{ c C }{}, C{})

	bad := func(a, b interface{}) {
		t.Helper()
		A := NewStruct(a)
		B := NewStruct(b)
		if A.ComposedOf(&B) {
			t.Fail()
		}
	}
	bad(struct{ c *C }{}, C{})
}

func TestVRecord_Aggregates(t *testing.T) {
	ok := func(a, b interface{}) {
		t.Helper()
		A := NewStruct(a)
		B := NewStruct(b)
		if !A.Aggregates(&B) {
			t.Fail()
		}
	}
	ok(struct{ c *C }{}, C{})

	bad := func(a, b interface{}) {
		t.Helper()
		A := NewStruct(a)
		B := NewStruct(b)
		if A.Aggregates(&B) {
			t.Fail()
		}
	}
	bad(struct{ c C }{}, C{})
}
