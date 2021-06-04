package util

import (
	"github.com/jinzhu/gorm"
)

// descargaPedidos.Response representa el modelo para una peticion del api Descargarpedidos
type Response struct {
	Messages []Message `json:"messages,omitempty"`
	IsValid  bool      `json:"isValid" example:"true"`
}

// Message represents the model for an Message in the descargaPedidos.Response
type Message struct {
	Type    int    `json:"type,omitempty" example:"0"`
	Code    string `json:"code,omitempty" example:"00"`
	Message string `json:"message,omitempty" example:"success"`
}

type ErrorDBStruct struct {
	Severity         string
	Code             string
	Message          string
	Detail           string
	Hint             string
	Position         string
	InternalPosition string
	InternalQuery    string
	Where            string
	Schema           string
	Table            string
	Column           string
	DataTypeName     string
	Constraint       string
	File             string
	Line             string
	Routine          string
}

type RowInserted struct {
	Value interface{}
}

type Pagination struct {
	Limit  string
	OffSet string
	Data   string
}

type ConnStruct struct {
	Name string
	DB   gorm.DB
}

type MessageResult struct {
	Message string `json:"message,omitempty" example:"success"`
}

// // Package errors implements functions to manipulate errors.
// package errors

// // New returns an error that formats as the given text.
// func New(text string) error {
// 	return &errorString{text}
// }

// // errorString is a trivial implementation of error.
// type errorString struct {
// 	s string
// }

// func (e *errorString) Error() string {
// 	return e.s
// }
