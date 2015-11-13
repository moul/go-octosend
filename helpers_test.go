package octosend

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

/*
func TestOctosendAPI_GetSpoolerMessage(t *testing.T) {
	Convey("Testing OctosendAPI.GetSpoolerMessage", t, func() {
		api := testGetAPI()
		spoolerMessage, err := api.GetSpoolerMessage(os.Getenv("OCTOSEND_SPOOLER_TOKEN"))
		So(err, ShouldBeNil)
		So(spoolerMessage, ShouldNotBeNil)
	})
}
*/

func TestOctosendAPI_GetDomain(t *testing.T) {
	Convey("Testing OctosendAPI.GetDomain", t, func() {
		api := testGetAPI()
		domainName := os.Getenv("OCTOSEND_DOMAIN")
		domain, err := api.GetDomain(domainName)
		So(err, ShouldBeNil)
		So(domain, ShouldNotBeNil)
		So(domain.Name, ShouldEqual, domainName)
	})
}
