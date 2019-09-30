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
```
curl -X GET localhost:8081/posts/14
```

```
curl -X POST -H "Content-type: application/json" -d '{"title": "A Blog Post", "author": "Melanie", "content": "May all beings be happy, healthy, and love with ease"}' localhost:8081/posts

```

curl -X PUT -H "Content-type: application/json" -d '{"title": "Update Post", "author": "Melanie", "content": "May all beings be happy, healthy, and love with ease"}' localhost:8081/posts
```

```
curl -X DELETE localhost:8081/posts/15
```
