package helpers_test

import (
	"errors"
	"fmt"

	plugin_models "code.cloudfoundry.org/cli/plugin/models"
	"code.cloudfoundry.org/cli/plugin/pluginfakes"
	"github.com/tstannard/set-weights-plugin/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("helpers", func() {
	var (
		cliConn *pluginfakes.FakeCliConnection
		client  *helpers.CliClient
	)

	BeforeEach(func() {
		cliConn = &pluginfakes.FakeCliConnection{}
		client = &helpers.CliClient{}
	})

	Describe("GetAppGUID", func() {
		var (
			appName string
		)

		BeforeEach(func() {
			appName = "testAppName"
			cliConn.GetAppReturns(plugin_models.GetAppModel{
				Guid: "testAppGuid",
			}, nil)
		})

		It("gets the app guid", func() {
			appGUID, err := client.GetAppGUID(cliConn, appName)
			Expect(err).NotTo(HaveOccurred())

			Expect(cliConn.GetAppArgsForCall(0)).To(Equal(appName))

			Expect(appGUID).To(Equal("testAppGuid"))
		})

		Context("when get app fails", func() {
			BeforeEach(func() {
				cliConn.GetAppReturns(plugin_models.GetAppModel{}, errors.New("bad things happened"))
			})

			It("returns an error", func() {
				_, err := client.GetAppGUID(cliConn, appName)
				Expect(err).To(MatchError("bad things happened"))
			})
		})
	})

	Describe("GetDomainGUID", func() {
		var (
			domainName string
		)

		BeforeEach(func() {
			domainName = "fake.com"
			cliConn.CliCommandWithoutTerminalOutputReturns([]string{`{
				 "resources": [
						{
							 "metadata": {
									"guid": "testDomainGuid"
							 }
						}
				 ]
			}`}, nil)
		})

		It("gets the domain guid", func() {
			domainGUID, err := client.GetDomainGUID(cliConn, domainName)
			Expect(err).NotTo(HaveOccurred())

			Expect(cliConn.CliCommandWithoutTerminalOutputArgsForCall(0)).To(Equal([]string{
				"curl",
				fmt.Sprintf("/v2/domains?q=name:%s", domainName),
			}))

			Expect(domainGUID).To(Equal("testDomainGuid"))
		})

		Context("when get domain fails", func() {
			BeforeEach(func() {
				cliConn.CliCommandWithoutTerminalOutputReturns([]string{}, errors.New("bad things happened"))
			})

			It("returns an error", func() {
				_, err := client.GetDomainGUID(cliConn, domainName)
				Expect(err).To(MatchError("bad things happened"))
			})
		})

		Context("when unmarshalling fails", func() {
			BeforeEach(func() {
				cliConn.CliCommandWithoutTerminalOutputReturns([]string{`%%%`}, nil)
			})

			It("returns an error", func() {
				_, err := client.GetDomainGUID(cliConn, domainName)
				Expect(err).To(MatchError(`invalid character '%' looking for beginning of value`))
			})
		})

		Context("when no resources are returned", func() {
			BeforeEach(func() {
				cliConn.CliCommandWithoutTerminalOutputReturns([]string{`{
					 "resources": []
				}`}, nil)
			})

			It("returns an error", func() {
				_, err := client.GetDomainGUID(cliConn, domainName)
				Expect(err).To(MatchError(fmt.Sprintf("no domain found for %s", domainName)))
			})
		})

		Context("when an error is returned from curl", func() {
			BeforeEach(func() {
				cliConn.CliCommandWithoutTerminalOutputReturns([]string{`{
					 "description": "something bad happened",
					 "error_code": "foo",
					 "code": 1
				}`}, nil)
			})

			It("returns an error", func() {
				_, err := client.GetDomainGUID(cliConn, domainName)
				Expect(err).To(MatchError("1: foo: something bad happened"))
			})
		})
	})

	Describe("GetRouteGUID", func() {
		var (
			hostname   string
			domainGUID string
		)

		BeforeEach(func() {
			hostname = "testHostname"
			domainGUID = "testDomainGuid"
			cliConn.CliCommandWithoutTerminalOutputReturns([]string{`{
				 "resources": [
						{
							 "metadata": {
									"guid": "testRouteGuid"
							 }
						}
				 ]
			}`}, nil)
		})

		It("gets the route guid", func() {
			routeGUID, err := client.GetRouteGUID(cliConn, hostname, domainGUID)
			Expect(err).NotTo(HaveOccurred())

			Expect(cliConn.CliCommandWithoutTerminalOutputArgsForCall(0)).To(Equal([]string{
				"curl",
				fmt.Sprintf("/v2/routes?q=host:%s;domain_guid:%s", hostname, domainGUID),
			}))

			Expect(routeGUID).To(Equal("testRouteGuid"))
		})

		Context("when get route fails", func() {
			BeforeEach(func() {
				cliConn.CliCommandWithoutTerminalOutputReturns([]string{}, errors.New("bad things happened"))
			})

			It("returns an error", func() {
				_, err := client.GetRouteGUID(cliConn, hostname, domainGUID)
				Expect(err).To(MatchError("bad things happened"))
			})
		})

		Context("when unmarshalling fails", func() {
			BeforeEach(func() {
				cliConn.CliCommandWithoutTerminalOutputReturns([]string{`%%%`}, nil)
			})

			It("returns an error", func() {
				_, err := client.GetRouteGUID(cliConn, hostname, domainGUID)
				Expect(err).To(MatchError(`invalid character '%' looking for beginning of value`))
			})
		})

		Context("when no resources are returned", func() {
			BeforeEach(func() {
				cliConn.CliCommandWithoutTerminalOutputReturns([]string{`{
					 "resources": []
				}`}, nil)
			})

			It("returns an error", func() {
				_, err := client.GetRouteGUID(cliConn, hostname, domainGUID)
				Expect(err).To(MatchError(fmt.Sprintf("no route found for %s and domain guid %s", hostname, domainGUID)))
			})
		})

		Context("when an error is returned from curl", func() {
			BeforeEach(func() {
				cliConn.CliCommandWithoutTerminalOutputReturns([]string{`{
					 "description": "something bad happened",
					 "error_code": "foo",
					 "code": 1
				}`}, nil)
			})

			It("returns an error", func() {
				_, err := client.GetRouteGUID(cliConn, hostname, domainGUID)
				Expect(err).To(MatchError("1: foo: something bad happened"))
			})
		})
	})

	Describe("GetRouteMappingGUID", func() {
		var (
			appGUID   string
			routeGUID string
		)

		BeforeEach(func() {
			appGUID = "testAppGuid"
			routeGUID = "testRouteGuid"
			cliConn.CliCommandWithoutTerminalOutputReturns([]string{`{
				 "resources": [
						{
							"guid": "testRouteMappingGuid"
						}
				 ]
			}`}, nil)
		})

		It("gets the route mapping guid", func() {
			routeMappingGUID, err := client.GetRouteMappingGUID(cliConn, appGUID, routeGUID)
			Expect(err).NotTo(HaveOccurred())

			Expect(cliConn.CliCommandWithoutTerminalOutputArgsForCall(0)).To(Equal([]string{
				"curl",
				fmt.Sprintf("/v3/route_mappings?app_guids=%s&route_guids=%s", appGUID, routeGUID),
			}))

			Expect(routeMappingGUID).To(Equal("testRouteMappingGuid"))
		})

		Context("when get route mapping fails", func() {
			BeforeEach(func() {
				cliConn.CliCommandWithoutTerminalOutputReturns([]string{}, errors.New("bad things happened"))
			})

			It("returns an error", func() {
				_, err := client.GetRouteMappingGUID(cliConn, appGUID, routeGUID)
				Expect(err).To(MatchError("bad things happened"))
			})
		})

		Context("when unmarshalling fails", func() {
			BeforeEach(func() {
				cliConn.CliCommandWithoutTerminalOutputReturns([]string{`%%%`}, nil)
			})

			It("returns an error", func() {
				_, err := client.GetRouteMappingGUID(cliConn, appGUID, routeGUID)
				Expect(err).To(MatchError(`invalid character '%' looking for beginning of value`))
			})
		})

		Context("when no resources are returned", func() {
			BeforeEach(func() {
				cliConn.CliCommandWithoutTerminalOutputReturns([]string{`{
					 "resources": []
				}`}, nil)
			})

			It("returns an error", func() {
				_, err := client.GetRouteMappingGUID(cliConn, appGUID, routeGUID)
				Expect(err).To(MatchError(fmt.Sprintf("no route mapping found for app guid %s and route guid %s", appGUID, routeGUID)))
			})
		})

		Context("when an error is returned from curl", func() {
			BeforeEach(func() {
				cliConn.CliCommandWithoutTerminalOutputReturns([]string{`{
					 "description": "something bad happened",
					 "error_code": "foo",
					 "code": 1
				}`}, nil)
			})

			It("returns an error", func() {
				_, err := client.GetRouteMappingGUID(cliConn, appGUID, routeGUID)
				Expect(err).To(MatchError("1: foo: something bad happened"))
			})
		})
	})

	Describe("SetRouteMappingWeight", func() {
		var (
			routeMappingGUID string
			routeWeight      int
		)

		BeforeEach(func() {
			routeMappingGUID = "testRouteMappingGuid"
			cliConn.CliCommandWithoutTerminalOutputReturns([]string{}, nil)
		})

		It("sets the route mapping weight", func() {
			err := client.SetRouteMappingWeight(cliConn, routeMappingGUID, routeWeight)
			Expect(err).NotTo(HaveOccurred())

			Expect(cliConn.CliCommandWithoutTerminalOutputArgsForCall(0)).To(Equal([]string{
				"curl",
				fmt.Sprintf("/v3/route_mappings/%s", routeMappingGUID),
				"-X", "PATCH",
				"-d", fmt.Sprintf(`{"weight": %d}`, routeWeight),
			}))
		})
	})
})
