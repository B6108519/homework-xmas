package schema

import "github.com/facebookincubator/ent"

// Classroom holds the schema definition for the Classroom entity.
type Classroom struct {
	ent.Schema
}

// Fields of the Classroom.
func (Classroom) Fields() []ent.Field {
	return nil
}

// Edges of the Classroom.
func (Classroom) Edges() []ent.Edge {
	return nil
}
