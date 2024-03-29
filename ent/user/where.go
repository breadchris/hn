// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"github.com/breadchris/hn/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldID, id))
}

// Created applies equality check predicate on the "created" field. It's identical to CreatedEQ.
func Created(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreated, v))
}

// Karma applies equality check predicate on the "karma" field. It's identical to KarmaEQ.
func Karma(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldKarma, v))
}

// About applies equality check predicate on the "about" field. It's identical to AboutEQ.
func About(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAbout, v))
}

// CreatedEQ applies the EQ predicate on the "created" field.
func CreatedEQ(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreated, v))
}

// CreatedNEQ applies the NEQ predicate on the "created" field.
func CreatedNEQ(v int64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreated, v))
}

// CreatedIn applies the In predicate on the "created" field.
func CreatedIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreated, vs...))
}

// CreatedNotIn applies the NotIn predicate on the "created" field.
func CreatedNotIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreated, vs...))
}

// CreatedGT applies the GT predicate on the "created" field.
func CreatedGT(v int64) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreated, v))
}

// CreatedGTE applies the GTE predicate on the "created" field.
func CreatedGTE(v int64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreated, v))
}

// CreatedLT applies the LT predicate on the "created" field.
func CreatedLT(v int64) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreated, v))
}

// CreatedLTE applies the LTE predicate on the "created" field.
func CreatedLTE(v int64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreated, v))
}

// CreatedIsNil applies the IsNil predicate on the "created" field.
func CreatedIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldCreated))
}

// CreatedNotNil applies the NotNil predicate on the "created" field.
func CreatedNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldCreated))
}

// KarmaEQ applies the EQ predicate on the "karma" field.
func KarmaEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldKarma, v))
}

// KarmaNEQ applies the NEQ predicate on the "karma" field.
func KarmaNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldKarma, v))
}

// KarmaIn applies the In predicate on the "karma" field.
func KarmaIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldKarma, vs...))
}

// KarmaNotIn applies the NotIn predicate on the "karma" field.
func KarmaNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldKarma, vs...))
}

// KarmaGT applies the GT predicate on the "karma" field.
func KarmaGT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldKarma, v))
}

// KarmaGTE applies the GTE predicate on the "karma" field.
func KarmaGTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldKarma, v))
}

// KarmaLT applies the LT predicate on the "karma" field.
func KarmaLT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldKarma, v))
}

// KarmaLTE applies the LTE predicate on the "karma" field.
func KarmaLTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldKarma, v))
}

// AboutEQ applies the EQ predicate on the "about" field.
func AboutEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAbout, v))
}

// AboutNEQ applies the NEQ predicate on the "about" field.
func AboutNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAbout, v))
}

// AboutIn applies the In predicate on the "about" field.
func AboutIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldAbout, vs...))
}

// AboutNotIn applies the NotIn predicate on the "about" field.
func AboutNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAbout, vs...))
}

// AboutGT applies the GT predicate on the "about" field.
func AboutGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldAbout, v))
}

// AboutGTE applies the GTE predicate on the "about" field.
func AboutGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAbout, v))
}

// AboutLT applies the LT predicate on the "about" field.
func AboutLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldAbout, v))
}

// AboutLTE applies the LTE predicate on the "about" field.
func AboutLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAbout, v))
}

// AboutContains applies the Contains predicate on the "about" field.
func AboutContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldAbout, v))
}

// AboutHasPrefix applies the HasPrefix predicate on the "about" field.
func AboutHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldAbout, v))
}

// AboutHasSuffix applies the HasSuffix predicate on the "about" field.
func AboutHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldAbout, v))
}

// AboutIsNil applies the IsNil predicate on the "about" field.
func AboutIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldAbout))
}

// AboutNotNil applies the NotNil predicate on the "about" field.
func AboutNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldAbout))
}

// AboutEqualFold applies the EqualFold predicate on the "about" field.
func AboutEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldAbout, v))
}

// AboutContainsFold applies the ContainsFold predicate on the "about" field.
func AboutContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldAbout, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
