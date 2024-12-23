package jobs

import (
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"github.com/joho/godotenv"
	"go-clean-arch/config"
	"log"
	"time"
)

func startScheduler() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	cfg := config.LoadConfig()

	s, err := gocron.NewScheduler()
	if err != nil {
		// handle error
	}

	jobs := NewTradeHistoryActivity(cfg)

	j, err := s.NewJob(
		gocron.DurationJob(1*time.Second),
		gocron.NewTask(jobs.Work2, "wwwwwwwwww"),
	)

	if err != nil {
		// handle error
	}
	// each job has a unique id
	fmt.Println(j.ID())
	// start the scheduler
	s.Start()
	// block until you are ready to shut down
	select {
	case <-time.After(time.Minute):
	}

	// when you're done, shut it down
	err = s.Shutdown()
	if err != nil {
		// handle error
	}
}
