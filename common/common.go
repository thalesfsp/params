package common

import (
	"time"

	"github.com/thalesfsp/status"
)

// Common params.
type Common struct {
	// CreatedAt is the time the record was created.
	CreatedAt time.Time `bson:"createdAt" default:"now" form:"createdAt" json:"createdAt,omitempty" query:"createdAt" validate:"omitempty"`

	// CreatedBy is the user who created the record.
	CreatedBy string `bson:"createdBy" form:"createdBy" json:"createdBy,omitempty" query:"createdBy" validate:"omitempty,gt=0"`

	// DeleteAt is the time the record was deleted.
	DeleteAt time.Time `bson:"deletedAt" form:"deleteAt" json:"deleteAt,omitempty" query:"deleteAt" validate:"omitempty"`

	// DeleteBy is the user who deleted the record.
	DeleteBy string `bson:"deletedBy" form:"deleteBy" json:"deleteBy,omitempty" query:"deleteBy" validate:"omitempty,gt=0"`

	// ID is the unique identifier for the record.
	//
	// NOTE: the `id:"uuid"` tag automatically sets with an UUID ONLY if the
	// field is empty.
	ID string `bson:"_id" db:"id" dbType:"varchar(255)" form:"id" id:"uuid" json:"id,omitempty" param:"id" query:"id" validate:"omitempty,gt=0"`

	// Status is the status of the record.
	Status status.Status `bson:"status" default:"active" form:"status" json:"status,omitempty" query:"status" validate:"omitempty,gt=0"`

	// UpdatedAt is the time the record was updated.
	UpdatedAt time.Time `bson:"updatedAt" default:"now" form:"updatedAt" json:"updatedAt,omitempty" query:"updatedAt" validate:"omitempty"`

	// UpdatedBy is the user who updated the record.
	UpdatedBy string `bson:"updatedBy" form:"updatedBy" json:"updatedBy,omitempty" query:"updatedBy" validate:"omitempty,gt=0"`
}
