package lesson4

import (
	"fmt"
	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func Output(in string) string {
	fmt.Println("333")
	if strings.Contains("ed", in) {
		return in
	}
	return in + "ed"
}

type User struct {
	Id   int
	Name string
}

func (a *User) GetSome(in string) string {
	fmt.Println("666")
	if strings.Contains("s", in) {
		return in
	}
	return in + "s"
}

var Bar = 123

func TestUser(t *testing.T) {
	t.Log(new(User).GetSome("test"))
}

func TestConvey2(t *testing.T) {
	PatchConvey("test1", t, func() {

		Reset(func() {
			UnPatchAll()
		})
		Mock((*User).GetSome).To(func(string) string {
			return "mock2"
		}).Build()
		So(new(User).GetSome("test"), ShouldEqual, "mock2")

		MockValue(&Bar).To(2)
		So(Bar, ShouldEqual, 2)

		//Reset(func() {
		// UnPatchAll()
		//})

		//Mock(Output).To(func(string) string {
		//	return "mock1"
		//}).Build()
		Mock(Output).Return("mock1").Build()
		So(Output("test"), ShouldEqual, "mock1")

	})
}
