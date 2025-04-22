package store

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

var (
	mu     sync.Mutex
	file   = "data/users.json"
	users  []User
	nextId = 1
)

func Load() error {
	mu.Lock()
	defer mu.Unlock()

	_, err := os.Stat(file) //can be replaced with "if _, err := os.Stat(file);os.IsNotExist(err)"
	if os.IsNotExist(err) {
		users = []User{}
		return save()
	}

	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &users)

}

func save() error {
	b, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(file, b, 0644)
}

func NewUser(u User) (User, error) {
	mu.Lock()
	defer mu.Unlock()

	u.Id = nextId
	nextId++
	users = append(users, u)
	return u, save()
}

func AllUsers() []User {
	mu.Lock()
	defer mu.Unlock()

	return append([]User(nil), users...)
}

func FindUser(id int) (User, error) {
	mu.Lock()
	defer mu.Unlock()

	for _, u := range users {
		if u.Id == id {
			return u, nil
		}
	}

	return User{}, fmt.Errorf("User not found")
}

func KillUser(i int) error {
	mu.Lock()
	defer mu.Unlock()

	index := -1
	for idx, u := range users {
		if u.Id == i {
			index = idx
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("User not found")
	}

	users = append(users[:index], users[index+1:]...)

	return save()
}
