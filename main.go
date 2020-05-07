package main

import (
	"encoding/json"
	"fmt"

	"github.com/mehtadev17/gorm-play/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=docker sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&models.File{}, &models.WorkflowRequest{})
	file := &models.File{Name: "File1"}
	db.Create(file)
	db.Create(&models.WorkflowRequest{
		WorkflowName:     "WORKFLOW",
		FileID:           file.ID,
	})
	files := []models.File{}
	db.Find(&files)
	filesJSON, err := json.Marshal(files)
	fmt.Println(string(filesJSON))

	workflowRequests := []models.WorkflowRequest{}
	db.Find(&workflowRequests)
	workflowRequestsJSON, err := json.Marshal(workflowRequests)
	fmt.Println(string(workflowRequestsJSON))

}
