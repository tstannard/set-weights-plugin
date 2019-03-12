package main

import (
	"fmt"

	"code.cloudfoundry.org/cli/plugin"
)

type UpdateRouteWeightPlugin struct{}

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
	appName, route, routeWeight := parseArgs(args)
	fmt.Println(route)
	fmt.Println(routeWeight)

	appObject, err := cliConnection.GetApp(appName)
	if err != nil {
		return err
	}
	appGUID := appObject.Guid
	fmt.Println(appGUID)
	//routeGUID := CliConnection.CliCommand()
	//routeMappingGUID := CliConnection.CliCommand("curl", fmt.Sprintf("/v3/route_mappings?app_guids=%s&route_guids=%s"), appGUID, routeGUID))
	//rawOutput, err2 := CliConnection.CliCommandWithoutTerminalOutput("curl", fmt.Sprintf("/v3/route_mappings/%s", routeMappingGUID), "-X", "PATCH", "-d", parsedWeight)
	return nil
}

func parseArgs(args []string) (appName string, route string, weight int) {
	return "", "", 0
}

func main() {
	plugin.Start(new(UpdateRouteWeightPlugin))
}
