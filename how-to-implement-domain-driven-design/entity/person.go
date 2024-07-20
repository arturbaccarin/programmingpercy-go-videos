// Package entities hold all the entities that are shared across the application
package entity

import "github.com/google/uuid"

// Person is an entity that represents a person in all domains
type Person struct {
	// ID and the identifier of the entity
	ID   uuid.UUID
	Name string
	Age  int
}
