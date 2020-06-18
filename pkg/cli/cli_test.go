package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApp(t *testing.T) {
	app := App()
	assert.NotNil(t, app)
}
