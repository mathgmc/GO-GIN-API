package models

type Student struct {
	Name     string `json: "name"`
	Document string `json: "document_number"`
	Email    string `json: "email"`
}

var Students []Student
