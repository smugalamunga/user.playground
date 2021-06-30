package main

import (
	"context"
	"log"
	"sync"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/smugalamunga/user.playground/pkg/models"
	"github.com/smugalamunga/user.playground/pkg/persistence/mongodb"
)

func main() {
	implementation = mongodb.NewPersistenceImplementation()
	ctx := context.Background()

	// for i := 0; i < 100000; i++ {
	// 	CreateUser(ctx)
	// }

	// var wg sync.WaitGroup
	// for i := 0; i < 100000; i++ {
	// 	wg.Add(1)
	// 	go CreateUser(ctx, &wg)
	// }
	// wg.Wait()

	var username = "Reilly6239"
	user, err := implementation.FindUserByUsername(ctx, username)
	if err != nil {
		panic(err)
	}
	log.Println(user)

	var id = "5379ac9e3993bed550bb08fc"
	user, err = implementation.FindUserByUserID(ctx, id)
	if err != nil {
		panic(err)
	}
	log.Println(user)

}

var implementation *mongodb.PersistenceImplementation

func CreateUser(ctx context.Context) {
	id, err := implementation.CreateUser(ctx, GenerateFakeUser())

	if err != nil {
		panic(err)
	}
	log.Println("[", id, "]")
}

func CreateUserWG(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	id, err := implementation.CreateUser(ctx, GenerateFakeUser())

	if err != nil {
		panic(err)
	}
	log.Println("[", id, "]")
}

func GenerateFakeUser() *models.UserModel {
	var u models.UserModel
	gofakeit.Struct(&u)
	return &u

	// return &models.UserModel{
	// 	Username:     gofakeit.Username(),
	// 	Firstname:    gofakeit.FirstName(),
	// 	Lastname:     gofakeit.LastName(),
	// 	EmailAddress: gofakeit.Email(),
	// 	Password:     gofakeit.Password(true, true, true, true, false, 8),

	// 	Saved:   time.Now(),
	// 	Created: time.Now(),
	// }
}
