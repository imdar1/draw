package shape

import (
	"io/ioutil"
	"testing"

	"github.com/gregoryv/go-design/style"
)

func BenchmarkWriteSvg(b *testing.B) {
	svg := &Svg{
		Content: []Shape{
			&Record{},
		},
	}
	style := style.NewStyler(ioutil.Discard)
	for i := 0; i < b.N; i++ {
		svg.WriteSvg(style)
	}
	b.StopTimer()
	b.ReportAllocs()
}
