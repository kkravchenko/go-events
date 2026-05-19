package main

import (
	"context"
	"log"

	"events/pkg/repository"
	"events/pkg/service"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	ctx := context.Background()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewEventRepository(db)
	eventService := service.NewEventService(repo)

	err = eventService.CreateEvent(ctx, "connector start", "INFO")
	if err != nil {
		log.Println("error:", err)
	}
}
