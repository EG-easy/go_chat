package trace

import (
	"testing"
	"bytes"
)

func TestNew(t *testing.T){
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil{
		t.Error("Newからの戻り値がnilです")
	}else{
		tracer.Trace("こんにちは、traceパッケージ")
		if buf.String() != "こんにちは、traceパッケージ\n"{
			t.Errorf("%sという誤った文字列が表示されました。", buf.String())
		}
	}
}



