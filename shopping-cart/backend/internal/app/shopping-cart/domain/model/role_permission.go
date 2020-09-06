package model

// RolePermission define the permission role model which is the mappings of the permissions to the roles
type RolePermission struct {
	// The role id to identify the role
	RoleID uint	`gorm:"type:bigint;primaryKey;auto_increment:false" validate:"required" json:"role_id"`

	// The permission id to identify the permission
	PermissionID uint `gorm:"type:bigint;primaryKey;auto_increment:false" validate:"required" json:"permission_id"`

	// The date and time at which the mapping is created
	//CreatedAt time.Time `gorm:"type:datetime;not null" json:"created_at"`

	// The date and time at which the mapping is last updated
	//UpdatedAt sql.NullTime `gorm:"type:datetime" json:"updated_at"`
}
