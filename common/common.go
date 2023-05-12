package common

import (
	"time"

	"github.com/thalesfsp/status"
)

// Common params.
type Common struct {
	// CreatedAt is the time the record was created.
	CreatedAt time.Time `bson:"createdAt" json:"createdAt,omitempty" form:"createdAt" query:"createdAt" validate:"omitempty" default:"now"`

	// CreatedBy is the user who created the record.
	CreatedBy string `bson:"createdBy" json:"createdBy,omitempty" form:"createdBy" query:"createdBy" validate:"omitempty,gt=0"`

	// DeleteAt is the time the record was deleted.
	DeleteAt time.Time `bson:"deleteAt" json:"deleteAt,omitempty" form:"deleteAt" query:"deleteAt" validate:"omitempty"`

	// DeleteBy is the user who deleted the record.
	DeleteBy string `bson:"deleteBy" json:"deleteBy,omitempty" form:"deleteBy" query:"deleteBy" validate:"omitempty,gt=0"`

	// ID is the unique identifier for the record.
	//
	// NOTE: the `id:"uuid"` tag automatically sets with an UUID ONLY if the
	// field is empty.
	ID string `bson:"_id" json:"id,omitempty" id:"uuid" form:"id" param:"id" query:"id" db:"id" dbType:"varchar(255)" validate:"omitempty,gt=0"`

	// Status is the status of the record.
	Status status.Status `bson:"status" json:"status" form:"status" query:"status" validate:"omitempty,gt=0" default:"active"`

	// UpdatedAt is the time the record was updated.
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt,omitempty" form:"updatedAt" query:"updatedAt" validate:"omitempty" default:"now"`

	// UpdatedBy is the user who updated the record.
	UpdatedBy string `bson:"updatedBy" json:"updatedBy,omitempty" form:"updatedBy" query:"updatedBy" validate:"omitempty,gt=0"`
}
