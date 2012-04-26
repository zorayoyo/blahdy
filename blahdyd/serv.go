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
	if urlPath == "/" {
		return
	}
	switch pathLevels[0] {
	case "blah":
		// Get  blah/all
		// Get  blah/actions
		// Post blah/members/destroy
		// Post blah/members/create
		// Get  blah/members
		// Get	blah/show
		HandleBlah(ctx, pathLevels)
	case "actions":
		HandleAction(ctx, pathLevels)
	}
}

func HandleBlah(ctx * webapp.Context, pathLevels []string) {
	if len(pathLevels) < 2 {
		return
	}
	switch pathLevels[1] {
	case "all":
		ctx.Writer.Write(RenderJson(ctx, "blah/all"))
	case "actions":
		ctx.Writer.Write([]byte("blah/actions"))
	case "members":
		ctx.Writer.Write([]byte("blah/members"))
	case "show":
		ctx.Writer.Write([]byte("blah/show"))
	}
}

func HandleAction(ctx * webapp.Context, pathLevels []string) {
	fmt.Fprintf(ctx.Writer, "OK, I am action.")
}



