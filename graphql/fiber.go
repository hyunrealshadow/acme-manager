package graphql

import (
	"acme-manager/config"
	"acme-manager/database"
	"acme-manager/graphql/model"
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	set "github.com/deckarep/golang-set/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"github.com/vektah/gqlparser/v2/ast"
	"net/http"
)

type FieldInterceptor struct{}

func wrapHandler(f func(http.ResponseWriter, *http.Request)) func(ctx *fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		model.SetGlobalIdContext(ctx)
		fasthttpadaptor.NewFastHTTPHandler(http.HandlerFunc(f))(ctx.Context())
	}
}

var excludeDirectives = set.NewSet[string]("goField", "goModel", "goTag", "model")

func removeExcludedDirectives(ec graphql.ExecutableSchema) {
	schema := ec.Schema()
	directives := schema.Directives
	newDirectives := make(map[string]*ast.DirectiveDefinition, len(directives)-excludeDirectives.Cardinality())
	for name, directive := range directives {
		if !excludeDirectives.Contains(name) {
			newDirectives[name] = directive
		}
	}
	schema.Directives = newDirectives
}

func MapGraphQLRoutes(app *fiber.App) {
	cfg := config.Get()
	ec := newSchema(database.Client)
	removeExcludedDirectives(ec)
	server := handler.NewDefaultServer(ec)
	server.Use(entgql.Transactioner{TxOpener: database.Client})
	playgroundHandler := playground.Handler("GraphiQL", "/graphql")
	app.Get("/graphql", func(ctx *fiber.Ctx) error {
		wrapHandler(server.ServeHTTP)(ctx)
		return nil
	})
	app.Post("/graphql", func(ctx *fiber.Ctx) error {
		wrapHandler(server.ServeHTTP)(ctx)
		return nil
	})
	if cfg.Server.Playground {
		app.Get("/", func(ctx *fiber.Ctx) error {
			wrapHandler(playgroundHandler)(ctx)
			return nil
		})
	}
}
