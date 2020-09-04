package model

import (
	"database/sql"
	"time"
)

// Permission define permission model
type Permission struct {
	// The unique id to identify the permission
	ID uint `gorm:"primaryKey" validate:"required" json:"id"`

	// The permission title
	Title string `gorm:"type:varchar(100);unique;not null" validate:"required" json:"title"`

	// The unique slug of search the permission
	Slug string `gorm:"type:varchar(100);unique;not null" validate:"required" json:"slug"`

	// The description to mention the permission
	Description sql.NullString `gorm:"type:tinytext" json:"description"`

	// The flag to check whether the permission is currently active
	Active bool `gorm:"type:tinyint(1);default false" json:"active"`

	// The date and time at which the permission is created
	CreatedAT time.Time `gorm:"type:datetime;not null" json:"created_at"`

	// The date and time at which the permission is last updated
	UpdatedAt sql.NullTime `gorm:"type:datetime" json:"updated_at"`

	// The complete details about the permission
	Content sql.NullString `gorm:"type:text" json:"content"`

	// The role list belongs to the permission
	Roles []*Role `gorm:"many2many:role_permissions"`
}
