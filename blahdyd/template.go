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

type T_MESSAGE struct {
	Id string
	Action int
	Text string
	BlahId string
	CreateTime int64
	Author * User
}

func (t * T_MESSAGE) Build(msg * Message, user * User) {
	t.Author = user
	t.Id = msg.Id
	t.BlahId = msg.BlahId
	t.Text = msg.Text
	t.Action = msg.Action
	t.CreateTime = msg.CreateTime
}
