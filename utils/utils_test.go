package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHomeHandler(t *testing.T) {
	uuid := GenShortUUID()
	assert.NotNil(t, uuid)
	assert.IsType(t, "", uuid)
}
