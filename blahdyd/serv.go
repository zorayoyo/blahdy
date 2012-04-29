package main

import (
	"fmt"
	"strings"
	"time"
	"github.com/shellex/tattoo/webapp"
)

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
		// Get  blah/actions
		// Get	blah/show
		// Post blah/create
		// Post blah/destroy
		HandleBlah(ctx, pathLevels)
	case "actions":
		HandleAction(ctx, pathLevels)
	}
}

func HandleBlah(ctx * webapp.Context, pathLevels []string) {
	if len(pathLevels) < 2 {
		return
	}
	if ctx.Request.Method == "GET" {
		switch pathLevels[1] {
		case "all":
			ctx.Writer.Write(RenderJson(ctx, BlahdyDB.GetBlahs()))
		case "actions":
			ctx.Writer.Write([]byte("blah/actions"))
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
			blah.AuthorId = "0"
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

func HandleBlahMember(ctx * webapp.Context, pathLevels []string) {
}

func HandleAction(ctx * webapp.Context, pathLevels []string) {
	fmt.Fprintf(ctx.Writer, "OK, I am action.")
}



