package main_test

import (
	"code.cloudfoundry.org/cli/plugin/pluginfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	//io_helpers "code.cloudfoundry.org/cli/testhelpers/io"
	. "github.com/tstannard/set-weights-plugin"
)

var _ = Describe("SetWeights", func() {
	var cliConn *pluginfakes.FakeCliConnection
	var plugin *UpdateRouteWeightPlugin

	BeforeEach(func() {
		cliConn = &pluginfakes.FakeCliConnection{}
		plugin = &UpdateRouteWeightPlugin{}
	})
	// simple test - a allowed weight  is set

	Context("given an app mapped to a route", func() {
		BeforeEach(func() {
			appName := "testAppName"
			route := "testRoute"
			routeWeight := "35"
			args := []string{"update-route-weight", appName, route, routeWeight}
			appObject := &plugin_models.GetAppModel{
				Guid: "testAppGuid",
			}

			cliConn.GetAppReturns(appObject, error)
			plugin.Run(cliConn, args)
		})
		It("will change the weight of the route mapping", func() {
			Expect(cliConn.GetAppCallCount()).To(Equal(1))
			err := plugin.SetWeight(cliConn, args)
			Expect(err).To(BeNil())
		})
	})
	It("successfully sets a weight", func() {
		//err := plugin.Run(cliConn, []string{"100"})
		// err := plugin.SetWeight("35")
		// Expect(err).To(BeNil())
	})
})
