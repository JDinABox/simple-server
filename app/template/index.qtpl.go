// Code generated by qtc from "index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/template/index.qtpl:1
package template

//line app/template/index.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/template/index.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/template/index.qtpl:2
type Index struct {
	Header   string
	BodyHtml string
}

//line app/template/index.qtpl:9
func (i *Index) StreamIndexTPL(qw422016 *qt422016.Writer) {
//line app/template/index.qtpl:9
	qw422016.N().S(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"><meta http-equiv="X-UA-Compatible" content="ie=edge">`)
//line app/template/index.qtpl:16
	qw422016.N().S(i.Header)
//line app/template/index.qtpl:16
	qw422016.N().S(`<title x-data x-text="$store.title">Index</title></head><body class="dark" x-data :class="{'dark': $store.theme.dark}"><div>`)
//line app/template/index.qtpl:21
	qw422016.N().S(i.BodyHtml)
//line app/template/index.qtpl:21
	qw422016.N().S(`</div></body></html>`)
//line app/template/index.qtpl:25
}

//line app/template/index.qtpl:25
func (i *Index) WriteIndexTPL(qq422016 qtio422016.Writer) {
//line app/template/index.qtpl:25
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/template/index.qtpl:25
	i.StreamIndexTPL(qw422016)
//line app/template/index.qtpl:25
	qt422016.ReleaseWriter(qw422016)
//line app/template/index.qtpl:25
}

//line app/template/index.qtpl:25
func (i *Index) IndexTPL() string {
//line app/template/index.qtpl:25
	qb422016 := qt422016.AcquireByteBuffer()
//line app/template/index.qtpl:25
	i.WriteIndexTPL(qb422016)
//line app/template/index.qtpl:25
	qs422016 := string(qb422016.B)
//line app/template/index.qtpl:25
	qt422016.ReleaseByteBuffer(qb422016)
//line app/template/index.qtpl:25
	return qs422016
//line app/template/index.qtpl:25
}
