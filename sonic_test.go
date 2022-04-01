package magic

import (
	"testing"

	"github.com/bytedance/sonic"
	"github.com/gookit/goutil/dump"
)

func TestUnmarshalToMap(t *testing.T) {
	js := []byte(`{"a":1, "b":"c", "d": null}`)

	var data map[string]any
	if	err := sonic.Unmarshal(js, &data); err != nil {
		t.Fatal(err)
	}

	dump.P(data)
}


