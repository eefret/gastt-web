package xml

import (
	"testing"

	"github.com/eefret/gastt-web/controllers/lang"
)

func TestBuilderNew(t *testing.T) {
	strs := make(map[string]string)
	strs["hi"] = "hola"
	strs["name"] = "nombre"
	builder, err := NewBuilder(strs)
	check(err)
	builder.BuildStr()
}

func TestBuilderBuildStr(t *testing.T) {
	strs := make(map[string]string)
	strs["hi"] = "hola"
	strs["name"] = "nombre"
	builder, err := NewBuilder(strs)
	check(err)
	builder.SetLang(lang.Languages.ES)

}
