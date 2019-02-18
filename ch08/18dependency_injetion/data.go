package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Post型のダブルを作りたい->Post型と同じメソッドを持つFakePost型を使いたい->インターフェースを使う！
type Text interface {
	fetch(id int) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
}

type Post struct {
	Db      *sql.DB
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// 以下でTextが持つべきメソッドを実装しているので、Post型はインターフェースTextの実装となる！
// Get a single post
func (post *Post) fetch(id int) (err error) {
	err = post.Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

// Ceate a new post
func (post *Post) create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := post.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

// Update a post
func (post *Post) update() (err error) {
	_, err = post.Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

// Delete a post
func (post *Post) delete() (err error) {
	_, err = post.Db.Exec("delete from posts where id = $1", post.Id)
	return
}
