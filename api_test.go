package octosend

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func testGetAPI() *OctosendAPI {
	api, err := NewAPI(os.Getenv("OCTOSEND_USERNAME"), os.Getenv("OCTOSEND_PASSWORD"))
	if err != nil {
		return nil
	}
	return api
}

func TestNewAPIByToken(t *testing.T) {
	Convey("Testing NewAPIByToken", t, func() {
		api := NewAPIByToken("dummy-token")
		So(api, ShouldNotBeNil)
		So(api.Token, ShouldEqual, "dummy-token")
	})
}

func TestNewAPI(t *testing.T) {
	Convey("Testing NewAPI", t, func() {
		api, err := NewAPI(os.Getenv("OCTOSEND_USERNAME"), os.Getenv("OCTOSEND_PASSWORD"))
		So(err, ShouldBeNil)
		So(api, ShouldNotBeNil)
		So(api.Token, ShouldNotBeEmpty)
	})
}
