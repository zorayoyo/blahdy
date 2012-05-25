package main

type Blah struct {
	Id string
	Text string
	AuthorId string
	CreateTime int64
	UpdateTime int64
}

func (b * Blah) BuildFromJson(json interface{}) {
	var jsonMap map[string]interface{}
	jsonMap = json.(map[string]interface{})
	for k, v := range jsonMap {
		switch vv := v.(type) {
		case string:
			switch k {
			case "Id":
				b.Id = vv
			case "Text":
				b.Text = vv
			case "AuthorId":
				b.AuthorId = vv
			}
		case float64:
			if k == "CreateTime" {
				b.CreateTime = int64(vv)
			} else if k == "UpdateTime" {
				b.UpdateTime = int64(vv)
			}
		}
	}
}

type Message struct {
	Id string
	Action int
	Text string
	AuthorId string
	BlahId string
	CreateTime int64
}

type MessageIndex struct {
	Id string
	MessageIds []string
}

func (m * Message) BuildFromJson(json interface{}) {
	var jsonMap map[string]interface{}
	jsonMap = json.(map[string]interface{})
	for k, v := range jsonMap {
		switch vv := v.(type) {
		case string:
			switch k {
			case "Id":
				m.Id = vv
			case "Text":
				m.Text = vv
			case "AuthorId":
				m.AuthorId = vv
			}
		case float64:
			if k == "CreateTime" {
				m.CreateTime = int64(vv)
			} else if k == "Action" {
				m.Action = int(vv)
			}
		}
	}
}

type User struct {
	Id string
	Name string
	Email string
	Phone string
	Bio string
}

func (u * User) BuildFromJson(json interface{}) {
	var jsonMap map[string]interface{}
	jsonMap = json.(map[string]interface{})
	for k, v := range jsonMap {
		switch vv := v.(type) {
		case string:
			switch k {
			case "Id":
				u.Id = vv
			case "Name":
				u.Name = vv
			case "Email":
				u.Email = vv
			case "Phone":
				u.Phone = vv
			case "Bio":
				u.Bio = vv
			}
		}
	}
}

type UserShadow struct {
	Certification string
	CertificationAlt1 string
	CertificationAlt2 string
}

func (u * UserShadow) BuildFromJson(json interface{}) {
	var jsonMap map[string]interface{}
	jsonMap = json.(map[string]interface{})
	for k, v := range jsonMap {
		switch vv := v.(type) {
		case string:
			switch k {
			case "Certification":
				u.Certification = vv
			case "CertificationAlt1":
				u.CertificationAlt1 = vv
			case "CertificationAlt2":
				u.CertificationAlt2 = vv
			}
		}
	}
}

