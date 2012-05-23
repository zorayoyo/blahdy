package main

type T_BLAH struct {
	Id string
	Text string
	CreateTime int64
	UpdateTime int64
	Author * User
}

func (t * T_BLAH) Build(blah * Blah, user * User) {
	t.Author = user
	t.Id = blah.Id
	t.Text = blah.Text
	t.CreateTime = blah.CreateTime
	t.UpdateTime = blah.UpdateTime
}

