package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"code.cloudfoundry.org/cli/plugin"
	"github.com/tstannard/set-weights-plugin/helpers"
)

type UpdateRouteWeightPlugin struct {
	CliClient cliClient
}

// Models
//go:generate counterfeiter -o fakes/cli_client.go --fake-name CliClient . cliClient
type cliClient interface {
	GetAppGUID(cliConnection plugin.CliConnection, appName string) (string, error)
	GetDomainGUID(cliConnection plugin.CliConnection, domainName string) (string, error)
	GetRouteGUID(cliConnection plugin.CliConnection, hostName string, domainGuid string) (string, error)
	GetRouteMappingGUID(cliConnection plugin.CliConnection, appGUID string, routeGUID string) (string, error)
	SetRouteMappingWeight(cliConnection plugin.CliConnection, routeMappingGUID string, routeWeight int) error
}

// Plugin
func (c *UpdateRouteWeightPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "update-route-weight" {
		err := c.SetWeight(cliConnection, args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func (c *UpdateRouteWeightPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "UpdateRouteWeightPlugin",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 1,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "update-route-weight",
				HelpText: "Update Route Weight help text",

				UsageDetails: plugin.Usage{
					Usage: "update-route-weight\n   cf update-route-weight APP_NAME ROUTE WEIGHT",
				},
			},
		},
	}
}

func (c *UpdateRouteWeightPlugin) SetWeight(cliConnection plugin.CliConnection, args []string) error {
	appName, host, domain, routeWeight, err := parseArgs(args)
	if err != nil {
		return err
	}

	appGUID, err := c.CliClient.GetAppGUID(cliConnection, appName)
	if err != nil {
		return err
	}

	domainGUID, err := c.CliClient.GetDomainGUID(cliConnection, domain)
	if err != nil {
		return err
	}

	routeGUID, err := c.CliClient.GetRouteGUID(cliConnection, host, domainGUID)
	if err != nil {
		return err
	}

	routeMappingGUID, err := c.CliClient.GetRouteMappingGUID(cliConnection, appGUID, routeGUID)
	if err != nil {
		return err
	}

	err = c.CliClient.SetRouteMappingWeight(cliConnection, routeMappingGUID, routeWeight)
	if err != nil {
		return err
	}
	fmt.Println("Successfully set route mapping weight.")

	return nil
}

func parseArgs(args []string) (appName string, host string, domain string, weight int, err error) {
	if len(args) < 3 {
		return "", "", "", 0, errors.New("not enough arguments provided")
	}

	parsedWeight, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		return "", "", "", 0, fmt.Errorf("could not parse route weight: %s", err)
	}

	route := args[2]
	idx := strings.Index(route, ".")
	if idx == -1 {
		return "", "", "", 0, fmt.Errorf("invalid route format (e.g hostname.domain)")
	}
	host = route[0:idx]
	domain = route[idx+1 : len(route)]

	return args[1], host, domain, int(parsedWeight), nil
}

func main() {
	plugin.Start(&UpdateRouteWeightPlugin{
		CliClient: &helpers.CliClient{},
	})
}
