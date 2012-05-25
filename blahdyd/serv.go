package main

import (
	"fmt"
	"strings"
	"time"
	"github.com/shellex/tattoo/webapp"
)

func Auth(ctx * webapp.Context) bool {
	username := ctx.Request.Header.Get("X-BLAHDY-NAME")
	token := ctx.Request.Header.Get("X-BLAHDY-TOKEN")
	return BlahdyDB.AuthUser(username, token)
}

// Root Handler.
func HandleRoot(ctx *webapp.Context) {
	ctx.Info.UseGZip = strings.Index(ctx.Request.Header.Get("Accept-Encoding"), "gzip") > -1
	ctx.Info.StartTime = time.Now()
	urlPath := ctx.Request.URL.Path
	pathLevels := strings.Split(strings.Trim(urlPath, "/"), "/")
	pathLevels = pathLevels[1:]
	// fmt.Printf("%v\n", pathLevels)
	if urlPath == "/api" {
		return
	}
	switch pathLevels[0] {
	case "blah":
		// Get  blah/all
		// Get	blah/show
		// Post blah/create
		// Post blah/destroy
		HandleBlah(ctx, pathLevels)
	case "message":
		HandleMessage(ctx, pathLevels)
	case "account":
		HandleAccount(ctx, pathLevels)
	}
}

func HandleBlah(ctx * webapp.Context, pathLevels []string) {
	if len(pathLevels) < 2 {
		return
	}
	if ! Auth(ctx) {
		ctx.Writer.Write([]byte("not authorization"))
		return
	}
	if ctx.Request.Method == "GET" {
		switch pathLevels[1] {
		case "all":
			ctx.Writer.Write(RenderAllBlahs())
		case "timeline":
			// blah/timeline?id=BLAH_ID
			id := ctx.Request.FormValue("id")
			ctx.Writer.Write(RenderTimeline(id))
		case "show":
			id := ctx.Request.FormValue("id")
			ctx.Writer.Write(RenderBlah(id))
		case "members":
			ctx.Writer.Write([]byte("blah/members"))
		}
	} else if ctx.Request.Method == "POST" {
		switch pathLevels[1] {
		case "create":
			// @TODO verify text
			text := ctx.Request.FormValue("text")
			var blah Blah
			blah.UpdateTime = time.Now().Unix()
			blah.CreateTime = blah.UpdateTime
			// @TODO verify author
			blah.AuthorId = ctx.Request.Header.Get("X-BLAHDY-NAME")
			blah.Text = text
			blahBytes, err := BlahdyDB.CreateBlah(&blah)
			if err == nil {
				ctx.Writer.Write(blahBytes)
			}
		case "destroy":
			id := ctx.Request.FormValue("id")
			blahBytes, err := BlahdyDB.DestroyBlah(id)
			if err == nil {
				ctx.Writer.Write(blahBytes)
			}
		}
	} else {
		// do nothing
	}
}

func HandleMessage(ctx * webapp.Context, pathLevels []string) {
	if len(pathLevels) < 2 {
		return
	}
	if ! Auth(ctx) {
		ctx.Writer.Write([]byte("not authorization"))
		return
	}
	if ctx.Request.Method == "GET" {

	} else if ctx.Request.Method == "POST" {
		switch pathLevels[1] {
		case "create":
			// @TODO verify 
			text := ctx.Request.FormValue("text")
			blahId := ctx.Request.FormValue("blah_id")
			var msg Message
			msg.CreateTime = time.Now().Unix()
			msg.BlahId = blahId
			// @TODO verify author
			msg.AuthorId = ctx.Request.Header.Get("X-BLAHDY-NAME")
			msg.Text = text
			msgBytes, err := BlahdyDB.CreateMessage(&msg)
			if err == nil {
				ctx.Writer.Write(msgBytes)
			}
		case "destroy":
			id := ctx.Request.FormValue("id")
			blahBytes, err := BlahdyDB.DestroyMessage(id)
			if err == nil {
				ctx.Writer.Write(blahBytes)
			}
		}
	} else {
		// do nothing
	}
}

func HandleBlahMember(ctx * webapp.Context, pathLevels []string) {
	fmt.Fprintf(ctx.Writer, "OK, I am members.")
}

func HandleAccount(ctx * webapp.Context, pathLevels []string) {
	if len(pathLevels) < 2 {
		return
	}
	if ctx.Request.Method == "GET" {
		return
	} else if ctx.Request.Method == "POST" {
		switch pathLevels[1] {
		case "auth":
			name := ctx.Request.FormValue("username")
			token := SHA256Sum(ctx.Request.FormValue("password"))
			if BlahdyDB.AuthUser(name, token) {
				ctx.Writer.Write([]byte(token))
			} else {
				ctx.Writer.Write([]byte("error"))
			}
		case "create":
			user := User{}
			user.Id = ctx.Request.FormValue("username")
			user.Name = ctx.Request.FormValue("name")
			user.Email = ctx.Request.FormValue("email")
			user.Bio = ctx.Request.FormValue("bio")
			user.Phone = ctx.Request.FormValue("phone")
			password := ctx.Request.FormValue("password")
			json, _ := BlahdyDB.CreateUser(&user, password)
			if json != nil {
				ctx.Writer.Write([]byte("ok"))
			} else {
				ctx.Writer.Write([]byte("error"))
			}
		}
	}

}




