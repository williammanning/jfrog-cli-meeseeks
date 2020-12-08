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

func GetRepoList() components.Command {
	return components.Command{
		Name:        "repolist",
		Description: "Get Artifactory Repo Info.",
		Aliases:     []string{"repo"},
		Arguments:   getRepoListArguments(),
		Flags:       getRepoListFlags(),
		EnvVars:     getRepoListEnvVar(),
		Action: func(c *components.Context) error {
			return getRepoListCmd(c)
		},
	}
}

func getRepoListArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "server-id",
			Description: "Default Server ID from JFrog CLI Config",
		},
	}
}

type RepoListConfiguration struct {
	server string
}

func getRepoListFlags() []components.Flag {
	return []components.Flag{
		components.BoolFlag{
			Name:         "test",
			Description:  "Test connection.",
			DefaultValue: false,
		},
	}
}

func getRepoListEnvVar() []components.EnvVar {
	return []components.EnvVar{
		{
			Name:        "HELLO_FROG_GREET_PREFIX",
			Default:     "A new greet from your plugin template: ",
			Description: "Adds a prefix to every greet.",
		},
	}
}

func getRepoListCmd(c *components.Context) error {
	if !(len(c.Arguments) == 1 || len(c.Arguments) == 0) {
		return errors.New("wrong number of arguments. Expected 1 arguments, or 0 with build details passed as environment variables")
	}
	var conf = new(ArtifactoryInfoConfiguration)
	conf.server = c.Arguments[0]
	rtDetails, err := utils.GetRtDetails(c)

	if err != nil {
		return err
	}

	fmt.Print(rtDetails)
	connectArtifactoryRepo(rtDetails)
	return nil
}

func connectArtifactoryRepo(rtDetails *config.ArtifactoryDetails) {
	fmt.Print("Get Details")
	artAuth, err := rtDetails.CreateArtAuthConfig()

	if err != nil {
		return
	}

	httpClientsDetails := artAuth.CreateHttpClientDetails()
	client, err := httpclient.ArtifactoryClientBuilder().SetServiceDetails(&artAuth).Build()

	if err != nil {
		return
	}

	fmt.Print(artAuth.GetUrl())
	restApi := path.Join("api", "repositories")

	requestFullUrl, err := servicesutils.BuildArtifactoryUrl(artAuth.GetUrl(), restApi, nil)

	fmt.Print("Getting system info from: ", requestFullUrl)

	resp, body, _, err := client.SendGet(requestFullUrl, true, &httpClientsDetails)

	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Print("Return Ok: " + resp.Status + " " + clientutils.IndentJson(body))
		errors.New("Artifactory response: " + resp.Status + "\n" + clientutils.IndentJson(body))
		return
	} else {
		fmt.Print("Return Ok: " + resp.Status + " " + clientutils.IndentJson(body))
	}

}
