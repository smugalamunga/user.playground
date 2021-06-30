package mongodb_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/smugalamunga/user.playground/pkg/models"
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
	p := mongodb.NewPersistenceImplementation()
	err := p.Initialize(context.Background())
	assert.NoError(t, err)
}

func TestPersistence_AddUser(t *testing.T) {
	p := mongodb.NewPersistenceImplementation()
	id, err := p.CreateUser(context.Background(), createTestUser())
	assert.NoError(t, err)
	assert.NotNil(t, id)
}

func createTestUser() *models.UserModel {
	var u models.UserModel
	gofakeit.Struct(&u)
	return &u

	// return &models.UserModel{
	// 	Username: gofakeit.Username(),
	// 	Password: gofakeit.Password(true, true, true, true, false, 8),
	// 	Saved:    time.Now(),
	// 	Created:  time.Now(),
	// }
}

func TestPersistence_Destroy(t *testing.T) {
	p := mongodb.NewPersistenceImplementation()
	err := p.Destroy(context.Background())
	assert.NoError(t, err)
}
