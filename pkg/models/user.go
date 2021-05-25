package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProfileID primitive.ObjectID `bson:"profile_id,omitempty"`

	Firstname    string `bson:"firstname,omitempty" fake:"{firstname}"`
	Lastname     string `bson:"lastname,omitempty" fake:"{lastname}"`
	EmailAddress string `bson:"email_address,omitempty" fake:"{email}"`

	Username string `bson:"username,omitempty" fake:"{username}"`
	Password string `bson:"password,omitempty" fake:"{password}"`
}
