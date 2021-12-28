package e2e

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("test example world", func(t *testing.T) {
		assert.Equal(t, 1, 1)
	})
}
