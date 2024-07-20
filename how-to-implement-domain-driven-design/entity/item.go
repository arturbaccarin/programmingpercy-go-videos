// Package entities hold all the entities that are shared across the application
package entity

import "github.com/google/uuid"

// Item is an entity that represents a item in all domains
type Item struct {
	// ID and the identifier of the entity
	ID          uuid.UUID
	Name        string
	Description string
}
