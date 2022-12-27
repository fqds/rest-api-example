package model_test

import (
	"testing"

	"github.com/fqds/rest-api-example/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BefforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
