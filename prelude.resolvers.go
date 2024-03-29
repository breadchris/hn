package main

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/breadchris/hn/ent/item"
)

// Type is the resolver for the type field.
func (r *__FieldResolver) Type(ctx context.Context, obj *introspection.Field) (item.Type, error) {
	panic(fmt.Errorf("not implemented: Type - type"))
}

// Type is the resolver for the type field.
func (r *__InputValueResolver) Type(ctx context.Context, obj *introspection.InputValue) (item.Type, error) {
	panic(fmt.Errorf("not implemented: Type - type"))
}

// Types is the resolver for the types field.
func (r *__SchemaResolver) Types(ctx context.Context, obj *introspection.Schema) ([]item.Type, error) {
	panic(fmt.Errorf("not implemented: Types - types"))
}

// QueryType is the resolver for the queryType field.
func (r *__SchemaResolver) QueryType(ctx context.Context, obj *introspection.Schema) (item.Type, error) {
	panic(fmt.Errorf("not implemented: QueryType - queryType"))
}

// MutationType is the resolver for the mutationType field.
func (r *__SchemaResolver) MutationType(ctx context.Context, obj *introspection.Schema) (*item.Type, error) {
	panic(fmt.Errorf("not implemented: MutationType - mutationType"))
}

// SubscriptionType is the resolver for the subscriptionType field.
func (r *__SchemaResolver) SubscriptionType(ctx context.Context, obj *introspection.Schema) (*item.Type, error) {
	panic(fmt.Errorf("not implemented: SubscriptionType - subscriptionType"))
}

// Kind is the resolver for the kind field.
func (r *__TypeResolver) Kind(ctx context.Context, obj *item.Type) (string, error) {
	panic(fmt.Errorf("not implemented: Kind - kind"))
}

// Name is the resolver for the name field.
func (r *__TypeResolver) Name(ctx context.Context, obj *item.Type) (*string, error) {
	panic(fmt.Errorf("not implemented: Name - name"))
}

// Description is the resolver for the description field.
func (r *__TypeResolver) Description(ctx context.Context, obj *item.Type) (*string, error) {
	panic(fmt.Errorf("not implemented: Description - description"))
}

// Fields is the resolver for the fields field.
func (r *__TypeResolver) Fields(ctx context.Context, obj *item.Type, includeDeprecated *bool) ([]*introspection.Field, error) {
	panic(fmt.Errorf("not implemented: Fields - fields"))
}

// Interfaces is the resolver for the interfaces field.
func (r *__TypeResolver) Interfaces(ctx context.Context, obj *item.Type) ([]item.Type, error) {
	panic(fmt.Errorf("not implemented: Interfaces - interfaces"))
}

// PossibleTypes is the resolver for the possibleTypes field.
func (r *__TypeResolver) PossibleTypes(ctx context.Context, obj *item.Type) ([]item.Type, error) {
	panic(fmt.Errorf("not implemented: PossibleTypes - possibleTypes"))
}

// EnumValues is the resolver for the enumValues field.
func (r *__TypeResolver) EnumValues(ctx context.Context, obj *item.Type, includeDeprecated *bool) ([]*introspection.EnumValue, error) {
	panic(fmt.Errorf("not implemented: EnumValues - enumValues"))
}

// InputFields is the resolver for the inputFields field.
func (r *__TypeResolver) InputFields(ctx context.Context, obj *item.Type) ([]*introspection.InputValue, error) {
	panic(fmt.Errorf("not implemented: InputFields - inputFields"))
}

// OfType is the resolver for the ofType field.
func (r *__TypeResolver) OfType(ctx context.Context, obj *item.Type) (*item.Type, error) {
	panic(fmt.Errorf("not implemented: OfType - ofType"))
}

// SpecifiedByURL is the resolver for the specifiedByURL field.
func (r *__TypeResolver) SpecifiedByURL(ctx context.Context, obj *item.Type) (*string, error) {
	panic(fmt.Errorf("not implemented: SpecifiedByURL - specifiedByURL"))
}

// __Field returns __FieldResolver implementation.
func (r *Resolver) __Field() __FieldResolver { return __FieldResolver{r} }

// __InputValue returns __InputValueResolver implementation.
func (r *Resolver) __InputValue() __InputValueResolver { return __InputValueResolver{r} }

// __Schema returns __SchemaResolver implementation.
func (r *Resolver) __Schema() __SchemaResolver { return __SchemaResolver{r} }

// __Type returns __TypeResolver implementation.
func (r *Resolver) __Type() __TypeResolver { return __TypeResolver{r} }

type __FieldResolver struct{ *Resolver }
type __InputValueResolver struct{ *Resolver }
type __SchemaResolver struct{ *Resolver }
type __TypeResolver struct{ *Resolver }
