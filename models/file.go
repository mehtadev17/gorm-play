package models

import "github.com/jinzhu/gorm"

type File struct {
	gorm.Model
	Name             string
	Processed        bool
	LastRowProcessed uint64
	WorkflowRequests []WorkflowRequest
}
