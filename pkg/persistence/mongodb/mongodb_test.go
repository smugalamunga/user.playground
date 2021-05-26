package mongodb_test

import (
	"github.com/smugalamunga/user.playground/pkg/persistence/mongodb"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestPersistence_GenerateDatabase(t *testing.T) {
	err := mongodb.Initialize()
	assert.NoError(t, err)
}

func TestPersistence_Destroy(t *testing.T) {
	err := mongodb.Destroy()
	assert.NoError(t, err)
}
