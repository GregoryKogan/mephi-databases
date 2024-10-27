package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uint
	Name string
}

func main() {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: os.Getenv("DSN")}), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to database: %v\n", err)
		os.Exit(1)
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		fmt.Fprintf(os.Stderr, "failed to migrate database: %v\n", err)
		os.Exit(1)
	}

	user := User{Name: "John Doe"}
	if err := db.Create(&user).Error; err != nil {
		fmt.Fprintf(os.Stderr, "failed to create user: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("User created: %v\n", user)

	var users []User
	if err := db.Find(&users).Error; err != nil {
		fmt.Fprintf(os.Stderr, "failed to find users: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Users found: %v\n", users)
}
