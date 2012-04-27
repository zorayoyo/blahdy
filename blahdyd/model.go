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

