package main

import (
	"fmt"
	"strconv"
	"strings"

	"code.cloudfoundry.org/cli/plugin"
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
		c.SetWeight(cliConnection, args)
	}
}

func (c *UpdateRouteWeightPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "UpdateRouteWeightPlugin",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
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
					Usage: "update-route-weight\n   cf update-route-weight weight",
				},
			},
		},
	}
}

func (c *UpdateRouteWeightPlugin) SetWeight(cliConnection plugin.CliConnection, args []string) error {
	appName, host, domain, routeWeight, err := parseArgs(args)
	if err != nil {
		panic(err)
		// return err
	}

	appGUID, err := c.CliClient.GetAppGUID(cliConnection, appName)
	if err != nil {
		panic(err)
		// return err
	}

	domainGUID, err := c.CliClient.GetDomainGUID(cliConnection, domain)
	if err != nil {
		panic(err)
		// return err
	}

	routeGUID, err := c.CliClient.GetRouteGUID(cliConnection, host, domainGUID)
	if err != nil {
		panic(err)
		// return err
	}

	routeMappingGUID, err := c.CliClient.GetRouteMappingGUID(cliConnection, appGUID, routeGUID)
	if err != nil {
		panic(err)
		// return err
	}

	err = c.CliClient.SetRouteMappingWeight(cliConnection, routeMappingGUID, routeWeight)
	if err != nil {
		panic(err)
		// return err
	}

	return nil
}

func parseArgs(args []string) (appName string, host string, domain string, weight int, err error) {
	parsedWeight, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		// TODO fix this
		panic(err)
	}

	route := args[2]

	// get host
	idx := strings.Index(route, ".")
	fmt.Println(route)
	if idx == -1 {
		panic("fuuuuck")
	}
	host = route[0:idx]

	// get domain
	idx = strings.Index(route, ".")
	if idx == -1 {
		panic("fuuuuck")
	}
	domain = route[idx+1 : len(route)]

	return args[1], host, domain, int(parsedWeight), nil
}

func main() {
	plugin.Start(new(UpdateRouteWeightPlugin))
}
