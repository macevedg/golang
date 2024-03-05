package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
}

type Storage struct {
	users []*User
}

func (s *Storage) saveUser(user *User) {
	s.users = append(s.users, user)
}

func (s *Storage) retrieveUser(index int) *User {
	s.internalProcess()
	return s.users[index]
}

func (s *Storage) internalProcess() {
	time.Sleep(10 * time.Second)
}

func main() {
	storage := Storage{}
	numOfUsers := 100
	for i := 0; i < numOfUsers; i++ {
		user := &User{
			name: fmt.Sprintf("my name is %d", i),
		}

		storage.saveUser(user)
	}

	for i := 0; i < numOfUsers; i++ {
		user := storage.retrieveUser(i)
		fmt.Println(user.name)
	}

}
