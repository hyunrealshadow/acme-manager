package directive

import (
	"entgo.io/contrib/entgql"
	"github.com/vektah/gqlparser/v2/ast"
)

func Model(name string) entgql.Directive {
	return entgql.NewDirective("model",
		&ast.Argument{
			Name: "name",
			Value: &ast.Value{
				Raw:  name,
				Kind: ast.StringValue,
			},
		})
}
