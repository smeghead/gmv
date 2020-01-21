package option

import (
	"testing"
)

func Test_NewOption(t *testing.T) {
	actual := NewOption()
	t.Logf("pass %p", actual)
}
