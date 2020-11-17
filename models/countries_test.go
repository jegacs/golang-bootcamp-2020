package models

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRetrieveCountry(t *testing.T) {
	Convey("when creating new country ", t, func() {
		country := NewCountry("mx")
		err := country.Fetch()
		So(err, ShouldBeNil)
		serializedCountry, err := json.Marshal(country)
		So(err, ShouldBeNil)
		fmt.Println(string(serializedCountry))
		So(err, ShouldBeNil)

		err = country.Store()
		So(err, ShouldBeNil)
	})

}
