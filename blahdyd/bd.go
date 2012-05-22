package main

import (
	"flag"
	"log"
	"os"
	"path"
	"github.com/shellex/tattoo/webapp"
)

var useFCGI = flag.Bool("fcgi", false, "Use FastCGI")

func main() {
	flag.Parse()

	rootPath, _ := os.Getwd()
	rootPath += ""
	app := webapp.App{}
	app.Log("App Starts", "OK")
	app.SetStaticPath("/static", path.Join(rootPath, "/static"))
	app.SetHandler("/api/", HandleRoot)

	BlahdyDB.Load(&app)
	LoadSamples()

	// Start Server.
	if *useFCGI {
		log.Printf("Server Starts(FastCGI): Listen on port %d\n", 4321)
		app.RunCGI(4321)
	} else {
		log.Printf("Server Starts: Listen on port %d\n", 8888)
		app.Run(8888)
	}
}
