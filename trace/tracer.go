package trace

import (
	"io"
	"fmt"
)

//tracerはコード内での出来事を記録することができるオブジェクトを表すインターフェースです
type Tracer interface{
	Trace(...interface{})
}

type tracer struct{
	out io.Writer
}

func (t *tracer)Trace(a ...interface{}){
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

func New(w io.Writer)Tracer{
	return & tracer{out:w}
}


