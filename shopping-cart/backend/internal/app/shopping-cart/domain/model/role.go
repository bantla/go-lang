// Package model define domain models that has wide enterprise business rules and
// can be a set of data structures and functions
package model

import (
	"database/sql"
	"time"
)

// Role defines the role model
type Role struct {
	// The unique id to identify the role
	ID uint `gorm:"primaryKey" validate:"required" json:"id"`

	// The role title
	Title string `gorm:"type:varchar(100);unique;not null" validate:"required" json:"title"`

	// The unique slug of search the role
	Slug string `gorm:"type:varchar(100);unique;not null" validate:"required" json:"slug"`

	// The description to mention the role
	Description sql.NullString `gorm:"type:tinytext" json:"description"`

	// The flag to check whether the role is currently active
	Active bool `gorm:"type:tinyint(1);default false" json:"active"`

	// The date and time at which the role is created
	CreatedAT time.Time `gorm:"type:datetime;not null" json:"created_at"`

	// The date and time at which the role is last updated
	UpdatedAt sql.NullTime `gorm:"type:datetime" json:"updated_at"`

	// The complete details about the role
	Content sql.NullString `gorm:"type:text" json:"content"`

	// The permission list belong to the role
	Permissions []*Permission `gorm:"many2many:role_permissions"`
}
