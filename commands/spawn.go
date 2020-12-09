package commands

import (
	"errors"

	"github.com/jfrog/jfrog-cli-core/plugins/components"
	"github.com/williammanning/jfrog-cli-meeseeks/utils"

	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Get Artifactory Information of current default instance in CLI

func SpawnMeeseekUI() components.Command {
	return components.Command{
		Name:        "spawn",
		Description: "Spawn a meeseek.",
		Aliases:     []string{"spawn"},
		Arguments:   spawnMeeseekArguments(),
		Flags:       spawnMeeseekFlags(),
		EnvVars:     spawnMeeseekEnvVar(),
		Action: func(c *components.Context) error {
			return spawnMeeseekCmd(c)
		},
	}
}

func spawnMeeseekArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "server-id",
			Description: "Default Server ID from JFrog CLI Config",
		},
	}
}

type spawnMeeseekConfiguration struct {
	server string
}

func spawnMeeseekFlags() []components.Flag {
	return []components.Flag{
		components.StringFlag{
			Name:        utils.ServerIdFlag,
			Description: "Artifactory server ID configured using the config command.",
		},
	}
}

func spawnMeeseekEnvVar() []components.EnvVar {
	return []components.EnvVar{
		{
			Name:        "HELLO_FROG_GREET_PREFIX",
			Default:     "A new greet from your plugin template: ",
			Description: "Adds a prefix to every greet.",
		},
	}
}

func spawnMeeseekCmd(c *components.Context) error {
	if !(len(c.Arguments) == 1 || len(c.Arguments) == 0) {
		return errors.New("wrong number of arguments. Expected 1 arguments, or 0 with build details passed as environment variables")
	}
	var conf = new(ArtifactoryInfoConfiguration)
	conf.server = c.Arguments[0]
	rtDetails, err := utils.GetRtDetails(c)
	fmt.Printf("conf.Server %s\n", conf.server)

	if err != nil {
		return err
	}

	fmt.Print(rtDetails)
//	connectArtifactoryRepo(rtDetails)

	//	fileServer := http.FileServer(http.Dir("./web"))
	http.HandleFunc("/", httpserver)
	fmt.Printf("Starting Meeseeks UI at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	return nil
}

// generate random data for bar chart
func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}

func barGraph(w http.ResponseWriter) {
	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Number of Artifacts",
		Subtitle: "Most recently published artifacts",
	}))

	// Put data into instance
	bar.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	// Where the magic happens
	//	f, _ := os.Create("bar.html")
	//	bar.Render(f) // Output to html file
	bar.Render(w) // Output to browser
}

func lineGraph(w http.ResponseWriter) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Cloud weekly transfers",
			Subtitle: "Line chart rendered by the http server this time",
		}))

	// Put data into instance
	line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

// generate random data for line chart
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func httpserver(w http.ResponseWriter, r *http.Request) {

//	repolist := connectArtifactoryRepo(rtDetails)

//	fmt.Printf("Processing chart metadata based on previous query..\n%s\n", storage)
	
	//	meeseeksWelcome(w)
	barGraph(w)
	lineGraph(w)

	//pieGraph(w)

	if r.URL.Path != "/bar.html" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.URL.Path != "/meeseeks.html" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

}
