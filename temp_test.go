package fw

import "testing"

type Foo struct {
	bar map[string]string
}

func TestTemp(t *testing.T) {
	f := new(Foo)
	f.bar["a"] = "b"
}
