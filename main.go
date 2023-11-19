package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type DB struct {
	counter int
	mu      sync.Mutex
	c       chan (int)
}

func (db *DB) Inc() {
	db.mu.Lock()
	// Some Expensive Computation
	db.counter++
	time.Sleep(time.Second * 2)
	db.c <- db.counter
	db.mu.Unlock()
}

func main() {
	db := &DB{
		counter: 0,
		mu:      sync.Mutex{},
		c:       make(chan int),
	}
	app := fiber.New(fiber.Config{
		Concurrency: 256 * 1024,
		AppName:     "Fiber",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/INC", func(c *fiber.Ctx) error {
		go db.Inc()

		val := <-db.c

		return c.SendString(fmt.Sprintf("%+v", val))
	})

	listener, err := net.Listen("tcp", ":3000")

	if err != nil {
		panic(err)
	}

	app.Listener(listener)
}
