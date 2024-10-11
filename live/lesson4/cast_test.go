package lesson4

import (
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCast(t *testing.T) {
	assert.Equal(t, int64(1), cast.ToInt64("1"))
	assert.Equal(t, 1.0, cast.ToFloat64("1"))
	assert.Equal(t, "1", cast.ToString(1))
	assert.Equal(t, true, cast.ToBool("true"))
	//cast.
}
