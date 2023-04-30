package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type UserRepository interface {
	Create(user string) error
	GetAll() ([]string, error)
}

type RAMUserRepository struct {
	users []string
}

func (p *RAMUserRepository) Create(user string) error {
	p.users = append(p.users, user)

	return nil
}

func (p *RAMUserRepository) GetAll() ([]string, error) {
	return p.users, nil
}

type LocalUserRepository struct {
	Path string
}

func (l *LocalUserRepository) Create(user string) error {
	file, openFileErr := os.OpenFile(l.Path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if openFileErr != nil {
		return fmt.Errorf("can't open file: %w", openFileErr)
	}
	defer file.Close()

	_, writeFileErr := file.WriteString(user + "\n")
	if writeFileErr != nil {
		return fmt.Errorf("can't write to file: %w", writeFileErr)
	}

	return nil
}

func (l *LocalUserRepository) GetAll() ([]string, error) {
	file, openFileErr := os.Open(l.Path)
	if openFileErr != nil {
		return nil, fmt.Errorf("can't open file: %w", openFileErr)
	}
	defer file.Close()

	usersBytes, readFileErr := io.ReadAll(file)
	if readFileErr != nil {
		return nil, fmt.Errorf("can't read file: %w", readFileErr)
	}

	users := strings.Split(string(usersBytes), "\n")

	return users, nil
}

type Service struct {
	repo UserRepository
}

func (s *Service) GetUsers() ([]string, error) {
	return s.repo.GetAll()
}

func (s *Service) CreateUser(user string) error {
	return s.repo.Create(user)
}

func main() {
	var repo UserRepository

	repo = &LocalUserRepository{"users.txt"}
	s := &Service{repo: repo}

	if createUserErr := s.CreateUser("user1"); createUserErr != nil {
		log.Fatalln(createUserErr)
	}
	if createUserErr := s.CreateUser("user2"); createUserErr != nil {
		log.Fatalln(createUserErr)
	}

	users, getUsersErr := s.GetUsers()
	if getUsersErr != nil {
		log.Fatalln(getUsersErr)
	}

	fmt.Println(users)

	repo = &RAMUserRepository{}
	s.repo = repo

	if createUserErr := s.CreateUser("new user1"); createUserErr != nil {
		log.Fatalln(createUserErr)
	}
	if createUserErr := s.CreateUser("new user1"); createUserErr != nil {
		log.Fatalln(createUserErr)
	}

	users, getUsersErr = s.GetUsers()
	if getUsersErr != nil {
		log.Fatalln(getUsersErr)
	}
	fmt.Println(users)
}
