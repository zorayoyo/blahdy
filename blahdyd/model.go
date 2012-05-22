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

