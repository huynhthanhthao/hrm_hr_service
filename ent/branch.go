// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/huynhthanhthao/hrm_hr_service/ent/branch"
	"github.com/huynhthanhthao/hrm_hr_service/ent/company"
)

// Branch is the model entity for the Branch schema.
type Branch struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// ContactInfo holds the value of the "contact_info" field.
	ContactInfo string `json:"contact_info,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BranchQuery when eager-loading is set.
	Edges            BranchEdges `json:"edges"`
	company_branches *uuid.UUID
	selectValues     sql.SelectValues
}

// BranchEdges holds the relations/edges for other nodes in the graph.
type BranchEdges struct {
	// Company holds the value of the company edge.
	Company *Company `json:"company,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BranchEdges) CompanyOrErr() (*Company, error) {
	if e.Company != nil {
		return e.Company, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: company.Label}
	}
	return nil, &NotLoadedError{edge: "company"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Branch) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case branch.FieldName, branch.FieldCode, branch.FieldAddress, branch.FieldContactInfo:
			values[i] = new(sql.NullString)
		case branch.FieldCreatedAt, branch.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case branch.FieldID:
			values[i] = new(uuid.UUID)
		case branch.ForeignKeys[0]: // company_branches
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Branch fields.
func (b *Branch) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case branch.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				b.ID = *value
			}
		case branch.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				b.Name = value.String
			}
		case branch.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				b.Code = value.String
			}
		case branch.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				b.Address = value.String
			}
		case branch.FieldContactInfo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contact_info", values[i])
			} else if value.Valid {
				b.ContactInfo = value.String
			}
		case branch.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case branch.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case branch.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field company_branches", values[i])
			} else if value.Valid {
				b.company_branches = new(uuid.UUID)
				*b.company_branches = *value.S.(*uuid.UUID)
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Branch.
// This includes values selected through modifiers, order, etc.
func (b *Branch) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryCompany queries the "company" edge of the Branch entity.
func (b *Branch) QueryCompany() *CompanyQuery {
	return NewBranchClient(b.config).QueryCompany(b)
}

// Update returns a builder for updating this Branch.
// Note that you need to call Branch.Unwrap() before calling this method if this Branch
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Branch) Update() *BranchUpdateOne {
	return NewBranchClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Branch entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Branch) Unwrap() *Branch {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Branch is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Branch) String() string {
	var builder strings.Builder
	builder.WriteString("Branch(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("name=")
	builder.WriteString(b.Name)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(b.Code)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(b.Address)
	builder.WriteString(", ")
	builder.WriteString("contact_info=")
	builder.WriteString(b.ContactInfo)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Branches is a parsable slice of Branch.
type Branches []*Branch
