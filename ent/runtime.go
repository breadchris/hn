// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/breadchris/hn/ent/schema"
	"github.com/breadchris/hn/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescKarma is the schema descriptor for karma field.
	userDescKarma := userFields[2].Descriptor()
	// user.DefaultKarma holds the default value on creation for the karma field.
	user.DefaultKarma = userDescKarma.Default.(int)
}