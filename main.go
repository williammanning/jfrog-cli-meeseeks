package main

import (
	"fmt"
	"net/http"

	"github.com/jfrog/jfrog-cli-core/plugins"
	"github.com/jfrog/jfrog-cli-core/plugins/components"
	"github.com/jfrog/jfrog-cli-plugin-template/commands"
)

func main() {
	plugins.PluginMain(getApp())
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func getApp() components.App {
	app := components.App{}
	app.Name = "meeseeks"
	app.Description = "Can Do!"
	app.Version = "v0.1.0"
	app.Commands = getCommands()
	return app
}

func getCommands() []components.Command {
	return []components.Command{
		commands.GetHelloCommand()}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
