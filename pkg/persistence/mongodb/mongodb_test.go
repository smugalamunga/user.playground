package mongodb_test

import (
	"context"
	"testing"

	"github.com/smugalamunga/user.playground/pkg/persistence/mongodb"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestPersistence_GenerateDatabase(t *testing.T) {
	// URL for Replica Set Testing
	// PersistenceConfiguration getUri(...) Does not handle Replica Set URL.
	// "mongodb://dbadmin:BkvFxU6Mp%25v%23@10.0.0.77:27017,10.0.0.77:27018,10.0.0.77:27019/test?authSource=admin&replicaSet=test-rs0&readPreference=primary&ssl=false"
	implementation := mongodb.NewPersistenceImplementation()
	err := implementation.Initialize(context.Background())
	//err := mongodb.Initialize()
	assert.NoError(t, err)
}

func TestPersistence_AddUsers(t *testing.T) {

}

func TestPersistence_Destroy(t *testing.T) {
	implementation := mongodb.NewPersistenceImplementation()
	err := implementation.Destroy(context.Background())
	//err := mongodb.Destroy()
	assert.NoError(t, err)
}
