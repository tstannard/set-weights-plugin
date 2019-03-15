package main_test

import (
	"code.cloudfoundry.org/cli/plugin/pluginfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	//io_helpers "code.cloudfoundry.org/cli/testhelpers/io"
	. "github.com/tstannard/set-weights-plugin"
	"github.com/tstannard/set-weights-plugin/fakes"
)

var _ = Describe("SetWeights", func() {
	var cliConn *pluginfakes.FakeCliConnection
	var cliClient *fakes.CliClient

	var plugin *UpdateRouteWeightPlugin
	BeforeEach(func() {
		cliConn = &pluginfakes.FakeCliConnection{}
		cliClient = &fakes.CliClient{}
		plugin = &UpdateRouteWeightPlugin{
			CliClient: cliClient,
		}
	})
	// simple test - a allowed weight  is set

	Context("given an app mapped to a route", func() {
		var (
			appName     string
			route       string
			routeWeight string
		)
		BeforeEach(func() {
			appName = "testAppName"
			route = "test.routing.cf-app.com"
			routeWeight = "35"
			args := []string{"update-route-weight", appName, route, routeWeight}

			cliClient.GetAppGUIDReturns("testAppGuid", nil)
			cliClient.GetDomainGUIDReturns("testDomainGuid", nil)
			cliClient.GetRouteGUIDReturns("testRouteGuid", nil)
			cliClient.GetRouteMappingGUIDReturns("testRouteMappingGuid", nil)

			plugin.Run(cliConn, args)
		})

		It("changes the weight of the route mapping", func() {
			Expect(cliClient.GetAppGUIDCallCount()).To(Equal(1))

			passedCliConn, passedAppName := cliClient.GetAppGUIDArgsForCall(0)
			Expect(passedCliConn).To(Equal(cliConn))
			Expect(passedAppName).To(Equal(appName))

			Expect(cliClient.GetDomainGUIDCallCount()).To(Equal(1))

			passedCliConn, passedDomainName := cliClient.GetDomainGUIDArgsForCall(0)
			Expect(passedCliConn).To(Equal(cliConn))
			Expect(passedDomainName).To(Equal("routing.cf-app.com"))

			Expect(cliClient.GetRouteGUIDCallCount()).To(Equal(1))

			passedCliConn, passedHostname, passedDomainGuid := cliClient.GetRouteGUIDArgsForCall(0)
			Expect(passedCliConn).To(Equal(cliConn))
			Expect(passedHostname).To(Equal("test"))
			Expect(passedDomainGuid).To(Equal("testDomainGuid"))

			Expect(cliClient.GetRouteMappingGUIDCallCount()).To(Equal(1))

			passedCliConn, passedAppGuid, passedRouteGuid := cliClient.GetRouteMappingGUIDArgsForCall(0)
			Expect(passedCliConn).To(Equal(cliConn))
			Expect(passedAppGuid).To(Equal("testAppGuid"))
			Expect(passedRouteGuid).To(Equal("testRouteGuid"))

			Expect(cliClient.SetRouteMappingWeightCallCount()).To(Equal(1))

			passedCliConn, passedRouteMappingGuid, passedRouteWeight := cliClient.SetRouteMappingWeightArgsForCall(0)
			Expect(passedCliConn).To(Equal(cliConn))
			Expect(passedRouteMappingGuid).To(Equal("testRouteMappingGuid"))
			Expect(passedRouteWeight).To(Equal(35))
		})
	})
})
