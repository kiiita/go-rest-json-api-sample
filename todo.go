package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Name      string    `json:"name"`
	Completed bool      `json:"complted"`
	Due       time.Time `json:"due"`
}

type Todos []Todo
