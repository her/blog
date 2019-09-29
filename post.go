package main

import "github.com/jinzhu/gorm"

// Post model
type Post struct {
	gorm.Model
	Title   string
	Author  string
	Content string `gorm:"type:varchar(255)"`
}
