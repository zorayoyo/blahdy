package main

type T_BLAH struct {
	Blah;
	Author *User;
}

func (t * T_BLAH) Build(blah * Blah, user * User) {
	t.Id = blah.Id
	t.AuthorId = blah.AuthorId
	t.Text = blah.Text
	t.CreateTime = blah.CreateTime
	t.UpdateTime = blah.UpdateTime
	t.Author = user
}

