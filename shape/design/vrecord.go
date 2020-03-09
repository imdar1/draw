package design

import (
	"fmt"
	"reflect"

	"github.com/gregoryv/draw/shape"
)

func NewVRecord(v interface{}) *VRecord {
	t := reflect.TypeOf(v)
	title := fmt.Sprintf("%s %s", t, t.Kind())
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		if t.Kind() == reflect.Interface {
			title = fmt.Sprintf("%s %s", t, t.Kind())
		}
	}
	rec := shape.NewRecord(title)
	// todo add methods and fields if any
	return &VRecord{
		Record: rec,
		t:      t,
	}
}

// VRecord represents a type struct or interface as a record shape.
type VRecord struct {
	*shape.Record
	t reflect.Type
}

// NewStruct returns a VRecord of the given object, panics if not
// struct.
func NewStruct(obj interface{}) VRecord {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		panic(fmt.Sprintf("Expected struct kind got %v", t.Kind()))
	}
	return VRecord{
		Record: shape.NewStructRecord(obj),
		t:      t,
	}
}

// TitleOnly hides fields and methods.
func (vr *VRecord) TitleOnly() {
	vr.HideFields()
	vr.HideMethods()
}

// NewInterface returns a VRecord of the given object, panics if not
// interface.
func NewInterface(obj interface{}) VRecord {
	t := reflect.TypeOf(obj).Elem()
	if t.Kind() != reflect.Interface {
		panic(fmt.Sprintf("Expected ptr kind got %v", t.Kind()))
	}
	return VRecord{
		Record: shape.NewInterfaceRecord(obj),
		t:      t,
	}
}

func (vr *VRecord) Implements(iface *VRecord) bool {
	return reflect.PtrTo(vr.t).Implements(iface.t)
}

func (vr *VRecord) ComposedOf(d *VRecord) bool {
	for i := 0; i < vr.t.NumField(); i++ {
		field := vr.t.Field(i)
		if field.Type == d.t {
			return true
		}
	}
	return false
}

func (vr *VRecord) Aggregates(d *VRecord) bool {
	for i := 0; i < vr.t.NumField(); i++ {
		field := vr.t.Field(i)
		if field.Type == reflect.PtrTo(d.t) {
			return true
		}
	}
	return false
}
