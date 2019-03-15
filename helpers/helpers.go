package helpers

import (
	"encoding/json"
	"fmt"
	"strings"

	"code.cloudfoundry.org/cli/plugin"
)

type V2Error struct {
	Description string `json:"description,omitempty"`
	ErrorCode   string `json:"error_code,omitempty"`
	Code        int    `json:"code,omitempty"`
}

type RouteMappingsModel struct {
	V2Error
	Resources []RouteMappingResourceModel `json:"resources"`
}

type DomainModel struct {
	V2Error
	Resources []DomainResourceModel `json:"resources"`
}

type RoutesModel struct {
	V2Error
	Resources []RouteResourceModel `json:"resources"`
}

type DomainResourceModel struct {
	Metadata DomainMetadataModel `json:"metadata"`
}

type RouteMappingResourceModel struct {
	Guid string `json:"guid"`
}

type RouteResourceModel struct {
	Metadata RouteMetadataModel `json:"metadata"`
}

type DomainMetadataModel struct {
	Guid string `json:"guid"`
}

type RouteMetadataModel struct {
	Guid string `json:"guid"`
}

type CliClient struct{}

func (c *CliClient) GetAppGUID(cliConnection plugin.CliConnection, appName string) (string, error) {
	appObject, err := cliConnection.GetApp(appName)
	if err != nil {
		return "", err
	}
	appGUID := appObject.Guid
	return appGUID, nil
}

func (c *CliClient) GetDomainGUID(cliConnection plugin.CliConnection, domainName string) (string, error) {
	domainJSON, err := cliConnection.CliCommandWithoutTerminalOutput("curl", fmt.Sprintf("/v2/domains?q=name:%s", domainName))
	if err != nil {
		return "", err
	}
	domain := DomainModel{}
	domainJSONString := strings.Join(domainJSON, "")
	err = json.Unmarshal([]byte(domainJSONString), &domain)
	if err != nil {
		return "", err
	}
	if domain.Description != "" {
		return "", fmt.Errorf("%d: %s: %s", domain.Code, domain.ErrorCode, domain.Description)
	}
	if len(domain.Resources) == 0 {
		return "", fmt.Errorf("no domain found for %s", domainName)
	}
	return domain.Resources[0].Metadata.Guid, nil
}

func (c *CliClient) GetRouteGUID(cliConnection plugin.CliConnection, hostName string, domainGuid string) (string, error) {
	routeObjectsJSON, err := cliConnection.CliCommandWithoutTerminalOutput("curl", fmt.Sprintf("/v2/routes?q=host:%s;domain_guid:%s", hostName, domainGuid))
	if err != nil {
		return "", err
	}
	routeObjectsJSONString := strings.Join(routeObjectsJSON, "")
	routes := RoutesModel{}
	err = json.Unmarshal([]byte(routeObjectsJSONString), &routes)
	if err != nil {
		return "", err
	}
	if routes.Description != "" {
		return "", fmt.Errorf("%d: %s: %s", routes.Code, routes.ErrorCode, routes.Description)
	}
	if len(routes.Resources) == 0 {
		return "", fmt.Errorf("no route found for %s and domain guid %s", hostName, domainGuid)
	}
	return routes.Resources[0].Metadata.Guid, nil
}

func (c *CliClient) GetRouteMappingGUID(cliConnection plugin.CliConnection, appGUID string, routeGUID string) (string, error) {
	routeMappingsJSON, err := cliConnection.CliCommandWithoutTerminalOutput("curl", fmt.Sprintf("/v3/route_mappings?app_guids=%s&route_guids=%s", appGUID, routeGUID))
	if err != nil {
		return "", err
	}
	routeMappingsJSONString := strings.Join(routeMappingsJSON, "")
	routeMappings := RouteMappingsModel{}
	err = json.Unmarshal([]byte(routeMappingsJSONString), &routeMappings)
	if err != nil {
		return "", err
	}
	if routeMappings.Description != "" {
		return "", fmt.Errorf("%d: %s: %s", routeMappings.Code, routeMappings.ErrorCode, routeMappings.Description)
	}
	if len(routeMappings.Resources) == 0 {
		return "", fmt.Errorf("no route mapping found for app guid %s and route guid %s", appGUID, routeGUID)
	}
	routeMappingGUID := routeMappings.Resources[0].Guid
	return routeMappingGUID, nil
}

func (c *CliClient) SetRouteMappingWeight(cliConnection plugin.CliConnection, routeMappingGUID string, weight int) error {
	routeMappingsJSON, err := cliConnection.CliCommandWithoutTerminalOutput("curl", fmt.Sprintf("/v3/route_mappings/%s", routeMappingGUID), "-X", "PATCH", "-d", fmt.Sprintf(`{"weight": %d}`, weight))
	if err != nil {
		return err
	}
	routeMappingsJSONString := strings.Join(routeMappingsJSON, "")
	v2Error := V2Error{}
	err = json.Unmarshal([]byte(routeMappingsJSONString), &v2Error)
	if err != nil {
		return err
	}
	if v2Error.Description != "" {
		return fmt.Errorf("%d: %s: %s", v2Error.Code, v2Error.ErrorCode, v2Error.Description)
	}
	return err
}
