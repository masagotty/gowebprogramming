package main

type FakePost struct {
	Id      int
	Content string
	Author  string
}

func (post *Fakepost) fetch(id int) (err error) {
	post.Id = id
	return
}

func (post *Fakepost) create() (err error) {
	return
}

func (post *Fakepost) update() (err error) {
	return
}

func (post *Fakepost) delete() (err error) {
	return
}
