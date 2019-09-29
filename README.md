# A playground blog written in Go

A super simple API ontop of a postgres database with a single model (Posts).
I wanted to play around with Golang and see what it felt like writing a basic CRUD-y service.
This is not ready for production (and may never be) and is just an example application.


### Requirements
 * Postgres with database `blog`
 * Go

TODO:
  * Make a frontend!
  * Write some tests 😅

#routes

## Index
`GET /posts`

## Show
`GET /posts/:id`

## Create
`POST /posts`

## Update
`PUT /posts/:id`

## Delete
`DELETE /posts/:id`

