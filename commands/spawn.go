package commands

import (
	"errors"
	"os"

	"github.com/jfrog/jfrog-cli-core/plugins/components"
	"github.com/jfrog/jfrog-cli-core/utils/config"
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
var newArtDetails *config.ArtifactoryDetails

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
	newArtDetails = rtDetails

	if err != nil {
		return err
	}

	//	fmt.Println(rtDetails)   // Debugging to validate correct RT config
	//	connectArtifactoryRepo(rtDetails)

	fmt.Printf("\nChecking connectivity to Artifactory before spawning UI agent..\n")
	//var pingInfo = connectArtifactoryPing(newArtDetails.Url)
	connectArtifactoryPing(newArtDetails)

	//	fileServer := http.FileServer(http.Dir("./web"))
	http.HandleFunc("/", httpserver)
	fmt.Println("Spawning Meeseeks Dashboard at http://localhost:9033/admin/hello\n")
	//fmt.Println("Spawning Meeseeks Dashboard at http://localhost:9033/admin/vue\n")
	fmt.Println("Spawning Meeseeks UI agent at http://localhost:8080/  <--- Click me!\n")
	fmt.Println("Generating reports from selected queries on", conf.server)

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

type Bins struct {
	BinariesCount int
	BinariesSize  int
}

func barGraph(w http.ResponseWriter, artInfo string) {
	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Repos Size By Type",
		Subtitle: "Displays Repos by the Type and Size",
	}))

	// in := []byte(artInfo)
	// var raw map[string]interface{}
	// // if err := json.Unmarshal(in, &raw); err != nil {
	// //     panic(err)
	// // }
	// raw["count"] = 1
	// out, err := json.Marshal(raw)
	// if err != nil {
	// 	panic(err)
	// }
	// println(artInfo)

	// var s Bins
	// json.Unmarshal([]byte(str), &s)
	// fmt.Println(s)

	// type Bins struct {
	// 		Name   string 'json:"binariesSummary"'
	// 		Values []struct {
	// 		BinariesCount int 'json:"binariesCount,omitempty"'
	// 		BinariesSize int 'json:"binariesSize,omitempty"'
	// 		ArtifactsSize    int 'json:"artifactsSize,omitempty"'
	// 		Optimization   int 'json:"optimization,omitempty"'
	// 	} 'json:"values"'

	// b = []Bins{} // Summary
	// json.Unmarshal(in, &b)
	// println(b)

	//var rtStorage rtData
	//var dataArr = parse(artInfo);
	//dataArr[0].fileStoreSummary

	// storageType = dataArr[1].storageType
	// fmt.Println("Storage type = %s", storageType)
	// numArtifacts = dataArr[0].binariesSummary.artifactsCount

	// fmt.Println("bxxxxx + %s", artInfo)
	// err := json.Unmarshal([]byte(artInfo), &rtStorage)
	// if err != nil {
	// 	log.Println(err)
	// }
	// //fmt.Println(rtStorage.Name, rtStorage.fileStoreSummary, rtStorage.Id)

	// Put data into instance
	bar.SetXAxis([]string{"Docker", "Maven", "NPM", "Nuget", "PyPi", "RPM", "Debain"}).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	// Where the magic happens
	f, _ := os.Create("artifacts.html")
	bar.Render(f) // Output to html file
	bar.Render(w) // Output to browser
}

func lineGraph(w http.ResponseWriter) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Daily transfers",
			Subtitle: "Line chart rendered by the http server this time",
		}))

	// Put data into instance
	line.SetXAxis([]string{"Docker", "Maven", "NPM", "Nuget", "PyPi", "RPM", "Debain"}).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	f, _ := os.Create("transfers.html")
	line.Render(f) // Output to html file
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

	//var artInfo = getArtifactoryInfo(newArtDetails)

	// Run the Storage query for parsing to generate a report/graph
	var storageinfo = getArtifactoryStorageAPI(newArtDetails)
	//	fmt.Printf(storageinfo)

	// Run the Repo query for parsing to generate a report/graph
	//var repoInfo = getArtifactoryRepo(newArtDetails)

	barGraph(w, storageinfo)
	//lineGraph(w)

	// if r.URL.Path != "/artifacts.html" {
	// 	http.Error(w, "404 not found.", http.StatusNotFound)
	// 	return
	// }

	// if r.URL.Path != "/transfers.html" {
	// 	http.Error(w, "404 not found.", http.StatusNotFound)
	// 	return
	// }

	// if r.Method != "GET" {
	// 	http.Error(w, "Method is not supported.", http.StatusNotFound)
	// 	return
	// }

}
