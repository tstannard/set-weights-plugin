package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"code.cloudfoundry.org/cli/plugin"
)

type UpdateRouteWeightPlugin struct{}

// Models
type V2Error struct {
	Description string `json:"description,omitempty"`
	ErrorCode   string `json:"error_code,omitempty"`
	Code        int    `json:"code,omitempty"`
}

type RouteEntityModel struct {
	Host      string `json:"host"`
	DomainURL string `json:"domain_url"`
}

type MetadataModel struct {
	Guid string `json:"guid"`
}

type DomainEntityModel struct {
	Name string `json: "name"`
}

type RouteMappingModel struct {
	V2Error
	Metadata MetadataModel     `json:"metadata"`
	Entity   DomainEntityModel `json:"entity"`
}

type DomainModel struct {
	V2Error
	Metadata MetadataModel     `json:"metadata"`
	Entity   DomainEntityModel `json:"entity"`
}

type RouteModel struct {
	V2Error
	Metadata MetadataModel    `json:"metadata"`
	Entity   RouteEntityModel `json:"entity"`
}

type RoutesModel struct {
	V2Error
	Routes []RouteModel `json:"resources"`
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
	appName, route, domain, host, routeWeight, err := parseArgs(args)
	if err != nil {
		panic(err)
		// return err
	}
	fmt.Println(route)
	fmt.Println(routeWeight)

	appObject, err := cliConnection.GetApp(appName)
	if err != nil {
		return err
	}
	appGUID := appObject.Guid
	fmt.Println(appGUID)
	routeObjectsJSON, err := cliConnection.CliCommandWithoutTerminalOutput("curl", fmt.Sprintf("v2/routes?q=host:%s;domain_guid:%s", host, domain))
	if err != nil {
		return err
	}
	routeObjectsJSONString := strings.Join(routeObjectsJSON, "")
	routes := RoutesModel{}
	err = json.Unmarshal([]byte(routeObjectsJSONString), &routes)
	if err != nil {
		return err
	}
	if routes.Description != "" {
		err = errors.New("Error unmarshalling JSON. booooooo")
	}
	routeGUID := ""
	routesLen := len(routes.Routes)
	for i := 0; i < routesLen; i++ {
		//curl domain endpoint
		route := routes.Routes[i]
		curlRouteGUID := routes.Routes[i].Metadata.Guid
		domainJSON, err := cliConnection.CliCommand("curl", route.Entity.DomainURL)
		if err != nil {
			return err
		}
		domainJSONString := strings.Join(domainJSON, "")
		curlDomain := DomainEntityModel{}
		err = json.Unmarshal([]byte(domainJSONString), &curlDomain)
		if err != nil {
			return err
		}
		if domain == curlDomain.Name {
			routeGUID = curlRouteGUID
		}
	}
	// will get multiple routes from the above query
	// json.unmarshall -- we will get an array of route objects
	// iterate through the array of route objects and curl the domain endpoints until the domain we get back matches the domain we got as input
	//		and get the route guid for that route object

	// then, curl the route mappings endpoint with the app guid and route guid to get the route mapping guid
	// use route mapping guid and update the weight

	routeMappingJSON, err := cliConnection.CliCommand("curl", fmt.Sprintf("/v3/route_mappings?app_guids=%s&route_guids=%s"), appGUID, routeGUID)
	if err != nil {
		return err
	}
	routeMappingJSONString := strings.Join(routeMappingJSON, "")
	routeMapping := RouteMappingModel{}
	err = json.Unmarshal([]byte(routeMappingJSONString), &routeMapping)
	if err != nil {
		return err
	}
	//, err = cliConnection.CliCommandWithoutTerminalOutput("curl", fmt.Sprintf("/v3/route_mappings/%s", routeMappingGUID), "-X", "PATCH", "-d", parsedWeight)
	return nil
}

func parseArgs(args []string) (appName string, route string, host string, domain string, weight int, err error) {
	parsedWeight, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		// TODO fix this
		panic(err)
		// return err
	}

	route = args[2]

	// get host
	idx := strings.Index(route, ".")
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

	return args[1], route, host, domain, int(parsedWeight), nil
}

func main() {
	plugin.Start(new(UpdateRouteWeightPlugin))
}
