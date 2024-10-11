package lesson4

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSum3(t *testing.T) {
	assert.Equal(t, Sum[int](1, 2, 3), 6)
	assert.Equal(t, Sum[int8](1, 2, 3), int8(6))
	assert.Equal(t, Sum[int16](1, 2, 3), int16(6))
	assert.Equal(t, Sum[int32](1, 2, 3), int32(6))
	assert.Equal(t, Sum[int64](1, 2, 3), int64(6))

	assert.Equal(t, Sum[float64](1.1, 2.2, 3.3), 6.6)
	assert.Equal(t, Sum[float32](1.2, 2.3, 3.4), float32(6.9))

	//assert.
}

func TestConvey(t *testing.T) {
	Convey("string", t, func() {
		So("hello", ShouldEqual, "hello")
		So("hello", ShouldNotEqual, "world")
		Convey("int", func() {
			So(1, ShouldBeGreaterThan, 0)
			So(1, ShouldBeLessThan, 2)
			So(1, ShouldBeLessThanOrEqualTo, 1)
		})
		Convey("time", func() {
			So(time.Now(), ShouldHappenBetween, time.Now().Add(-1), time.Now().Add(1))
			Convey("time2", func() {
				So(time.Now(), ShouldHappenOnOrBetween, time.Now().Add(-1), time.Now().Add(1))
			})
		})
	})
}
