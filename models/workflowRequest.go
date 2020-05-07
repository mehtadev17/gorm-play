package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type WorkflowRequest struct {
	gorm.Model
	WorkflowName     string
	WorkflowID       sql.NullInt64
	FileID           uint `sql:"type:integer REFERENCES files(id)"`
}
