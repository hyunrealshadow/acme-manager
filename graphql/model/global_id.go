package model

import (
	"context"
	"encoding/base64"
	"github.com/99designs/gqlgen/graphql"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"io"
	"strings"
)

var GlobalIdTypeMap = map[string]string{
	"AcmeServer":  "acme_server",
	"AcmeAccount": "acme_account",
	"DnsProvider": "dns_provider",
	"Certificate": "certificate",
	"User":        "user",
}

func getModel(field graphql.CollectedField) *string {
	directive := field.Field.Definition.Directives.ForName("model")
	if directive == nil {
		return nil
	}
	model := directive.Arguments.ForName("name")
	if model == nil {
		return nil
	}
	return &model.Value.Raw
}

// MarshalUUID marshals a UUID to a global id
func MarshalUUID(id uuid.UUID) graphql.ContextMarshaler {
	return graphql.ContextWriterFunc(func(ctx context.Context, writer io.Writer) error {
		fieldContext := graphql.GetFieldContext(ctx)
		if fieldContext == nil {
			return errors.Errorf("no field context found")
		}
		objectType := getModel(fieldContext.Field)
		if objectType == nil {
			return errors.Errorf("unknown object type")
		}
		part1 := []byte(*objectType + ":")
		part2 := id[:]
		globalId := base64.StdEncoding.EncodeToString(append(part1, part2...))
		graphql.MarshalString(globalId).MarshalGQL(writer)
		return nil
	})
}

func globalIdError(id string) error {
	return errors.Errorf("invalid global id: %s", id)
}

type key string

const globalIdCtx key = "global_id_context"
const globalIdModelCtx key = "global_id_model_context"

func SetGlobalIdContext(ctx *fiber.Ctx) {
	globalIdTableMap := make(map[uuid.UUID]string)
	globalIdModelMap := make(map[string]string)
	ctx.Context().SetUserValue(globalIdCtx, globalIdTableMap)
	ctx.Context().SetUserValue(globalIdModelCtx, globalIdModelMap)
}

func GetGlobalIdContext(ctx context.Context) map[uuid.UUID]string {
	if val, ok := ctx.Value(globalIdCtx).(map[uuid.UUID]string); ok {
		return val
	}
	return nil
}

func getGlobalIdModelContext(ctx context.Context) map[string]string {
	if val, ok := ctx.Value(globalIdModelCtx).(map[string]string); ok {
		return val
	}
	return nil
}

func getParamModel(ctx context.Context, pathContext graphql.PathContext) *string {
	path := pathContext.Path().String()
	globalIdModelMap := getGlobalIdModelContext(ctx)
	if globalIdModelMap == nil {
		return nil
	}
	if model, ok := globalIdModelMap[path]; ok {
		return &model
	}
	return nil
}

// UnmarshalUUID unmarshal a global id to a UUID
func UnmarshalUUID(ctx context.Context, v interface{}) (uuid.UUID, error) {
	pathContext := graphql.GetPathContext(ctx)
	if pathContext == nil {
		return uuid.Nil, errors.Errorf("no path context found")
	}
	var id string
	switch v := v.(type) {
	case string:
		id = v
	default:
		return uuid.Nil, errors.Errorf("unsupport type %T", v)
	}
	decoded, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		return uuid.Nil, globalIdError(id)
	}
	if len(decoded) < 16 {
		return uuid.Nil, globalIdError(id)
	}

	decodedLen := len(decoded)
	head := string(decoded[0 : decodedLen-16])
	parts := strings.Split(head, ":")
	if len(parts) != 2 {
		return uuid.Nil, globalIdError(id)
	}

	tableName, ok := GlobalIdTypeMap[parts[0]]
	if !ok {
		return uuid.Nil, globalIdError(id)
	}
	model := getParamModel(ctx, *pathContext)
	if model != nil && *model != parts[0] {
		return uuid.Nil, errors.Errorf("invalid object type: %s, expected: %s", parts[0], *model)
	}

	objectID, err := uuid.FromBytes(decoded[decodedLen-16:])
	if err != nil {
		return uuid.Nil, err
	}

	GetGlobalIdContext(ctx)[objectID] = tableName

	return objectID, nil
}

func ModelDirective(ctx context.Context, obj interface{}, next graphql.Resolver, name *string) (interface{}, error) {
	pathCtx := graphql.GetPathContext(ctx)
	globalIdCtx := getGlobalIdModelContext(ctx)
	if pathCtx != nil && globalIdCtx != nil {
		path := pathCtx.Path().String()
		globalIdCtx[path] = *name
	}
	return next(ctx)
}
