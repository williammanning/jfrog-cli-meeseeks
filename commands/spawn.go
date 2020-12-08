package commands

import (
	"errors"

	"github.com/jfrog/jfrog-cli-core/plugins/components"
	"github.com/williammanning/jfrog-cli-meeseeks/utils"
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

	// ADD YOU WEB CODE HERE

	return nil
}
