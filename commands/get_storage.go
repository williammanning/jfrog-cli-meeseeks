package commands

import (
	"errors"
	"fmt"
	"net/http"
	"path"

	"github.com/jfrog/jfrog-cli-core/plugins/components"
	"github.com/jfrog/jfrog-cli-core/utils/config"
	"github.com/jfrog/jfrog-client-go/artifactory/httpclient"
	servicesutils "github.com/jfrog/jfrog-client-go/artifactory/services/utils"
	clientutils "github.com/jfrog/jfrog-client-go/utils"
	"github.com/williammanning/jfrog-cli-meeseeks/utils"
)

func GetStorage() components.Command {
	return components.Command{
		Name:        "storage",
		Description: "Get Artifactory Storage Info.",
		Aliases:     []string{"storage"},
		Arguments:   getStorageArguments(),
		Flags:       getStorageFlags(),
		EnvVars:     getStorageEnvVar(),
		Action: func(c *components.Context) error {
			return getStorageCmd(c)
		},
	}
}

func getStorageArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "server-id",
			Description: "Default Server ID from JFrog CLI Config",
		},
	}
}

type StorageConfiguration struct {
	server string
}

func getStorageFlags() []components.Flag {
	return []components.Flag{
		components.BoolFlag{
			Name:         "test",
			Description:  "Test connection.",
			DefaultValue: false,
		},
	}
}

func getStorageEnvVar() []components.EnvVar {
	return []components.EnvVar{
		{
			Name:        "HELLO_FROG_GREET_PREFIX",
			Default:     "A new greet from your plugin template: ",
			Description: "Adds a prefix to every greet.",
		},
	}
}

func getStorageCmd(c *components.Context) error {
	if !(len(c.Arguments) == 1 || len(c.Arguments) == 0) {
		return errors.New("wrong number of arguments. Expected 1 arguments, or 0 with build details passed as environment variables")
	}
	var conf = new(ArtifactoryInfoConfiguration)
	conf.server = c.Arguments[0]
	rtDetails, err := utils.GetRtDetails(c)

	if err != nil {
		return err
	}

	fmt.Println(rtDetails)
	getArtifactoryStorageAPI(rtDetails)
	return nil
}

func getArtifactoryStorageAPI(rtDetails *config.ArtifactoryDetails) string {
	fmt.Print("Get Details from ")

	artAuth, err := rtDetails.CreateArtAuthConfig()
	if err != nil {
		return ""
	}

	httpClientsDetails := artAuth.CreateHttpClientDetails()
	client, err := httpclient.ArtifactoryClientBuilder().SetServiceDetails(&artAuth).Build()

	if err != nil {
		return ""
	}

	fmt.Println(artAuth.GetUrl())
	restApi := path.Join("api", "storageinfo")

	requestFullUrl, err := servicesutils.BuildArtifactoryUrl(artAuth.GetUrl(), restApi, nil)

	fmt.Println("Getting system info from: ", requestFullUrl)

	resp, body, _, err := client.SendGet(requestFullUrl, true, &httpClientsDetails)

	if err != nil {
		return ""
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Print("Return Ok: " + resp.Status + " " + clientutils.IndentJson(body))
		errors.New("Artifactory response: " + resp.Status + "\n" + clientutils.IndentJson(body))
		return ""
	} else {
		//fmt.Print("Return Ok: " + resp.Status + " " + clientutils.IndentJson(body))
		jsonbody = clientutils.IndentJson(body)
		fmt.Print(jsonbody)
		return jsonbody
	}

}
