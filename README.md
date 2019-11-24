# A playground blog written in Go

A small API written in Go, using Postgres, with a single model (Posts).
I wanted to play around with Golang and see what it felt like writing a CRUD-y service.
This is not ready for production (and may never be) and is just an example application.

### Requirements
 * Postgres with database `blog`
 * Go
 * Some Go packages listed below

TODO:
  * Make a frontend!
  * Write some tests 😅

## Installation and Setup
You need these packages:

```
go get -u github.com/gorilla/handlers
go get -u github.com/gorilla/mux
go get -u github.com/jinzhu/gorm
go get -u github.com/joho/godotenv
go get -u github.com/lib/pq
```

# Posts API

A post is a single entity. Posts contain the contents of a blog entry, a title, and an author.

## Posts Object
```
{
  "ID": 1,
  "Title": "A Blog Post",
  "Author": "Melanie Berkley",
  "Content": "Lorem ipsum dolor sit amet..."
}
```

### Index
`GET /posts`

* **Path Parameters**
  * _None_
* **Query Parameters**
  * _None_
* **Body Parameters**
  * _None_

**Response**
```
[
  {<Post>}
]
```

### Show
`GET /posts/:id`

Fetches the record for a specific post.

* **Path Parameters**
  * **ID** `[Integer]` The ID of the post
* **Query Parameters**
  * _None_
* **Body Parameters**
  * _None_

**Response**
```
{<Post>}
```

### Create
`POST /posts`

### Update
`PUT /posts/:id`

### Delete
`DELETE /posts/:id`


#### Notes
**Create a post**
```
curl localhost:8081/posts \
  -X POST \
  -H "Content-type: application/json" \
  -d '{"title": "A Blog Post", "author": "Melanie", "content": "May all beings be happy, healthy, and love with ease"}'
```

**Get that post**
```
curl localhost:8081/posts/1
```

**Edit that post**
```
curl localhost:8081/posts/1 \
  -X PUT \
  -H "Content-type: application/json" \
  -d '{"title": "A Post", "author": "Melanie", "content": "May all beings be happy, healthy, and love with ease"}'
```

**Delete that post**
```
curl -X DELETE localhost:8081/posts/1
```
