package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type post struct {
	Title   string
	Author  string
	Content string
}

func logs(loc string, r *http.Request) string {
	return fmt.Sprintf("%s %s %s %s", r.Proto, r.Method, r.Host, loc)
}

// Index GET /posts
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(logs("/posts", r), http.StatusOK)

	// TODO Add pagination
	var posts []Post
	db.Find(&posts)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&posts)
}

// Show GET /posts/{id}
func Show(w http.ResponseWriter, r *http.Request) {
	fmt.Println(logs("/posts/{id}", r), http.StatusOK)

	params := mux.Vars(r)
	id := params["id"]
	var post []Post
	db.Debug().Where("id = ?", id).Find(&post)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// Create POST /posts
func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println(logs("/posts", r), http.StatusOK)

	decoder := json.NewDecoder(r.Body)
	var post post
	err := decoder.Decode(&post)
	if err != nil {
		fmt.Println(err)
	}
	db.Debug().Create(&Post{Title: post.Title, Author: post.Author, Content: post.Content})
	json.NewEncoder(w).Encode(post)
}

// Update PUT /posts/{id}
func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println(logs("/posts/{id}", r), http.StatusOK)

	params := mux.Vars(r)
	id := params["id"]

	decoder := json.NewDecoder(r.Body)
	var bodyParams post
	err := decoder.Decode(&bodyParams)
	if err != nil {
		fmt.Println(err)
	}

	var p Post
	db.Debug().Where("id = ?", id).Find(&p)
	p.Title = bodyParams.Title
	p.Author = bodyParams.Author
	p.Content = bodyParams.Content
	db.Save(&p)
}

// Delete DELETE /posts/{id}
func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println(logs("/posts/{id}", r), http.StatusOK)

	params := mux.Vars(r)
	id := params["id"]
	var post Post
	db.Debug().Unscoped().Where("id = ?", id).Delete(&post)
}
