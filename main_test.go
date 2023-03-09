package main

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"testing"
)

var agent = fiber.AcquireAgent()

func TestRouter(t *testing.T) {
	req := agent.Request()
	req.Header.SetMethod("GET")
	req.SetRequestURI("http://localhost:3000")

	if err := agent.Parse(); err != nil {
		panic(err)
	}

	code, body, errs := agent.Bytes()
	fmt.Println(code, bytes.NewBuffer(body).String(), errs)
}
