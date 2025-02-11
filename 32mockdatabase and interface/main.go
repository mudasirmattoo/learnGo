package main

import (
	"fmt"
	"log"
)

// create a user
type User struct {
	ID   int
	name string
}

// create a mockdatabase using map for storing retrievable data, we can also connect any other databse
type MockDataStore struct {
	Users map[int]User //int here would be ID
}

// create a function to get user data  --> pass the MockDataStore user as receiver
func (md MockDataStore) GetUser(id int) (User, error) {
	user, ok := md.Users[id] //check if user ID is present and return the user

	if !ok {
		return User{}, fmt.Errorf("User with ID %v not found", id)
	}

	return user, nil
}

// save the user if not already present in the mockDataStore map
func (md MockDataStore) SaveUser(u User) error {
	_, ok := md.Users[u.ID] //check if user ID exists in map mockDataStore
	if ok {
		return fmt.Errorf("user ID %v already exists", u.ID)
	}
	md.Users[u.ID] = u
	return nil
}

// Datastore defines an interface for storing retrievable data.
// Any TYPE that implements this interface (has these two methods) is also of TYPE `Datastore`.
// This means any value of TYPE `MockDatastore` is also of TYPE `Datastore`
// This means we could have a value of TYPE `*sql.DB` and it can also be of TYPE `Datastore`
// This means we can write functions to take TYPE `Datastore` and use either of these:
// -- `MockDatastore`
// -- `*sql.DB`

type Datastore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

/*
A service is a long-running process that performs specific tasks in the background, typically without direct user interaction. Services are commonly used for handling requests, processing data, managing resources, or communicating between different parts of a system.
*/

// Service represents a service that stores retrievable data.
// We will attach methods to `Service` so that we can use either of these:
// -- `MockDatastore`
// -- `*sql.DB`
type Service struct {
	ds Datastore
}

func (s Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}

func (s Service) SaveUser(u User) error {
	return s.ds.SaveUser((u))
}

func main() {

	db := MockDataStore{
		Users: make(map[int]User),
	}

	srvc := Service{
		ds: db,
	}
	u1 := User{
		ID:   1,
		name: "James",
	}

	err := srvc.SaveUser(u1)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	u1Returned, err := srvc.GetUser(u1.ID)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	fmt.Println(u1)
	fmt.Println(u1Returned)
}
