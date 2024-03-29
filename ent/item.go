// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/breadchris/hn/ent/item"
)

// Item is the model entity for the Item schema.
type Item struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted bool `json:"deleted,omitempty"`
	// Type holds the value of the "type" field.
	Type item.Type `json:"type,omitempty"`
	// By holds the value of the "by" field.
	By string `json:"by,omitempty"`
	// Text holds the value of the "text" field.
	Text *string `json:"text,omitempty"`
	// Dead holds the value of the "dead" field.
	Dead bool `json:"dead,omitempty"`
	// Parent holds the value of the "parent" field.
	Parent *int `json:"parent,omitempty"`
	// Poll holds the value of the "poll" field.
	Poll *int `json:"poll,omitempty"`
	// Kids holds the value of the "kids" field.
	Kids []int `json:"kids,omitempty"`
	// URL holds the value of the "url" field.
	URL *string `json:"url,omitempty"`
	// Score holds the value of the "score" field.
	Score int `json:"score,omitempty"`
	// Title holds the value of the "title" field.
	Title *string `json:"title,omitempty"`
	// Parts holds the value of the "parts" field.
	Parts []int `json:"parts,omitempty"`
	// Descendants holds the value of the "descendants" field.
	Descendants int `json:"descendants,omitempty"`
	// Time holds the value of the "time" field.
	Time int `json:"time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ItemQuery when eager-loading is set.
	Edges        ItemEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ItemEdges holds the relations/edges for other nodes in the graph.
type ItemEdges struct {
	// Children holds the value of the children edge.
	Children []*Item `json:"children,omitempty"`
	// Parents holds the value of the parents edge.
	Parents []*Item `json:"parents,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedChildren map[string][]*Item
	namedParents  map[string][]*Item
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e ItemEdges) ChildrenOrErr() ([]*Item, error) {
	if e.loadedTypes[0] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ParentsOrErr returns the Parents value or an error if the edge
// was not loaded in eager-loading.
func (e ItemEdges) ParentsOrErr() ([]*Item, error) {
	if e.loadedTypes[1] {
		return e.Parents, nil
	}
	return nil, &NotLoadedError{edge: "parents"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Item) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case item.FieldKids, item.FieldParts:
			values[i] = new([]byte)
		case item.FieldDeleted, item.FieldDead:
			values[i] = new(sql.NullBool)
		case item.FieldParent, item.FieldPoll, item.FieldScore, item.FieldDescendants, item.FieldTime:
			values[i] = new(sql.NullInt64)
		case item.FieldID, item.FieldType, item.FieldBy, item.FieldText, item.FieldURL, item.FieldTitle:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Item fields.
func (i *Item) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case item.FieldID:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value.Valid {
				i.ID = value.String
			}
		case item.FieldDeleted:
			if value, ok := values[j].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[j])
			} else if value.Valid {
				i.Deleted = value.Bool
			}
		case item.FieldType:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[j])
			} else if value.Valid {
				i.Type = item.Type(value.String)
			}
		case item.FieldBy:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field by", values[j])
			} else if value.Valid {
				i.By = value.String
			}
		case item.FieldText:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[j])
			} else if value.Valid {
				i.Text = new(string)
				*i.Text = value.String
			}
		case item.FieldDead:
			if value, ok := values[j].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field dead", values[j])
			} else if value.Valid {
				i.Dead = value.Bool
			}
		case item.FieldParent:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent", values[j])
			} else if value.Valid {
				i.Parent = new(int)
				*i.Parent = int(value.Int64)
			}
		case item.FieldPoll:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field poll", values[j])
			} else if value.Valid {
				i.Poll = new(int)
				*i.Poll = int(value.Int64)
			}
		case item.FieldKids:
			if value, ok := values[j].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field kids", values[j])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &i.Kids); err != nil {
					return fmt.Errorf("unmarshal field kids: %w", err)
				}
			}
		case item.FieldURL:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[j])
			} else if value.Valid {
				i.URL = new(string)
				*i.URL = value.String
			}
		case item.FieldScore:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field score", values[j])
			} else if value.Valid {
				i.Score = int(value.Int64)
			}
		case item.FieldTitle:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[j])
			} else if value.Valid {
				i.Title = new(string)
				*i.Title = value.String
			}
		case item.FieldParts:
			if value, ok := values[j].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field parts", values[j])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &i.Parts); err != nil {
					return fmt.Errorf("unmarshal field parts: %w", err)
				}
			}
		case item.FieldDescendants:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field descendants", values[j])
			} else if value.Valid {
				i.Descendants = int(value.Int64)
			}
		case item.FieldTime:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field time", values[j])
			} else if value.Valid {
				i.Time = int(value.Int64)
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Item.
// This includes values selected through modifiers, order, etc.
func (i *Item) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// QueryChildren queries the "children" edge of the Item entity.
func (i *Item) QueryChildren() *ItemQuery {
	return NewItemClient(i.config).QueryChildren(i)
}

// QueryParents queries the "parents" edge of the Item entity.
func (i *Item) QueryParents() *ItemQuery {
	return NewItemClient(i.config).QueryParents(i)
}

// Update returns a builder for updating this Item.
// Note that you need to call Item.Unwrap() before calling this method if this Item
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Item) Update() *ItemUpdateOne {
	return NewItemClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Item entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Item) Unwrap() *Item {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Item is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Item) String() string {
	var builder strings.Builder
	builder.WriteString("Item(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("deleted=")
	builder.WriteString(fmt.Sprintf("%v", i.Deleted))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", i.Type))
	builder.WriteString(", ")
	builder.WriteString("by=")
	builder.WriteString(i.By)
	builder.WriteString(", ")
	if v := i.Text; v != nil {
		builder.WriteString("text=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("dead=")
	builder.WriteString(fmt.Sprintf("%v", i.Dead))
	builder.WriteString(", ")
	if v := i.Parent; v != nil {
		builder.WriteString("parent=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := i.Poll; v != nil {
		builder.WriteString("poll=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("kids=")
	builder.WriteString(fmt.Sprintf("%v", i.Kids))
	builder.WriteString(", ")
	if v := i.URL; v != nil {
		builder.WriteString("url=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("score=")
	builder.WriteString(fmt.Sprintf("%v", i.Score))
	builder.WriteString(", ")
	if v := i.Title; v != nil {
		builder.WriteString("title=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("parts=")
	builder.WriteString(fmt.Sprintf("%v", i.Parts))
	builder.WriteString(", ")
	builder.WriteString("descendants=")
	builder.WriteString(fmt.Sprintf("%v", i.Descendants))
	builder.WriteString(", ")
	builder.WriteString("time=")
	builder.WriteString(fmt.Sprintf("%v", i.Time))
	builder.WriteByte(')')
	return builder.String()
}

// NamedChildren returns the Children named value or an error if the edge was not
// loaded in eager-loading with this name.
func (i *Item) NamedChildren(name string) ([]*Item, error) {
	if i.Edges.namedChildren == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := i.Edges.namedChildren[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (i *Item) appendNamedChildren(name string, edges ...*Item) {
	if i.Edges.namedChildren == nil {
		i.Edges.namedChildren = make(map[string][]*Item)
	}
	if len(edges) == 0 {
		i.Edges.namedChildren[name] = []*Item{}
	} else {
		i.Edges.namedChildren[name] = append(i.Edges.namedChildren[name], edges...)
	}
}

// NamedParents returns the Parents named value or an error if the edge was not
// loaded in eager-loading with this name.
func (i *Item) NamedParents(name string) ([]*Item, error) {
	if i.Edges.namedParents == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := i.Edges.namedParents[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (i *Item) appendNamedParents(name string, edges ...*Item) {
	if i.Edges.namedParents == nil {
		i.Edges.namedParents = make(map[string][]*Item)
	}
	if len(edges) == 0 {
		i.Edges.namedParents[name] = []*Item{}
	} else {
		i.Edges.namedParents[name] = append(i.Edges.namedParents[name], edges...)
	}
}

// Items is a parsable slice of Item.
type Items []*Item
