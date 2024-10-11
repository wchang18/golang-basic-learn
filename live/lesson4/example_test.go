package lesson4

import (
	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMockey(t *testing.T) {
	PatchConvey("TestProfileTask_CheckBudget", t, func() {
		Convey("TestProfileTask_CheckBudget_US", func() {
			So(CheckBudget(30, 30.1, "US"), ShouldBeFalse)
			So(CheckBudget(30, 32, "US"), ShouldBeTrue)
			So(CheckBudget(20, 19.1, "US"), ShouldBeFalse)
			So(CheckBudget(20, 19, "US"), ShouldBeTrue)
		})

		Convey("TestProfileTask_CheckBudget_JP", func() {
			So(CheckBudget(3000, 3010, "JP"), ShouldBeFalse)
			So(CheckBudget(3000, 3200, "JP"), ShouldBeTrue)
			So(CheckBudget(2000, 1910, "JP"), ShouldBeFalse)
			So(CheckBudget(2000, 1850, "JP"), ShouldBeTrue)
			So(CheckBudget(2000, 2000, "JP"), ShouldBeFalse)
		})
	})
}
